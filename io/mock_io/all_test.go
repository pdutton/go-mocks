package mock_io

import (
	"errors"
	"io"
	"testing"

	"github.com/pdutton/go-mocks/internal/testutil"
	"go.uber.org/mock/gomock"
)

// TestMockReader_BasicRead tests a single Read() call.
func TestMockReader_BasicRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReader := NewMockReader(ctrl)

	buf := make([]byte, 10)
	mockReader.EXPECT().Read(buf).Return(5, nil)

	n, err := mockReader.Read(buf)
	testutil.AssertEqual(t, 5, n)
	testutil.AssertNil(t, err)
}

// TestMockReader_MultipleReads tests ordered Read() calls.
func TestMockReader_MultipleReads(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReader := NewMockReader(ctrl)

	buf := make([]byte, 10)

	gomock.InOrder(
		mockReader.EXPECT().Read(gomock.Any()).Return(5, nil),
		mockReader.EXPECT().Read(gomock.Any()).Return(3, nil),
		mockReader.EXPECT().Read(gomock.Any()).Return(0, io.EOF),
	)

	n1, err1 := mockReader.Read(buf)
	testutil.AssertEqual(t, 5, n1)
	testutil.AssertNil(t, err1)

	n2, err2 := mockReader.Read(buf)
	testutil.AssertEqual(t, 3, n2)
	testutil.AssertNil(t, err2)

	n3, err3 := mockReader.Read(buf)
	testutil.AssertEqual(t, 0, n3)
	testutil.AssertError(t, io.EOF, err3)
}

// TestMockReader_ReadError tests error handling.
func TestMockReader_ReadError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReader := NewMockReader(ctrl)

	expectedErr := errors.New("read error")
	mockReader.EXPECT().Read(gomock.Any()).Return(0, expectedErr)

	n, err := mockReader.Read(make([]byte, 10))
	testutil.AssertEqual(t, 0, n)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockReader_EOF tests EOF scenario.
func TestMockReader_EOF(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReader := NewMockReader(ctrl)

	mockReader.EXPECT().Read(gomock.Any()).Return(0, io.EOF)

	_, err := mockReader.Read(make([]byte, 10))
	testutil.AssertError(t, io.EOF, err)
}

// TestMockReader_WithTimes tests Times() modifier.
func TestMockReader_WithTimes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReader := NewMockReader(ctrl)

	mockReader.EXPECT().Read(gomock.Any()).Return(10, nil).Times(3)

	for i := 0; i < 3; i++ {
		n, err := mockReader.Read(make([]byte, 10))
		testutil.AssertEqual(t, 10, n)
		testutil.AssertNil(t, err)
	}
}

// TestMockWriter_BasicWrite tests a single Write() call.
func TestMockWriter_BasicWrite(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWriter := NewMockWriter(ctrl)

	data := []byte("hello")
	mockWriter.EXPECT().Write(data).Return(5, nil)

	n, err := mockWriter.Write(data)
	testutil.AssertEqual(t, 5, n)
	testutil.AssertNil(t, err)
}

// TestMockWriter_OrderedWrites tests multiple writes in order.
func TestMockWriter_OrderedWrites(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWriter := NewMockWriter(ctrl)

	data1 := []byte("hello")
	data2 := []byte("world")

	gomock.InOrder(
		mockWriter.EXPECT().Write(data1).Return(5, nil),
		mockWriter.EXPECT().Write(data2).Return(5, nil),
	)

	n1, err1 := mockWriter.Write(data1)
	testutil.AssertEqual(t, 5, n1)
	testutil.AssertNil(t, err1)

	n2, err2 := mockWriter.Write(data2)
	testutil.AssertEqual(t, 5, n2)
	testutil.AssertNil(t, err2)
}

// TestMockWriter_WriteError tests write error handling.
func TestMockWriter_WriteError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWriter := NewMockWriter(ctrl)

	expectedErr := errors.New("write error")
	mockWriter.EXPECT().Write(gomock.Any()).Return(0, expectedErr)

	n, err := mockWriter.Write([]byte("test"))
	testutil.AssertEqual(t, 0, n)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockWriter_PartialWrite tests partial write scenario.
func TestMockWriter_PartialWrite(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWriter := NewMockWriter(ctrl)

	data := []byte("hello world")
	mockWriter.EXPECT().Write(data).Return(5, nil)

	n, err := mockWriter.Write(data)
	testutil.AssertEqual(t, 5, n)
	testutil.AssertNil(t, err)
}

// TestMockCloser_Close tests Close() method.
func TestMockCloser_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCloser := NewMockCloser(ctrl)

	mockCloser.EXPECT().Close().Return(nil)

	err := mockCloser.Close()
	testutil.AssertNil(t, err)
}

// TestMockCloser_CloseError tests Close() error handling.
func TestMockCloser_CloseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCloser := NewMockCloser(ctrl)

	expectedErr := errors.New("close error")
	mockCloser.EXPECT().Close().Return(expectedErr)

	err := mockCloser.Close()
	testutil.AssertError(t, expectedErr, err)
}

