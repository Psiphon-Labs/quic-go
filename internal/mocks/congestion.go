// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go/internal/congestion (interfaces: SendAlgorithmWithDebugInfos)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package mocks -destination congestion.go github.com/quic-go/quic-go/internal/congestion SendAlgorithmWithDebugInfos
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	protocol "github.com/Psiphon-Labs/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockSendAlgorithmWithDebugInfos is a mock of SendAlgorithmWithDebugInfos interface.
type MockSendAlgorithmWithDebugInfos struct {
	ctrl     *gomock.Controller
	recorder *MockSendAlgorithmWithDebugInfosMockRecorder
}

// MockSendAlgorithmWithDebugInfosMockRecorder is the mock recorder for MockSendAlgorithmWithDebugInfos.
type MockSendAlgorithmWithDebugInfosMockRecorder struct {
	mock *MockSendAlgorithmWithDebugInfos
}

// NewMockSendAlgorithmWithDebugInfos creates a new mock instance.
func NewMockSendAlgorithmWithDebugInfos(ctrl *gomock.Controller) *MockSendAlgorithmWithDebugInfos {
	mock := &MockSendAlgorithmWithDebugInfos{ctrl: ctrl}
	mock.recorder = &MockSendAlgorithmWithDebugInfosMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSendAlgorithmWithDebugInfos) EXPECT() *MockSendAlgorithmWithDebugInfosMockRecorder {
	return m.recorder
}

// CanSend mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) CanSend(arg0 protocol.ByteCount) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSend", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanSend indicates an expected call of CanSend.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) CanSend(arg0 any) *SendAlgorithmWithDebugInfosCanSendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSend", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).CanSend), arg0)
	return &SendAlgorithmWithDebugInfosCanSendCall{Call: call}
}

// SendAlgorithmWithDebugInfosCanSendCall wrap *gomock.Call
type SendAlgorithmWithDebugInfosCanSendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendAlgorithmWithDebugInfosCanSendCall) Return(arg0 bool) *SendAlgorithmWithDebugInfosCanSendCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendAlgorithmWithDebugInfosCanSendCall) Do(f func(protocol.ByteCount) bool) *SendAlgorithmWithDebugInfosCanSendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendAlgorithmWithDebugInfosCanSendCall) DoAndReturn(f func(protocol.ByteCount) bool) *SendAlgorithmWithDebugInfosCanSendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCongestionWindow mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) GetCongestionWindow() protocol.ByteCount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCongestionWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetCongestionWindow indicates an expected call of GetCongestionWindow.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) GetCongestionWindow() *SendAlgorithmWithDebugInfosGetCongestionWindowCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCongestionWindow", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).GetCongestionWindow))
	return &SendAlgorithmWithDebugInfosGetCongestionWindowCall{Call: call}
}

// SendAlgorithmWithDebugInfosGetCongestionWindowCall wrap *gomock.Call
type SendAlgorithmWithDebugInfosGetCongestionWindowCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendAlgorithmWithDebugInfosGetCongestionWindowCall) Return(arg0 protocol.ByteCount) *SendAlgorithmWithDebugInfosGetCongestionWindowCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendAlgorithmWithDebugInfosGetCongestionWindowCall) Do(f func() protocol.ByteCount) *SendAlgorithmWithDebugInfosGetCongestionWindowCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendAlgorithmWithDebugInfosGetCongestionWindowCall) DoAndReturn(f func() protocol.ByteCount) *SendAlgorithmWithDebugInfosGetCongestionWindowCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HasPacingBudget mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) HasPacingBudget(arg0 time.Time) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPacingBudget", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPacingBudget indicates an expected call of HasPacingBudget.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) HasPacingBudget(arg0 any) *SendAlgorithmWithDebugInfosHasPacingBudgetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPacingBudget", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).HasPacingBudget), arg0)
	return &SendAlgorithmWithDebugInfosHasPacingBudgetCall{Call: call}
}

