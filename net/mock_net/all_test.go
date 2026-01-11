package mock_net

import (
	"errors"
	"io"
	"testing"
	"time"

	"github.com/pdutton/go-mocks/internal/testutil"
	"go.uber.org/mock/gomock"
)

// TestMockConn_Read tests connection read.
func TestMockConn_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)

	buf := make([]byte, 10)
	mockConn.EXPECT().Read(buf).Return(10, nil)

	n, err := mockConn.Read(buf)
	testutil.AssertEqual(t, 10, n)
	testutil.AssertNil(t, err)
}

// TestMockConn_Write tests connection write.
func TestMockConn_Write(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)

	data := []byte("hello")
	mockConn.EXPECT().Write(data).Return(5, nil)

	n, err := mockConn.Write(data)
	testutil.AssertEqual(t, 5, n)
	testutil.AssertNil(t, err)
}

// TestMockConn_Close tests connection close.
func TestMockConn_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)

	mockConn.EXPECT().Close().Return(nil)

	err := mockConn.Close()
	testutil.AssertNil(t, err)
}

// TestMockConn_Lifecycle tests full connection lifecycle.
func TestMockConn_Lifecycle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)

	gomock.InOrder(
		mockConn.EXPECT().Write(gomock.Any()).Return(5, nil),
		mockConn.EXPECT().Read(gomock.Any()).Return(10, nil),
		mockConn.EXPECT().Close().Return(nil),
	)

	n1, err1 := mockConn.Write([]byte("hello"))
	testutil.AssertEqual(t, 5, n1)
	testutil.AssertNil(t, err1)

	n2, err2 := mockConn.Read(make([]byte, 10))
	testutil.AssertEqual(t, 10, n2)
	testutil.AssertNil(t, err2)

	err3 := mockConn.Close()
	testutil.AssertNil(t, err3)
}

// TestMockConn_SetDeadline tests setting deadline.
func TestMockConn_SetDeadline(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)

	deadline := time.Now().Add(5 * time.Second)
	mockConn.EXPECT().SetDeadline(deadline).Return(nil)

	err := mockConn.SetDeadline(deadline)
	testutil.AssertNil(t, err)
}

// TestMockConn_SetReadDeadline tests setting read deadline.
func TestMockConn_SetReadDeadline(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)

	deadline := time.Now().Add(5 * time.Second)
	mockConn.EXPECT().SetReadDeadline(deadline).Return(nil)

	err := mockConn.SetReadDeadline(deadline)
	testutil.AssertNil(t, err)
}

// TestMockConn_SetWriteDeadline tests setting write deadline.
func TestMockConn_SetWriteDeadline(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)

	deadline := time.Now().Add(5 * time.Second)
	mockConn.EXPECT().SetWriteDeadline(deadline).Return(nil)

	err := mockConn.SetWriteDeadline(deadline)
	testutil.AssertNil(t, err)
}

// TestMockConn_LocalAddr tests getting local address.
func TestMockConn_LocalAddr(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)
	mockAddr := NewMockAddr(ctrl)

	mockConn.EXPECT().LocalAddr().Return(mockAddr)

	addr := mockConn.LocalAddr()
	testutil.AssertNotNil(t, addr)
}

// TestMockConn_RemoteAddr tests getting remote address.
func TestMockConn_RemoteAddr(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)
	mockAddr := NewMockAddr(ctrl)

	mockConn.EXPECT().RemoteAddr().Return(mockAddr)

	addr := mockConn.RemoteAddr()
	testutil.AssertNotNil(t, addr)
}

// TestMockConn_ReadError tests read error handling.
func TestMockConn_ReadError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)

	expectedErr := errors.New("connection reset")
	mockConn.EXPECT().Read(gomock.Any()).Return(0, expectedErr)

	n, err := mockConn.Read(make([]byte, 10))
	testutil.AssertEqual(t, 0, n)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockConn_WriteError tests write error handling.
func TestMockConn_WriteError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)

	expectedErr := errors.New("broken pipe")
	mockConn.EXPECT().Write(gomock.Any()).Return(0, expectedErr)

	n, err := mockConn.Write([]byte("test"))
	testutil.AssertEqual(t, 0, n)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockConn_EOF tests EOF handling.
func TestMockConn_EOF(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)

	mockConn.EXPECT().Read(gomock.Any()).Return(0, io.EOF)

	n, err := mockConn.Read(make([]byte, 10))
	testutil.AssertEqual(t, 0, n)
	testutil.AssertError(t, io.EOF, err)
}

