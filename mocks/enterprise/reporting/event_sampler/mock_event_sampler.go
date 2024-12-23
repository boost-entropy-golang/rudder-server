// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/enterprise/reporting/event_sampler (interfaces: EventSampler)
//
// Generated by this command:
//
//	mockgen -destination=../../../mocks/enterprise/reporting/event_sampler/mock_event_sampler.go -package=mocks github.com/rudderlabs/rudder-server/enterprise/reporting/event_sampler EventSampler
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockEventSampler is a mock of EventSampler interface.
type MockEventSampler struct {
	ctrl     *gomock.Controller
	recorder *MockEventSamplerMockRecorder
	isgomock struct{}
}

// MockEventSamplerMockRecorder is the mock recorder for MockEventSampler.
type MockEventSamplerMockRecorder struct {
	mock *MockEventSampler
}

// NewMockEventSampler creates a new mock instance.
func NewMockEventSampler(ctrl *gomock.Controller) *MockEventSampler {
	mock := &MockEventSampler{ctrl: ctrl}
	mock.recorder = &MockEventSamplerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventSampler) EXPECT() *MockEventSamplerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockEventSampler) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockEventSamplerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEventSampler)(nil).Close))
}

// Get mocks base method.
func (m *MockEventSampler) Get(key string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockEventSamplerMockRecorder) Get(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEventSampler)(nil).Get), key)
}

// Put mocks base method.
func (m *MockEventSampler) Put(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockEventSamplerMockRecorder) Put(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockEventSampler)(nil).Put), key)
}
