// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/modules/user/domain/types/types.go
//
// Generated by this command:
//
//	mockgen -source=./app/modules/user/domain/types/types.go -destination=./app/modules/user/domain/types/mocks/mock_types.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	usertypes "github.com/Black-And-White-Club/tcr-bot/app/modules/user/domain/types"
	gomock "go.uber.org/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
	isgomock struct{}
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// GetDiscordID mocks base method.
func (m *MockUser) GetDiscordID() usertypes.DiscordID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiscordID")
	ret0, _ := ret[0].(usertypes.DiscordID)
	return ret0
}

// GetDiscordID indicates an expected call of GetDiscordID.
func (mr *MockUserMockRecorder) GetDiscordID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiscordID", reflect.TypeOf((*MockUser)(nil).GetDiscordID))
}

// GetID mocks base method.
func (m *MockUser) GetID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockUserMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockUser)(nil).GetID))
}

// GetName mocks base method.
func (m *MockUser) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockUserMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockUser)(nil).GetName))
}

// GetRole mocks base method.
func (m *MockUser) GetRole() usertypes.UserRoleEnum {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole")
	ret0, _ := ret[0].(usertypes.UserRoleEnum)
	return ret0
}

// GetRole indicates an expected call of GetRole.
func (mr *MockUserMockRecorder) GetRole() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockUser)(nil).GetRole))
}
