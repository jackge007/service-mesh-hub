// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_csr_generator is a generated GoMock package.
package mock_csr_generator

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1"
	types "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1/types"
	cert_secrets "github.com/solo-io/service-mesh-hub/pkg/common/csr/certgen/secrets"
)

// MockCertClient is a mock of CertClient interface.
type MockCertClient struct {
	ctrl     *gomock.Controller
	recorder *MockCertClientMockRecorder
}

// MockCertClientMockRecorder is the mock recorder for MockCertClient.
type MockCertClientMockRecorder struct {
	mock *MockCertClient
}

// NewMockCertClient creates a new mock instance.
func NewMockCertClient(ctrl *gomock.Controller) *MockCertClient {
	mock := &MockCertClient{ctrl: ctrl}
	mock.recorder = &MockCertClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCertClient) EXPECT() *MockCertClientMockRecorder {
	return m.recorder
}

// EnsureSecretKey mocks base method.
func (m *MockCertClient) EnsureSecretKey(ctx context.Context, obj *v1alpha1.VirtualMeshCertificateSigningRequest) (*cert_secrets.IntermediateCAData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureSecretKey", ctx, obj)
	ret0, _ := ret[0].(*cert_secrets.IntermediateCAData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureSecretKey indicates an expected call of EnsureSecretKey.
func (mr *MockCertClientMockRecorder) EnsureSecretKey(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureSecretKey", reflect.TypeOf((*MockCertClient)(nil).EnsureSecretKey), ctx, obj)
}

// MockPrivateKeyGenerator is a mock of PrivateKeyGenerator interface.
type MockPrivateKeyGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateKeyGeneratorMockRecorder
}

// MockPrivateKeyGeneratorMockRecorder is the mock recorder for MockPrivateKeyGenerator.
type MockPrivateKeyGeneratorMockRecorder struct {
	mock *MockPrivateKeyGenerator
}

// NewMockPrivateKeyGenerator creates a new mock instance.
func NewMockPrivateKeyGenerator(ctrl *gomock.Controller) *MockPrivateKeyGenerator {
	mock := &MockPrivateKeyGenerator{ctrl: ctrl}
	mock.recorder = &MockPrivateKeyGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateKeyGenerator) EXPECT() *MockPrivateKeyGeneratorMockRecorder {
	return m.recorder
}

// GenerateRSA mocks base method.
func (m *MockPrivateKeyGenerator) GenerateRSA(keySize int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRSA", keySize)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRSA indicates an expected call of GenerateRSA.
func (mr *MockPrivateKeyGeneratorMockRecorder) GenerateRSA(keySize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRSA", reflect.TypeOf((*MockPrivateKeyGenerator)(nil).GenerateRSA), keySize)
}

// MockIstioCSRGenerator is a mock of IstioCSRGenerator interface.
type MockIstioCSRGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockIstioCSRGeneratorMockRecorder
}

// MockIstioCSRGeneratorMockRecorder is the mock recorder for MockIstioCSRGenerator.
type MockIstioCSRGeneratorMockRecorder struct {
	mock *MockIstioCSRGenerator
}

// NewMockIstioCSRGenerator creates a new mock instance.
func NewMockIstioCSRGenerator(ctrl *gomock.Controller) *MockIstioCSRGenerator {
	mock := &MockIstioCSRGenerator{ctrl: ctrl}
	mock.recorder = &MockIstioCSRGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIstioCSRGenerator) EXPECT() *MockIstioCSRGeneratorMockRecorder {
	return m.recorder
}

// GenerateIstioCSR mocks base method.
func (m *MockIstioCSRGenerator) GenerateIstioCSR(ctx context.Context, obj *v1alpha1.VirtualMeshCertificateSigningRequest) *types.VirtualMeshCertificateSigningRequestStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIstioCSR", ctx, obj)
	ret0, _ := ret[0].(*types.VirtualMeshCertificateSigningRequestStatus)
	return ret0
}

// GenerateIstioCSR indicates an expected call of GenerateIstioCSR.
func (mr *MockIstioCSRGeneratorMockRecorder) GenerateIstioCSR(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIstioCSR", reflect.TypeOf((*MockIstioCSRGenerator)(nil).GenerateIstioCSR), ctx, obj)
}

// MockVirtualMeshCSRProcessor is a mock of VirtualMeshCSRProcessor interface.
type MockVirtualMeshCSRProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshCSRProcessorMockRecorder
}

// MockVirtualMeshCSRProcessorMockRecorder is the mock recorder for MockVirtualMeshCSRProcessor.
type MockVirtualMeshCSRProcessorMockRecorder struct {
	mock *MockVirtualMeshCSRProcessor
}

// NewMockVirtualMeshCSRProcessor creates a new mock instance.
func NewMockVirtualMeshCSRProcessor(ctrl *gomock.Controller) *MockVirtualMeshCSRProcessor {
	mock := &MockVirtualMeshCSRProcessor{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshCSRProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshCSRProcessor) EXPECT() *MockVirtualMeshCSRProcessorMockRecorder {
	return m.recorder
}

// ProcessUpsert mocks base method.
func (m *MockVirtualMeshCSRProcessor) ProcessUpsert(ctx context.Context, csr *v1alpha1.VirtualMeshCertificateSigningRequest) *types.VirtualMeshCertificateSigningRequestStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessUpsert", ctx, csr)
	ret0, _ := ret[0].(*types.VirtualMeshCertificateSigningRequestStatus)
	return ret0
}

// ProcessUpsert indicates an expected call of ProcessUpsert.
func (mr *MockVirtualMeshCSRProcessorMockRecorder) ProcessUpsert(ctx, csr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessUpsert", reflect.TypeOf((*MockVirtualMeshCSRProcessor)(nil).ProcessUpsert), ctx, csr)
}

// ProcessDelete mocks base method.
func (m *MockVirtualMeshCSRProcessor) ProcessDelete(ctx context.Context, csr *v1alpha1.VirtualMeshCertificateSigningRequest) *types.VirtualMeshCertificateSigningRequestStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessDelete", ctx, csr)
	ret0, _ := ret[0].(*types.VirtualMeshCertificateSigningRequestStatus)
	return ret0
}

// ProcessDelete indicates an expected call of ProcessDelete.
func (mr *MockVirtualMeshCSRProcessorMockRecorder) ProcessDelete(ctx, csr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessDelete", reflect.TypeOf((*MockVirtualMeshCSRProcessor)(nil).ProcessDelete), ctx, csr)
}
