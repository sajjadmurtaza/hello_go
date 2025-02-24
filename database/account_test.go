package database

import (
	"context"
	"database/sql"
	"fmt"
	"hello_go/database"
	"hello_go/util"
	"log"
	"testing"

	_ "github.com/lib/pq" // keep it for linting even not use
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
)

var testDBQueries *database.Queries

const (
	testDbDriver = "postgres"
	testDbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

// Test _prefix for every test start with T
func TestCreateAccount(t *testing.T) {
	fmt.Println(testDBQueries)
	// Open a connection to the database
	connection, err := sql.Open(testDbDriver, testDbSource)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	// Initialize the testQueries variable using db.New
	testDBQueries = database.New(connection)

	arg := database.CreateAccountParams{
		Owner:    util.RandomString(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	fmt.Println(arg)

	account, err := testDBQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)

	require.NotZero(t, account.ID)

}
