package mock_fs

import (
	"errors"
	"io"
	"io/fs"
	"testing"

	"github.com/pdutton/go-mocks/internal/testutil"
	"go.uber.org/mock/gomock"
)

// TestMockFS_Open tests basic file open operation.
func TestMockFS_Open(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFS := NewMockFS(ctrl)
	mockFile := NewMockFile(ctrl)

	mockFS.EXPECT().Open("test.txt").Return(mockFile, nil)

	file, err := mockFS.Open("test.txt")
	testutil.AssertNotNil(t, file)
	testutil.AssertNil(t, err)
}

// TestMockFS_OpenNotFound tests file not found error.
func TestMockFS_OpenNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFS := NewMockFS(ctrl)

	expectedErr := fs.ErrNotExist
	mockFS.EXPECT().Open("missing.txt").Return(nil, expectedErr)

	file, err := mockFS.Open("missing.txt")
	testutil.AssertNil(t, file)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockFS_OpenWithClose tests Open followed by Close.
func TestMockFS_OpenWithClose(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFS := NewMockFS(ctrl)
	mockFile := NewMockFile(ctrl)

	gomock.InOrder(
		mockFS.EXPECT().Open("test.txt").Return(mockFile, nil),
		mockFile.EXPECT().Close().Return(nil),
	)

	file, err := mockFS.Open("test.txt")
	testutil.AssertNil(t, err)

	err = file.Close()
	testutil.AssertNil(t, err)
}

// TestMockFile_Read tests file read operations.
func TestMockFile_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	buf := make([]byte, 10)
	mockFile.EXPECT().Read(buf).Return(10, nil)

	n, err := mockFile.Read(buf)
	testutil.AssertEqual(t, 10, n)
	testutil.AssertNil(t, err)
}

// TestMockFile_ReadEOF tests EOF scenario.
func TestMockFile_ReadEOF(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	gomock.InOrder(
		mockFile.EXPECT().Read(gomock.Any()).Return(10, nil),
		mockFile.EXPECT().Read(gomock.Any()).Return(0, io.EOF),
	)

	n1, err1 := mockFile.Read(make([]byte, 10))
	testutil.AssertEqual(t, 10, n1)
	testutil.AssertNil(t, err1)

	n2, err2 := mockFile.Read(make([]byte, 10))
	testutil.AssertEqual(t, 0, n2)
	testutil.AssertError(t, io.EOF, err2)
}

// TestMockFile_Close tests file close.
func TestMockFile_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	mockFile.EXPECT().Close().Return(nil)

	err := mockFile.Close()
	testutil.AssertNil(t, err)
}

// TestMockFile_CloseError tests close error handling.
func TestMockFile_CloseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	expectedErr := errors.New("close error")
	mockFile.EXPECT().Close().Return(expectedErr)

	err := mockFile.Close()
	testutil.AssertError(t, expectedErr, err)
}

// TestMockFileInfo_Properties tests FileInfo property methods.
func TestMockFileInfo_Properties(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFileInfo := NewMockFileInfo(ctrl)
	mockFileMode := NewMockFileMode(ctrl)

	mockFileInfo.EXPECT().Name().Return("test.txt")
	mockFileInfo.EXPECT().Size().Return(int64(1024))
	mockFileInfo.EXPECT().Mode().Return(mockFileMode)
	mockFileInfo.EXPECT().IsDir().Return(false)

	name := mockFileInfo.Name()
	testutil.AssertEqual(t, "test.txt", name)

	size := mockFileInfo.Size()
	testutil.AssertEqual(t, int64(1024), size)

	mode := mockFileInfo.Mode()
	testutil.AssertNotNil(t, mode)

	isDir := mockFileInfo.IsDir()
	testutil.AssertEqual(t, false, isDir)
}

// TestMockDirEntry_Properties tests directory entry properties.
func TestMockDirEntry_Properties(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDirEntry := NewMockDirEntry(ctrl)
	mockFileMode := NewMockFileMode(ctrl)

	mockDirEntry.EXPECT().Name().Return("file.txt")
	mockDirEntry.EXPECT().IsDir().Return(false)
	mockDirEntry.EXPECT().Type().Return(mockFileMode)

	name := mockDirEntry.Name()
	testutil.AssertEqual(t, "file.txt", name)

	isDir := mockDirEntry.IsDir()
	testutil.AssertEqual(t, false, isDir)

	fileType := mockDirEntry.Type()
	testutil.AssertNotNil(t, fileType)
}

// TestMockDirEntry_Info tests directory entry info.
func TestMockDirEntry_Info(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDirEntry := NewMockDirEntry(ctrl)
	mockFileInfo := NewMockFileInfo(ctrl)

	mockDirEntry.EXPECT().Info().Return(mockFileInfo, nil)

	info, err := mockDirEntry.Info()
	testutil.AssertNotNil(t, info)
	testutil.AssertNil(t, err)
}

// TestMockGlobFS_Glob tests pattern matching.
func TestMockGlobFS_Glob(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockGlobFS := NewMockGlobFS(ctrl)

	matches := []string{"test1.txt", "test2.txt"}
	mockGlobFS.EXPECT().Glob("*.txt").Return(matches, nil)

	result, err := mockGlobFS.Glob("*.txt")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, 2, len(result))
	testutil.AssertEqual(t, "test1.txt", result[0])
	testutil.AssertEqual(t, "test2.txt", result[1])
}

