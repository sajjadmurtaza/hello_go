package database

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %w, rb err: %w", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

type AccountTxParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type AccountTxResult struct {
	Account Account `json:"account"`
}

func (store *Store) AccountTx(ctx context.Context, arg AccountTxParams) (AccountTxResult, error) {
	var result AccountTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Account, err = q.CreateAccount(ctx, CreateAccountParams{
			Owner:    arg.Owner,
			Balance:  arg.Balance,
			Currency: arg.Owner,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
