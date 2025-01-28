// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/modules/score/infrastructure/handlers/interface.go
//
// Generated by this command:
//
//	mockgen -source=./app/modules/score/infrastructure/handlers/interface.go -destination=./app/modules/score/infrastructure/handlers/mocks/mock_handlers.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	message "github.com/ThreeDotsLabs/watermill/message"
	gomock "go.uber.org/mock/gomock"
)

// MockHandlers is a mock of Handlers interface.
type MockHandlers struct {
	ctrl     *gomock.Controller
	recorder *MockHandlersMockRecorder
	isgomock struct{}
}

// MockHandlersMockRecorder is the mock recorder for MockHandlers.
type MockHandlersMockRecorder struct {
	mock *MockHandlers
}

// NewMockHandlers creates a new mock instance.
func NewMockHandlers(ctrl *gomock.Controller) *MockHandlers {
	mock := &MockHandlers{ctrl: ctrl}
	mock.recorder = &MockHandlersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandlers) EXPECT() *MockHandlersMockRecorder {
	return m.recorder
}

// HandleProcessRoundScoresRequest mocks base method.
func (m *MockHandlers) HandleProcessRoundScoresRequest(msg *message.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleProcessRoundScoresRequest", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleProcessRoundScoresRequest indicates an expected call of HandleProcessRoundScoresRequest.
func (mr *MockHandlersMockRecorder) HandleProcessRoundScoresRequest(msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleProcessRoundScoresRequest", reflect.TypeOf((*MockHandlers)(nil).HandleProcessRoundScoresRequest), msg)
}

// HandleScoreUpdateRequest mocks base method.
func (m *MockHandlers) HandleScoreUpdateRequest(msg *message.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleScoreUpdateRequest", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleScoreUpdateRequest indicates an expected call of HandleScoreUpdateRequest.
func (mr *MockHandlersMockRecorder) HandleScoreUpdateRequest(msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleScoreUpdateRequest", reflect.TypeOf((*MockHandlers)(nil).HandleScoreUpdateRequest), msg)
}
