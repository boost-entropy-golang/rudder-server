// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/app (interfaces: Interface)

// Package mock_app is a generated GoMock package.
package mock_app

import (
	gomock "github.com/golang/mock/gomock"
	app "github.com/rudderlabs/rudder-server/app"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Features mocks base method
func (m *MockInterface) Features() *app.Features {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Features")
	ret0, _ := ret[0].(*app.Features)
	return ret0
}

// Features indicates an expected call of Features
func (mr *MockInterfaceMockRecorder) Features() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Features", reflect.TypeOf((*MockInterface)(nil).Features))
}

// Options mocks base method
func (m *MockInterface) Options() *app.Options {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Options")
	ret0, _ := ret[0].(*app.Options)
	return ret0
}

// Options indicates an expected call of Options
func (mr *MockInterfaceMockRecorder) Options() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockInterface)(nil).Options))
}

// Setup mocks base method
func (m *MockInterface) Setup() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Setup")
}

// Setup indicates an expected call of Setup
func (mr *MockInterfaceMockRecorder) Setup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockInterface)(nil).Setup))
}

// Stop mocks base method
func (m *MockInterface) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockInterfaceMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockInterface)(nil).Stop))
}
