// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Psiphon-Labs/quic-go (interfaces: SendConn)

// Package quic is a generated GoMock package.
package quic

import (
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSendConn is a mock of SendConn interface.
type MockSendConn struct {
	ctrl     *gomock.Controller
	recorder *MockSendConnMockRecorder
}

// MockSendConnMockRecorder is the mock recorder for MockSendConn.
type MockSendConnMockRecorder struct {
	mock *MockSendConn
}

// NewMockSendConn creates a new mock instance.
func NewMockSendConn(ctrl *gomock.Controller) *MockSendConn {
	mock := &MockSendConn{ctrl: ctrl}
	mock.recorder = &MockSendConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSendConn) EXPECT() *MockSendConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSendConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSendConnMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSendConn)(nil).Close))
}

// LocalAddr mocks base method.
func (m *MockSendConn) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockSendConnMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockSendConn)(nil).LocalAddr))
}

// RemoteAddr mocks base method.
func (m *MockSendConn) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockSendConnMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockSendConn)(nil).RemoteAddr))
}

// Write mocks base method.
func (m *MockSendConn) Write(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockSendConnMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSendConn)(nil).Write), arg0)
}
