package database

import (
	"context"
	"database/sql"
	"fmt"
	"hello_go/database"
	"hello_go/util"
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

var testStoreDBQueries *database.Queries

const (
	testStoreDbDriver = "postgres"
	testSoreDbSource  = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestAccountTx(t *testing.T) {

	connection, err := sql.Open(testStoreDbDriver, testSoreDbSource)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	store := database.NewStore(connection)

	fmt.Println(store)

	n := 5

	errors := make(chan error)
	results := make(chan database.AccountTxResult)

	for i := 0; i < n; i++ {
		go func() {
			// this function is running into different go routine
			// so no gruntee
			// that's why use channel
			//  channel use to connect cuncurrent go routine
			// and share data
			result, err := store.AccountTx(context.Background(), database.AccountTxParams{
				Owner:    util.RandomString(),
				Balance:  util.RandomMoney(),
				Currency: util.RandomCurrency(),
			})

			errors <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errors
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		fmt.Println(result)
	}
}