// SendAlgorithmWithDebugInfosHasPacingBudgetCall wrap *gomock.Call
type SendAlgorithmWithDebugInfosHasPacingBudgetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendAlgorithmWithDebugInfosHasPacingBudgetCall) Return(arg0 bool) *SendAlgorithmWithDebugInfosHasPacingBudgetCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendAlgorithmWithDebugInfosHasPacingBudgetCall) Do(f func(time.Time) bool) *SendAlgorithmWithDebugInfosHasPacingBudgetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendAlgorithmWithDebugInfosHasPacingBudgetCall) DoAndReturn(f func(time.Time) bool) *SendAlgorithmWithDebugInfosHasPacingBudgetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InRecovery mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) InRecovery() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InRecovery")
	ret0, _ := ret[0].(bool)
	return ret0
}

// InRecovery indicates an expected call of InRecovery.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) InRecovery() *SendAlgorithmWithDebugInfosInRecoveryCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InRecovery", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).InRecovery))
	return &SendAlgorithmWithDebugInfosInRecoveryCall{Call: call}
}

// SendAlgorithmWithDebugInfosInRecoveryCall wrap *gomock.Call
type SendAlgorithmWithDebugInfosInRecoveryCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendAlgorithmWithDebugInfosInRecoveryCall) Return(arg0 bool) *SendAlgorithmWithDebugInfosInRecoveryCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendAlgorithmWithDebugInfosInRecoveryCall) Do(f func() bool) *SendAlgorithmWithDebugInfosInRecoveryCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendAlgorithmWithDebugInfosInRecoveryCall) DoAndReturn(f func() bool) *SendAlgorithmWithDebugInfosInRecoveryCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InSlowStart mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) InSlowStart() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InSlowStart")
	ret0, _ := ret[0].(bool)
	return ret0
}

// InSlowStart indicates an expected call of InSlowStart.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) InSlowStart() *SendAlgorithmWithDebugInfosInSlowStartCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InSlowStart", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).InSlowStart))
	return &SendAlgorithmWithDebugInfosInSlowStartCall{Call: call}
}

// SendAlgorithmWithDebugInfosInSlowStartCall wrap *gomock.Call
type SendAlgorithmWithDebugInfosInSlowStartCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendAlgorithmWithDebugInfosInSlowStartCall) Return(arg0 bool) *SendAlgorithmWithDebugInfosInSlowStartCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendAlgorithmWithDebugInfosInSlowStartCall) Do(f func() bool) *SendAlgorithmWithDebugInfosInSlowStartCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendAlgorithmWithDebugInfosInSlowStartCall) DoAndReturn(f func() bool) *SendAlgorithmWithDebugInfosInSlowStartCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MaybeExitSlowStart mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) MaybeExitSlowStart() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MaybeExitSlowStart")
}

// MaybeExitSlowStart indicates an expected call of MaybeExitSlowStart.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) MaybeExitSlowStart() *SendAlgorithmWithDebugInfosMaybeExitSlowStartCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeExitSlowStart", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).MaybeExitSlowStart))
	return &SendAlgorithmWithDebugInfosMaybeExitSlowStartCall{Call: call}
}

// SendAlgorithmWithDebugInfosMaybeExitSlowStartCall wrap *gomock.Call
type SendAlgorithmWithDebugInfosMaybeExitSlowStartCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendAlgorithmWithDebugInfosMaybeExitSlowStartCall) Return() *SendAlgorithmWithDebugInfosMaybeExitSlowStartCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendAlgorithmWithDebugInfosMaybeExitSlowStartCall) Do(f func()) *SendAlgorithmWithDebugInfosMaybeExitSlowStartCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendAlgorithmWithDebugInfosMaybeExitSlowStartCall) DoAndReturn(f func()) *SendAlgorithmWithDebugInfosMaybeExitSlowStartCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OnCongestionEvent mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) OnCongestionEvent(arg0 protocol.PacketNumber, arg1, arg2 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnCongestionEvent", arg0, arg1, arg2)
}