// TestMockListener_Accept tests accepting connections.
func TestMockListener_Accept(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockListener := NewMockListener(ctrl)
	mockConn := NewMockConn(ctrl)

	mockListener.EXPECT().Accept().Return(mockConn, nil)

	conn, err := mockListener.Accept()
	testutil.AssertNotNil(t, conn)
	testutil.AssertNil(t, err)
}

// TestMockListener_AcceptError tests accept error handling.
func TestMockListener_AcceptError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockListener := NewMockListener(ctrl)

	expectedErr := errors.New("listener closed")
	mockListener.EXPECT().Accept().Return(nil, expectedErr)

	conn, err := mockListener.Accept()
	testutil.AssertNil(t, conn)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockListener_Close tests closing listener.
func TestMockListener_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockListener := NewMockListener(ctrl)

	mockListener.EXPECT().Close().Return(nil)

	err := mockListener.Close()
	testutil.AssertNil(t, err)
}

// TestMockListener_Addr tests getting listener address.
func TestMockListener_Addr(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockListener := NewMockListener(ctrl)
	mockAddr := NewMockAddr(ctrl)

	mockListener.EXPECT().Addr().Return(mockAddr)

	addr := mockListener.Addr()
	testutil.AssertNotNil(t, addr)
}

// TestMockListener_AcceptMultiple tests multiple accepts.
func TestMockListener_AcceptMultiple(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockListener := NewMockListener(ctrl)
	mockConn1 := NewMockConn(ctrl)
	mockConn2 := NewMockConn(ctrl)

	gomock.InOrder(
		mockListener.EXPECT().Accept().Return(mockConn1, nil),
		mockListener.EXPECT().Accept().Return(mockConn2, nil),
		mockListener.EXPECT().Close().Return(nil),
	)

	conn1, err1 := mockListener.Accept()
	testutil.AssertNotNil(t, conn1)
	testutil.AssertNil(t, err1)

	conn2, err2 := mockListener.Accept()
	testutil.AssertNotNil(t, conn2)
	testutil.AssertNil(t, err2)

	err3 := mockListener.Close()
	testutil.AssertNil(t, err3)
}

// TestMockDialer_Dial tests dialing a connection.
func TestMockDialer_Dial(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDialer := NewMockDialer(ctrl)
	mockConn := NewMockConn(ctrl)

	mockDialer.EXPECT().Dial("tcp", "localhost:8080").Return(mockConn, nil)

	conn, err := mockDialer.Dial("tcp", "localhost:8080")
	testutil.AssertNotNil(t, conn)
	testutil.AssertNil(t, err)
}

// TestMockDialer_DialError tests dial error handling.
func TestMockDialer_DialError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDialer := NewMockDialer(ctrl)

	expectedErr := errors.New("connection refused")
	mockDialer.EXPECT().Dial("tcp", "localhost:8080").Return(nil, expectedErr)

	conn, err := mockDialer.Dial("tcp", "localhost:8080")
	testutil.AssertNil(t, conn)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockAddr_Network tests getting network type.
func TestMockAddr_Network(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAddr := NewMockAddr(ctrl)

	mockAddr.EXPECT().Network().Return("tcp")

	network := mockAddr.Network()
	testutil.AssertEqual(t, "tcp", network)
}

// TestMockAddr_String tests address string representation.
func TestMockAddr_String(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAddr := NewMockAddr(ctrl)

	mockAddr.EXPECT().String().Return("127.0.0.1:8080")

	addr := mockAddr.String()
	testutil.AssertEqual(t, "127.0.0.1:8080", addr)
}

// TestMockPacketConn_ReadFrom tests reading from packet connection.
func TestMockPacketConn_ReadFrom(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPacketConn := NewMockPacketConn(ctrl)
	mockAddr := NewMockAddr(ctrl)

	buf := make([]byte, 1024)
	mockPacketConn.EXPECT().ReadFrom(buf).Return(10, mockAddr, nil)

	n, addr, err := mockPacketConn.ReadFrom(buf)
	testutil.AssertEqual(t, 10, n)
	testutil.AssertNotNil(t, addr)
	testutil.AssertNil(t, err)
}

// TestMockPacketConn_WriteTo tests writing to packet connection.
func TestMockPacketConn_WriteTo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPacketConn := NewMockPacketConn(ctrl)
	mockAddr := NewMockAddr(ctrl)

	data := []byte("hello")
	mockPacketConn.EXPECT().WriteTo(data, mockAddr).Return(5, nil)

	n, err := mockPacketConn.WriteTo(data, mockAddr)
	testutil.AssertEqual(t, 5, n)
	testutil.AssertNil(t, err)
}

