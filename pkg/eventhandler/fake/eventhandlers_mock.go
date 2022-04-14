// Code generated by MockGen. DO NOT EDIT.
// Source: keptn-contrib/job-executor-service/pkg/eventhandler (interfaces: ImageFilter,EventMapper,JobConfigReader,K8s)

// Package fake is a generated GoMock package.
package fake

import (
	config "keptn-contrib/job-executor-service/pkg/config"
	k8sutils "keptn-contrib/job-executor-service/pkg/k8sutils"
	reflect "reflect"
	time "time"

	event "github.com/cloudevents/sdk-go/v2/event"
	gomock "github.com/golang/mock/gomock"
	keptn "github.com/keptn/go-utils/pkg/lib/keptn"
)

// MockImageFilter is a mock of ImageFilter interface.
type MockImageFilter struct {
	ctrl     *gomock.Controller
	recorder *MockImageFilterMockRecorder
}

// MockImageFilterMockRecorder is the mock recorder for MockImageFilter.
type MockImageFilterMockRecorder struct {
	mock *MockImageFilter
}

// NewMockImageFilter creates a new mock instance.
func NewMockImageFilter(ctrl *gomock.Controller) *MockImageFilter {
	mock := &MockImageFilter{ctrl: ctrl}
	mock.recorder = &MockImageFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageFilter) EXPECT() *MockImageFilterMockRecorder {
	return m.recorder
}

// IsImageAllowed mocks base method.
func (m *MockImageFilter) IsImageAllowed(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsImageAllowed", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsImageAllowed indicates an expected call of IsImageAllowed.
func (mr *MockImageFilterMockRecorder) IsImageAllowed(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsImageAllowed", reflect.TypeOf((*MockImageFilter)(nil).IsImageAllowed), arg0)
}

// MockEventMapper is a mock of EventMapper interface.
type MockEventMapper struct {
	ctrl     *gomock.Controller
	recorder *MockEventMapperMockRecorder
}

// MockEventMapperMockRecorder is the mock recorder for MockEventMapper.
type MockEventMapperMockRecorder struct {
	mock *MockEventMapper
}

// NewMockEventMapper creates a new mock instance.
func NewMockEventMapper(ctrl *gomock.Controller) *MockEventMapper {
	mock := &MockEventMapper{ctrl: ctrl}
	mock.recorder = &MockEventMapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventMapper) EXPECT() *MockEventMapperMockRecorder {
	return m.recorder
}

// Map mocks base method.
func (m *MockEventMapper) Map(arg0 event.Event) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Map indicates an expected call of Map.
func (mr *MockEventMapperMockRecorder) Map(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockEventMapper)(nil).Map), arg0)
}

// MockJobConfigReader is a mock of JobConfigReader interface.
type MockJobConfigReader struct {
	ctrl     *gomock.Controller
	recorder *MockJobConfigReaderMockRecorder
}

// MockJobConfigReaderMockRecorder is the mock recorder for MockJobConfigReader.
type MockJobConfigReaderMockRecorder struct {
	mock *MockJobConfigReader
}

// NewMockJobConfigReader creates a new mock instance.
func NewMockJobConfigReader(ctrl *gomock.Controller) *MockJobConfigReader {
	mock := &MockJobConfigReader{ctrl: ctrl}
	mock.recorder = &MockJobConfigReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobConfigReader) EXPECT() *MockJobConfigReaderMockRecorder {
	return m.recorder
}

// GetJobConfig mocks base method.
func (m *MockJobConfigReader) GetJobConfig() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobConfig")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobConfig indicates an expected call of GetJobConfig.
func (mr *MockJobConfigReaderMockRecorder) GetJobConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobConfig", reflect.TypeOf((*MockJobConfigReader)(nil).GetJobConfig))
}

// MockK8s is a mock of K8s interface.
type MockK8s struct {
	ctrl     *gomock.Controller
	recorder *MockK8sMockRecorder
}

// MockK8sMockRecorder is the mock recorder for MockK8s.
type MockK8sMockRecorder struct {
	mock *MockK8s
}

// NewMockK8s creates a new mock instance.
func NewMockK8s(ctrl *gomock.Controller) *MockK8s {
	mock := &MockK8s{ctrl: ctrl}
	mock.recorder = &MockK8sMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockK8s) EXPECT() *MockK8sMockRecorder {
	return m.recorder
}

// AwaitK8sJobDone mocks base method.
func (m *MockK8s) AwaitK8sJobDone(arg0 string, arg1, arg2 time.Duration, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AwaitK8sJobDone", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AwaitK8sJobDone indicates an expected call of AwaitK8sJobDone.
func (mr *MockK8sMockRecorder) AwaitK8sJobDone(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AwaitK8sJobDone", reflect.TypeOf((*MockK8s)(nil).AwaitK8sJobDone), arg0, arg1, arg2, arg3)
}

// ConnectToCluster mocks base method.
func (m *MockK8s) ConnectToCluster() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectToCluster")
	ret0, _ := ret[0].(error)
	return ret0
}

// ConnectToCluster indicates an expected call of ConnectToCluster.
func (mr *MockK8sMockRecorder) ConnectToCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectToCluster", reflect.TypeOf((*MockK8s)(nil).ConnectToCluster))
}

// CreateK8sJob mocks base method.
func (m *MockK8s) CreateK8sJob(arg0 string, arg1 *config.Action, arg2 config.Task, arg3 keptn.EventProperties, arg4 k8sutils.JobSettings, arg5 interface{}, arg6 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateK8sJob", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateK8sJob indicates an expected call of CreateK8sJob.
func (mr *MockK8sMockRecorder) CreateK8sJob(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateK8sJob", reflect.TypeOf((*MockK8s)(nil).CreateK8sJob), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// GetLogsOfPod mocks base method.
func (m *MockK8s) GetLogsOfPod(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogsOfPod", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogsOfPod indicates an expected call of GetLogsOfPod.
func (mr *MockK8sMockRecorder) GetLogsOfPod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogsOfPod", reflect.TypeOf((*MockK8s)(nil).GetLogsOfPod), arg0, arg1)
}
