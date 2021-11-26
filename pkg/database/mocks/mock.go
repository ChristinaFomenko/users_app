// Code generated by MockGen. DO NOT EDIT.
// Source: db.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	reflect "reflect"

	model "github.com/ChristinaFomenko/users_app/pkg/model"
	gomock "github.com/golang/mock/gomock"
)

// MockUserDB is a mock of UserDB interface.
type MockUserDB struct {
	ctrl     *gomock.Controller
	recorder *MockUserDBMockRecorder
}

// MockUserDBMockRecorder is the mock recorder for MockUserDB.
type MockUserDBMockRecorder struct {
	mock *MockUserDB
}

// NewMockUserDB creates a new mock instance.
func NewMockUserDB(ctrl *gomock.Controller) *MockUserDB {
	mock := &MockUserDB{ctrl: ctrl}
	mock.recorder = &MockUserDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDB) EXPECT() *MockUserDBMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserDB) CreateUser(u *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", u)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserDBMockRecorder) CreateUser(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserDB)(nil).CreateUser), u)
}

// DeleteAllUsers mocks base method.
func (m *MockUserDB) DeleteAllUsers() ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllUsers")
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllUsers indicates an expected call of DeleteAllUsers.
func (mr *MockUserDBMockRecorder) DeleteAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllUsers", reflect.TypeOf((*MockUserDB)(nil).DeleteAllUsers))
}

// GetAllUsers mocks base method.
func (m *MockUserDB) GetAllUsers() ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserDBMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserDB)(nil).GetAllUsers))
}

// GetUser mocks base method.
func (m *MockUserDB) GetUser() ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser")
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserDBMockRecorder) GetUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserDB)(nil).GetUser))
}