// TestMockReadCloser_Lifecycle tests Read + Close sequence.
func TestMockReadCloser_Lifecycle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRC := NewMockReadCloser(ctrl)

	gomock.InOrder(
		mockRC.EXPECT().Read(gomock.Any()).Return(10, nil),
		mockRC.EXPECT().Read(gomock.Any()).Return(0, io.EOF),
		mockRC.EXPECT().Close().Return(nil),
	)

	buf := make([]byte, 10)
	n, err := mockRC.Read(buf)
	testutil.AssertEqual(t, 10, n)
	testutil.AssertNil(t, err)

	_, err = mockRC.Read(buf)
	testutil.AssertError(t, io.EOF, err)

	err = mockRC.Close()
	testutil.AssertNil(t, err)
}

// TestMockSeeker_SeekOperations tests Seek with different whence values.
func TestMockSeeker_SeekOperations(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSeeker := NewMockSeeker(ctrl)

	gomock.InOrder(
		mockSeeker.EXPECT().Seek(int64(0), io.SeekStart).Return(int64(0), nil),
		mockSeeker.EXPECT().Seek(int64(10), io.SeekCurrent).Return(int64(10), nil),
		mockSeeker.EXPECT().Seek(int64(0), io.SeekEnd).Return(int64(100), nil),
	)

	pos1, err1 := mockSeeker.Seek(0, io.SeekStart)
	testutil.AssertEqual(t, int64(0), pos1)
	testutil.AssertNil(t, err1)

	pos2, err2 := mockSeeker.Seek(10, io.SeekCurrent)
	testutil.AssertEqual(t, int64(10), pos2)
	testutil.AssertNil(t, err2)

	pos3, err3 := mockSeeker.Seek(0, io.SeekEnd)
	testutil.AssertEqual(t, int64(100), pos3)
	testutil.AssertNil(t, err3)
}

// TestMockSeeker_SeekError tests Seek error handling.
func TestMockSeeker_SeekError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSeeker := NewMockSeeker(ctrl)

	expectedErr := errors.New("seek error")
	mockSeeker.EXPECT().Seek(gomock.Any(), gomock.Any()).Return(int64(0), expectedErr)

	pos, err := mockSeeker.Seek(0, io.SeekStart)
	testutil.AssertEqual(t, int64(0), pos)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockReaderAt_ReadAt tests ReadAt with offset.
func TestMockReaderAt_ReadAt(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReaderAt := NewMockReaderAt(ctrl)

	buf := make([]byte, 10)
	mockReaderAt.EXPECT().ReadAt(buf, int64(5)).Return(10, nil)

	n, err := mockReaderAt.ReadAt(buf, 5)
	testutil.AssertEqual(t, 10, n)
	testutil.AssertNil(t, err)
}

// TestMockReaderAt_MultipleReadAt tests multiple ReadAt calls at different offsets.
func TestMockReaderAt_MultipleReadAt(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReaderAt := NewMockReaderAt(ctrl)

	buf := make([]byte, 10)

	gomock.InOrder(
		mockReaderAt.EXPECT().ReadAt(gomock.Any(), int64(0)).Return(10, nil),
		mockReaderAt.EXPECT().ReadAt(gomock.Any(), int64(10)).Return(10, nil),
		mockReaderAt.EXPECT().ReadAt(gomock.Any(), int64(20)).Return(5, io.EOF),
	)

	n1, err1 := mockReaderAt.ReadAt(buf, 0)
	testutil.AssertEqual(t, 10, n1)
	testutil.AssertNil(t, err1)

	n2, err2 := mockReaderAt.ReadAt(buf, 10)
	testutil.AssertEqual(t, 10, n2)
	testutil.AssertNil(t, err2)

	n3, err3 := mockReaderAt.ReadAt(buf, 20)
	testutil.AssertEqual(t, 5, n3)
	testutil.AssertError(t, io.EOF, err3)
}

// TestMockWriterTo_WriteTo tests WriteTo operation.
func TestMockWriterTo_WriteTo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWriterTo := NewMockWriterTo(ctrl)
	mockWriter := NewMockWriter(ctrl)

	mockWriterTo.EXPECT().WriteTo(mockWriter).Return(int64(100), nil)

	n, err := mockWriterTo.WriteTo(mockWriter)
	testutil.AssertEqual(t, int64(100), n)
	testutil.AssertNil(t, err)
}

// TestMockWriterTo_WriteToError tests WriteTo error handling.
func TestMockWriterTo_WriteToError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWriterTo := NewMockWriterTo(ctrl)
	mockWriter := NewMockWriter(ctrl)

	expectedErr := errors.New("write to error")
	mockWriterTo.EXPECT().WriteTo(mockWriter).Return(int64(0), expectedErr)

	n, err := mockWriterTo.WriteTo(mockWriter)
	testutil.AssertEqual(t, int64(0), n)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockReadWriter_Combined tests ReadWriter interface.
