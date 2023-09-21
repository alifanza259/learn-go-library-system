package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	member db.Member,
	duration time.Duration,
) {
	token, payload, err := tokenMaker.CreateToken(member, duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	// Create new server instance, for testing
	server := newTestServer(t, nil)

	// Define new route "/auth" with method "GET", the controller is only returning 200 OK, but there is middleware authMiddleware that will be run
	server.router.GET(
		"/",
		authMiddleware(server.tokenMaker),
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{})
		},
	)

	// Create response recorder instance
	recorder := httptest.NewRecorder()

	// Create http request instance
	request, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err)
	member := db.Member{
		ID:        uuid.New(),
		Email:     "alifanza259@gmail.com",
		FirstName: "Alif",
	}
	func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
		addAuthorization(t, request, tokenMaker, authorizationTypeBearer, member, time.Minute)
	}(t, request, server.tokenMaker)

	server.router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
}
