// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/discord/service.go

// Package mock_discord is a generated GoMock package.
package mock_discord

import (
	reflect "reflect"

	model "github.com/defipod/mochi/pkg/model"
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

// NotifyGmStreak mocks base method.
func (m *MockService) NotifyGmStreak(channelID, userDiscordID string, streakCount int, podTownXps model.CreateUserTxResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyGmStreak", channelID, userDiscordID, streakCount, podTownXps)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyGmStreak indicates an expected call of NotifyGmStreak.
func (mr *MockServiceMockRecorder) NotifyGmStreak(channelID, userDiscordID, streakCount, podTownXps interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyGmStreak", reflect.TypeOf((*MockService)(nil).NotifyGmStreak), channelID, userDiscordID, streakCount, podTownXps)
}

// NotifyGuildDelete mocks base method.
func (m *MockService) NotifyGuildDelete(guildID, guildName, iconURL string, guildsLeft int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyGuildDelete", guildID, guildName, iconURL, guildsLeft)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyGuildDelete indicates an expected call of NotifyGuildDelete.
func (mr *MockServiceMockRecorder) NotifyGuildDelete(guildID, guildName, iconURL, guildsLeft interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyGuildDelete", reflect.TypeOf((*MockService)(nil).NotifyGuildDelete), guildID, guildName, iconURL, guildsLeft)
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

// NotifyStealAveragePrice mocks base method.
func (m *MockService) NotifyStealAveragePrice(price, floor float64, url, name, image string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyStealAveragePrice", price, floor, url, name, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyStealAveragePrice indicates an expected call of NotifyStealAveragePrice.
func (mr *MockServiceMockRecorder) NotifyStealAveragePrice(price, floor, url, name, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyStealAveragePrice", reflect.TypeOf((*MockService)(nil).NotifyStealAveragePrice), price, floor, url, name, image)
}

// NotifyStealFloorPrice mocks base method.
func (m *MockService) NotifyStealFloorPrice(price, floor float64, url, name, image string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyStealFloorPrice", price, floor, url, name, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyStealFloorPrice indicates an expected call of NotifyStealFloorPrice.
func (mr *MockServiceMockRecorder) NotifyStealFloorPrice(price, floor, url, name, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyStealFloorPrice", reflect.TypeOf((*MockService)(nil).NotifyStealFloorPrice), price, floor, url, name, image)
}

// ReplyUpvoteMessage mocks base method.
func (m *MockService) ReplyUpvoteMessage(msg *response.SetUpvoteMessageCacheResponse, source string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplyUpvoteMessage", msg, source)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplyUpvoteMessage indicates an expected call of ReplyUpvoteMessage.
func (mr *MockServiceMockRecorder) ReplyUpvoteMessage(msg, source interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyUpvoteMessage", reflect.TypeOf((*MockService)(nil).ReplyUpvoteMessage), msg, source)
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

// SendUpdateRolesLog mocks base method.
func (m *MockService) SendUpdateRolesLog(guildID, logChannelID, userID, roleID, _type string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendUpdateRolesLog", guildID, logChannelID, userID, roleID, _type)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendUpdateRolesLog indicates an expected call of SendUpdateRolesLog.
func (mr *MockServiceMockRecorder) SendUpdateRolesLog(guildID, logChannelID, userID, roleID, _type interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendUpdateRolesLog", reflect.TypeOf((*MockService)(nil).SendUpdateRolesLog), guildID, logChannelID, userID, roleID, _type)
}

// SendUpvoteMessage mocks base method.
func (m *MockService) SendUpvoteMessage(discordID, source string, isStranger bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendUpvoteMessage", discordID, source, isStranger)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendUpvoteMessage indicates an expected call of SendUpvoteMessage.
func (mr *MockServiceMockRecorder) SendUpvoteMessage(discordID, source, isStranger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendUpvoteMessage", reflect.TypeOf((*MockService)(nil).SendUpvoteMessage), discordID, source, isStranger)
}
