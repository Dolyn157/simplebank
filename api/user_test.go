package api

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
	"time"

	db "dolyn157.dev/simplebank/db/sqlc"
	"dolyn157.dev/simplebank/utils"
	"github.com/stretchr/testify/require"
)

// randomuser is a helper function to create a random user for testing

func createRandomUser(t *testing.T) (user db.User, password string) {

	password = utils.RandomString(6)
	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.WithinDuration(t, user.PasswordChangedAt, gotUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, gotUser.CreatedAt, time.Second)
}

/*
func TestLoginUserAPI(t *testing.T) {
	user, password := createRandomUser(t)

	testCases := []struct {
		name          string
		reqBody       gin.H
		buildStubs    func(store *db.Store)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			reqBody: gin.H{
				"username": user.Username,
				"password": password,
			},
			buildStubs: func(store *db.Store) {
				store.GetUser()
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "InvalidUsername",
			reqBody: gin.H{
				"username": "",
				"password": password,
			},
			buildStubs: func(store *db.Store) {

	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			conn, err := sql.Open("postgres", dbSource)
			if err != nil {
				log.Fatal("cannot connected to the server.", err)
			}

			store := db.NewStore(conn)
			server := newTestServer(t, store)

			loginPath := "/users/login"
			server.router.GET(
				authPath,

				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)
			tc.buildStubs(store)

			url := fmt.Sprintf("/login")
			req, err := http.NewRequest(http.MethodPost, url, nil)
			require.NoError(t, err)

			if tc.reqBody != nil {
				jsonBody, err := json.Marshal(tc.reqBody)
				require.NoError(t, err)
				req.Body = io.NopCloser(bytes.NewReader(jsonBody))
			}

			recorder := httptest.NewRecorder()
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}
}
*/
