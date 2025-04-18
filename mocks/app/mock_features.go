// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/app (interfaces: SuppressUserFeature)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/app/mock_features.go -package=mock_app github.com/rudderlabs/rudder-server/app SuppressUserFeature
//

// Package mock_app is a generated GoMock package.
package mock_app

import (
	context "context"
	reflect "reflect"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	types "github.com/rudderlabs/rudder-server/utils/types"
	gomock "go.uber.org/mock/gomock"
)

// MockSuppressUserFeature is a mock of SuppressUserFeature interface.
type MockSuppressUserFeature struct {
	ctrl     *gomock.Controller
	recorder *MockSuppressUserFeatureMockRecorder
	isgomock struct{}
}

// MockSuppressUserFeatureMockRecorder is the mock recorder for MockSuppressUserFeature.
type MockSuppressUserFeatureMockRecorder struct {
	mock *MockSuppressUserFeature
}

// NewMockSuppressUserFeature creates a new mock instance.
func NewMockSuppressUserFeature(ctrl *gomock.Controller) *MockSuppressUserFeature {
	mock := &MockSuppressUserFeature{ctrl: ctrl}
	mock.recorder = &MockSuppressUserFeatureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSuppressUserFeature) EXPECT() *MockSuppressUserFeatureMockRecorder {
	return m.recorder
}

// Setup mocks base method.
func (m *MockSuppressUserFeature) Setup(ctx context.Context, backendConfig backendconfig.BackendConfig) (types.UserSuppression, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup", ctx, backendConfig)
	ret0, _ := ret[0].(types.UserSuppression)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Setup indicates an expected call of Setup.
func (mr *MockSuppressUserFeatureMockRecorder) Setup(ctx, backendConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockSuppressUserFeature)(nil).Setup), ctx, backendConfig)
}
