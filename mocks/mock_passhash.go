// Code generated by MockGen. DO NOT EDIT.
// Source: notary-public-online/internal/pkg/passwordHash (interfaces: PasswordHash)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPasswordHash is a mock of PasswordHash interface.
type MockPasswordHash struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordHashMockRecorder
}

// MockPasswordHashMockRecorder is the mock recorder for MockPasswordHash.
type MockPasswordHashMockRecorder struct {
	mock *MockPasswordHash
}

// NewMockPasswordHash creates a new mock instance.
func NewMockPasswordHash(ctrl *gomock.Controller) *MockPasswordHash {
	mock := &MockPasswordHash{ctrl: ctrl}
	mock.recorder = &MockPasswordHashMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordHash) EXPECT() *MockPasswordHashMockRecorder {
	return m.recorder
}

// CheckPasswordHash mocks base method.
func (m *MockPasswordHash) CheckPasswordHash(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPasswordHash", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckPasswordHash indicates an expected call of CheckPasswordHash.
func (mr *MockPasswordHashMockRecorder) CheckPasswordHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPasswordHash", reflect.TypeOf((*MockPasswordHash)(nil).CheckPasswordHash), arg0, arg1)
}

// HashPassword mocks base method.
func (m *MockPasswordHash) HashPassword(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockPasswordHashMockRecorder) HashPassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockPasswordHash)(nil).HashPassword), arg0)
}
