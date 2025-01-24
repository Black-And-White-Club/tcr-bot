// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/shared/eventbus.go
//
// Generated by this command:
//
//	mockgen -source=./app/shared/eventbus.go -destination=./app/eventbus/mocks/mock_eventbus.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	message "github.com/ThreeDotsLabs/watermill/message"
	gomock "go.uber.org/mock/gomock"
)

// MockEventBus is a mock of EventBus interface.
type MockEventBus struct {
	ctrl     *gomock.Controller
	recorder *MockEventBusMockRecorder
	isgomock struct{}
}

// MockEventBusMockRecorder is the mock recorder for MockEventBus.
type MockEventBusMockRecorder struct {
	mock *MockEventBus
}

// NewMockEventBus creates a new mock instance.
func NewMockEventBus(ctrl *gomock.Controller) *MockEventBus {
	mock := &MockEventBus{ctrl: ctrl}
	mock.recorder = &MockEventBusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventBus) EXPECT() *MockEventBusMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockEventBus) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockEventBusMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEventBus)(nil).Close))
}

// Publish mocks base method.
func (m *MockEventBus) Publish(topic string, messages ...*message.Message) error {
	m.ctrl.T.Helper()
	varargs := []any{topic}
	for _, a := range messages {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Publish", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockEventBusMockRecorder) Publish(topic any, messages ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{topic}, messages...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockEventBus)(nil).Publish), varargs...)
}

// Subscribe mocks base method.
func (m *MockEventBus) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, topic)
	ret0, _ := ret[0].(<-chan *message.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockEventBusMockRecorder) Subscribe(ctx, topic any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockEventBus)(nil).Subscribe), ctx, topic)
}
