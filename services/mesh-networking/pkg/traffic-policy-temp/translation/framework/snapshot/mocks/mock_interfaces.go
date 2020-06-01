// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_snapshot is a generated GoMock package.
package mock_snapshot

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	snapshot "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/traffic-policy-temp/translation/framework/snapshot"
)

// MockTranslationSnapshotAccumulator is a mock of TranslationSnapshotAccumulator interface.
type MockTranslationSnapshotAccumulator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslationSnapshotAccumulatorMockRecorder
}

// MockTranslationSnapshotAccumulatorMockRecorder is the mock recorder for MockTranslationSnapshotAccumulator.
type MockTranslationSnapshotAccumulatorMockRecorder struct {
	mock *MockTranslationSnapshotAccumulator
}

// NewMockTranslationSnapshotAccumulator creates a new mock instance.
func NewMockTranslationSnapshotAccumulator(ctrl *gomock.Controller) *MockTranslationSnapshotAccumulator {
	mock := &MockTranslationSnapshotAccumulator{ctrl: ctrl}
	mock.recorder = &MockTranslationSnapshotAccumulatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslationSnapshotAccumulator) EXPECT() *MockTranslationSnapshotAccumulatorMockRecorder {
	return m.recorder
}

// AccumulateFromTranslation mocks base method.
func (m *MockTranslationSnapshotAccumulator) AccumulateFromTranslation(snapshotInProgress *snapshot.TranslatedSnapshot, meshService *v1alpha1.MeshService, allMeshServices []*v1alpha1.MeshService, mesh *v1alpha1.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccumulateFromTranslation", snapshotInProgress, meshService, allMeshServices, mesh)
	ret0, _ := ret[0].(error)
	return ret0
}

// AccumulateFromTranslation indicates an expected call of AccumulateFromTranslation.
func (mr *MockTranslationSnapshotAccumulatorMockRecorder) AccumulateFromTranslation(snapshotInProgress, meshService, allMeshServices, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccumulateFromTranslation", reflect.TypeOf((*MockTranslationSnapshotAccumulator)(nil).AccumulateFromTranslation), snapshotInProgress, meshService, allMeshServices, mesh)
}

// MockTranslationSnapshotReconciler is a mock of TranslationSnapshotReconciler interface.
type MockTranslationSnapshotReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockTranslationSnapshotReconcilerMockRecorder
}

// MockTranslationSnapshotReconcilerMockRecorder is the mock recorder for MockTranslationSnapshotReconciler.
type MockTranslationSnapshotReconcilerMockRecorder struct {
	mock *MockTranslationSnapshotReconciler
}

// NewMockTranslationSnapshotReconciler creates a new mock instance.
func NewMockTranslationSnapshotReconciler(ctrl *gomock.Controller) *MockTranslationSnapshotReconciler {
	mock := &MockTranslationSnapshotReconciler{ctrl: ctrl}
	mock.recorder = &MockTranslationSnapshotReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslationSnapshotReconciler) EXPECT() *MockTranslationSnapshotReconcilerMockRecorder {
	return m.recorder
}

// InitializeClusterNameToSnapshot mocks base method.
func (m *MockTranslationSnapshotReconciler) InitializeClusterNameToSnapshot(knownMeshes []*v1alpha1.Mesh) map[string]*snapshot.TranslatedSnapshot {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeClusterNameToSnapshot", knownMeshes)
	ret0, _ := ret[0].(map[string]*snapshot.TranslatedSnapshot)
	return ret0
}

// InitializeClusterNameToSnapshot indicates an expected call of InitializeClusterNameToSnapshot.
func (mr *MockTranslationSnapshotReconcilerMockRecorder) InitializeClusterNameToSnapshot(knownMeshes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeClusterNameToSnapshot", reflect.TypeOf((*MockTranslationSnapshotReconciler)(nil).InitializeClusterNameToSnapshot), knownMeshes)
}

// ReconcileAllSnapshots mocks base method.
func (m *MockTranslationSnapshotReconciler) ReconcileAllSnapshots(ctx context.Context, clusterNameToSnapshot map[string]*snapshot.TranslatedSnapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAllSnapshots", ctx, clusterNameToSnapshot)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileAllSnapshots indicates an expected call of ReconcileAllSnapshots.
func (mr *MockTranslationSnapshotReconcilerMockRecorder) ReconcileAllSnapshots(ctx, clusterNameToSnapshot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAllSnapshots", reflect.TypeOf((*MockTranslationSnapshotReconciler)(nil).ReconcileAllSnapshots), ctx, clusterNameToSnapshot)
}