// TestMockGlobFS_GlobNoMatches tests glob with no matches.
func TestMockGlobFS_GlobNoMatches(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockGlobFS := NewMockGlobFS(ctrl)

	mockGlobFS.EXPECT().Glob("*.xyz").Return([]string{}, nil)

	result, err := mockGlobFS.Glob("*.xyz")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, 0, len(result))
}

// TestMockSubFS_Sub tests subdirectory access.
func TestMockSubFS_Sub(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSubFS := NewMockSubFS(ctrl)
	mockFS := NewMockFS(ctrl)

	mockSubFS.EXPECT().Sub("subdir").Return(mockFS, nil)

	subFS, err := mockSubFS.Sub("subdir")
	testutil.AssertNotNil(t, subFS)
	testutil.AssertNil(t, err)
}

// TestMockSubFS_SubError tests Sub error handling.
func TestMockSubFS_SubError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSubFS := NewMockSubFS(ctrl)

	expectedErr := fs.ErrNotExist
	mockSubFS.EXPECT().Sub("invalid").Return(nil, expectedErr)

	subFS, err := mockSubFS.Sub("invalid")
	testutil.AssertNil(t, subFS)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockReadFileFS_ReadFile tests reading entire file.
func TestMockReadFileFS_ReadFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReadFileFS := NewMockReadFileFS(ctrl)

	content := []byte("file contents")
	mockReadFileFS.EXPECT().ReadFile("test.txt").Return(content, nil)

	data, err := mockReadFileFS.ReadFile("test.txt")
	testutil.AssertNil(t, err)
	testutil.AssertBytes(t, content, data)
}

// TestMockStatFS_StatError tests Stat error handling.
func TestMockStatFS_StatError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStatFS := NewMockStatFS(ctrl)

	expectedErr := fs.ErrNotExist
	mockStatFS.EXPECT().Stat("missing.txt").Return(nil, expectedErr)

	info, err := mockStatFS.Stat("missing.txt")
	testutil.AssertNil(t, info)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockFileMode_String tests FileMode string representation.
func TestMockFileMode_String(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFileMode := NewMockFileMode(ctrl)

	mockFileMode.EXPECT().String().Return("-rw-r--r--")

	str := mockFileMode.String()
	testutil.AssertEqual(t, "-rw-r--r--", str)
}

// TestMockFileMode_IsDir tests FileMode directory check.
func TestMockFileMode_IsDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFileMode := NewMockFileMode(ctrl)

	mockFileMode.EXPECT().IsDir().Return(false)

	isDir := mockFileMode.IsDir()
	testutil.AssertEqual(t, false, isDir)
}

// TestMockFileMode_IsRegular tests FileMode regular file check.
func TestMockFileMode_IsRegular(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFileMode := NewMockFileMode(ctrl)

	mockFileMode.EXPECT().IsRegular().Return(true)

	isRegular := mockFileMode.IsRegular()
	testutil.AssertEqual(t, true, isRegular)
}

// TestMockFile_ReadAndClose tests read followed by close.
func TestMockFile_ReadAndClose(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	gomock.InOrder(
		mockFile.EXPECT().Read(gomock.Any()).Return(10, nil),
		mockFile.EXPECT().Read(gomock.Any()).Return(0, io.EOF),
		mockFile.EXPECT().Close().Return(nil),
	)

	n, err := mockFile.Read(make([]byte, 10))
	testutil.AssertEqual(t, 10, n)
	testutil.AssertNil(t, err)

	_, err = mockFile.Read(make([]byte, 10))
	testutil.AssertError(t, io.EOF, err)

	err = mockFile.Close()
	testutil.AssertNil(t, err)
}

// TestMockGlobFS_MultiplePatterns tests multiple glob patterns.
func TestMockGlobFS_MultiplePatterns(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockGlobFS := NewMockGlobFS(ctrl)

	gomock.InOrder(
		mockGlobFS.EXPECT().Glob("*.txt").Return([]string{"a.txt", "b.txt"}, nil),
		mockGlobFS.EXPECT().Glob("*.md").Return([]string{"README.md"}, nil),
	)

	txtFiles, err := mockGlobFS.Glob("*.txt")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, 2, len(txtFiles))

	mdFiles, err := mockGlobFS.Glob("*.md")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, 1, len(mdFiles))
}

// TestMockReadFileFS_MultipleReads tests reading multiple files.
func TestMockReadFileFS_MultipleReads(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReadFileFS := NewMockReadFileFS(ctrl)

	file1Content := []byte("content1")
	file2Content := []byte("content2")

	gomock.InOrder(
		mockReadFileFS.EXPECT().ReadFile("file1.txt").Return(file1Content, nil),
		mockReadFileFS.EXPECT().ReadFile("file2.txt").Return(file2Content, nil),
	)

	data1, err1 := mockReadFileFS.ReadFile("file1.txt")
	testutil.AssertNil(t, err1)
	testutil.AssertBytes(t, file1Content, data1)

	data2, err2 := mockReadFileFS.ReadFile("file2.txt")
	testutil.AssertNil(t, err2)
	testutil.AssertBytes(t, file2Content, data2)
}
