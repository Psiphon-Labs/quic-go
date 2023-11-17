// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Psiphon-Labs/quic-go (interfaces: SendStreamI)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/Psiphon-Labs/quic-go -destination mock_send_stream_internal_test.go github.com/Psiphon-Labs/quic-go SendStreamI
//
// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	reflect "reflect"
	time "time"

	ackhandler "github.com/Psiphon-Labs/quic-go/internal/ackhandler"
	protocol "github.com/Psiphon-Labs/quic-go/internal/protocol"
	qerr "github.com/Psiphon-Labs/quic-go/internal/qerr"
	wire "github.com/Psiphon-Labs/quic-go/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockSendStreamI is a mock of SendStreamI interface.
type MockSendStreamI struct {
	ctrl     *gomock.Controller
	recorder *MockSendStreamIMockRecorder
}

// MockSendStreamIMockRecorder is the mock recorder for MockSendStreamI.
type MockSendStreamIMockRecorder struct {
	mock *MockSendStreamI
}

// NewMockSendStreamI creates a new mock instance.
func NewMockSendStreamI(ctrl *gomock.Controller) *MockSendStreamI {
	mock := &MockSendStreamI{ctrl: ctrl}
	mock.recorder = &MockSendStreamIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSendStreamI) EXPECT() *MockSendStreamIMockRecorder {
	return m.recorder
}

// CancelWrite mocks base method.
func (m *MockSendStreamI) CancelWrite(arg0 qerr.StreamErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelWrite", arg0)
}

// CancelWrite indicates an expected call of CancelWrite.
func (mr *MockSendStreamIMockRecorder) CancelWrite(arg0 any) *SendStreamICancelWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelWrite", reflect.TypeOf((*MockSendStreamI)(nil).CancelWrite), arg0)
	return &SendStreamICancelWriteCall{Call: call}
}

// SendStreamICancelWriteCall wrap *gomock.Call
type SendStreamICancelWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendStreamICancelWriteCall) Return() *SendStreamICancelWriteCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendStreamICancelWriteCall) Do(f func(qerr.StreamErrorCode)) *SendStreamICancelWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendStreamICancelWriteCall) DoAndReturn(f func(qerr.StreamErrorCode)) *SendStreamICancelWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockSendStreamI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSendStreamIMockRecorder) Close() *SendStreamICloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSendStreamI)(nil).Close))
	return &SendStreamICloseCall{Call: call}
}

// SendStreamICloseCall wrap *gomock.Call
type SendStreamICloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendStreamICloseCall) Return(arg0 error) *SendStreamICloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendStreamICloseCall) Do(f func() error) *SendStreamICloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendStreamICloseCall) DoAndReturn(f func() error) *SendStreamICloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Context mocks base method.
func (m *MockSendStreamI) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockSendStreamIMockRecorder) Context() *SendStreamIContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSendStreamI)(nil).Context))
	return &SendStreamIContextCall{Call: call}
}

// SendStreamIContextCall wrap *gomock.Call
type SendStreamIContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendStreamIContextCall) Return(arg0 context.Context) *SendStreamIContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendStreamIContextCall) Do(f func() context.Context) *SendStreamIContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendStreamIContextCall) DoAndReturn(f func() context.Context) *SendStreamIContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetWriteDeadline mocks base method.
func (m *MockSendStreamI) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockSendStreamIMockRecorder) SetWriteDeadline(arg0 any) *SendStreamISetWriteDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockSendStreamI)(nil).SetWriteDeadline), arg0)
	return &SendStreamISetWriteDeadlineCall{Call: call}
}

// SendStreamISetWriteDeadlineCall wrap *gomock.Call
type SendStreamISetWriteDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendStreamISetWriteDeadlineCall) Return(arg0 error) *SendStreamISetWriteDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendStreamISetWriteDeadlineCall) Do(f func(time.Time) error) *SendStreamISetWriteDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendStreamISetWriteDeadlineCall) DoAndReturn(f func(time.Time) error) *SendStreamISetWriteDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StreamID mocks base method.
func (m *MockSendStreamI) StreamID() protocol.StreamID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamID")
	ret0, _ := ret[0].(protocol.StreamID)
	return ret0
}

// StreamID indicates an expected call of StreamID.
func (mr *MockSendStreamIMockRecorder) StreamID() *SendStreamIStreamIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamID", reflect.TypeOf((*MockSendStreamI)(nil).StreamID))
	return &SendStreamIStreamIDCall{Call: call}
}

// SendStreamIStreamIDCall wrap *gomock.Call
type SendStreamIStreamIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendStreamIStreamIDCall) Return(arg0 protocol.StreamID) *SendStreamIStreamIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendStreamIStreamIDCall) Do(f func() protocol.StreamID) *SendStreamIStreamIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendStreamIStreamIDCall) DoAndReturn(f func() protocol.StreamID) *SendStreamIStreamIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Write mocks base method.
func (m *MockSendStreamI) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockSendStreamIMockRecorder) Write(arg0 any) *SendStreamIWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSendStreamI)(nil).Write), arg0)
	return &SendStreamIWriteCall{Call: call}
}

