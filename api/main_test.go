package api

import (
	"os"
	"testing"
	"time"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, library db.Library) *Server {
	config := util.Config{AccessTokenDuration: time.Minute}
	server, err := NewServer(library, config, nil)
	require.NoError(t, err)

	return server
}

// According to docs, TestMain will be invoked before running the main test. property m can be accessed in the test
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
