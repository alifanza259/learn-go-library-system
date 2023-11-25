// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/alifanza259/learn-go-library-system/db/sqlc (interfaces: Library)

// Package mock_sqlc is a generated GoMock package.
package mock_sqlc

import (
	context "context"
	reflect "reflect"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockLibrary is a mock of Library interface.
type MockLibrary struct {
	ctrl     *gomock.Controller
	recorder *MockLibraryMockRecorder
}

// MockLibraryMockRecorder is the mock recorder for MockLibrary.
type MockLibraryMockRecorder struct {
	mock *MockLibrary
}

// NewMockLibrary creates a new mock instance.
func NewMockLibrary(ctrl *gomock.Controller) *MockLibrary {
	mock := &MockLibrary{ctrl: ctrl}
	mock.recorder = &MockLibraryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLibrary) EXPECT() *MockLibraryMockRecorder {
	return m.recorder
}

// BorrowTx mocks base method.
func (m *MockLibrary) BorrowTx(arg0 context.Context, arg1 db.BorrowTxParams) (db.BorrowTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BorrowTx", arg0, arg1)
	ret0, _ := ret[0].(db.BorrowTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BorrowTx indicates an expected call of BorrowTx.
func (mr *MockLibraryMockRecorder) BorrowTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BorrowTx", reflect.TypeOf((*MockLibrary)(nil).BorrowTx), arg0, arg1)
}

// CreateBook mocks base method.
func (m *MockLibrary) CreateBook(arg0 context.Context, arg1 db.CreateBookParams) (db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", arg0, arg1)
	ret0, _ := ret[0].(db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockLibraryMockRecorder) CreateBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockLibrary)(nil).CreateBook), arg0, arg1)
}

// CreateBorrow mocks base method.
func (m *MockLibrary) CreateBorrow(arg0 context.Context, arg1 db.CreateBorrowParams) (db.BorrowDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBorrow", arg0, arg1)
	ret0, _ := ret[0].(db.BorrowDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBorrow indicates an expected call of CreateBorrow.
func (mr *MockLibraryMockRecorder) CreateBorrow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBorrow", reflect.TypeOf((*MockLibrary)(nil).CreateBorrow), arg0, arg1)
}

// CreateEmailVerification mocks base method.
func (m *MockLibrary) CreateEmailVerification(arg0 context.Context, arg1 db.CreateEmailVerificationParams) (db.EmailVerification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmailVerification", arg0, arg1)
	ret0, _ := ret[0].(db.EmailVerification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmailVerification indicates an expected call of CreateEmailVerification.
func (mr *MockLibraryMockRecorder) CreateEmailVerification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmailVerification", reflect.TypeOf((*MockLibrary)(nil).CreateEmailVerification), arg0, arg1)
}

// CreateMember mocks base method.
func (m *MockLibrary) CreateMember(arg0 context.Context, arg1 db.CreateMemberParams) (db.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMember", arg0, arg1)
	ret0, _ := ret[0].(db.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMember indicates an expected call of CreateMember.
func (mr *MockLibraryMockRecorder) CreateMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMember", reflect.TypeOf((*MockLibrary)(nil).CreateMember), arg0, arg1)
}

// CreateMemberTx mocks base method.
func (m *MockLibrary) CreateMemberTx(arg0 context.Context, arg1 db.CreateMemberTxParams) (db.CreateMemberTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMemberTx", arg0, arg1)
	ret0, _ := ret[0].(db.CreateMemberTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMemberTx indicates an expected call of CreateMemberTx.
func (mr *MockLibraryMockRecorder) CreateMemberTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMemberTx", reflect.TypeOf((*MockLibrary)(nil).CreateMemberTx), arg0, arg1)
}

// CreateTransaction mocks base method.
func (m *MockLibrary) CreateTransaction(arg0 context.Context, arg1 db.CreateTransactionParams) (db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", arg0, arg1)
	ret0, _ := ret[0].(db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockLibraryMockRecorder) CreateTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockLibrary)(nil).CreateTransaction), arg0, arg1)
}

// DeleteBook mocks base method.
func (m *MockLibrary) DeleteBook(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockLibraryMockRecorder) DeleteBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockLibrary)(nil).DeleteBook), arg0, arg1)
}

