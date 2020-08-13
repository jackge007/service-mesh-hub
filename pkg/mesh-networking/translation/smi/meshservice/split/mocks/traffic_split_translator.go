// Code generated by MockGen. DO NOT EDIT.
// Source: ./traffic_split_translator.go

// Package mock_split is a generated GoMock package.
package mock_split

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha3"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	input "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/input"
	reporting "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/reporting"
	reflect "reflect"
)

// MockTranslator is a mock of Translator interface
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorMockRecorder
}

// MockTranslatorMockRecorder is the mock recorder for MockTranslator
type MockTranslatorMockRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTranslator) EXPECT() *MockTranslatorMockRecorder {
	return m.recorder
}

// Translate mocks base method
func (m *MockTranslator) Translate(ctx context.Context, in input.Snapshot, meshService *v1alpha2.MeshService, reporter reporting.Reporter) *v1alpha3.TrafficSplit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", ctx, in, meshService, reporter)
	ret0, _ := ret[0].(*v1alpha3.TrafficSplit)
	return ret0
}

// Translate indicates an expected call of Translate
func (mr *MockTranslatorMockRecorder) Translate(ctx, in, meshService, reporter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockTranslator)(nil).Translate), ctx, in, meshService, reporter)
}
