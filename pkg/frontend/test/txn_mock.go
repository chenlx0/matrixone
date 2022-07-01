// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package mock_moengine is a generated GoMock package.
package mock_frontend

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	engine "github.com/matrixorigin/matrixone/pkg/vm/engine"
	moengine "github.com/matrixorigin/matrixone/pkg/vm/engine/tae/moengine"
)

// MockTxn is a mock of Txn interface.
type MockTxn struct {
	ctrl     *gomock.Controller
	recorder *MockTxnMockRecorder
}

// MockTxnMockRecorder is the mock recorder for MockTxn.
type MockTxnMockRecorder struct {
	mock *MockTxn
}

// NewMockTxn creates a new mock instance.
func NewMockTxn(ctrl *gomock.Controller) *MockTxn {
	mock := &MockTxn{ctrl: ctrl}
	mock.recorder = &MockTxnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxn) EXPECT() *MockTxnMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockTxn) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTxnMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTxn)(nil).Commit))
}

// GetCtx mocks base method.
func (m *MockTxn) GetCtx() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCtx")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetCtx indicates an expected call of GetCtx.
func (mr *MockTxnMockRecorder) GetCtx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCtx", reflect.TypeOf((*MockTxn)(nil).GetCtx))
}

// GetError mocks base method.
func (m *MockTxn) GetError() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetError")
	ret0, _ := ret[0].(error)
	return ret0
}

// GetError indicates an expected call of GetError.
func (mr *MockTxnMockRecorder) GetError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetError", reflect.TypeOf((*MockTxn)(nil).GetError))
}

// GetID mocks base method.
func (m *MockTxn) GetID() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockTxnMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockTxn)(nil).GetID))
}

// Repr mocks base method.
func (m *MockTxn) Repr() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Repr")
	ret0, _ := ret[0].(string)
	return ret0
}

// Repr indicates an expected call of Repr.
func (mr *MockTxnMockRecorder) Repr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Repr", reflect.TypeOf((*MockTxn)(nil).Repr))
}

// Rollback mocks base method.
func (m *MockTxn) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTxnMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTxn)(nil).Rollback))
}

// String mocks base method.
func (m *MockTxn) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockTxnMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockTxn)(nil).String))
}

// MockTxnEngine is a mock of TxnEngine interface.
type MockTxnEngine struct {
	ctrl     *gomock.Controller
	recorder *MockTxnEngineMockRecorder
}

// MockTxnEngineMockRecorder is the mock recorder for MockTxnEngine.
type MockTxnEngineMockRecorder struct {
	mock *MockTxnEngine
}

// NewMockTxnEngine creates a new mock instance.
func NewMockTxnEngine(ctrl *gomock.Controller) *MockTxnEngine {
	mock := &MockTxnEngine{ctrl: ctrl}
	mock.recorder = &MockTxnEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxnEngine) EXPECT() *MockTxnEngineMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTxnEngine) Create(arg0 uint64, arg1 string, arg2 int, arg3 engine.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTxnEngineMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTxnEngine)(nil).Create), arg0, arg1, arg2, arg3)
}

// Database mocks base method.
func (m *MockTxnEngine) Database(arg0 string, arg1 engine.Snapshot) (engine.Database, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Database", arg0, arg1)
	ret0, _ := ret[0].(engine.Database)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Database indicates an expected call of Database.
func (mr *MockTxnEngineMockRecorder) Database(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Database", reflect.TypeOf((*MockTxnEngine)(nil).Database), arg0, arg1)
}

// Databases mocks base method.
func (m *MockTxnEngine) Databases(arg0 engine.Snapshot) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Databases", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// Databases indicates an expected call of Databases.
func (mr *MockTxnEngineMockRecorder) Databases(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Databases", reflect.TypeOf((*MockTxnEngine)(nil).Databases), arg0)
}

// Delete mocks base method.
func (m *MockTxnEngine) Delete(arg0 uint64, arg1 string, arg2 engine.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTxnEngineMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTxnEngine)(nil).Delete), arg0, arg1, arg2)
}

// Node mocks base method.
func (m *MockTxnEngine) Node(arg0 string, arg1 engine.Snapshot) *engine.NodeInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Node", arg0, arg1)
	ret0, _ := ret[0].(*engine.NodeInfo)
	return ret0
}

// Node indicates an expected call of Node.
func (mr *MockTxnEngineMockRecorder) Node(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*MockTxnEngine)(nil).Node), arg0, arg1)
}

// StartTxn mocks base method.
func (m *MockTxnEngine) StartTxn(info []byte) (moengine.Txn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTxn", info)
	ret0, _ := ret[0].(moengine.Txn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartTxn indicates an expected call of StartTxn.
func (mr *MockTxnEngineMockRecorder) StartTxn(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTxn", reflect.TypeOf((*MockTxnEngine)(nil).StartTxn), info)
}
