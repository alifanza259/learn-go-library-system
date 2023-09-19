package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/alifanza259/learn-go-library-system/db/mock"
	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/token"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// Unit Test controller only
func TestGetMember(t *testing.T) {
	member := db.Member{
		ID:        uuid.MustParse("051cfcb3-e699-43aa-b9fa-f4d68abaafd0"),
		Email:     "alif@gmail.com",
		FirstName: "Alif",
	}

	ctrl := gomock.NewController(t)
	library := mockdb.NewMockLibrary(ctrl)

	library.EXPECT().
		GetMember(gomock.Any(), gomock.Any()).
		Times(1).
		Return(member, nil)

	server := newTestServer(t, library)
	recorder := httptest.NewRecorder()

	// url := fmt.Sprintf("/member/%d", member.ID)
	// request, err := http.NewRequest(http.MethodGet, url, nil)
	// require.NoError(t, err)

	// server.router.ServeHTTP(recorder)

	c := gin.CreateTestContextOnly(recorder, &gin.Engine{})
	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: member.ID.String(),
		},
	}
	c.Set(authorizationPayloadKey, &token.Payload{ID: member.ID.String(), Email: member.Email})
	server.getMember(c)

	require.Equal(t, http.StatusOK, recorder.Code)

}
