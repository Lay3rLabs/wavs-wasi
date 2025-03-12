// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package tcp represents the imported interface "wasi:sockets/tcp@0.2.0".
package tcp

import (
	monotonicclock "github.com/Lay3rLabs/wavs-wasi/go/wasi/clocks/monotonic-clock"
	"github.com/Lay3rLabs/wavs-wasi/go/wasi/io/poll"
	"github.com/Lay3rLabs/wavs-wasi/go/wasi/io/streams"
	"github.com/Lay3rLabs/wavs-wasi/go/wasi/sockets/network"
	"go.bytecodealliance.org/cm"
)

// InputStream represents the imported type alias "wasi:sockets/tcp@0.2.0#input-stream".
//
// See [streams.InputStream] for more information.
type InputStream = streams.InputStream

// OutputStream represents the imported type alias "wasi:sockets/tcp@0.2.0#output-stream".
//
// See [streams.OutputStream] for more information.
type OutputStream = streams.OutputStream

// Pollable represents the imported type alias "wasi:sockets/tcp@0.2.0#pollable".
//
// See [poll.Pollable] for more information.
type Pollable = poll.Pollable

// Duration represents the type alias "wasi:sockets/tcp@0.2.0#duration".
//
// See [monotonicclock.Duration] for more information.
type Duration = monotonicclock.Duration

// Network represents the imported type alias "wasi:sockets/tcp@0.2.0#network".
//
// See [network.Network] for more information.
type Network = network.Network

// ErrorCode represents the type alias "wasi:sockets/tcp@0.2.0#error-code".
//
// See [network.ErrorCode] for more information.
type ErrorCode = network.ErrorCode

// IPSocketAddress represents the type alias "wasi:sockets/tcp@0.2.0#ip-socket-address".
//
// See [network.IPSocketAddress] for more information.
type IPSocketAddress = network.IPSocketAddress

// IPAddressFamily represents the type alias "wasi:sockets/tcp@0.2.0#ip-address-family".
//
// See [network.IPAddressFamily] for more information.
type IPAddressFamily = network.IPAddressFamily

// ShutdownType represents the enum "wasi:sockets/tcp@0.2.0#shutdown-type".
//
//	enum shutdown-type {
//		receive,
//		send,
//		both
//	}
type ShutdownType uint8

const (
	ShutdownTypeReceive ShutdownType = iota
	ShutdownTypeSend
	ShutdownTypeBoth
)

var _ShutdownTypeStrings = [3]string{
	"receive",
	"send",
	"both",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e ShutdownType) String() string {
	return _ShutdownTypeStrings[e]
}

// MarshalText implements [encoding.TextMarshaler].
func (e ShutdownType) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler], unmarshaling into an enum
// case. Returns an error if the supplied text is not one of the enum cases.
func (e *ShutdownType) UnmarshalText(text []byte) error {
	return _ShutdownTypeUnmarshalCase(e, text)
}

var _ShutdownTypeUnmarshalCase = cm.CaseUnmarshaler[ShutdownType](_ShutdownTypeStrings[:])

// TCPSocket represents the imported resource "wasi:sockets/tcp@0.2.0#tcp-socket".
//
//	resource tcp-socket
type TCPSocket cm.Resource

// ResourceDrop represents the imported resource-drop for resource "tcp-socket".
//
// Drops a resource handle.
//
//go:nosplit
func (self TCPSocket) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketResourceDrop((uint32)(self0))
	return
}

// Accept represents the imported method "accept".
//
//	accept: func() -> result<tuple<tcp-socket, input-stream, output-stream>, error-code>
//
//go:nosplit
func (self TCPSocket) Accept() (result cm.Result[TupleTCPSocketInputStreamOutputStreamShape, cm.Tuple3[TCPSocket, InputStream, OutputStream], ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketAccept((uint32)(self0), &result)
	return
}

// AddressFamily represents the imported method "address-family".
//
//	address-family: func() -> ip-address-family
//
//go:nosplit
func (self TCPSocket) AddressFamily() (result IPAddressFamily) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_TCPSocketAddressFamily((uint32)(self0))
	result = (network.IPAddressFamily)((uint32)(result0))
	return
}

// FinishBind represents the imported method "finish-bind".
//
//	finish-bind: func() -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) FinishBind() (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketFinishBind((uint32)(self0), &result)
	return
}

