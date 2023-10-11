package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"dolyn157.dev/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	user := CreateRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	returnedAccInfo, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, returnedAccInfo)

	require.Equal(t, account1.ID, returnedAccInfo.ID)
	require.Equal(t, account1.Owner, returnedAccInfo.Owner)
	require.Equal(t, account1.Balance, returnedAccInfo.Balance)
	require.Equal(t, account1.Currency, returnedAccInfo.Currency)
	require.WithinDuration(t, account1.CreatedAt, returnedAccInfo.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: account1.Balance,
	}

	returnedAccInfo, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, returnedAccInfo)

	require.Equal(t, account1.ID, returnedAccInfo.ID)
	require.Equal(t, account1.Owner, returnedAccInfo.Owner)
	require.Equal(t, arg.Balance, returnedAccInfo.Balance)
	require.Equal(t, account1.Currency, returnedAccInfo.Currency)
	require.WithinDuration(t, account1.CreatedAt, returnedAccInfo.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	returnedAccInfo, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, returnedAccInfo)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = CreateRandomAccount(t)
	}
	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}
	accountList1, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accountList1)
	fmt.Printf("accoutList: %v", accountList1)

	for _, account := range accountList1 {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
