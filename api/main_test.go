package api

import (
	"testing"
	"time"

	"dolyn157.dev/simplebank/config"
	db "dolyn157.dev/simplebank/db/sqlc"
	"github.com/stretchr/testify/require"
)

const (
	dbSource      = "postgresql://root:secret@localhost:15432/simple_bank?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
)

func newTestServer(t *testing.T, store *db.Store) *Server {

	config := config.Config{
		TokenSymmetricKey:   "11451419198103641145141919810364",
		AccessTokenDuration: time.Minute * 3,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}
