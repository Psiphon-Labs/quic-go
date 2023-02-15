// Code generated by MockGen. DO NOT EDIT.
// Source: connection.go

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"

	protocol "github.com/Psiphon-Labs/quic-go/internal/protocol"
	wire "github.com/Psiphon-Labs/quic-go/internal/wire"
)

// MockUnpacker is a mock of Unpacker interface.
type MockUnpacker struct {
	ctrl     *gomock.Controller
	recorder *MockUnpackerMockRecorder
}

// MockUnpackerMockRecorder is the mock recorder for MockUnpacker.
type MockUnpackerMockRecorder struct {
	mock *MockUnpacker
}

// NewMockUnpacker creates a new mock instance.
func NewMockUnpacker(ctrl *gomock.Controller) *MockUnpacker {
	mock := &MockUnpacker{ctrl: ctrl}
	mock.recorder = &MockUnpackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnpacker) EXPECT() *MockUnpackerMockRecorder {
	return m.recorder
}

// UnpackLongHeader mocks base method.
func (m *MockUnpacker) UnpackLongHeader(hdr *wire.Header, rcvTime time.Time, data []byte, v protocol.VersionNumber) (*unpackedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpackLongHeader", hdr, rcvTime, data, v)
	ret0, _ := ret[0].(*unpackedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpackLongHeader indicates an expected call of UnpackLongHeader.
func (mr *MockUnpackerMockRecorder) UnpackLongHeader(hdr, rcvTime, data, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpackLongHeader", reflect.TypeOf((*MockUnpacker)(nil).UnpackLongHeader), hdr, rcvTime, data, v)
}

// UnpackShortHeader mocks base method.
func (m *MockUnpacker) UnpackShortHeader(rcvTime time.Time, data []byte) (protocol.PacketNumber, protocol.PacketNumberLen, protocol.KeyPhaseBit, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpackShortHeader", rcvTime, data)
	ret0, _ := ret[0].(protocol.PacketNumber)
	ret1, _ := ret[1].(protocol.PacketNumberLen)
	ret2, _ := ret[2].(protocol.KeyPhaseBit)
	ret3, _ := ret[3].([]byte)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// UnpackShortHeader indicates an expected call of UnpackShortHeader.
func (mr *MockUnpackerMockRecorder) UnpackShortHeader(rcvTime, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpackShortHeader", reflect.TypeOf((*MockUnpacker)(nil).UnpackShortHeader), rcvTime, data)
}