// OnCongestionEvent indicates an expected call of OnCongestionEvent.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) OnCongestionEvent(arg0, arg1, arg2 any) *SendAlgorithmWithDebugInfosOnCongestionEventCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnCongestionEvent", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).OnCongestionEvent), arg0, arg1, arg2)
	return &SendAlgorithmWithDebugInfosOnCongestionEventCall{Call: call}
}

// SendAlgorithmWithDebugInfosOnCongestionEventCall wrap *gomock.Call
type SendAlgorithmWithDebugInfosOnCongestionEventCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendAlgorithmWithDebugInfosOnCongestionEventCall) Return() *SendAlgorithmWithDebugInfosOnCongestionEventCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendAlgorithmWithDebugInfosOnCongestionEventCall) Do(f func(protocol.PacketNumber, protocol.ByteCount, protocol.ByteCount)) *SendAlgorithmWithDebugInfosOnCongestionEventCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendAlgorithmWithDebugInfosOnCongestionEventCall) DoAndReturn(f func(protocol.PacketNumber, protocol.ByteCount, protocol.ByteCount)) *SendAlgorithmWithDebugInfosOnCongestionEventCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OnPacketAcked mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) OnPacketAcked(arg0 protocol.PacketNumber, arg1, arg2 protocol.ByteCount, arg3 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPacketAcked", arg0, arg1, arg2, arg3)
}

// OnPacketAcked indicates an expected call of OnPacketAcked.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) OnPacketAcked(arg0, arg1, arg2, arg3 any) *SendAlgorithmWithDebugInfosOnPacketAckedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPacketAcked", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).OnPacketAcked), arg0, arg1, arg2, arg3)
	return &SendAlgorithmWithDebugInfosOnPacketAckedCall{Call: call}
}

// SendAlgorithmWithDebugInfosOnPacketAckedCall wrap *gomock.Call
type SendAlgorithmWithDebugInfosOnPacketAckedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendAlgorithmWithDebugInfosOnPacketAckedCall) Return() *SendAlgorithmWithDebugInfosOnPacketAckedCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendAlgorithmWithDebugInfosOnPacketAckedCall) Do(f func(protocol.PacketNumber, protocol.ByteCount, protocol.ByteCount, time.Time)) *SendAlgorithmWithDebugInfosOnPacketAckedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendAlgorithmWithDebugInfosOnPacketAckedCall) DoAndReturn(f func(protocol.PacketNumber, protocol.ByteCount, protocol.ByteCount, time.Time)) *SendAlgorithmWithDebugInfosOnPacketAckedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OnPacketSent mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) OnPacketSent(arg0 time.Time, arg1 protocol.ByteCount, arg2 protocol.PacketNumber, arg3 protocol.ByteCount, arg4 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPacketSent", arg0, arg1, arg2, arg3, arg4)
}

// OnPacketSent indicates an expected call of OnPacketSent.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) OnPacketSent(arg0, arg1, arg2, arg3, arg4 any) *SendAlgorithmWithDebugInfosOnPacketSentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPacketSent", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).OnPacketSent), arg0, arg1, arg2, arg3, arg4)
	return &SendAlgorithmWithDebugInfosOnPacketSentCall{Call: call}
}

// SendAlgorithmWithDebugInfosOnPacketSentCall wrap *gomock.Call
type SendAlgorithmWithDebugInfosOnPacketSentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendAlgorithmWithDebugInfosOnPacketSentCall) Return() *SendAlgorithmWithDebugInfosOnPacketSentCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendAlgorithmWithDebugInfosOnPacketSentCall) Do(f func(time.Time, protocol.ByteCount, protocol.PacketNumber, protocol.ByteCount, bool)) *SendAlgorithmWithDebugInfosOnPacketSentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendAlgorithmWithDebugInfosOnPacketSentCall) DoAndReturn(f func(time.Time, protocol.ByteCount, protocol.PacketNumber, protocol.ByteCount, bool)) *SendAlgorithmWithDebugInfosOnPacketSentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OnRetransmissionTimeout mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) OnRetransmissionTimeout(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnRetransmissionTimeout", arg0)
}

