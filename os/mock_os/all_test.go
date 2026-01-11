package mock_os

import (
	"errors"
	"io"
	"testing"
	"time"

	"github.com/pdutton/go-mocks/internal/testutil"
	"go.uber.org/mock/gomock"
)

// TestMockFile_Read tests file read.
func TestMockFile_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	buf := make([]byte, 10)
	mockFile.EXPECT().Read(buf).Return(10, nil)

	n, err := mockFile.Read(buf)
	testutil.AssertEqual(t, 10, n)
	testutil.AssertNil(t, err)
}

// TestMockFile_ReadError tests read error handling.
func TestMockFile_ReadError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	expectedErr := errors.New("read error")
	mockFile.EXPECT().Read(gomock.Any()).Return(0, expectedErr)

	n, err := mockFile.Read(make([]byte, 10))
	testutil.AssertEqual(t, 0, n)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockFile_Write tests file write.
func TestMockFile_Write(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	data := []byte("hello world")
	mockFile.EXPECT().Write(data).Return(11, nil)

	n, err := mockFile.Write(data)
	testutil.AssertEqual(t, 11, n)
	testutil.AssertNil(t, err)
}

// TestMockFile_WriteError tests write error handling.
func TestMockFile_WriteError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	expectedErr := errors.New("disk full")
	mockFile.EXPECT().Write(gomock.Any()).Return(0, expectedErr)

	n, err := mockFile.Write([]byte("test"))
	testutil.AssertEqual(t, 0, n)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockFile_Close tests file close.
func TestMockFile_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	mockFile.EXPECT().Close().Return(nil)

	err := mockFile.Close()
	testutil.AssertNil(t, err)
}

// TestMockFile_Lifecycle tests full file lifecycle.
func TestMockFile_Lifecycle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	gomock.InOrder(
		mockFile.EXPECT().Write(gomock.Any()).Return(5, nil),
		mockFile.EXPECT().Seek(int64(0), 0).Return(int64(0), nil),
		mockFile.EXPECT().Read(gomock.Any()).Return(5, nil),
		mockFile.EXPECT().Close().Return(nil),
	)

	n1, err1 := mockFile.Write([]byte("hello"))
	testutil.AssertEqual(t, 5, n1)
	testutil.AssertNil(t, err1)

	pos, err2 := mockFile.Seek(0, 0)
	testutil.AssertEqual(t, int64(0), pos)
	testutil.AssertNil(t, err2)

	n2, err3 := mockFile.Read(make([]byte, 10))
	testutil.AssertEqual(t, 5, n2)
	testutil.AssertNil(t, err3)

	err4 := mockFile.Close()
	testutil.AssertNil(t, err4)
}

// TestMockFile_Seek tests seeking in file.
func TestMockFile_Seek(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	mockFile.EXPECT().Seek(int64(10), 0).Return(int64(10), nil)

	pos, err := mockFile.Seek(10, 0)
	testutil.AssertEqual(t, int64(10), pos)
	testutil.AssertNil(t, err)
}

// TestMockFile_Name tests getting file name.
func TestMockFile_Name(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	mockFile.EXPECT().Name().Return("/tmp/test.txt")

	name := mockFile.Name()
	testutil.AssertEqual(t, "/tmp/test.txt", name)
}

// TestMockFile_Stat tests getting file info.
func TestMockFile_Stat(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)
	mockFileInfo := NewMockFileInfo(ctrl)

	mockFile.EXPECT().Stat().Return(mockFileInfo, nil)

	info, err := mockFile.Stat()
	testutil.AssertNotNil(t, info)
	testutil.AssertNil(t, err)
}

// TestMockFile_Sync tests syncing file to disk.
func TestMockFile_Sync(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	mockFile.EXPECT().Sync().Return(nil)

	err := mockFile.Sync()
	testutil.AssertNil(t, err)
}

// TestMockFile_Truncate tests truncating file.
func TestMockFile_Truncate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	mockFile.EXPECT().Truncate(int64(100)).Return(nil)

	err := mockFile.Truncate(100)
	testutil.AssertNil(t, err)
}

// TestMockFile_ReadAt tests reading at offset.
func TestMockFile_ReadAt(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	buf := make([]byte, 5)
	mockFile.EXPECT().ReadAt(buf, int64(10)).Return(5, nil)

	n, err := mockFile.ReadAt(buf, 10)
	testutil.AssertEqual(t, 5, n)
	testutil.AssertNil(t, err)
}

