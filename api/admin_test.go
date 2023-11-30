package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/alifanza259/learn-go-library-system/db/mock"
	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/token"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
)

func TestListMembers(t *testing.T) {
	members := []db.ListMembersRow{
		{
			ID: uuid.New(),
		},
		{
			ID: uuid.New(),
		},
	}
	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(libraryMock *mockdb.MockLibrary)
		checkResponse func(t *testing.T, recorder httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, db.Member{ID: uuid.New(), Email: "diff@gmail.com"}.Email, db.Member{ID: uuid.New(), Email: "diff@gmail.com"}.ID, time.Minute, "member")
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().GetAdmin(gomock.Any(), gomock.Any()).Times(1).Return(db.Admin{Permission: "super"}, nil)
				libraryMock.EXPECT().ListMembers(gomock.Any()).Times(1).Return(members, nil)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Internal Error",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, db.Member{ID: uuid.New(), Email: "diff@gmail.com"}.Email, db.Member{ID: uuid.New(), Email: "diff@gmail.com"}.ID, time.Minute, "member")
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().GetAdmin(gomock.Any(), gomock.Any()).Times(1).Return(db.Admin{Permission: "super"}, nil)
				libraryMock.EXPECT().ListMembers(gomock.Any()).Times(1).Return([]db.ListMembersRow{}, &pgconn.PgError{Code: "123"})
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			libraryMock := mockdb.NewMockLibrary(ctrl)
			tc.buildStubs(libraryMock)
			request, err := http.NewRequest(http.MethodGet, "/admin/members", nil)
			require.NoError(t, err)

			recorder := httptest.NewRecorder()

			server := newTestServer(t, libraryMock)
			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t, *recorder)
		})
	}
}
