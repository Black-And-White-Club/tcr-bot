// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/modules/user/application/interface.go
//
// Generated by this command:
//
//	mockgen -source=./app/modules/user/application/interface.go -destination=./app/modules/user/application/mocks/mock_service.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	usertypes "github.com/Black-And-White-Club/frolf-bot/app/modules/user/domain/types"
	message "github.com/ThreeDotsLabs/watermill/message"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
	isgomock struct{}
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CheckTagAvailability mocks base method.
func (m *MockService) CheckTagAvailability(ctx context.Context, msg *message.Message, tagNumber int, discordID usertypes.DiscordID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTagAvailability", ctx, msg, tagNumber, discordID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckTagAvailability indicates an expected call of CheckTagAvailability.
func (mr *MockServiceMockRecorder) CheckTagAvailability(ctx, msg, tagNumber, discordID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTagAvailability", reflect.TypeOf((*MockService)(nil).CheckTagAvailability), ctx, msg, tagNumber, discordID)
}

// CheckUserPermissions mocks base method.
func (m *MockService) CheckUserPermissions(ctx context.Context, msg *message.Message, userID usertypes.DiscordID, role usertypes.UserRoleEnum, requesterID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserPermissions", ctx, msg, userID, role, requesterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUserPermissions indicates an expected call of CheckUserPermissions.
func (mr *MockServiceMockRecorder) CheckUserPermissions(ctx, msg, userID, role, requesterID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserPermissions", reflect.TypeOf((*MockService)(nil).CheckUserPermissions), ctx, msg, userID, role, requesterID)
}

// CheckUserPermissionsInDB mocks base method.
func (m *MockService) CheckUserPermissionsInDB(ctx context.Context, msg *message.Message, discordID usertypes.DiscordID, role usertypes.UserRoleEnum, requesterID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserPermissionsInDB", ctx, msg, discordID, role, requesterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUserPermissionsInDB indicates an expected call of CheckUserPermissionsInDB.
func (mr *MockServiceMockRecorder) CheckUserPermissionsInDB(ctx, msg, discordID, role, requesterID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserPermissionsInDB", reflect.TypeOf((*MockService)(nil).CheckUserPermissionsInDB), ctx, msg, discordID, role, requesterID)
}

// CreateUser mocks base method.
func (m *MockService) CreateUser(arg0 context.Context, arg1 *message.Message, arg2 usertypes.DiscordID, arg3 *int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServiceMockRecorder) CreateUser(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockService)(nil).CreateUser), arg0, arg1, arg2, arg3)
}

// GetUser mocks base method.
func (m *MockService) GetUser(ctx context.Context, msg *message.Message, discordID usertypes.DiscordID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, msg, discordID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUser indicates an expected call of GetUser.
func (mr *MockServiceMockRecorder) GetUser(ctx, msg, discordID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockService)(nil).GetUser), ctx, msg, discordID)
}

