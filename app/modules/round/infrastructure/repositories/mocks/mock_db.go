// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/modules/round/infrastructure/repositories/interface.go
//
// Generated by this command:
//
//	mockgen -source=./app/modules/round/infrastructure/repositories/interface.go -destination=./app/modules/round/infrastructure/repositories/mocks/mock_db.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	roundtypes "github.com/Black-And-White-Club/frolf-bot/app/modules/round/domain/types"
	gomock "go.uber.org/mock/gomock"
)

// MockRoundDB is a mock of RoundDB interface.
type MockRoundDB struct {
	ctrl     *gomock.Controller
	recorder *MockRoundDBMockRecorder
	isgomock struct{}
}

// MockRoundDBMockRecorder is the mock recorder for MockRoundDB.
type MockRoundDBMockRecorder struct {
	mock *MockRoundDB
}

// NewMockRoundDB creates a new mock instance.
func NewMockRoundDB(ctrl *gomock.Controller) *MockRoundDB {
	mock := &MockRoundDB{ctrl: ctrl}
	mock.recorder = &MockRoundDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoundDB) EXPECT() *MockRoundDBMockRecorder {
	return m.recorder
}

// CreateRound mocks base method.
func (m *MockRoundDB) CreateRound(ctx context.Context, round *roundtypes.Round) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRound", ctx, round)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRound indicates an expected call of CreateRound.
func (mr *MockRoundDBMockRecorder) CreateRound(ctx, round any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRound", reflect.TypeOf((*MockRoundDB)(nil).CreateRound), ctx, round)
}

// DeleteRound mocks base method.
func (m *MockRoundDB) DeleteRound(ctx context.Context, roundID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRound", ctx, roundID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRound indicates an expected call of DeleteRound.
func (mr *MockRoundDBMockRecorder) DeleteRound(ctx, roundID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRound", reflect.TypeOf((*MockRoundDB)(nil).DeleteRound), ctx, roundID)
}

// GetParticipants mocks base method.
func (m *MockRoundDB) GetParticipants(ctx context.Context, roundID string) ([]roundtypes.RoundParticipant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParticipants", ctx, roundID)
	ret0, _ := ret[0].([]roundtypes.RoundParticipant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParticipants indicates an expected call of GetParticipants.
func (mr *MockRoundDBMockRecorder) GetParticipants(ctx, roundID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParticipants", reflect.TypeOf((*MockRoundDB)(nil).GetParticipants), ctx, roundID)
}

// GetParticipantsWithResponses mocks base method.
func (m *MockRoundDB) GetParticipantsWithResponses(ctx context.Context, roundID string, responses ...roundtypes.Response) ([]roundtypes.RoundParticipant, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, roundID}
	for _, a := range responses {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetParticipantsWithResponses", varargs...)
	ret0, _ := ret[0].([]roundtypes.RoundParticipant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParticipantsWithResponses indicates an expected call of GetParticipantsWithResponses.
func (mr *MockRoundDBMockRecorder) GetParticipantsWithResponses(ctx, roundID any, responses ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, roundID}, responses...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParticipantsWithResponses", reflect.TypeOf((*MockRoundDB)(nil).GetParticipantsWithResponses), varargs...)
}

// GetRound mocks base method.
func (m *MockRoundDB) GetRound(ctx context.Context, roundID string) (*roundtypes.Round, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRound", ctx, roundID)
	ret0, _ := ret[0].(*roundtypes.Round)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRound indicates an expected call of GetRound.
func (mr *MockRoundDBMockRecorder) GetRound(ctx, roundID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRound", reflect.TypeOf((*MockRoundDB)(nil).GetRound), ctx, roundID)
}

// GetRoundState mocks base method.
func (m *MockRoundDB) GetRoundState(ctx context.Context, roundID string) (roundtypes.RoundState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoundState", ctx, roundID)
	ret0, _ := ret[0].(roundtypes.RoundState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoundState indicates an expected call of GetRoundState.
func (mr *MockRoundDBMockRecorder) GetRoundState(ctx, roundID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoundState", reflect.TypeOf((*MockRoundDB)(nil).GetRoundState), ctx, roundID)
}

// GetUpcomingRounds mocks base method.
func (m *MockRoundDB) GetUpcomingRounds(ctx context.Context, now, oneHourFromNow time.Time) ([]*roundtypes.Round, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpcomingRounds", ctx, now, oneHourFromNow)
	ret0, _ := ret[0].([]*roundtypes.Round)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpcomingRounds indicates an expected call of GetUpcomingRounds.
func (mr *MockRoundDBMockRecorder) GetUpcomingRounds(ctx, now, oneHourFromNow any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpcomingRounds", reflect.TypeOf((*MockRoundDB)(nil).GetUpcomingRounds), ctx, now, oneHourFromNow)
}

// LogRound mocks base method.
func (m *MockRoundDB) LogRound(ctx context.Context, round *roundtypes.Round) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogRound", ctx, round)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogRound indicates an expected call of LogRound.
func (mr *MockRoundDBMockRecorder) LogRound(ctx, round any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogRound", reflect.TypeOf((*MockRoundDB)(nil).LogRound), ctx, round)
}

// UpdateDiscordEventID mocks base method.
func (m *MockRoundDB) UpdateDiscordEventID(ctx context.Context, roundID, discordEventID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDiscordEventID", ctx, roundID, discordEventID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDiscordEventID indicates an expected call of UpdateDiscordEventID.
func (mr *MockRoundDBMockRecorder) UpdateDiscordEventID(ctx, roundID, discordEventID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDiscordEventID", reflect.TypeOf((*MockRoundDB)(nil).UpdateDiscordEventID), ctx, roundID, discordEventID)
}

// UpdateParticipant mocks base method.
func (m *MockRoundDB) UpdateParticipant(ctx context.Context, roundID string, participant roundtypes.RoundParticipant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateParticipant", ctx, roundID, participant)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateParticipant indicates an expected call of UpdateParticipant.
func (mr *MockRoundDBMockRecorder) UpdateParticipant(ctx, roundID, participant any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateParticipant", reflect.TypeOf((*MockRoundDB)(nil).UpdateParticipant), ctx, roundID, participant)
}

// UpdateParticipantScore mocks base method.
func (m *MockRoundDB) UpdateParticipantScore(ctx context.Context, roundID, participantID string, score int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateParticipantScore", ctx, roundID, participantID, score)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateParticipantScore indicates an expected call of UpdateParticipantScore.
func (mr *MockRoundDBMockRecorder) UpdateParticipantScore(ctx, roundID, participantID, score any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateParticipantScore", reflect.TypeOf((*MockRoundDB)(nil).UpdateParticipantScore), ctx, roundID, participantID, score)
}

// UpdateRound mocks base method.
func (m *MockRoundDB) UpdateRound(ctx context.Context, roundID string, round *roundtypes.Round) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRound", ctx, roundID, round)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRound indicates an expected call of UpdateRound.
func (mr *MockRoundDBMockRecorder) UpdateRound(ctx, roundID, round any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRound", reflect.TypeOf((*MockRoundDB)(nil).UpdateRound), ctx, roundID, round)
}

// UpdateRoundState mocks base method.
func (m *MockRoundDB) UpdateRoundState(ctx context.Context, roundID string, state roundtypes.RoundState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoundState", ctx, roundID, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoundState indicates an expected call of UpdateRoundState.
func (mr *MockRoundDBMockRecorder) UpdateRoundState(ctx, roundID, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoundState", reflect.TypeOf((*MockRoundDB)(nil).UpdateRoundState), ctx, roundID, state)
}