// FinishConnect represents the imported method "finish-connect".
//
//	finish-connect: func() -> result<tuple<input-stream, output-stream>, error-code>
//
//go:nosplit
func (self TCPSocket) FinishConnect() (result cm.Result[TupleInputStreamOutputStreamShape, cm.Tuple[InputStream, OutputStream], ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketFinishConnect((uint32)(self0), &result)
	return
}

// FinishListen represents the imported method "finish-listen".
//
//	finish-listen: func() -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) FinishListen() (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketFinishListen((uint32)(self0), &result)
	return
}

// HopLimit represents the imported method "hop-limit".
//
//	hop-limit: func() -> result<u8, error-code>
//
//go:nosplit
func (self TCPSocket) HopLimit() (result cm.Result[uint8, uint8, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketHopLimit((uint32)(self0), &result)
	return
}

// IsListening represents the imported method "is-listening".
//
//	is-listening: func() -> bool
//
//go:nosplit
func (self TCPSocket) IsListening() (result bool) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_TCPSocketIsListening((uint32)(self0))
	result = (bool)(cm.U32ToBool((uint32)(result0)))
	return
}

// KeepAliveCount represents the imported method "keep-alive-count".
//
//	keep-alive-count: func() -> result<u32, error-code>
//
//go:nosplit
func (self TCPSocket) KeepAliveCount() (result cm.Result[uint32, uint32, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketKeepAliveCount((uint32)(self0), &result)
	return
}

// KeepAliveEnabled represents the imported method "keep-alive-enabled".
//
//	keep-alive-enabled: func() -> result<bool, error-code>
//
//go:nosplit
func (self TCPSocket) KeepAliveEnabled() (result cm.Result[ErrorCode, bool, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketKeepAliveEnabled((uint32)(self0), &result)
	return
}

// KeepAliveIdleTime represents the imported method "keep-alive-idle-time".
//
//	keep-alive-idle-time: func() -> result<duration, error-code>
//
//go:nosplit
func (self TCPSocket) KeepAliveIdleTime() (result cm.Result[uint64, Duration, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketKeepAliveIdleTime((uint32)(self0), &result)
	return
}

// KeepAliveInterval represents the imported method "keep-alive-interval".
//
//	keep-alive-interval: func() -> result<duration, error-code>
//
//go:nosplit
func (self TCPSocket) KeepAliveInterval() (result cm.Result[uint64, Duration, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketKeepAliveInterval((uint32)(self0), &result)
	return
}

// LocalAddress represents the imported method "local-address".
//
//	local-address: func() -> result<ip-socket-address, error-code>
//
//go:nosplit
func (self TCPSocket) LocalAddress() (result cm.Result[IPSocketAddressShape, IPSocketAddress, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketLocalAddress((uint32)(self0), &result)
	return
}

// ReceiveBufferSize represents the imported method "receive-buffer-size".
//
//	receive-buffer-size: func() -> result<u64, error-code>
//
//go:nosplit
func (self TCPSocket) ReceiveBufferSize() (result cm.Result[uint64, uint64, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketReceiveBufferSize((uint32)(self0), &result)
	return
}

// RemoteAddress represents the imported method "remote-address".
//
//	remote-address: func() -> result<ip-socket-address, error-code>
//
//go:nosplit
func (self TCPSocket) RemoteAddress() (result cm.Result[IPSocketAddressShape, IPSocketAddress, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketRemoteAddress((uint32)(self0), &result)
	return
}

// SendBufferSize represents the imported method "send-buffer-size".
//
//	send-buffer-size: func() -> result<u64, error-code>
//
//go:nosplit
func (self TCPSocket) SendBufferSize() (result cm.Result[uint64, uint64, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketSendBufferSize((uint32)(self0), &result)
	return
}

// SetHopLimit represents the imported method "set-hop-limit".
//
//	set-hop-limit: func(value: u8) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetHopLimit(value uint8) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	value0 := (uint32)(value)
	wasmimport_TCPSocketSetHopLimit((uint32)(self0), (uint32)(value0), &result)
	return
}

// SetKeepAliveCount represents the imported method "set-keep-alive-count".
//
//	set-keep-alive-count: func(value: u32) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetKeepAliveCount(value uint32) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	value0 := (uint32)(value)
	wasmimport_TCPSocketSetKeepAliveCount((uint32)(self0), (uint32)(value0), &result)
	return
}

// SetKeepAliveEnabled represents the imported method "set-keep-alive-enabled".
//
//	set-keep-alive-enabled: func(value: bool) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetKeepAliveEnabled(value bool) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	value0 := (uint32)(cm.BoolToU32(value))
	wasmimport_TCPSocketSetKeepAliveEnabled((uint32)(self0), (uint32)(value0), &result)
	return
}

// SetKeepAliveIdleTime represents the imported method "set-keep-alive-idle-time".
//
//	set-keep-alive-idle-time: func(value: duration) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetKeepAliveIdleTime(value Duration) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	value0 := (uint64)(value)
	wasmimport_TCPSocketSetKeepAliveIdleTime((uint32)(self0), (uint64)(value0), &result)
	return
}

// SetKeepAliveInterval represents the imported method "set-keep-alive-interval".
//
//	set-keep-alive-interval: func(value: duration) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetKeepAliveInterval(value Duration) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	value0 := (uint64)(value)
	wasmimport_TCPSocketSetKeepAliveInterval((uint32)(self0), (uint64)(value0), &result)
	return
}

// SetListenBacklogSize represents the imported method "set-listen-backlog-size".
//
//	set-listen-backlog-size: func(value: u64) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetListenBacklogSize(value uint64) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	value0 := (uint64)(value)
	wasmimport_TCPSocketSetListenBacklogSize((uint32)(self0), (uint64)(value0), &result)
	return
}

// SetReceiveBufferSize represents the imported method "set-receive-buffer-size".
//
//	set-receive-buffer-size: func(value: u64) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetReceiveBufferSize(value uint64) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	value0 := (uint64)(value)
	wasmimport_TCPSocketSetReceiveBufferSize((uint32)(self0), (uint64)(value0), &result)
	return
}

// SetSendBufferSize represents the imported method "set-send-buffer-size".
//
//	set-send-buffer-size: func(value: u64) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) SetSendBufferSize(value uint64) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	value0 := (uint64)(value)
	wasmimport_TCPSocketSetSendBufferSize((uint32)(self0), (uint64)(value0), &result)
	return
}

