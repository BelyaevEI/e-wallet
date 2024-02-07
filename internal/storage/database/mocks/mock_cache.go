// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/cache/cache.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/BelyaevEI/e-wallet/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// AddWallet mocks base method.
func (m *MockCache) AddWallet(wallet models.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWallet", wallet)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWallet indicates an expected call of AddWallet.
func (mr *MockCacheMockRecorder) AddWallet(wallet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWallet", reflect.TypeOf((*MockCache)(nil).AddWallet), wallet)
}

// FillCache mocks base method.
func (m *MockCache) FillCache(wallets []models.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FillCache", wallets)
	ret0, _ := ret[0].(error)
	return ret0
}

// FillCache indicates an expected call of FillCache.
func (mr *MockCacheMockRecorder) FillCache(wallets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FillCache", reflect.TypeOf((*MockCache)(nil).FillCache), wallets)
}

// GetBalance mocks base method.
func (m *MockCache) GetBalance(walletID uint32) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", walletID)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockCacheMockRecorder) GetBalance(walletID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockCache)(nil).GetBalance), walletID)
}

// ModifyWallet mocks base method.
func (m *MockCache) ModifyWallet(walletID uint32, amount float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyWallet", walletID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyWallet indicates an expected call of ModifyWallet.
func (mr *MockCacheMockRecorder) ModifyWallet(walletID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyWallet", reflect.TypeOf((*MockCache)(nil).ModifyWallet), walletID, amount)
}