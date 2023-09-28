package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/alifanza259/learn-go-library-system/db/mock"
	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/token"
	"github.com/alifanza259/learn-go-library-system/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
)

// Unit Test controller only
// func TestGetMember(t *testing.T) {
// 	member := db.Member{
// 		ID:        uuid.MustParse("051cfcb3-e699-43aa-b9fa-f4d68abaafd0"),
// 		Email:     "alif@gmail.com",
// 		FirstName: "Alif",
// 	}

// 	ctrl := gomock.NewController(t)
// 	library := mockdb.NewMockLibrary(ctrl)

// 	library.EXPECT().
// 		GetMember(gomock.Any(), gomock.Any()).
// 		Times(1).
// 		Return(member, nil)

// 	server := newTestServer(t, library)
// 	recorder := httptest.NewRecorder()

// 	// url := fmt.Sprintf("/member/%d", member.ID)
// 	// request, err := http.NewRequest(http.MethodGet, url, nil)
// 	// require.NoError(t, err)

// 	// server.router.ServeHTTP(recorder)

// 	c := gin.CreateTestContextOnly(recorder, &gin.Engine{})
// 	c.Params = []gin.Param{
// 		{
// 			Key:   "id",
// 			Value: member.ID.String(),
// 		},
// 	}
// 	c.Set(authorizationPayloadKey, &token.Payload{ID: member.ID.String(), Email: member.Email})
// 	server.getMember(c)

// 	require.Equal(t, http.StatusOK, recorder.Code)
// }

func TestGetMember(t *testing.T) {
	member := db.Member{
		ID:        uuid.New(),
		Email:     "alifanza259@gmail.com",
		FirstName: "Alif",
	}

	testCases := []struct {
		name          string
		member        db.Member
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(libraryMock *mockdb.MockLibrary)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			member: member,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, member.Email, member.ID, time.Minute, "member")
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().
					GetMember(gomock.Any(), gomock.Any()).
					Times(1).
					Return(member, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:   "No Member Found",
			member: member,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, member.Email, member.ID, time.Minute, "member")
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().
					GetMember(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Member{}, pgx.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "Internal Error",
			member: member,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, member.Email, member.ID, time.Minute, "member")
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().
					GetMember(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Member{}, &pgconn.PgError{})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "Unauthorized Access",
			member: member,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, db.Member{ID: uuid.New(), Email: "diff@gmail.com"}.Email, db.Member{ID: uuid.New(), Email: "diff@gmail.com"}.ID, time.Minute, "membeer")
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().
					GetMember(gomock.Any(), gomock.Any()).
					Times(1).
					Return(member, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			libraryMock := mockdb.NewMockLibrary(ctrl)

			server := newTestServer(t, libraryMock)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/v1/member/%s", tc.member.ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			tc.buildStubs(libraryMock)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})

	}
}

func TestCreateMember(t *testing.T) {
	member := db.Member{
		ID:        uuid.New(),
		Email:     "dummy@gmail.com",
		FirstName: "Dummy",
		Gender:    "male",
		Password:  "Password",
	}

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(libraryMock *mockdb.MockLibrary)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"email":      member.Email,
				"first_name": member.FirstName,
				"gender":     string(member.Gender),
				"password":   member.Password,
				"dob":        1695293619,
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().CreateMember(gomock.Any(), gomock.Any()).Times(1).Return(member, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Bad Parameter",
			body: gin.H{
				"email":      member.Email,
				"first_name": member.FirstName,
				"gender":     string(member.Gender),
				"password":   "Pass",
				"dob":        1695293619,
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().CreateMember(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Error Create Member",
			body: gin.H{
				"email":      member.Email,
				"first_name": member.FirstName,
				"gender":     string(member.Gender),
				"password":   "Password",
				"dob":        1695293619,
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().CreateMember(gomock.Any(), gomock.Any()).Times(1).Return(db.Member{}, &pgconn.PgError{})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
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

			server := newTestServer(t, libraryMock)
			recorder := httptest.NewRecorder()

			url := "/v1/members"

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}

}

func TestLoginMember(t *testing.T) {
	hashedPassword, err := util.HashPassword("Password")
	require.NoError(t, err)
	member := db.Member{
		Email:    "a@gmail.com",
		Password: hashedPassword,
	}

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(libraryMock *mockdb.MockLibrary)
		checkResponse func(t *testing.T, recorder httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"email":    "a@gmail.com",
				"password": "Password",
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().GetMemberByEmail(gomock.Any(), gomock.Any()).Times(1).Return(member, nil)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Bad Parameters",
			body: gin.H{
				"email":    "a@gmail.com",
				"password": "Passw",
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().GetMemberByEmail(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Invalid Credentials (Email)",
			body: gin.H{
				"email":    "diff@gmail.com",
				"password": "Password",
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().GetMemberByEmail(gomock.Any(), gomock.Any()).Times(1).Return(db.Member{}, pgx.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Invalid Credentials (Password)",
			body: gin.H{
				"email":    "a@gmail.com",
				"password": "password",
			},
			buildStubs: func(libraryMock *mockdb.MockLibrary) {
				libraryMock.EXPECT().GetMemberByEmail(gomock.Any(), gomock.Any()).Times(1).Return(member, nil)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			libraryMock := mockdb.NewMockLibrary(ctrl)
			tc.buildStubs(libraryMock)
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/v1/members/login", bytes.NewBuffer(data))
			require.NoError(t, err)

			server := newTestServer(t, libraryMock)

			recorder := httptest.NewRecorder()

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t, *recorder)
		})
	}

}
