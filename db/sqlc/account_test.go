package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// function for creating an account with random parameters
func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	fmt.Println(arg)
	account, err := testQueries.CreateAccount(context.Background(), arg) // testQueries has db connection
	log.Println("err", err)
	log.Println("err", account)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// testing account creation
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

// testing account creation
func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

//testing update account

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg1 := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg1)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account1.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.NotEqual(t, account1.Balance, account2.Balance)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// testing delete account
func TestDeleteAccount(t *testing.T) {

	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {

	// for i := 0; i < 10; i++ {
	// 	createRandomAccount(t)
	// }

	arg := ListAccountsParams{
		// Owner:  "dedafe",
		Limit:  1,
		Offset: 1,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	fmt.Println(accounts)

	require.NoError(t, err)
	require.Len(t, accounts, 1)

}
