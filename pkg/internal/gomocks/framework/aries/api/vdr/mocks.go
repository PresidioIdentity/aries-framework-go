// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdr (interfaces: Registry)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	did "github.com/hyperledger/aries-framework-go/pkg/doc/did"
	vdr "github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdr"
	reflect "reflect"
)

// MockRegistry is a mock of Registry interface
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockRegistry) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockRegistryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRegistry)(nil).Close))
}

// Create mocks base method
func (m *MockRegistry) Create(arg0 string, arg1 *did.Doc, arg2 ...vdr.DIDMethodOption) (*did.DocResolution, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*did.DocResolution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRegistryMockRecorder) Create(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRegistry)(nil).Create), varargs...)
}

// Resolve mocks base method
func (m *MockRegistry) Resolve(arg0 string, arg1 ...vdr.ResolveOption) (*did.DocResolution, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Resolve", varargs...)
	ret0, _ := ret[0].(*did.DocResolution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve
func (mr *MockRegistryMockRecorder) Resolve(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockRegistry)(nil).Resolve), varargs...)
}