// SendStreamIWriteCall wrap *gomock.Call
type SendStreamIWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendStreamIWriteCall) Return(arg0 int, arg1 error) *SendStreamIWriteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendStreamIWriteCall) Do(f func([]byte) (int, error)) *SendStreamIWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendStreamIWriteCall) DoAndReturn(f func([]byte) (int, error)) *SendStreamIWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// closeForShutdown mocks base method.
func (m *MockSendStreamI) closeForShutdown(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "closeForShutdown", arg0)
}

// closeForShutdown indicates an expected call of closeForShutdown.
func (mr *MockSendStreamIMockRecorder) closeForShutdown(arg0 any) *SendStreamIcloseForShutdownCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeForShutdown", reflect.TypeOf((*MockSendStreamI)(nil).closeForShutdown), arg0)
	return &SendStreamIcloseForShutdownCall{Call: call}
}

// SendStreamIcloseForShutdownCall wrap *gomock.Call
type SendStreamIcloseForShutdownCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendStreamIcloseForShutdownCall) Return() *SendStreamIcloseForShutdownCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendStreamIcloseForShutdownCall) Do(f func(error)) *SendStreamIcloseForShutdownCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendStreamIcloseForShutdownCall) DoAndReturn(f func(error)) *SendStreamIcloseForShutdownCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// handleStopSendingFrame mocks base method.
func (m *MockSendStreamI) handleStopSendingFrame(arg0 *wire.StopSendingFrame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handleStopSendingFrame", arg0)
}

// handleStopSendingFrame indicates an expected call of handleStopSendingFrame.
func (mr *MockSendStreamIMockRecorder) handleStopSendingFrame(arg0 any) *SendStreamIhandleStopSendingFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleStopSendingFrame", reflect.TypeOf((*MockSendStreamI)(nil).handleStopSendingFrame), arg0)
	return &SendStreamIhandleStopSendingFrameCall{Call: call}
}

// SendStreamIhandleStopSendingFrameCall wrap *gomock.Call
type SendStreamIhandleStopSendingFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendStreamIhandleStopSendingFrameCall) Return() *SendStreamIhandleStopSendingFrameCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendStreamIhandleStopSendingFrameCall) Do(f func(*wire.StopSendingFrame)) *SendStreamIhandleStopSendingFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendStreamIhandleStopSendingFrameCall) DoAndReturn(f func(*wire.StopSendingFrame)) *SendStreamIhandleStopSendingFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// hasData mocks base method.
func (m *MockSendStreamI) hasData() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "hasData")
	ret0, _ := ret[0].(bool)
	return ret0
}

// hasData indicates an expected call of hasData.
func (mr *MockSendStreamIMockRecorder) hasData() *SendStreamIhasDataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "hasData", reflect.TypeOf((*MockSendStreamI)(nil).hasData))
	return &SendStreamIhasDataCall{Call: call}
}

// SendStreamIhasDataCall wrap *gomock.Call
type SendStreamIhasDataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendStreamIhasDataCall) Return(arg0 bool) *SendStreamIhasDataCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendStreamIhasDataCall) Do(f func() bool) *SendStreamIhasDataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendStreamIhasDataCall) DoAndReturn(f func() bool) *SendStreamIhasDataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// popStreamFrame mocks base method.
func (m *MockSendStreamI) popStreamFrame(arg0 protocol.ByteCount, arg1 protocol.VersionNumber) (ackhandler.StreamFrame, bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "popStreamFrame", arg0, arg1)
	ret0, _ := ret[0].(ackhandler.StreamFrame)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// popStreamFrame indicates an expected call of popStreamFrame.
func (mr *MockSendStreamIMockRecorder) popStreamFrame(arg0, arg1 any) *SendStreamIpopStreamFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "popStreamFrame", reflect.TypeOf((*MockSendStreamI)(nil).popStreamFrame), arg0, arg1)
	return &SendStreamIpopStreamFrameCall{Call: call}
}

// SendStreamIpopStreamFrameCall wrap *gomock.Call
type SendStreamIpopStreamFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendStreamIpopStreamFrameCall) Return(arg0 ackhandler.StreamFrame, arg1, arg2 bool) *SendStreamIpopStreamFrameCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendStreamIpopStreamFrameCall) Do(f func(protocol.ByteCount, protocol.VersionNumber) (ackhandler.StreamFrame, bool, bool)) *SendStreamIpopStreamFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendStreamIpopStreamFrameCall) DoAndReturn(f func(protocol.ByteCount, protocol.VersionNumber) (ackhandler.StreamFrame, bool, bool)) *SendStreamIpopStreamFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// updateSendWindow mocks base method.
func (m *MockSendStreamI) updateSendWindow(arg0 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "updateSendWindow", arg0)
}

// updateSendWindow indicates an expected call of updateSendWindow.
func (mr *MockSendStreamIMockRecorder) updateSendWindow(arg0 any) *SendStreamIupdateSendWindowCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateSendWindow", reflect.TypeOf((*MockSendStreamI)(nil).updateSendWindow), arg0)
	return &SendStreamIupdateSendWindowCall{Call: call}
}

// SendStreamIupdateSendWindowCall wrap *gomock.Call
type SendStreamIupdateSendWindowCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SendStreamIupdateSendWindowCall) Return() *SendStreamIupdateSendWindowCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SendStreamIupdateSendWindowCall) Do(f func(protocol.ByteCount)) *SendStreamIupdateSendWindowCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SendStreamIupdateSendWindowCall) DoAndReturn(f func(protocol.ByteCount)) *SendStreamIupdateSendWindowCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
