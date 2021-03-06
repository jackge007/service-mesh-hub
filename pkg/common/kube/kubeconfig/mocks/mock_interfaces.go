// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_kubeconfig is a generated GoMock package.
package mock_kubeconfig

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kubeconfig "github.com/solo-io/service-mesh-hub/pkg/common/kube/kubeconfig"
	v1 "k8s.io/api/core/v1"
	resource "k8s.io/cli-runtime/pkg/resource"
	rest "k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	api "k8s.io/client-go/tools/clientcmd/api"
)

// MockConverter is a mock of Converter interface.
type MockConverter struct {
	ctrl     *gomock.Controller
	recorder *MockConverterMockRecorder
}

// MockConverterMockRecorder is the mock recorder for MockConverter.
type MockConverterMockRecorder struct {
	mock *MockConverter
}

// NewMockConverter creates a new mock instance.
func NewMockConverter(ctrl *gomock.Controller) *MockConverter {
	mock := &MockConverter{ctrl: ctrl}
	mock.recorder = &MockConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConverter) EXPECT() *MockConverterMockRecorder {
	return m.recorder
}

// ConfigToSecret mocks base method.
func (m *MockConverter) ConfigToSecret(secretName, secretNamespace string, config *kubeconfig.KubeConfig) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigToSecret", secretName, secretNamespace, config)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigToSecret indicates an expected call of ConfigToSecret.
func (mr *MockConverterMockRecorder) ConfigToSecret(secretName, secretNamespace, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigToSecret", reflect.TypeOf((*MockConverter)(nil).ConfigToSecret), secretName, secretNamespace, config)
}

// SecretToConfig mocks base method.
func (m *MockConverter) SecretToConfig(secret *v1.Secret) (string, *kubeconfig.ConvertedConfigs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretToConfig", secret)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*kubeconfig.ConvertedConfigs)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SecretToConfig indicates an expected call of SecretToConfig.
func (mr *MockConverterMockRecorder) SecretToConfig(secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretToConfig", reflect.TypeOf((*MockConverter)(nil).SecretToConfig), secret)
}

// MockKubeLoader is a mock of KubeLoader interface.
type MockKubeLoader struct {
	ctrl     *gomock.Controller
	recorder *MockKubeLoaderMockRecorder
}

// MockKubeLoaderMockRecorder is the mock recorder for MockKubeLoader.
type MockKubeLoaderMockRecorder struct {
	mock *MockKubeLoader
}

// NewMockKubeLoader creates a new mock instance.
func NewMockKubeLoader(ctrl *gomock.Controller) *MockKubeLoader {
	mock := &MockKubeLoader{ctrl: ctrl}
	mock.recorder = &MockKubeLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeLoader) EXPECT() *MockKubeLoaderMockRecorder {
	return m.recorder
}

// GetConfigWithContext mocks base method.
func (m *MockKubeLoader) GetConfigWithContext(masterURL, kubeconfigPath, context string) (clientcmd.ClientConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigWithContext", masterURL, kubeconfigPath, context)
	ret0, _ := ret[0].(clientcmd.ClientConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigWithContext indicates an expected call of GetConfigWithContext.
func (mr *MockKubeLoaderMockRecorder) GetConfigWithContext(masterURL, kubeconfigPath, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigWithContext", reflect.TypeOf((*MockKubeLoader)(nil).GetConfigWithContext), masterURL, kubeconfigPath, context)
}

// GetRestConfigForContext mocks base method.
func (m *MockKubeLoader) GetRestConfigForContext(path, context string) (*rest.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestConfigForContext", path, context)
	ret0, _ := ret[0].(*rest.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestConfigForContext indicates an expected call of GetRestConfigForContext.
func (mr *MockKubeLoaderMockRecorder) GetRestConfigForContext(path, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestConfigForContext", reflect.TypeOf((*MockKubeLoader)(nil).GetRestConfigForContext), path, context)
}

// GetRawConfigForContext mocks base method.
func (m *MockKubeLoader) GetRawConfigForContext(path, context string) (api.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawConfigForContext", path, context)
	ret0, _ := ret[0].(api.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawConfigForContext indicates an expected call of GetRawConfigForContext.
func (mr *MockKubeLoaderMockRecorder) GetRawConfigForContext(path, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawConfigForContext", reflect.TypeOf((*MockKubeLoader)(nil).GetRawConfigForContext), path, context)
}

// RESTClientGetter mocks base method.
func (m *MockKubeLoader) RESTClientGetter(path, context string) resource.RESTClientGetter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClientGetter", path, context)
	ret0, _ := ret[0].(resource.RESTClientGetter)
	return ret0
}

// RESTClientGetter indicates an expected call of RESTClientGetter.
func (mr *MockKubeLoaderMockRecorder) RESTClientGetter(path, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClientGetter", reflect.TypeOf((*MockKubeLoader)(nil).RESTClientGetter), path, context)
}

// GetRestConfigFromBytes mocks base method.
func (m *MockKubeLoader) GetRestConfigFromBytes(config []byte) (*rest.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestConfigFromBytes", config)
	ret0, _ := ret[0].(*rest.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestConfigFromBytes indicates an expected call of GetRestConfigFromBytes.
func (mr *MockKubeLoaderMockRecorder) GetRestConfigFromBytes(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestConfigFromBytes", reflect.TypeOf((*MockKubeLoader)(nil).GetRestConfigFromBytes), config)
}

// MockKubeConfigLookup is a mock of KubeConfigLookup interface.
type MockKubeConfigLookup struct {
	ctrl     *gomock.Controller
	recorder *MockKubeConfigLookupMockRecorder
}

// MockKubeConfigLookupMockRecorder is the mock recorder for MockKubeConfigLookup.
type MockKubeConfigLookupMockRecorder struct {
	mock *MockKubeConfigLookup
}

// NewMockKubeConfigLookup creates a new mock instance.
func NewMockKubeConfigLookup(ctrl *gomock.Controller) *MockKubeConfigLookup {
	mock := &MockKubeConfigLookup{ctrl: ctrl}
	mock.recorder = &MockKubeConfigLookupMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeConfigLookup) EXPECT() *MockKubeConfigLookupMockRecorder {
	return m.recorder
}

// FromCluster mocks base method.
func (m *MockKubeConfigLookup) FromCluster(ctx context.Context, clusterName string) (*kubeconfig.ConvertedConfigs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromCluster", ctx, clusterName)
	ret0, _ := ret[0].(*kubeconfig.ConvertedConfigs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FromCluster indicates an expected call of FromCluster.
func (mr *MockKubeConfigLookupMockRecorder) FromCluster(ctx, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromCluster", reflect.TypeOf((*MockKubeConfigLookup)(nil).FromCluster), ctx, clusterName)
}