// GetUserRole mocks base method.
func (m *MockService) GetUserRole(ctx context.Context, msg *message.Message, discordID usertypes.DiscordID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRole", ctx, msg, discordID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUserRole indicates an expected call of GetUserRole.
func (mr *MockServiceMockRecorder) GetUserRole(ctx, msg, discordID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRole", reflect.TypeOf((*MockService)(nil).GetUserRole), ctx, msg, discordID)
}

// PublishUserCreated mocks base method.
func (m *MockService) PublishUserCreated(arg0 context.Context, arg1 *message.Message, arg2 string, arg3 *int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishUserCreated", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishUserCreated indicates an expected call of PublishUserCreated.
func (mr *MockServiceMockRecorder) PublishUserCreated(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishUserCreated", reflect.TypeOf((*MockService)(nil).PublishUserCreated), arg0, arg1, arg2, arg3)
}

// PublishUserCreationFailed mocks base method.
func (m *MockService) PublishUserCreationFailed(arg0 context.Context, arg1 *message.Message, arg2 usertypes.DiscordID, arg3 *int, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishUserCreationFailed", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishUserCreationFailed indicates an expected call of PublishUserCreationFailed.
func (mr *MockServiceMockRecorder) PublishUserCreationFailed(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishUserCreationFailed", reflect.TypeOf((*MockService)(nil).PublishUserCreationFailed), arg0, arg1, arg2, arg3, arg4)
}

// PublishUserPermissionsCheckFailed mocks base method.
func (m *MockService) PublishUserPermissionsCheckFailed(ctx context.Context, msg *message.Message, discordID usertypes.DiscordID, role usertypes.UserRoleEnum, requesterID, reason string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishUserPermissionsCheckFailed", ctx, msg, discordID, role, requesterID, reason)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishUserPermissionsCheckFailed indicates an expected call of PublishUserPermissionsCheckFailed.
func (mr *MockServiceMockRecorder) PublishUserPermissionsCheckFailed(ctx, msg, discordID, role, requesterID, reason any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishUserPermissionsCheckFailed", reflect.TypeOf((*MockService)(nil).PublishUserPermissionsCheckFailed), ctx, msg, discordID, role, requesterID, reason)
}

// PublishUserPermissionsCheckResponse mocks base method.
func (m *MockService) PublishUserPermissionsCheckResponse(ctx context.Context, msg *message.Message, discordID usertypes.DiscordID, role usertypes.UserRoleEnum, requesterID string, hasPermission bool, reason string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishUserPermissionsCheckResponse", ctx, msg, discordID, role, requesterID, hasPermission, reason)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishUserPermissionsCheckResponse indicates an expected call of PublishUserPermissionsCheckResponse.
func (mr *MockServiceMockRecorder) PublishUserPermissionsCheckResponse(ctx, msg, discordID, role, requesterID, hasPermission, reason any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishUserPermissionsCheckResponse", reflect.TypeOf((*MockService)(nil).PublishUserPermissionsCheckResponse), ctx, msg, discordID, role, requesterID, hasPermission, reason)
}

// PublishUserRoleUpdateFailed mocks base method.
func (m *MockService) PublishUserRoleUpdateFailed(arg0 context.Context, arg1 *message.Message, arg2 usertypes.DiscordID, arg3 usertypes.UserRoleEnum, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishUserRoleUpdateFailed", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishUserRoleUpdateFailed indicates an expected call of PublishUserRoleUpdateFailed.
func (mr *MockServiceMockRecorder) PublishUserRoleUpdateFailed(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishUserRoleUpdateFailed", reflect.TypeOf((*MockService)(nil).PublishUserRoleUpdateFailed), arg0, arg1, arg2, arg3, arg4)
}

// PublishUserRoleUpdated mocks base method.
func (m *MockService) PublishUserRoleUpdated(arg0 context.Context, arg1 *message.Message, arg2 usertypes.DiscordID, arg3 usertypes.UserRoleEnum) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishUserRoleUpdated", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishUserRoleUpdated indicates an expected call of PublishUserRoleUpdated.
func (mr *MockServiceMockRecorder) PublishUserRoleUpdated(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishUserRoleUpdated", reflect.TypeOf((*MockService)(nil).PublishUserRoleUpdated), arg0, arg1, arg2, arg3)
}

// TagUnavailable mocks base method.
func (m *MockService) TagUnavailable(ctx context.Context, msg *message.Message, tagNumber int, discordID usertypes.DiscordID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagUnavailable", ctx, msg, tagNumber, discordID)
	ret0, _ := ret[0].(error)
	return ret0
}

// TagUnavailable indicates an expected call of TagUnavailable.
func (mr *MockServiceMockRecorder) TagUnavailable(ctx, msg, tagNumber, discordID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagUnavailable", reflect.TypeOf((*MockService)(nil).TagUnavailable), ctx, msg, tagNumber, discordID)
}

// UpdateUserRole mocks base method.
func (m *MockService) UpdateUserRole(arg0 context.Context, arg1 *message.Message, arg2 usertypes.DiscordID, arg3 usertypes.UserRoleEnum, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserRole", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserRole indicates an expected call of UpdateUserRole.
func (mr *MockServiceMockRecorder) UpdateUserRole(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserRole", reflect.TypeOf((*MockService)(nil).UpdateUserRole), arg0, arg1, arg2, arg3, arg4)
}

// UpdateUserRoleInDatabase mocks base method.
func (m *MockService) UpdateUserRoleInDatabase(arg0 context.Context, arg1 *message.Message, arg2 usertypes.DiscordID, arg3 usertypes.UserRoleEnum) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserRoleInDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserRoleInDatabase indicates an expected call of UpdateUserRoleInDatabase.
func (mr *MockServiceMockRecorder) UpdateUserRoleInDatabase(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserRoleInDatabase", reflect.TypeOf((*MockService)(nil).UpdateUserRoleInDatabase), arg0, arg1, arg2, arg3)
}
