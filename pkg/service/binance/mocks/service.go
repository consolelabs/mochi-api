// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/coingecko/service.go

// Package mock_coingecko is a generated GoMock package.
package mock_coingecko

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	request "github.com/defipod/mochi/pkg/request"
	response "github.com/defipod/mochi/pkg/response"
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

// GetCoin mocks base method.
func (m *MockService) GetCoin(coinID string) (*response.GetCoinResponse, error, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoin", coinID)
	ret0, _ := ret[0].(*response.GetCoinResponse)
	ret1, _ := ret[1].(error)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// GetCoin indicates an expected call of GetCoin.
func (mr *MockServiceMockRecorder) GetCoin(coinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoin", reflect.TypeOf((*MockService)(nil).GetCoin), coinID)
}

// GetCoinPrice mocks base method.
func (m *MockService) GetCoinPrice(coinIDs []string, currency string) (map[string]float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoinPrice", coinIDs, currency)
	ret0, _ := ret[0].(map[string]float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCoinPrice indicates an expected call of GetCoinPrice.
func (mr *MockServiceMockRecorder) GetCoinPrice(coinIDs, currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoinPrice", reflect.TypeOf((*MockService)(nil).GetCoinPrice), coinIDs, currency)
}

// GetCoinsMarketData mocks base method.
func (m *MockService) GetCoinsMarketData(ids []string) ([]response.CoinMarketItemData, error, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoinsMarketData", ids)
	ret0, _ := ret[0].([]response.CoinMarketItemData)
	ret1, _ := ret[1].(error)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// GetCoinsMarketData indicates an expected call of GetCoinsMarketData.
func (mr *MockServiceMockRecorder) GetCoinsMarketData(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoinsMarketData", reflect.TypeOf((*MockService)(nil).GetCoinsMarketData), ids)
}

// GetHistoricalMarketData mocks base method.
func (m *MockService) GetHistoricalMarketData(req *request.GetMarketChartRequest) (*response.CoinPriceHistoryResponse, error, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoricalMarketData", req)
	ret0, _ := ret[0].(*response.CoinPriceHistoryResponse)
	ret1, _ := ret[1].(error)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// GetHistoricalMarketData indicates an expected call of GetHistoricalMarketData.
func (mr *MockServiceMockRecorder) GetHistoricalMarketData(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoricalMarketData", reflect.TypeOf((*MockService)(nil).GetHistoricalMarketData), req)
}

// GetHistoryCoinInfo mocks base method.
func (m *MockService) GetHistoryCoinInfo(sourceSymbol, interval string) ([][]float32, error, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryCoinInfo", sourceSymbol, interval)
	ret0, _ := ret[0].([][]float32)
	ret1, _ := ret[1].(error)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// GetHistoryCoinInfo indicates an expected call of GetHistoryCoinInfo.
func (mr *MockServiceMockRecorder) GetHistoryCoinInfo(sourceSymbol, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryCoinInfo", reflect.TypeOf((*MockService)(nil).GetHistoryCoinInfo), sourceSymbol, interval)
}

// GetSupportedCoins mocks base method.
func (m *MockService) GetSupportedCoins() ([]response.CoingeckoSupportedTokenResponse, error, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportedCoins")
	ret0, _ := ret[0].([]response.CoingeckoSupportedTokenResponse)
	ret1, _ := ret[1].(error)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// GetSupportedCoins indicates an expected call of GetSupportedCoins.
func (mr *MockServiceMockRecorder) GetSupportedCoins() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportedCoins", reflect.TypeOf((*MockService)(nil).GetSupportedCoins))
}
