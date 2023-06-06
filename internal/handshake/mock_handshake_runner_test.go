// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Psiphon-Labs/quic-go/internal/handshake (interfaces: HandshakeRunner)

// Package handshake is a generated GoMock package.
package handshake

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	protocol "github.com/Psiphon-Labs/quic-go/internal/protocol"
	wire "github.com/Psiphon-Labs/quic-go/internal/wire"
)

// MockHandshakeRunner is a mock of HandshakeRunner interface.
type MockHandshakeRunner struct {
	ctrl     *gomock.Controller
	recorder *MockHandshakeRunnerMockRecorder
}

// MockHandshakeRunnerMockRecorder is the mock recorder for MockHandshakeRunner.
type MockHandshakeRunnerMockRecorder struct {
	mock *MockHandshakeRunner
}

// NewMockHandshakeRunner creates a new mock instance.
func NewMockHandshakeRunner(ctrl *gomock.Controller) *MockHandshakeRunner {
	mock := &MockHandshakeRunner{ctrl: ctrl}
	mock.recorder = &MockHandshakeRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandshakeRunner) EXPECT() *MockHandshakeRunnerMockRecorder {
	return m.recorder
}

// DropKeys mocks base method.
func (m *MockHandshakeRunner) DropKeys(arg0 protocol.EncryptionLevel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DropKeys", arg0)
}

// DropKeys indicates an expected call of DropKeys.
func (mr *MockHandshakeRunnerMockRecorder) DropKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropKeys", reflect.TypeOf((*MockHandshakeRunner)(nil).DropKeys), arg0)
}

// OnError mocks base method.
func (m *MockHandshakeRunner) OnError(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnError", arg0)
}

// OnError indicates an expected call of OnError.
func (mr *MockHandshakeRunnerMockRecorder) OnError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnError", reflect.TypeOf((*MockHandshakeRunner)(nil).OnError), arg0)
}

// OnHandshakeComplete mocks base method.
func (m *MockHandshakeRunner) OnHandshakeComplete() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnHandshakeComplete")
}

// OnHandshakeComplete indicates an expected call of OnHandshakeComplete.
func (mr *MockHandshakeRunnerMockRecorder) OnHandshakeComplete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnHandshakeComplete", reflect.TypeOf((*MockHandshakeRunner)(nil).OnHandshakeComplete))
}

// OnReceivedParams mocks base method.
func (m *MockHandshakeRunner) OnReceivedParams(arg0 *wire.TransportParameters) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnReceivedParams", arg0)
}

// OnReceivedParams indicates an expected call of OnReceivedParams.
func (mr *MockHandshakeRunnerMockRecorder) OnReceivedParams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnReceivedParams", reflect.TypeOf((*MockHandshakeRunner)(nil).OnReceivedParams), arg0)
}
