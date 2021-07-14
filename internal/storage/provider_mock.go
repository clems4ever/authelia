// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/authelia/authelia/internal/storage (interfaces: Provider)

// Package storage is a generated GoMock package.
package storage

import (
	models "github.com/authelia/authelia/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// AppendAuthenticationLog mocks base method
func (m *MockProvider) AppendAuthenticationLog(arg0 models.AuthenticationAttempt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendAuthenticationLog", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendAuthenticationLog indicates an expected call of AppendAuthenticationLog
func (mr *MockProviderMockRecorder) AppendAuthenticationLog(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendAuthenticationLog", reflect.TypeOf((*MockProvider)(nil).AppendAuthenticationLog), arg0)
}

// DeleteTOTPSecret mocks base method
func (m *MockProvider) DeleteTOTPSecret(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTOTPSecret", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTOTPSecret indicates an expected call of DeleteTOTPSecret
func (mr *MockProviderMockRecorder) DeleteTOTPSecret(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTOTPSecret", reflect.TypeOf((*MockProvider)(nil).DeleteTOTPSecret), arg0)
}

// FindIdentityVerificationToken mocks base method
func (m *MockProvider) FindIdentityVerificationToken(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIdentityVerificationToken", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIdentityVerificationToken indicates an expected call of FindIdentityVerificationToken
func (mr *MockProviderMockRecorder) FindIdentityVerificationToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIdentityVerificationToken", reflect.TypeOf((*MockProvider)(nil).FindIdentityVerificationToken), arg0)
}

// LoadLatestAuthenticationLogs mocks base method
func (m *MockProvider) LoadLatestAuthenticationLogs(arg0 string, arg1 time.Time) ([]models.AuthenticationAttempt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadLatestAuthenticationLogs", arg0, arg1)
	ret0, _ := ret[0].([]models.AuthenticationAttempt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadLatestAuthenticationLogs indicates an expected call of LoadLatestAuthenticationLogs
func (mr *MockProviderMockRecorder) LoadLatestAuthenticationLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadLatestAuthenticationLogs", reflect.TypeOf((*MockProvider)(nil).LoadLatestAuthenticationLogs), arg0, arg1)
}

// LoadPreferred2FAMethod mocks base method
func (m *MockProvider) LoadPreferred2FAMethod(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadPreferred2FAMethod", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadPreferred2FAMethod indicates an expected call of LoadPreferred2FAMethod
func (mr *MockProviderMockRecorder) LoadPreferred2FAMethod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadPreferred2FAMethod", reflect.TypeOf((*MockProvider)(nil).LoadPreferred2FAMethod), arg0)
}

// LoadPreferredDuoDevice mocks base method
func (m *MockProvider) LoadPreferredDuoDevice(arg0 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadPreferredDuoDevice", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadPreferredDuoDevice indicates an expected call of LoadPreferredDuoDevice
func (mr *MockProviderMockRecorder) LoadPreferredDuoDevice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadPreferredDuoDevice", reflect.TypeOf((*MockProvider)(nil).LoadPreferredDuoDevice), arg0)
}

// LoadTOTPSecret mocks base method
func (m *MockProvider) LoadTOTPSecret(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadTOTPSecret", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTOTPSecret indicates an expected call of LoadTOTPSecret
func (mr *MockProviderMockRecorder) LoadTOTPSecret(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTOTPSecret", reflect.TypeOf((*MockProvider)(nil).LoadTOTPSecret), arg0)
}

// LoadU2FDeviceHandle mocks base method
func (m *MockProvider) LoadU2FDeviceHandle(arg0 string) ([]byte, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadU2FDeviceHandle", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadU2FDeviceHandle indicates an expected call of LoadU2FDeviceHandle
func (mr *MockProviderMockRecorder) LoadU2FDeviceHandle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadU2FDeviceHandle", reflect.TypeOf((*MockProvider)(nil).LoadU2FDeviceHandle), arg0)
}

// RemoveIdentityVerificationToken mocks base method
func (m *MockProvider) RemoveIdentityVerificationToken(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIdentityVerificationToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveIdentityVerificationToken indicates an expected call of RemoveIdentityVerificationToken
func (mr *MockProviderMockRecorder) RemoveIdentityVerificationToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIdentityVerificationToken", reflect.TypeOf((*MockProvider)(nil).RemoveIdentityVerificationToken), arg0)
}

// SaveIdentityVerificationToken mocks base method
func (m *MockProvider) SaveIdentityVerificationToken(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveIdentityVerificationToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveIdentityVerificationToken indicates an expected call of SaveIdentityVerificationToken
func (mr *MockProviderMockRecorder) SaveIdentityVerificationToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveIdentityVerificationToken", reflect.TypeOf((*MockProvider)(nil).SaveIdentityVerificationToken), arg0)
}

// SavePreferred2FAMethod mocks base method
func (m *MockProvider) SavePreferred2FAMethod(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePreferred2FAMethod", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePreferred2FAMethod indicates an expected call of SavePreferred2FAMethod
func (mr *MockProviderMockRecorder) SavePreferred2FAMethod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePreferred2FAMethod", reflect.TypeOf((*MockProvider)(nil).SavePreferred2FAMethod), arg0, arg1)
}

// SavePreferredDuoDevice mocks base method
func (m *MockProvider) SavePreferredDuoDevice(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePreferredDuoDevice", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePreferredDuoDevice indicates an expected call of SavePreferredDuoDevice
func (mr *MockProviderMockRecorder) SavePreferredDuoDevice(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePreferredDuoDevice", reflect.TypeOf((*MockProvider)(nil).SavePreferredDuoDevice), arg0, arg1, arg2)
}

// SaveTOTPSecret mocks base method
func (m *MockProvider) SaveTOTPSecret(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTOTPSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTOTPSecret indicates an expected call of SaveTOTPSecret
func (mr *MockProviderMockRecorder) SaveTOTPSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTOTPSecret", reflect.TypeOf((*MockProvider)(nil).SaveTOTPSecret), arg0, arg1)
}

// SaveU2FDeviceHandle mocks base method
func (m *MockProvider) SaveU2FDeviceHandle(arg0 string, arg1, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveU2FDeviceHandle", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveU2FDeviceHandle indicates an expected call of SaveU2FDeviceHandle
func (mr *MockProviderMockRecorder) SaveU2FDeviceHandle(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveU2FDeviceHandle", reflect.TypeOf((*MockProvider)(nil).SaveU2FDeviceHandle), arg0, arg1, arg2)
}