// Shutdown represents the imported method "shutdown".
//
//	shutdown: func(shutdown-type: shutdown-type) -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) Shutdown(shutdownType ShutdownType) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	shutdownType0 := (uint32)(shutdownType)
	wasmimport_TCPSocketShutdown((uint32)(self0), (uint32)(shutdownType0), &result)
	return
}

// StartBind represents the imported method "start-bind".
//
//	start-bind: func(network: borrow<network>, local-address: ip-socket-address) ->
//	result<_, error-code>
//
//go:nosplit
func (self TCPSocket) StartBind(network_ Network, localAddress IPSocketAddress) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	network0 := cm.Reinterpret[uint32](network_)
	localAddress0, localAddress1, localAddress2, localAddress3, localAddress4, localAddress5, localAddress6, localAddress7, localAddress8, localAddress9, localAddress10, localAddress11 := lower_IPSocketAddress(localAddress)
	wasmimport_TCPSocketStartBind((uint32)(self0), (uint32)(network0), (uint32)(localAddress0), (uint32)(localAddress1), (uint32)(localAddress2), (uint32)(localAddress3), (uint32)(localAddress4), (uint32)(localAddress5), (uint32)(localAddress6), (uint32)(localAddress7), (uint32)(localAddress8), (uint32)(localAddress9), (uint32)(localAddress10), (uint32)(localAddress11), &result)
	return
}

// StartConnect represents the imported method "start-connect".
//
//	start-connect: func(network: borrow<network>, remote-address: ip-socket-address)
//	-> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) StartConnect(network_ Network, remoteAddress IPSocketAddress) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	network0 := cm.Reinterpret[uint32](network_)
	remoteAddress0, remoteAddress1, remoteAddress2, remoteAddress3, remoteAddress4, remoteAddress5, remoteAddress6, remoteAddress7, remoteAddress8, remoteAddress9, remoteAddress10, remoteAddress11 := lower_IPSocketAddress(remoteAddress)
	wasmimport_TCPSocketStartConnect((uint32)(self0), (uint32)(network0), (uint32)(remoteAddress0), (uint32)(remoteAddress1), (uint32)(remoteAddress2), (uint32)(remoteAddress3), (uint32)(remoteAddress4), (uint32)(remoteAddress5), (uint32)(remoteAddress6), (uint32)(remoteAddress7), (uint32)(remoteAddress8), (uint32)(remoteAddress9), (uint32)(remoteAddress10), (uint32)(remoteAddress11), &result)
	return
}

// StartListen represents the imported method "start-listen".
//
//	start-listen: func() -> result<_, error-code>
//
//go:nosplit
func (self TCPSocket) StartListen() (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TCPSocketStartListen((uint32)(self0), &result)
	return
}

// Subscribe represents the imported method "subscribe".
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self TCPSocket) Subscribe() (result Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_TCPSocketSubscribe((uint32)(self0))
	result = cm.Reinterpret[Pollable]((uint32)(result0))
	return
}
