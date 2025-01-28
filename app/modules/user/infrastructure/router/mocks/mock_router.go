// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/modules/user/infrastructure/router/interface.go
//
// Generated by this command:
//
//	mockgen -source=./app/modules/user/infrastructure/router/interface.go -destination=./app/modules/user/infrastructure/router/mocks/mock_router.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	userservice "github.com/Black-And-White-Club/frolf-bot/app/modules/user/application"
	message "github.com/ThreeDotsLabs/watermill/message"
	gomock "go.uber.org/mock/gomock"
)

// MockRouter is a mock of Router interface.
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
	isgomock struct{}
}

// MockRouterMockRecorder is the mock recorder for MockRouter.
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance.
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRouter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRouterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRouter)(nil).Close))
}

// Configure mocks base method.
func (m *MockRouter) Configure(router *message.Router, userService userservice.Service, subscriber message.Subscriber) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", router, userService, subscriber)
	ret0, _ := ret[0].(error)
	return ret0
}

// Configure indicates an expected call of Configure.
func (mr *MockRouterMockRecorder) Configure(router, userService, subscriber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockRouter)(nil).Configure), router, userService, subscriber)
}

// Run mocks base method.
func (m *MockRouter) Run(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockRouterMockRecorder) Run(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRouter)(nil).Run), ctx)
}
