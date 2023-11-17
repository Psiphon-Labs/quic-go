// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Psiphon-Labs/quic-go (interfaces: StreamI)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/Psiphon-Labs/quic-go -destination mock_stream_internal_test.go github.com/Psiphon-Labs/quic-go StreamI
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

// MockStreamI is a mock of StreamI interface.
type MockStreamI struct {
	ctrl     *gomock.Controller
	recorder *MockStreamIMockRecorder
}

// MockStreamIMockRecorder is the mock recorder for MockStreamI.
type MockStreamIMockRecorder struct {
	mock *MockStreamI
}

// NewMockStreamI creates a new mock instance.
func NewMockStreamI(ctrl *gomock.Controller) *MockStreamI {
	mock := &MockStreamI{ctrl: ctrl}
	mock.recorder = &MockStreamIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamI) EXPECT() *MockStreamIMockRecorder {
	return m.recorder
}

// CancelRead mocks base method.
func (m *MockStreamI) CancelRead(arg0 qerr.StreamErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelRead", arg0)
}

// CancelRead indicates an expected call of CancelRead.
func (mr *MockStreamIMockRecorder) CancelRead(arg0 any) *StreamICancelReadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRead", reflect.TypeOf((*MockStreamI)(nil).CancelRead), arg0)
	return &StreamICancelReadCall{Call: call}
}

// StreamICancelReadCall wrap *gomock.Call
type StreamICancelReadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamICancelReadCall) Return() *StreamICancelReadCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamICancelReadCall) Do(f func(qerr.StreamErrorCode)) *StreamICancelReadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamICancelReadCall) DoAndReturn(f func(qerr.StreamErrorCode)) *StreamICancelReadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CancelWrite mocks base method.
func (m *MockStreamI) CancelWrite(arg0 qerr.StreamErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelWrite", arg0)
}

// CancelWrite indicates an expected call of CancelWrite.
func (mr *MockStreamIMockRecorder) CancelWrite(arg0 any) *StreamICancelWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelWrite", reflect.TypeOf((*MockStreamI)(nil).CancelWrite), arg0)
	return &StreamICancelWriteCall{Call: call}
}

// StreamICancelWriteCall wrap *gomock.Call
type StreamICancelWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamICancelWriteCall) Return() *StreamICancelWriteCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamICancelWriteCall) Do(f func(qerr.StreamErrorCode)) *StreamICancelWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamICancelWriteCall) DoAndReturn(f func(qerr.StreamErrorCode)) *StreamICancelWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockStreamI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStreamIMockRecorder) Close() *StreamICloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStreamI)(nil).Close))
	return &StreamICloseCall{Call: call}
}

// StreamICloseCall wrap *gomock.Call
type StreamICloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamICloseCall) Return(arg0 error) *StreamICloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamICloseCall) Do(f func() error) *StreamICloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamICloseCall) DoAndReturn(f func() error) *StreamICloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Context mocks base method.
func (m *MockStreamI) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockStreamIMockRecorder) Context() *StreamIContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockStreamI)(nil).Context))
	return &StreamIContextCall{Call: call}
}

// StreamIContextCall wrap *gomock.Call
type StreamIContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamIContextCall) Return(arg0 context.Context) *StreamIContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamIContextCall) Do(f func() context.Context) *StreamIContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamIContextCall) DoAndReturn(f func() context.Context) *StreamIContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Read mocks base method.
func (m *MockStreamI) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockStreamIMockRecorder) Read(arg0 any) *StreamIReadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockStreamI)(nil).Read), arg0)
	return &StreamIReadCall{Call: call}
}

// StreamIReadCall wrap *gomock.Call
type StreamIReadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamIReadCall) Return(arg0 int, arg1 error) *StreamIReadCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamIReadCall) Do(f func([]byte) (int, error)) *StreamIReadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamIReadCall) DoAndReturn(f func([]byte) (int, error)) *StreamIReadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetDeadline mocks base method.
func (m *MockStreamI) SetDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline.
func (mr *MockStreamIMockRecorder) SetDeadline(arg0 any) *StreamISetDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockStreamI)(nil).SetDeadline), arg0)
	return &StreamISetDeadlineCall{Call: call}
}

// StreamISetDeadlineCall wrap *gomock.Call
type StreamISetDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamISetDeadlineCall) Return(arg0 error) *StreamISetDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamISetDeadlineCall) Do(f func(time.Time) error) *StreamISetDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamISetDeadlineCall) DoAndReturn(f func(time.Time) error) *StreamISetDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetReadDeadline mocks base method.
func (m *MockStreamI) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockStreamIMockRecorder) SetReadDeadline(arg0 any) *StreamISetReadDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockStreamI)(nil).SetReadDeadline), arg0)
	return &StreamISetReadDeadlineCall{Call: call}
}

