// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_mesh_translation is a generated GoMock package.
package mock_mesh_translation

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	types "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/types"
	mesh_translation "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/traffic-policy-temp/translation/translators"
)

// MockTranslationValidator is a mock of TranslationValidator interface.
type MockTranslationValidator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslationValidatorMockRecorder
}

// MockTranslationValidatorMockRecorder is the mock recorder for MockTranslationValidator.
type MockTranslationValidatorMockRecorder struct {
	mock *MockTranslationValidator
}

// NewMockTranslationValidator creates a new mock instance.
func NewMockTranslationValidator(ctrl *gomock.Controller) *MockTranslationValidator {
	mock := &MockTranslationValidator{ctrl: ctrl}
	mock.recorder = &MockTranslationValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslationValidator) EXPECT() *MockTranslationValidatorMockRecorder {
	return m.recorder
}

// GetTranslationErrors mocks base method.
func (m *MockTranslationValidator) GetTranslationErrors(meshService *v1alpha1.MeshService, allMeshServices []*v1alpha1.MeshService, mesh *v1alpha1.Mesh, trafficPolicies []*types.MeshServiceStatus_ValidatedTrafficPolicy) []*mesh_translation.TranslationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTranslationErrors", meshService, allMeshServices, mesh, trafficPolicies)
	ret0, _ := ret[0].([]*mesh_translation.TranslationError)
	return ret0
}

// GetTranslationErrors indicates an expected call of GetTranslationErrors.
func (mr *MockTranslationValidatorMockRecorder) GetTranslationErrors(meshService, allMeshServices, mesh, trafficPolicies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTranslationErrors", reflect.TypeOf((*MockTranslationValidator)(nil).GetTranslationErrors), meshService, allMeshServices, mesh, trafficPolicies)
}

// MockDiscoveryLabelsGetter is a mock of DiscoveryLabelsGetter interface.
type MockDiscoveryLabelsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockDiscoveryLabelsGetterMockRecorder
}

// MockDiscoveryLabelsGetterMockRecorder is the mock recorder for MockDiscoveryLabelsGetter.
type MockDiscoveryLabelsGetterMockRecorder struct {
	mock *MockDiscoveryLabelsGetter
}

// NewMockDiscoveryLabelsGetter creates a new mock instance.
func NewMockDiscoveryLabelsGetter(ctrl *gomock.Controller) *MockDiscoveryLabelsGetter {
	mock := &MockDiscoveryLabelsGetter{ctrl: ctrl}
	mock.recorder = &MockDiscoveryLabelsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiscoveryLabelsGetter) EXPECT() *MockDiscoveryLabelsGetterMockRecorder {
	return m.recorder
}

// GetTranslationLabels mocks base method.
func (m *MockDiscoveryLabelsGetter) GetTranslationLabels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTranslationLabels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetTranslationLabels indicates an expected call of GetTranslationLabels.
func (mr *MockDiscoveryLabelsGetterMockRecorder) GetTranslationLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTranslationLabels", reflect.TypeOf((*MockDiscoveryLabelsGetter)(nil).GetTranslationLabels))
}

// MockNamedTranslator is a mock of NamedTranslator interface.
type MockNamedTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockNamedTranslatorMockRecorder
}

// MockNamedTranslatorMockRecorder is the mock recorder for MockNamedTranslator.
type MockNamedTranslatorMockRecorder struct {
	mock *MockNamedTranslator
}

// NewMockNamedTranslator creates a new mock instance.
func NewMockNamedTranslator(ctrl *gomock.Controller) *MockNamedTranslator {
	mock := &MockNamedTranslator{ctrl: ctrl}
	mock.recorder = &MockNamedTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamedTranslator) EXPECT() *MockNamedTranslatorMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockNamedTranslator) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockNamedTranslatorMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockNamedTranslator)(nil).Name))
}

// MockIstioTranslator is a mock of IstioTranslator interface.
type MockIstioTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockIstioTranslatorMockRecorder
}

// MockIstioTranslatorMockRecorder is the mock recorder for MockIstioTranslator.
type MockIstioTranslatorMockRecorder struct {
	mock *MockIstioTranslator
}

// NewMockIstioTranslator creates a new mock instance.
func NewMockIstioTranslator(ctrl *gomock.Controller) *MockIstioTranslator {
	mock := &MockIstioTranslator{ctrl: ctrl}
	mock.recorder = &MockIstioTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIstioTranslator) EXPECT() *MockIstioTranslatorMockRecorder {
	return m.recorder
}

// GetTranslationErrors mocks base method.
func (m *MockIstioTranslator) GetTranslationErrors(meshService *v1alpha1.MeshService, allMeshServices []*v1alpha1.MeshService, mesh *v1alpha1.Mesh, trafficPolicies []*types.MeshServiceStatus_ValidatedTrafficPolicy) []*mesh_translation.TranslationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTranslationErrors", meshService, allMeshServices, mesh, trafficPolicies)
	ret0, _ := ret[0].([]*mesh_translation.TranslationError)
	return ret0
}

// GetTranslationErrors indicates an expected call of GetTranslationErrors.
func (mr *MockIstioTranslatorMockRecorder) GetTranslationErrors(meshService, allMeshServices, mesh, trafficPolicies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTranslationErrors", reflect.TypeOf((*MockIstioTranslator)(nil).GetTranslationErrors), meshService, allMeshServices, mesh, trafficPolicies)
}

// GetTranslationLabels mocks base method.
func (m *MockIstioTranslator) GetTranslationLabels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTranslationLabels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetTranslationLabels indicates an expected call of GetTranslationLabels.
func (mr *MockIstioTranslatorMockRecorder) GetTranslationLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTranslationLabels", reflect.TypeOf((*MockIstioTranslator)(nil).GetTranslationLabels))
}

// Name mocks base method.
func (m *MockIstioTranslator) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockIstioTranslatorMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockIstioTranslator)(nil).Name))
}

// Translate mocks base method.
func (m *MockIstioTranslator) Translate(meshService *v1alpha1.MeshService, allMeshServices []*v1alpha1.MeshService, mesh *v1alpha1.Mesh, trafficPolicies []*types.MeshServiceStatus_ValidatedTrafficPolicy) (*mesh_translation.IstioTranslationOutput, []*mesh_translation.TranslationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", meshService, allMeshServices, mesh, trafficPolicies)
	ret0, _ := ret[0].(*mesh_translation.IstioTranslationOutput)
	ret1, _ := ret[1].([]*mesh_translation.TranslationError)
	return ret0, ret1
}

// Translate indicates an expected call of Translate.
func (mr *MockIstioTranslatorMockRecorder) Translate(meshService, allMeshServices, mesh, trafficPolicies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockIstioTranslator)(nil).Translate), meshService, allMeshServices, mesh, trafficPolicies)
}