func TestMockReadWriter_Combined(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRW := NewMockReadWriter(ctrl)

	gomock.InOrder(
		mockRW.EXPECT().Write(gomock.Any()).Return(5, nil),
		mockRW.EXPECT().Read(gomock.Any()).Return(5, nil),
	)

	n1, err1 := mockRW.Write([]byte("hello"))
	testutil.AssertEqual(t, 5, n1)
	testutil.AssertNil(t, err1)

	n2, err2 := mockRW.Read(make([]byte, 10))
	testutil.AssertEqual(t, 5, n2)
	testutil.AssertNil(t, err2)
}

// TestMockReadSeeker_SeekAndRead tests combined Seek and Read operations.
func TestMockReadSeeker_SeekAndRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRS := NewMockReadSeeker(ctrl)

	gomock.InOrder(
		mockRS.EXPECT().Seek(int64(10), io.SeekStart).Return(int64(10), nil),
		mockRS.EXPECT().Read(gomock.Any()).Return(5, nil),
	)

	pos, err := mockRS.Seek(10, io.SeekStart)
	testutil.AssertEqual(t, int64(10), pos)
	testutil.AssertNil(t, err)

	n, err := mockRS.Read(make([]byte, 10))
	testutil.AssertEqual(t, 5, n)
	testutil.AssertNil(t, err)
}

// TestMockByteWriter_WriteByte tests WriteByte operation.
func TestMockByteWriter_WriteByte(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockBW := NewMockByteWriter(ctrl)

	mockBW.EXPECT().WriteByte(byte('A')).Return(nil)

	err := mockBW.WriteByte('A')
	testutil.AssertNil(t, err)
}

// TestMockByteWriter_WriteByteError tests WriteByte error handling.
func TestMockByteWriter_WriteByteError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockBW := NewMockByteWriter(ctrl)

	expectedErr := errors.New("write byte error")
	mockBW.EXPECT().WriteByte(gomock.Any()).Return(expectedErr)

	err := mockBW.WriteByte('A')
	testutil.AssertError(t, expectedErr, err)
}

// TestMockStringWriter_WriteString tests WriteString operation.
func TestMockStringWriter_WriteString(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSW := NewMockStringWriter(ctrl)

	mockSW.EXPECT().WriteString("hello").Return(5, nil)

	n, err := mockSW.WriteString("hello")
	testutil.AssertEqual(t, 5, n)
	testutil.AssertNil(t, err)
}

// TestMockStringWriter_WriteStringError tests WriteString error handling.
func TestMockStringWriter_WriteStringError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSW := NewMockStringWriter(ctrl)

	expectedErr := errors.New("write string error")
	mockSW.EXPECT().WriteString(gomock.Any()).Return(0, expectedErr)

	n, err := mockSW.WriteString("test")
	testutil.AssertEqual(t, 0, n)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockReadWriteCloser_FullLifecycle tests a complete read-write-close sequence.
func TestMockReadWriteCloser_FullLifecycle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRWC := NewMockReadWriteCloser(ctrl)

	gomock.InOrder(
		mockRWC.EXPECT().Write(gomock.Any()).Return(5, nil),
		mockRWC.EXPECT().Read(gomock.Any()).Return(10, nil),
		mockRWC.EXPECT().Read(gomock.Any()).Return(0, io.EOF),
		mockRWC.EXPECT().Close().Return(nil),
	)

	n1, err1 := mockRWC.Write([]byte("hello"))
	testutil.AssertEqual(t, 5, n1)
	testutil.AssertNil(t, err1)

	n2, err2 := mockRWC.Read(make([]byte, 10))
	testutil.AssertEqual(t, 10, n2)
	testutil.AssertNil(t, err2)

	_, err3 := mockRWC.Read(make([]byte, 10))
	testutil.AssertError(t, io.EOF, err3)

	err4 := mockRWC.Close()
	testutil.AssertNil(t, err4)
}

// TestMockReader_AnyTimes tests AnyTimes() modifier.
func TestMockReader_AnyTimes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReader := NewMockReader(ctrl)

	mockReader.EXPECT().Read(gomock.Any()).Return(10, nil).AnyTimes()

	// Can call any number of times
	for i := 0; i < 5; i++ {
		n, err := mockReader.Read(make([]byte, 10))
		testutil.AssertEqual(t, 10, n)
		testutil.AssertNil(t, err)
	}
}

// TestMockWriteCloser_Lifecycle tests Write + Close sequence.
func TestMockWriteCloser_Lifecycle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWC := NewMockWriteCloser(ctrl)

	gomock.InOrder(
		mockWC.EXPECT().Write(gomock.Any()).Return(5, nil),
		mockWC.EXPECT().Write(gomock.Any()).Return(5, nil),
		mockWC.EXPECT().Close().Return(nil),
	)

	n1, err1 := mockWC.Write([]byte("hello"))
	testutil.AssertEqual(t, 5, n1)
	testutil.AssertNil(t, err1)

	n2, err2 := mockWC.Write([]byte("world"))
	testutil.AssertEqual(t, 5, n2)
	testutil.AssertNil(t, err2)

	err3 := mockWC.Close()
	testutil.AssertNil(t, err3)
}
