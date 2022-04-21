// Code generated by MockGen. DO NOT EDIT.
// Source: notary-public-online/internal/repository (interfaces: DB)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	model "notary-public-online/internal/entity/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// CreateDocument mocks base method.
func (m *MockDB) CreateDocument(arg0 context.Context, arg1, arg2, arg3 string, arg4 []byte, arg5 int, arg6 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDocument", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDocument indicates an expected call of CreateDocument.
func (mr *MockDBMockRecorder) CreateDocument(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDocument", reflect.TypeOf((*MockDB)(nil).CreateDocument), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// CreateNoatry mocks base method.
func (m *MockDB) CreateNoatry(arg0 context.Context, arg1, arg2, arg3 int, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNoatry", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNoatry indicates an expected call of CreateNoatry.
func (mr *MockDBMockRecorder) CreateNoatry(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNoatry", reflect.TypeOf((*MockDB)(nil).CreateNoatry), arg0, arg1, arg2, arg3, arg4)
}

// CreateSignature mocks base method.
func (m *MockDB) CreateSignature(arg0 context.Context, arg1, arg2 int, arg3 *[]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSignature", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSignature indicates an expected call of CreateSignature.
func (mr *MockDBMockRecorder) CreateSignature(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSignature", reflect.TypeOf((*MockDB)(nil).CreateSignature), arg0, arg1, arg2, arg3)
}

// GetDocument mocks base method.
func (m *MockDB) GetDocument(arg0 context.Context, arg1 int) (*model.Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocument", arg0, arg1)
	ret0, _ := ret[0].(*model.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDocument indicates an expected call of GetDocument.
func (mr *MockDBMockRecorder) GetDocument(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocument", reflect.TypeOf((*MockDB)(nil).GetDocument), arg0, arg1)
}

// GetDocumentAddress mocks base method.
func (m *MockDB) GetDocumentAddress(arg0 context.Context, arg1 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocumentAddress", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDocumentAddress indicates an expected call of GetDocumentAddress.
func (mr *MockDBMockRecorder) GetDocumentAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocumentAddress", reflect.TypeOf((*MockDB)(nil).GetDocumentAddress), arg0, arg1)
}

// GetDocumentHash mocks base method.
func (m *MockDB) GetDocumentHash(arg0 context.Context, arg1 int) (*[]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocumentHash", arg0, arg1)
	ret0, _ := ret[0].(*[]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDocumentHash indicates an expected call of GetDocumentHash.
func (mr *MockDBMockRecorder) GetDocumentHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocumentHash", reflect.TypeOf((*MockDB)(nil).GetDocumentHash), arg0, arg1)
}

// GetNoatry mocks base method.
func (m *MockDB) GetNoatry(arg0 context.Context, arg1 int) (model.Notary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNoatry", arg0, arg1)
	ret0, _ := ret[0].(model.Notary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNoatry indicates an expected call of GetNoatry.
func (mr *MockDBMockRecorder) GetNoatry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNoatry", reflect.TypeOf((*MockDB)(nil).GetNoatry), arg0, arg1)
}

// GetSignatures mocks base method.
func (m *MockDB) GetSignatures(arg0 context.Context, arg1, arg2 int) (*[]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSignatures", arg0, arg1, arg2)
	ret0, _ := ret[0].(*[]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSignatures indicates an expected call of GetSignatures.
func (mr *MockDBMockRecorder) GetSignatures(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSignatures", reflect.TypeOf((*MockDB)(nil).GetSignatures), arg0, arg1, arg2)
}

// GetUserKeys mocks base method.
func (m *MockDB) GetUserKeys(arg0 context.Context, arg1 int) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserKeys", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserKeys indicates an expected call of GetUserKeys.
func (mr *MockDBMockRecorder) GetUserKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserKeys", reflect.TypeOf((*MockDB)(nil).GetUserKeys), arg0, arg1)
}