// TestMockFile_WriteAt tests writing at offset.
func TestMockFile_WriteAt(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	data := []byte("hello")
	mockFile.EXPECT().WriteAt(data, int64(10)).Return(5, nil)

	n, err := mockFile.WriteAt(data, 10)
	testutil.AssertEqual(t, 5, n)
	testutil.AssertNil(t, err)
}

// TestMockFile_EOF tests EOF handling.
func TestMockFile_EOF(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFile := NewMockFile(ctrl)

	mockFile.EXPECT().Read(gomock.Any()).Return(0, io.EOF)

	n, err := mockFile.Read(make([]byte, 10))
	testutil.AssertEqual(t, 0, n)
	testutil.AssertError(t, io.EOF, err)
}

// TestMockOS_Create tests creating a file.
func TestMockOS_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)
	mockFile := NewMockFile(ctrl)

	mockOS.EXPECT().Create("/tmp/test.txt").Return(mockFile, nil)

	file, err := mockOS.Create("/tmp/test.txt")
	testutil.AssertNotNil(t, file)
	testutil.AssertNil(t, err)
}

// TestMockOS_CreateError tests create error handling.
func TestMockOS_CreateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	expectedErr := errors.New("permission denied")
	mockOS.EXPECT().Create("/root/test.txt").Return(nil, expectedErr)

	file, err := mockOS.Create("/root/test.txt")
	testutil.AssertNil(t, file)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockOS_Open tests opening a file.
func TestMockOS_Open(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)
	mockFile := NewMockFile(ctrl)

	mockOS.EXPECT().Open("/tmp/test.txt").Return(mockFile, nil)

	file, err := mockOS.Open("/tmp/test.txt")
	testutil.AssertNotNil(t, file)
	testutil.AssertNil(t, err)
}

// TestMockOS_OpenError tests open error handling.
func TestMockOS_OpenError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	expectedErr := errors.New("file not found")
	mockOS.EXPECT().Open("/nonexistent").Return(nil, expectedErr)

	file, err := mockOS.Open("/nonexistent")
	testutil.AssertNil(t, file)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockOS_Remove tests removing a file.
func TestMockOS_Remove(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().Remove("/tmp/test.txt").Return(nil)

	err := mockOS.Remove("/tmp/test.txt")
	testutil.AssertNil(t, err)
}

// TestMockOS_RemoveError tests remove error handling.
func TestMockOS_RemoveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	expectedErr := errors.New("permission denied")
	mockOS.EXPECT().Remove("/protected/file").Return(expectedErr)

	err := mockOS.Remove("/protected/file")
	testutil.AssertError(t, expectedErr, err)
}

// TestMockOS_Stat tests getting file info.
func TestMockOS_Stat(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)
	mockFileInfo := NewMockFileInfo(ctrl)

	mockOS.EXPECT().Stat("/tmp/test.txt").Return(mockFileInfo, nil)

	info, err := mockOS.Stat("/tmp/test.txt")
	testutil.AssertNotNil(t, info)
	testutil.AssertNil(t, err)
}

// TestMockOS_Getenv tests getting environment variable.
func TestMockOS_Getenv(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().Getenv("HOME").Return("/home/user")

	value := mockOS.Getenv("HOME")
	testutil.AssertEqual(t, "/home/user", value)
}

// TestMockOS_Setenv tests setting environment variable.
func TestMockOS_Setenv(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().Setenv("TEST", "value").Return(nil)

	err := mockOS.Setenv("TEST", "value")
	testutil.AssertNil(t, err)
}

// TestMockOS_Unsetenv tests unsetting environment variable.
func TestMockOS_Unsetenv(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().Unsetenv("TEST").Return(nil)

	err := mockOS.Unsetenv("TEST")
	testutil.AssertNil(t, err)
}

// TestMockOS_Environ tests getting all environment variables.
func TestMockOS_Environ(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	env := []string{"PATH=/usr/bin", "HOME=/home/user"}
	mockOS.EXPECT().Environ().Return(env)

	result := mockOS.Environ()
	testutil.AssertEqual(t, 2, len(result))
	testutil.AssertEqual(t, "PATH=/usr/bin", result[0])
}

// TestMockOS_Getwd tests getting working directory.
func TestMockOS_Getwd(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().Getwd().Return("/home/user/project", nil)

	wd, err := mockOS.Getwd()
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, "/home/user/project", wd)
}

// TestMockOS_Chdir tests changing directory.
func TestMockOS_Chdir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().Chdir("/tmp").Return(nil)

	err := mockOS.Chdir("/tmp")
	testutil.AssertNil(t, err)
}