// GetAdmin mocks base method.
func (m *MockLibrary) GetAdmin(arg0 context.Context, arg1 uuid.UUID) (db.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmin", arg0, arg1)
	ret0, _ := ret[0].(db.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdmin indicates an expected call of GetAdmin.
func (mr *MockLibraryMockRecorder) GetAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmin", reflect.TypeOf((*MockLibrary)(nil).GetAdmin), arg0, arg1)
}

// GetAdminByEmail mocks base method.
func (m *MockLibrary) GetAdminByEmail(arg0 context.Context, arg1 string) (db.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdminByEmail", arg0, arg1)
	ret0, _ := ret[0].(db.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdminByEmail indicates an expected call of GetAdminByEmail.
func (mr *MockLibraryMockRecorder) GetAdminByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdminByEmail", reflect.TypeOf((*MockLibrary)(nil).GetAdminByEmail), arg0, arg1)
}

// GetBook mocks base method.
func (m *MockLibrary) GetBook(arg0 context.Context, arg1 int32) (db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBook", arg0, arg1)
	ret0, _ := ret[0].(db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBook indicates an expected call of GetBook.
func (mr *MockLibraryMockRecorder) GetBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockLibrary)(nil).GetBook), arg0, arg1)
}

// GetBorrow mocks base method.
func (m *MockLibrary) GetBorrow(arg0 context.Context, arg1 uuid.UUID) (db.BorrowDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBorrow", arg0, arg1)
	ret0, _ := ret[0].(db.BorrowDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBorrow indicates an expected call of GetBorrow.
func (mr *MockLibraryMockRecorder) GetBorrow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBorrow", reflect.TypeOf((*MockLibrary)(nil).GetBorrow), arg0, arg1)
}

// GetEmailVerification mocks base method.
func (m *MockLibrary) GetEmailVerification(arg0 context.Context, arg1 string) (db.EmailVerification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmailVerification", arg0, arg1)
	ret0, _ := ret[0].(db.EmailVerification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmailVerification indicates an expected call of GetEmailVerification.
func (mr *MockLibraryMockRecorder) GetEmailVerification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailVerification", reflect.TypeOf((*MockLibrary)(nil).GetEmailVerification), arg0, arg1)
}

// GetMember mocks base method.
func (m *MockLibrary) GetMember(arg0 context.Context, arg1 uuid.UUID) (db.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMember", arg0, arg1)
	ret0, _ := ret[0].(db.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMember indicates an expected call of GetMember.
func (mr *MockLibraryMockRecorder) GetMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMember", reflect.TypeOf((*MockLibrary)(nil).GetMember), arg0, arg1)
}

// GetMemberByEmail mocks base method.
func (m *MockLibrary) GetMemberByEmail(arg0 context.Context, arg1 string) (db.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberByEmail", arg0, arg1)
	ret0, _ := ret[0].(db.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberByEmail indicates an expected call of GetMemberByEmail.
func (mr *MockLibraryMockRecorder) GetMemberByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberByEmail", reflect.TypeOf((*MockLibrary)(nil).GetMemberByEmail), arg0, arg1)
}

// GetTransaction mocks base method.
func (m *MockLibrary) GetTransaction(arg0 context.Context, arg1 uuid.UUID) (db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", arg0, arg1)
	ret0, _ := ret[0].(db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockLibraryMockRecorder) GetTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockLibrary)(nil).GetTransaction), arg0, arg1)
}

// GetTransactionAndBorrowDetail mocks base method.
func (m *MockLibrary) GetTransactionAndBorrowDetail(arg0 context.Context, arg1 uuid.UUID) (db.GetTransactionAndBorrowDetailRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionAndBorrowDetail", arg0, arg1)
	ret0, _ := ret[0].(db.GetTransactionAndBorrowDetailRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionAndBorrowDetail indicates an expected call of GetTransactionAndBorrowDetail.
func (mr *MockLibraryMockRecorder) GetTransactionAndBorrowDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionAndBorrowDetail", reflect.TypeOf((*MockLibrary)(nil).GetTransactionAndBorrowDetail), arg0, arg1)
}

// GetTransactionAssociatedDetail mocks base method.
func (m *MockLibrary) GetTransactionAssociatedDetail(arg0 context.Context, arg1 uuid.UUID) (db.GetTransactionAssociatedDetailRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionAssociatedDetail", arg0, arg1)
	ret0, _ := ret[0].(db.GetTransactionAssociatedDetailRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionAssociatedDetail indicates an expected call of GetTransactionAssociatedDetail.
func (mr *MockLibraryMockRecorder) GetTransactionAssociatedDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionAssociatedDetail", reflect.TypeOf((*MockLibrary)(nil).GetTransactionAssociatedDetail), arg0, arg1)
}

// ListAdmin mocks base method.
func (m *MockLibrary) ListAdmin(arg0 context.Context) ([]db.ListAdminRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAdmin", arg0)
	ret0, _ := ret[0].([]db.ListAdminRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAdmin indicates an expected call of ListAdmin.
func (mr *MockLibraryMockRecorder) ListAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAdmin", reflect.TypeOf((*MockLibrary)(nil).ListAdmin), arg0)
}

// ListBooks mocks base method.
func (m *MockLibrary) ListBooks(arg0 context.Context) ([]db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBooks", arg0)
	ret0, _ := ret[0].([]db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBooks indicates an expected call of ListBooks.
func (mr *MockLibraryMockRecorder) ListBooks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBooks", reflect.TypeOf((*MockLibrary)(nil).ListBooks), arg0)
}

// ListMembers mocks base method.
func (m *MockLibrary) ListMembers(arg0 context.Context) ([]db.ListMembersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMembers", arg0)
	ret0, _ := ret[0].([]db.ListMembersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMembers indicates an expected call of ListMembers.
func (mr *MockLibraryMockRecorder) ListMembers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembers", reflect.TypeOf((*MockLibrary)(nil).ListMembers), arg0)
}

// ProcessBorrowTx mocks base method.
func (m *MockLibrary) ProcessBorrowTx(arg0 context.Context, arg1 db.ProcessBorrowTxParams) (db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessBorrowTx", arg0, arg1)
	ret0, _ := ret[0].(db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessBorrowTx indicates an expected call of ProcessBorrowTx.
func (mr *MockLibraryMockRecorder) ProcessBorrowTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBorrowTx", reflect.TypeOf((*MockLibrary)(nil).ProcessBorrowTx), arg0, arg1)
}

// UpdateBook mocks base method.
func (m *MockLibrary) UpdateBook(arg0 context.Context, arg1 db.UpdateBookParams) (db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", arg0, arg1)
	ret0, _ := ret[0].(db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockLibraryMockRecorder) UpdateBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockLibrary)(nil).UpdateBook), arg0, arg1)
}

// UpdateEmailVerification mocks base method.
func (m *MockLibrary) UpdateEmailVerification(arg0 context.Context, arg1 db.UpdateEmailVerificationParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmailVerification", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmailVerification indicates an expected call of UpdateEmailVerification.
func (mr *MockLibraryMockRecorder) UpdateEmailVerification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmailVerification", reflect.TypeOf((*MockLibrary)(nil).UpdateEmailVerification), arg0, arg1)
}

// UpdateMember mocks base method.
func (m *MockLibrary) UpdateMember(arg0 context.Context, arg1 db.UpdateMemberParams) (db.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMember", arg0, arg1)
	ret0, _ := ret[0].(db.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMember indicates an expected call of UpdateMember.
func (mr *MockLibraryMockRecorder) UpdateMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMember", reflect.TypeOf((*MockLibrary)(nil).UpdateMember), arg0, arg1)
}

// UpdateTransaction mocks base method.
func (m *MockLibrary) UpdateTransaction(arg0 context.Context, arg1 db.UpdateTransactionParams) (db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransaction", arg0, arg1)
	ret0, _ := ret[0].(db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransaction indicates an expected call of UpdateTransaction.
func (mr *MockLibraryMockRecorder) UpdateTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransaction", reflect.TypeOf((*MockLibrary)(nil).UpdateTransaction), arg0, arg1)
}

// VerifyEmailTx mocks base method.
func (m *MockLibrary) VerifyEmailTx(arg0 context.Context, arg1 db.VerifyEmailTxParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyEmailTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyEmailTx indicates an expected call of VerifyEmailTx.
func (mr *MockLibraryMockRecorder) VerifyEmailTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyEmailTx", reflect.TypeOf((*MockLibrary)(nil).VerifyEmailTx), arg0, arg1)
}