// OnRetransmissionTimeout indicates an expected call of OnRetransmissionTimeout.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) OnRetransmissionTimeout(arg0 any) *SendAlgorithmWithDebugInfosOnRetransmissionTimeoutCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnRetransmissionTimeout", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).OnRetransmissionTimeout), arg0)
	return &SendAlgorithmWithDebugInfosOnRetransmissionTimeoutCall{Call: call}
}

// SendAlgorithmWithDebugInfosOnRetransmissionTimeoutCall wrap *gomock.Call
type SendAlgorithmWithDebugInfosOnRetransmissionTimeoutCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendAlgorithmWithDebugInfosOnRetransmissionTimeoutCall) Return() *SendAlgorithmWithDebugInfosOnRetransmissionTimeoutCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendAlgorithmWithDebugInfosOnRetransmissionTimeoutCall) Do(f func(bool)) *SendAlgorithmWithDebugInfosOnRetransmissionTimeoutCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendAlgorithmWithDebugInfosOnRetransmissionTimeoutCall) DoAndReturn(f func(bool)) *SendAlgorithmWithDebugInfosOnRetransmissionTimeoutCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetMaxDatagramSize mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) SetMaxDatagramSize(arg0 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxDatagramSize", arg0)
}

// SetMaxDatagramSize indicates an expected call of SetMaxDatagramSize.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) SetMaxDatagramSize(arg0 any) *SendAlgorithmWithDebugInfosSetMaxDatagramSizeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxDatagramSize", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).SetMaxDatagramSize), arg0)
	return &SendAlgorithmWithDebugInfosSetMaxDatagramSizeCall{Call: call}
}

// SendAlgorithmWithDebugInfosSetMaxDatagramSizeCall wrap *gomock.Call
type SendAlgorithmWithDebugInfosSetMaxDatagramSizeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendAlgorithmWithDebugInfosSetMaxDatagramSizeCall) Return() *SendAlgorithmWithDebugInfosSetMaxDatagramSizeCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendAlgorithmWithDebugInfosSetMaxDatagramSizeCall) Do(f func(protocol.ByteCount)) *SendAlgorithmWithDebugInfosSetMaxDatagramSizeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendAlgorithmWithDebugInfosSetMaxDatagramSizeCall) DoAndReturn(f func(protocol.ByteCount)) *SendAlgorithmWithDebugInfosSetMaxDatagramSizeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// TimeUntilSend mocks base method.
func (m *MockSendAlgorithmWithDebugInfos) TimeUntilSend(arg0 protocol.ByteCount) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeUntilSend", arg0)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// TimeUntilSend indicates an expected call of TimeUntilSend.
func (mr *MockSendAlgorithmWithDebugInfosMockRecorder) TimeUntilSend(arg0 any) *SendAlgorithmWithDebugInfosTimeUntilSendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeUntilSend", reflect.TypeOf((*MockSendAlgorithmWithDebugInfos)(nil).TimeUntilSend), arg0)
	return &SendAlgorithmWithDebugInfosTimeUntilSendCall{Call: call}
}

// SendAlgorithmWithDebugInfosTimeUntilSendCall wrap *gomock.Call
type SendAlgorithmWithDebugInfosTimeUntilSendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendAlgorithmWithDebugInfosTimeUntilSendCall) Return(arg0 time.Time) *SendAlgorithmWithDebugInfosTimeUntilSendCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendAlgorithmWithDebugInfosTimeUntilSendCall) Do(f func(protocol.ByteCount) time.Time) *SendAlgorithmWithDebugInfosTimeUntilSendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendAlgorithmWithDebugInfosTimeUntilSendCall) DoAndReturn(f func(protocol.ByteCount) time.Time) *SendAlgorithmWithDebugInfosTimeUntilSendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