// TestMockPacketConn_Close tests closing packet connection.
func TestMockPacketConn_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPacketConn := NewMockPacketConn(ctrl)

	mockPacketConn.EXPECT().Close().Return(nil)

	err := mockPacketConn.Close()
	testutil.AssertNil(t, err)
}

// TestMockPacketConn_SetDeadline tests setting packet connection deadline.
func TestMockPacketConn_SetDeadline(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPacketConn := NewMockPacketConn(ctrl)

	deadline := time.Now().Add(5 * time.Second)
	mockPacketConn.EXPECT().SetDeadline(deadline).Return(nil)

	err := mockPacketConn.SetDeadline(deadline)
	testutil.AssertNil(t, err)
}

// TestMockResolver_LookupHost tests host lookup.
func TestMockResolver_LookupHost(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockResolver := NewMockResolver(ctrl)

	addrs := []string{"192.168.1.1", "192.168.1.2"}
	mockResolver.EXPECT().LookupHost(gomock.Any(), "example.com").Return(addrs, nil)

	result, err := mockResolver.LookupHost(nil, "example.com")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, 2, len(result))
	testutil.AssertEqual(t, "192.168.1.1", result[0])
}

// TestMockResolver_LookupHostError tests lookup error.
func TestMockResolver_LookupHostError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockResolver := NewMockResolver(ctrl)

	expectedErr := errors.New("no such host")
	mockResolver.EXPECT().LookupHost(gomock.Any(), "invalid.example").Return([]string(nil), expectedErr)

	result, err := mockResolver.LookupHost(nil, "invalid.example")
	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
	testutil.AssertError(t, expectedErr, err)
}

// TestMockTCPConn_CloseRead tests closing read side of TCP connection.
func TestMockTCPConn_CloseRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTCPConn := NewMockTCPConn(ctrl)

	mockTCPConn.EXPECT().CloseRead().Return(nil)

	err := mockTCPConn.CloseRead()
	testutil.AssertNil(t, err)
}

// TestMockTCPConn_CloseWrite tests closing write side of TCP connection.
func TestMockTCPConn_CloseWrite(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTCPConn := NewMockTCPConn(ctrl)

	mockTCPConn.EXPECT().CloseWrite().Return(nil)

	err := mockTCPConn.CloseWrite()
	testutil.AssertNil(t, err)
}

// TestMockTCPListener_AcceptTCP tests accepting TCP connection.
func TestMockTCPListener_AcceptTCP(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTCPListener := NewMockTCPListener(ctrl)
	mockTCPConn := NewMockTCPConn(ctrl)

	mockTCPListener.EXPECT().AcceptTCP().Return(mockTCPConn, nil)

	conn, err := mockTCPListener.AcceptTCP()
	testutil.AssertNotNil(t, conn)
	testutil.AssertNil(t, err)
}

// TestMockNet_Listen tests creating a listener.
func TestMockNet_Listen(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockNet := NewMockNet(ctrl)
	mockListener := NewMockListener(ctrl)

	mockNet.EXPECT().Listen("tcp", ":8080").Return(mockListener, nil)

	listener, err := mockNet.Listen("tcp", ":8080")
	testutil.AssertNotNil(t, listener)
	testutil.AssertNil(t, err)
}

// TestMockNet_Dial tests dialing via Net interface.
func TestMockNet_Dial(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockNet := NewMockNet(ctrl)
	mockConn := NewMockConn(ctrl)

	mockNet.EXPECT().Dial("tcp", "localhost:8080").Return(mockConn, nil)

	conn, err := mockNet.Dial("tcp", "localhost:8080")
	testutil.AssertNotNil(t, conn)
	testutil.AssertNil(t, err)
}

// TestMockConn_ReadWriteSequence tests alternating read/write.
func TestMockConn_ReadWriteSequence(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConn := NewMockConn(ctrl)

	gomock.InOrder(
		mockConn.EXPECT().Write([]byte("request")).Return(7, nil),
		mockConn.EXPECT().Read(gomock.Any()).Return(8, nil),
		mockConn.EXPECT().Write([]byte("ack")).Return(3, nil),
		mockConn.EXPECT().Close().Return(nil),
	)

	n1, _ := mockConn.Write([]byte("request"))
	testutil.AssertEqual(t, 7, n1)

	n2, _ := mockConn.Read(make([]byte, 10))
	testutil.AssertEqual(t, 8, n2)

	n3, _ := mockConn.Write([]byte("ack"))
	testutil.AssertEqual(t, 3, n3)

	mockConn.Close()
}
