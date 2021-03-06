// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/discord/service.go

// Package mock_discord is a generated GoMock package.
package mock_discord

import (
	reflect "reflect"

	response "github.com/defipod/mochi/pkg/response"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
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

// NotifyAddNewCollection mocks base method.
func (m *MockService) NotifyAddNewCollection(guildID, collectionName, symbol, chain, image string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyAddNewCollection", guildID, collectionName, symbol, chain, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyAddNewCollection indicates an expected call of NotifyAddNewCollection.
func (mr *MockServiceMockRecorder) NotifyAddNewCollection(guildID, collectionName, symbol, chain, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyAddNewCollection", reflect.TypeOf((*MockService)(nil).NotifyAddNewCollection), guildID, collectionName, symbol, chain, image)
}

// NotifyNewGuild mocks base method.
func (m *MockService) NotifyNewGuild(newGuildID string, count int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyNewGuild", newGuildID, count)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyNewGuild indicates an expected call of NotifyNewGuild.
func (mr *MockServiceMockRecorder) NotifyNewGuild(newGuildID, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewGuild", reflect.TypeOf((*MockService)(nil).NotifyNewGuild), newGuildID, count)
}

// SendGuildActivityLogs mocks base method.
func (m *MockService) SendGuildActivityLogs(channelID, userID, title, description string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendGuildActivityLogs", channelID, userID, title, description)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendGuildActivityLogs indicates an expected call of SendGuildActivityLogs.
func (mr *MockServiceMockRecorder) SendGuildActivityLogs(channelID, userID, title, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGuildActivityLogs", reflect.TypeOf((*MockService)(nil).SendGuildActivityLogs), channelID, userID, title, description)
}

// SendLevelUpMessage mocks base method.
func (m *MockService) SendLevelUpMessage(logChannelID, role string, uActivity *response.HandleUserActivityResponse) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendLevelUpMessage", logChannelID, role, uActivity)
}

// SendLevelUpMessage indicates an expected call of SendLevelUpMessage.
func (mr *MockServiceMockRecorder) SendLevelUpMessage(logChannelID, role, uActivity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendLevelUpMessage", reflect.TypeOf((*MockService)(nil).SendLevelUpMessage), logChannelID, role, uActivity)
}
