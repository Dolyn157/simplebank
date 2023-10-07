package db

import (
	"context"
	"testing"
	"time"

	"dolyn157.dev/simplebank/db/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)

	require.NotZero(t, user.Username)
	require.NotZero(t, user.CreatedAt)

	return user
}
func TestCreateUser(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	returnedUserInfo, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, returnedUserInfo)

	require.Equal(t, user1.Username, returnedUserInfo.Username)
	require.Equal(t, user1.HashedPassword, returnedUserInfo.HashedPassword)
	require.Equal(t, user1.FullName, returnedUserInfo.FullName)
	require.Equal(t, user1.Email, returnedUserInfo.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, returnedUserInfo.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, returnedUserInfo.CreatedAt, time.Second)
}