// TestMockOS_TempDir tests getting temp directory.
func TestMockOS_TempDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().TempDir().Return("/tmp")

	dir := mockOS.TempDir()
	testutil.AssertEqual(t, "/tmp", dir)
}

// TestMockOS_UserCacheDir tests getting user cache directory.
func TestMockOS_UserCacheDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().UserCacheDir().Return("/home/user/.cache", nil)

	dir, err := mockOS.UserCacheDir()
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, "/home/user/.cache", dir)
}

// TestMockOS_UserConfigDir tests getting user config directory.
func TestMockOS_UserConfigDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().UserConfigDir().Return("/home/user/.config", nil)

	dir, err := mockOS.UserConfigDir()
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, "/home/user/.config", dir)
}

// TestMockOS_UserHomeDir tests getting user home directory.
func TestMockOS_UserHomeDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().UserHomeDir().Return("/home/user", nil)

	dir, err := mockOS.UserHomeDir()
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, "/home/user", dir)
}

// TestMockOS_Hostname tests getting hostname.
func TestMockOS_Hostname(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().Hostname().Return("localhost", nil)

	hostname, err := mockOS.Hostname()
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, "localhost", hostname)
}

// TestMockOS_Getpid tests getting process ID.
func TestMockOS_Getpid(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().Getpid().Return(12345)

	pid := mockOS.Getpid()
	testutil.AssertEqual(t, 12345, pid)
}

// TestMockOS_Getppid tests getting parent process ID.
func TestMockOS_Getppid(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)

	mockOS.EXPECT().Getppid().Return(1)

	ppid := mockOS.Getppid()
	testutil.AssertEqual(t, 1, ppid)
}

// TestMockFileInfo_Name tests FileInfo name.
func TestMockFileInfo_Name(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockInfo := NewMockFileInfo(ctrl)

	mockInfo.EXPECT().Name().Return("test.txt")

	name := mockInfo.Name()
	testutil.AssertEqual(t, "test.txt", name)
}

// TestMockFileInfo_Size tests file size.
func TestMockFileInfo_Size(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockInfo := NewMockFileInfo(ctrl)

	mockInfo.EXPECT().Size().Return(int64(1024))

	size := mockInfo.Size()
	testutil.AssertEqual(t, int64(1024), size)
}

// TestMockFileInfo_ModTime tests modification time.
func TestMockFileInfo_ModTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockInfo := NewMockFileInfo(ctrl)

	now := time.Now()
	mockInfo.EXPECT().ModTime().Return(now)

	modTime := mockInfo.ModTime()
	testutil.AssertEqual(t, now, modTime)
}

// TestMockFileInfo_IsDir tests directory check.
func TestMockFileInfo_IsDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockInfo := NewMockFileInfo(ctrl)

	mockInfo.EXPECT().IsDir().Return(false)

	isDir := mockInfo.IsDir()
	testutil.AssertEqual(t, false, isDir)
}

// TestMockProcess_Wait tests process wait.
func TestMockProcess_Wait(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockProcess := NewMockProcess(ctrl)

	mockProcess.EXPECT().Wait().Return(nil, nil)

	_, err := mockProcess.Wait()
	testutil.AssertNil(t, err)
}

// TestMockProcess_Kill tests killing a process.
func TestMockProcess_Kill(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockProcess := NewMockProcess(ctrl)

	mockProcess.EXPECT().Kill().Return(nil)

	err := mockProcess.Kill()
	testutil.AssertNil(t, err)
}

// TestMockProcess_Release tests releasing process resources.
func TestMockProcess_Release(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockProcess := NewMockProcess(ctrl)

	mockProcess.EXPECT().Release().Return(nil)

	err := mockProcess.Release()
	testutil.AssertNil(t, err)
}

// TestMockOS_CreateOpenCloseSequence tests file lifecycle via OS.
func TestMockOS_CreateOpenCloseSequence(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOS := NewMockOS(ctrl)
	mockFile := NewMockFile(ctrl)

	gomock.InOrder(
		mockOS.EXPECT().Create("/tmp/test.txt").Return(mockFile, nil),
		mockFile.EXPECT().Write(gomock.Any()).Return(5, nil),
		mockFile.EXPECT().Close().Return(nil),
	)

	file, err := mockOS.Create("/tmp/test.txt")
	testutil.AssertNotNil(t, file)
	testutil.AssertNil(t, err)

	n, err := file.Write([]byte("hello"))
	testutil.AssertEqual(t, 5, n)
	testutil.AssertNil(t, err)

	err = file.Close()
	testutil.AssertNil(t, err)
}
