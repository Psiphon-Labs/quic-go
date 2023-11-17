// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Psiphon-Labs/quic-go/http3 (interfaces: RoundTripCloser)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package http3 -destination mock_roundtripcloser_test.go github.com/Psiphon-Labs/quic-go/http3 RoundTripCloser
//
// Package http3 is a generated GoMock package.
package http3

import (
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRoundTripCloser is a mock of RoundTripCloser interface.
type MockRoundTripCloser struct {
	ctrl     *gomock.Controller
	recorder *MockRoundTripCloserMockRecorder
}

// MockRoundTripCloserMockRecorder is the mock recorder for MockRoundTripCloser.
type MockRoundTripCloserMockRecorder struct {
	mock *MockRoundTripCloser
}

// NewMockRoundTripCloser creates a new mock instance.
func NewMockRoundTripCloser(ctrl *gomock.Controller) *MockRoundTripCloser {
	mock := &MockRoundTripCloser{ctrl: ctrl}
	mock.recorder = &MockRoundTripCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoundTripCloser) EXPECT() *MockRoundTripCloserMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRoundTripCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRoundTripCloserMockRecorder) Close() *RoundTripCloserCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRoundTripCloser)(nil).Close))
	return &RoundTripCloserCloseCall{Call: call}
}

// RoundTripCloserCloseCall wrap *gomock.Call
type RoundTripCloserCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RoundTripCloserCloseCall) Return(arg0 error) *RoundTripCloserCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RoundTripCloserCloseCall) Do(f func() error) *RoundTripCloserCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RoundTripCloserCloseCall) DoAndReturn(f func() error) *RoundTripCloserCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HandshakeComplete mocks base method.
func (m *MockRoundTripCloser) HandshakeComplete() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandshakeComplete")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HandshakeComplete indicates an expected call of HandshakeComplete.
func (mr *MockRoundTripCloserMockRecorder) HandshakeComplete() *RoundTripCloserHandshakeCompleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandshakeComplete", reflect.TypeOf((*MockRoundTripCloser)(nil).HandshakeComplete))
	return &RoundTripCloserHandshakeCompleteCall{Call: call}
}

// RoundTripCloserHandshakeCompleteCall wrap *gomock.Call
type RoundTripCloserHandshakeCompleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RoundTripCloserHandshakeCompleteCall) Return(arg0 bool) *RoundTripCloserHandshakeCompleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RoundTripCloserHandshakeCompleteCall) Do(f func() bool) *RoundTripCloserHandshakeCompleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RoundTripCloserHandshakeCompleteCall) DoAndReturn(f func() bool) *RoundTripCloserHandshakeCompleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RoundTripOpt mocks base method.
func (m *MockRoundTripCloser) RoundTripOpt(arg0 *http.Request, arg1 RoundTripOpt) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoundTripOpt", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoundTripOpt indicates an expected call of RoundTripOpt.
func (mr *MockRoundTripCloserMockRecorder) RoundTripOpt(arg0, arg1 any) *RoundTripCloserRoundTripOptCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoundTripOpt", reflect.TypeOf((*MockRoundTripCloser)(nil).RoundTripOpt), arg0, arg1)
	return &RoundTripCloserRoundTripOptCall{Call: call}
}

// RoundTripCloserRoundTripOptCall wrap *gomock.Call
type RoundTripCloserRoundTripOptCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RoundTripCloserRoundTripOptCall) Return(arg0 *http.Response, arg1 error) *RoundTripCloserRoundTripOptCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RoundTripCloserRoundTripOptCall) Do(f func(*http.Request, RoundTripOpt) (*http.Response, error)) *RoundTripCloserRoundTripOptCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RoundTripCloserRoundTripOptCall) DoAndReturn(f func(*http.Request, RoundTripOpt) (*http.Response, error)) *RoundTripCloserRoundTripOptCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
