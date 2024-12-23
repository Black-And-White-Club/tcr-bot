// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -destination=mocks/message_handler_mock.go -package=user_mocks -source=interface.go MessageHandler
//

// Package user_mocks is a generated GoMock package.
package user_mocks

import (
	reflect "reflect"

	message "github.com/ThreeDotsLabs/watermill/message"
	gomock "go.uber.org/mock/gomock"
)

// MockMessageHandler is a mock of MessageHandler interface.
type MockMessageHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMessageHandlerMockRecorder
	isgomock struct{}
}

// MockMessageHandlerMockRecorder is the mock recorder for MockMessageHandler.
type MockMessageHandlerMockRecorder struct {
	mock *MockMessageHandler
}

// NewMockMessageHandler creates a new mock instance.
func NewMockMessageHandler(ctrl *gomock.Controller) *MockMessageHandler {
	mock := &MockMessageHandler{ctrl: ctrl}
	mock.recorder = &MockMessageHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageHandler) EXPECT() *MockMessageHandlerMockRecorder {
	return m.recorder
}

// HandleMessage mocks base method.
func (m *MockMessageHandler) HandleMessage(msg *message.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleMessage indicates an expected call of HandleMessage.
func (mr *MockMessageHandlerMockRecorder) HandleMessage(msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockMessageHandler)(nil).HandleMessage), msg)
}