// StreamISetReadDeadlineCall wrap *gomock.Call
type StreamISetReadDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamISetReadDeadlineCall) Return(arg0 error) *StreamISetReadDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamISetReadDeadlineCall) Do(f func(time.Time) error) *StreamISetReadDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamISetReadDeadlineCall) DoAndReturn(f func(time.Time) error) *StreamISetReadDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetWriteDeadline mocks base method.
func (m *MockStreamI) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockStreamIMockRecorder) SetWriteDeadline(arg0 any) *StreamISetWriteDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockStreamI)(nil).SetWriteDeadline), arg0)
	return &StreamISetWriteDeadlineCall{Call: call}
}

// StreamISetWriteDeadlineCall wrap *gomock.Call
type StreamISetWriteDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamISetWriteDeadlineCall) Return(arg0 error) *StreamISetWriteDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamISetWriteDeadlineCall) Do(f func(time.Time) error) *StreamISetWriteDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamISetWriteDeadlineCall) DoAndReturn(f func(time.Time) error) *StreamISetWriteDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StreamID mocks base method.
func (m *MockStreamI) StreamID() protocol.StreamID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamID")
	ret0, _ := ret[0].(protocol.StreamID)
	return ret0
}

// StreamID indicates an expected call of StreamID.
func (mr *MockStreamIMockRecorder) StreamID() *StreamIStreamIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamID", reflect.TypeOf((*MockStreamI)(nil).StreamID))
	return &StreamIStreamIDCall{Call: call}
}

// StreamIStreamIDCall wrap *gomock.Call
type StreamIStreamIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamIStreamIDCall) Return(arg0 protocol.StreamID) *StreamIStreamIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamIStreamIDCall) Do(f func() protocol.StreamID) *StreamIStreamIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamIStreamIDCall) DoAndReturn(f func() protocol.StreamID) *StreamIStreamIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Write mocks base method.
func (m *MockStreamI) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockStreamIMockRecorder) Write(arg0 any) *StreamIWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStreamI)(nil).Write), arg0)
	return &StreamIWriteCall{Call: call}
}

// StreamIWriteCall wrap *gomock.Call
type StreamIWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamIWriteCall) Return(arg0 int, arg1 error) *StreamIWriteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamIWriteCall) Do(f func([]byte) (int, error)) *StreamIWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamIWriteCall) DoAndReturn(f func([]byte) (int, error)) *StreamIWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// closeForShutdown mocks base method.
func (m *MockStreamI) closeForShutdown(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "closeForShutdown", arg0)
}

// closeForShutdown indicates an expected call of closeForShutdown.
func (mr *MockStreamIMockRecorder) closeForShutdown(arg0 any) *StreamIcloseForShutdownCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeForShutdown", reflect.TypeOf((*MockStreamI)(nil).closeForShutdown), arg0)
	return &StreamIcloseForShutdownCall{Call: call}
}

// StreamIcloseForShutdownCall wrap *gomock.Call
type StreamIcloseForShutdownCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamIcloseForShutdownCall) Return() *StreamIcloseForShutdownCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamIcloseForShutdownCall) Do(f func(error)) *StreamIcloseForShutdownCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamIcloseForShutdownCall) DoAndReturn(f func(error)) *StreamIcloseForShutdownCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// getWindowUpdate mocks base method.
func (m *MockStreamI) getWindowUpdate() protocol.ByteCount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getWindowUpdate")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// getWindowUpdate indicates an expected call of getWindowUpdate.
func (mr *MockStreamIMockRecorder) getWindowUpdate() *StreamIgetWindowUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getWindowUpdate", reflect.TypeOf((*MockStreamI)(nil).getWindowUpdate))
	return &StreamIgetWindowUpdateCall{Call: call}
}

// StreamIgetWindowUpdateCall wrap *gomock.Call
type StreamIgetWindowUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamIgetWindowUpdateCall) Return(arg0 protocol.ByteCount) *StreamIgetWindowUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamIgetWindowUpdateCall) Do(f func() protocol.ByteCount) *StreamIgetWindowUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamIgetWindowUpdateCall) DoAndReturn(f func() protocol.ByteCount) *StreamIgetWindowUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// handleResetStreamFrame mocks base method.
func (m *MockStreamI) handleResetStreamFrame(arg0 *wire.ResetStreamFrame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleResetStreamFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleResetStreamFrame indicates an expected call of handleResetStreamFrame.
func (mr *MockStreamIMockRecorder) handleResetStreamFrame(arg0 any) *StreamIhandleResetStreamFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleResetStreamFrame", reflect.TypeOf((*MockStreamI)(nil).handleResetStreamFrame), arg0)
	return &StreamIhandleResetStreamFrameCall{Call: call}
}

// StreamIhandleResetStreamFrameCall wrap *gomock.Call
type StreamIhandleResetStreamFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamIhandleResetStreamFrameCall) Return(arg0 error) *StreamIhandleResetStreamFrameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamIhandleResetStreamFrameCall) Do(f func(*wire.ResetStreamFrame) error) *StreamIhandleResetStreamFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamIhandleResetStreamFrameCall) DoAndReturn(f func(*wire.ResetStreamFrame) error) *StreamIhandleResetStreamFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// handleStopSendingFrame mocks base method.
func (m *MockStreamI) handleStopSendingFrame(arg0 *wire.StopSendingFrame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handleStopSendingFrame", arg0)
}

