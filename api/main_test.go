package api

import (
	"testing"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, library db.Library) *Server {
	config := util.Config{}
	server, err := NewServer(library, config)
	require.NoError(t, err)

	return server
}