// handleStopSendingFrame indicates an expected call of handleStopSendingFrame.
func (mr *MockStreamIMockRecorder) handleStopSendingFrame(arg0 any) *StreamIhandleStopSendingFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleStopSendingFrame", reflect.TypeOf((*MockStreamI)(nil).handleStopSendingFrame), arg0)
	return &StreamIhandleStopSendingFrameCall{Call: call}
}

// StreamIhandleStopSendingFrameCall wrap *gomock.Call
type StreamIhandleStopSendingFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamIhandleStopSendingFrameCall) Return() *StreamIhandleStopSendingFrameCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamIhandleStopSendingFrameCall) Do(f func(*wire.StopSendingFrame)) *StreamIhandleStopSendingFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamIhandleStopSendingFrameCall) DoAndReturn(f func(*wire.StopSendingFrame)) *StreamIhandleStopSendingFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// handleStreamFrame mocks base method.
func (m *MockStreamI) handleStreamFrame(arg0 *wire.StreamFrame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleStreamFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleStreamFrame indicates an expected call of handleStreamFrame.
func (mr *MockStreamIMockRecorder) handleStreamFrame(arg0 any) *StreamIhandleStreamFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleStreamFrame", reflect.TypeOf((*MockStreamI)(nil).handleStreamFrame), arg0)
	return &StreamIhandleStreamFrameCall{Call: call}
}

// StreamIhandleStreamFrameCall wrap *gomock.Call
type StreamIhandleStreamFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamIhandleStreamFrameCall) Return(arg0 error) *StreamIhandleStreamFrameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamIhandleStreamFrameCall) Do(f func(*wire.StreamFrame) error) *StreamIhandleStreamFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamIhandleStreamFrameCall) DoAndReturn(f func(*wire.StreamFrame) error) *StreamIhandleStreamFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// hasData mocks base method.
func (m *MockStreamI) hasData() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "hasData")
	ret0, _ := ret[0].(bool)
	return ret0
}

// hasData indicates an expected call of hasData.
func (mr *MockStreamIMockRecorder) hasData() *StreamIhasDataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "hasData", reflect.TypeOf((*MockStreamI)(nil).hasData))
	return &StreamIhasDataCall{Call: call}
}

// StreamIhasDataCall wrap *gomock.Call
type StreamIhasDataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamIhasDataCall) Return(arg0 bool) *StreamIhasDataCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamIhasDataCall) Do(f func() bool) *StreamIhasDataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamIhasDataCall) DoAndReturn(f func() bool) *StreamIhasDataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// popStreamFrame mocks base method.
func (m *MockStreamI) popStreamFrame(arg0 protocol.ByteCount, arg1 protocol.VersionNumber) (ackhandler.StreamFrame, bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "popStreamFrame", arg0, arg1)
	ret0, _ := ret[0].(ackhandler.StreamFrame)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// popStreamFrame indicates an expected call of popStreamFrame.
func (mr *MockStreamIMockRecorder) popStreamFrame(arg0, arg1 any) *StreamIpopStreamFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "popStreamFrame", reflect.TypeOf((*MockStreamI)(nil).popStreamFrame), arg0, arg1)
	return &StreamIpopStreamFrameCall{Call: call}
}

// StreamIpopStreamFrameCall wrap *gomock.Call
type StreamIpopStreamFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamIpopStreamFrameCall) Return(arg0 ackhandler.StreamFrame, arg1, arg2 bool) *StreamIpopStreamFrameCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamIpopStreamFrameCall) Do(f func(protocol.ByteCount, protocol.VersionNumber) (ackhandler.StreamFrame, bool, bool)) *StreamIpopStreamFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamIpopStreamFrameCall) DoAndReturn(f func(protocol.ByteCount, protocol.VersionNumber) (ackhandler.StreamFrame, bool, bool)) *StreamIpopStreamFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// updateSendWindow mocks base method.
func (m *MockStreamI) updateSendWindow(arg0 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "updateSendWindow", arg0)
}

// updateSendWindow indicates an expected call of updateSendWindow.
func (mr *MockStreamIMockRecorder) updateSendWindow(arg0 any) *StreamIupdateSendWindowCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateSendWindow", reflect.TypeOf((*MockStreamI)(nil).updateSendWindow), arg0)
	return &StreamIupdateSendWindowCall{Call: call}
}

// StreamIupdateSendWindowCall wrap *gomock.Call
type StreamIupdateSendWindowCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamIupdateSendWindowCall) Return() *StreamIupdateSendWindowCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamIupdateSendWindowCall) Do(f func(protocol.ByteCount)) *StreamIupdateSendWindowCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamIupdateSendWindowCall) DoAndReturn(f func(protocol.ByteCount)) *StreamIupdateSendWindowCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
