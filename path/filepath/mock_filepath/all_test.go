package mock_filepath

import (
	"errors"
	"testing"

	"github.com/pdutton/go-mocks/internal/testutil"
	"go.uber.org/mock/gomock"
)

// TestMockFilePath_Abs tests absolute path resolution.
func TestMockFilePath_Abs(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().Abs("file.txt").Return("/home/user/file.txt", nil)

	result, err := mockFP.Abs("file.txt")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, "/home/user/file.txt", result)
}

// TestMockFilePath_AbsError tests Abs error handling.
func TestMockFilePath_AbsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	expectedErr := errors.New("abs error")
	mockFP.EXPECT().Abs("invalid").Return("", expectedErr)

	result, err := mockFP.Abs("invalid")
	testutil.AssertEqual(t, "", result)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockFilePath_Base tests base name extraction.
func TestMockFilePath_Base(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().Base("/foo/bar/file.txt").Return("file.txt")

	result := mockFP.Base("/foo/bar/file.txt")
	testutil.AssertEqual(t, "file.txt", result)
}

// TestMockFilePath_Clean tests path cleaning.
func TestMockFilePath_Clean(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().Clean("/foo//bar/../baz").Return("/foo/baz")

	result := mockFP.Clean("/foo//bar/../baz")
	testutil.AssertEqual(t, "/foo/baz", result)
}

// TestMockFilePath_Dir tests directory extraction.
func TestMockFilePath_Dir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().Dir("/foo/bar/file.txt").Return("/foo/bar")

	result := mockFP.Dir("/foo/bar/file.txt")
	testutil.AssertEqual(t, "/foo/bar", result)
}

// TestMockFilePath_Ext tests extension extraction.
func TestMockFilePath_Ext(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().Ext("file.txt").Return(".txt")

	result := mockFP.Ext("file.txt")
	testutil.AssertEqual(t, ".txt", result)
}

// TestMockFilePath_Glob tests pattern matching.
func TestMockFilePath_Glob(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	matches := []string{"file1.txt", "file2.txt"}
	mockFP.EXPECT().Glob("*.txt").Return(matches, nil)

	result, err := mockFP.Glob("*.txt")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, 2, len(result))
	testutil.AssertEqual(t, "file1.txt", result[0])
}

// TestMockFilePath_GlobError tests Glob error handling.
func TestMockFilePath_GlobError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	expectedErr := errors.New("glob error")
	mockFP.EXPECT().Glob("[").Return([]string(nil), expectedErr)

	result, err := mockFP.Glob("[")
	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
	testutil.AssertError(t, expectedErr, err)
}

// TestMockFilePath_IsAbs tests absolute path check.
func TestMockFilePath_IsAbs(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().IsAbs("/foo/bar").Return(true)

	result := mockFP.IsAbs("/foo/bar")
	testutil.AssertEqual(t, true, result)
}

// TestMockFilePath_IsLocal tests local path check.
func TestMockFilePath_IsLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().IsLocal("foo/bar").Return(true)

	result := mockFP.IsLocal("foo/bar")
	testutil.AssertEqual(t, true, result)
}

// TestMockFilePath_Join tests path joining.
func TestMockFilePath_Join(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().Join("foo", "bar", "baz").Return("foo/bar/baz")

	result := mockFP.Join("foo", "bar", "baz")
	testutil.AssertEqual(t, "foo/bar/baz", result)
}

// TestMockFilePath_Match tests pattern matching.
func TestMockFilePath_Match(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().Match("*.txt", "file.txt").Return(true, nil)

	matched, err := mockFP.Match("*.txt", "file.txt")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, true, matched)
}

// TestMockFilePath_Rel tests relative path calculation.
func TestMockFilePath_Rel(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().Rel("/foo", "/foo/bar/baz").Return("bar/baz", nil)

	result, err := mockFP.Rel("/foo", "/foo/bar/baz")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, "bar/baz", result)
}

// TestMockFilePath_RelError tests Rel error handling.
func TestMockFilePath_RelError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	expectedErr := errors.New("rel error")
	mockFP.EXPECT().Rel("/foo", "/bar").Return("", expectedErr)

	result, err := mockFP.Rel("/foo", "/bar")
	testutil.AssertEqual(t, "", result)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockFilePath_Split tests path splitting.
func TestMockFilePath_Split(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().Split("/foo/bar/file.txt").Return("/foo/bar/", "file.txt")

	dir, file := mockFP.Split("/foo/bar/file.txt")
	testutil.AssertEqual(t, "/foo/bar/", dir)
	testutil.AssertEqual(t, "file.txt", file)
}

// TestMockFilePath_SplitList tests path list splitting.
func TestMockFilePath_SplitList(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().SplitList("/foo:/bar:/baz").Return([]string{"/foo", "/bar", "/baz"})

	result := mockFP.SplitList("/foo:/bar:/baz")
	testutil.AssertEqual(t, 3, len(result))
	testutil.AssertEqual(t, "/foo", result[0])
}

// TestMockFilePath_ToSlash tests converting to forward slashes.
func TestMockFilePath_ToSlash(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().ToSlash("foo\\bar\\baz").Return("foo/bar/baz")

	result := mockFP.ToSlash("foo\\bar\\baz")
	testutil.AssertEqual(t, "foo/bar/baz", result)
}

// TestMockFilePath_FromSlash tests converting from forward slashes.
func TestMockFilePath_FromSlash(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().FromSlash("foo/bar/baz").Return("foo\\bar\\baz")

	result := mockFP.FromSlash("foo/bar/baz")
	testutil.AssertEqual(t, "foo\\bar\\baz", result)
}

// TestMockFilePath_VolumeName tests volume name extraction.
func TestMockFilePath_VolumeName(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().VolumeName("C:\\foo\\bar").Return("C:")

	result := mockFP.VolumeName("C:\\foo\\bar")
	testutil.AssertEqual(t, "C:", result)
}

// TestMockFilePath_EvalSymlinks tests symlink evaluation.
func TestMockFilePath_EvalSymlinks(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().EvalSymlinks("/link").Return("/target", nil)

	result, err := mockFP.EvalSymlinks("/link")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, "/target", result)
}

// TestMockFilePath_Walk tests directory walking.
func TestMockFilePath_Walk(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().Walk("/foo", gomock.Any()).Return(nil)

	err := mockFP.Walk("/foo", nil)
	testutil.AssertNil(t, err)
}

// TestMockFilePath_WalkError tests Walk error handling.
func TestMockFilePath_WalkError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	expectedErr := errors.New("walk error")
	mockFP.EXPECT().Walk("/foo", gomock.Any()).Return(expectedErr)

	err := mockFP.Walk("/foo", nil)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockFilePath_WalkDir tests WalkDir operation.
func TestMockFilePath_WalkDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().WalkDir("/foo", gomock.Any()).Return(nil)

	err := mockFP.WalkDir("/foo", nil)
	testutil.AssertNil(t, err)
}

// TestMockFilePath_Localize tests path localization.
func TestMockFilePath_Localize(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	mockFP.EXPECT().Localize("foo/bar").Return("foo\\bar", nil)

	result, err := mockFP.Localize("foo/bar")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, "foo\\bar", result)
}

// TestMockDirEntry_Name tests directory entry name.
func TestMockDirEntry_Name(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockEntry := NewMockDirEntry(ctrl)

	mockEntry.EXPECT().Name().Return("file.txt")

	result := mockEntry.Name()
	testutil.AssertEqual(t, "file.txt", result)
}

// TestMockDirEntry_IsDir tests directory check.
func TestMockDirEntry_IsDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockEntry := NewMockDirEntry(ctrl)

	mockEntry.EXPECT().IsDir().Return(false)

	result := mockEntry.IsDir()
	testutil.AssertEqual(t, false, result)
}

// TestMockDirEntry_Info tests directory entry info.
func TestMockDirEntry_Info(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockEntry := NewMockDirEntry(ctrl)
	mockFileInfo := NewMockFileInfo(ctrl)

	mockEntry.EXPECT().Info().Return(mockFileInfo, nil)

	info, err := mockEntry.Info()
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, info)
}

// TestMockFileInfo_Name tests FileInfo name.
func TestMockFileInfo_Name(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockInfo := NewMockFileInfo(ctrl)

	mockInfo.EXPECT().Name().Return("file.txt")

	result := mockInfo.Name()
	testutil.AssertEqual(t, "file.txt", result)
}

// TestMockFileInfo_Size tests file size.
func TestMockFileInfo_Size(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockInfo := NewMockFileInfo(ctrl)

	mockInfo.EXPECT().Size().Return(int64(1024))

	result := mockInfo.Size()
	testutil.AssertEqual(t, int64(1024), result)
}

// TestMockFileInfo_IsDir tests directory check.
func TestMockFileInfo_IsDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockInfo := NewMockFileInfo(ctrl)

	mockInfo.EXPECT().IsDir().Return(false)

	result := mockInfo.IsDir()
	testutil.AssertEqual(t, false, result)
}

// TestMockFilePath_PathOperationsSequence tests sequence of operations.
func TestMockFilePath_PathOperationsSequence(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFP := NewMockFilePath(ctrl)

	gomock.InOrder(
		mockFP.EXPECT().Join("foo", "bar").Return("foo/bar"),
		mockFP.EXPECT().Clean("foo/bar").Return("foo/bar"),
		mockFP.EXPECT().Abs("foo/bar").Return("/home/user/foo/bar", nil),
		mockFP.EXPECT().Base("/home/user/foo/bar").Return("bar"),
	)

	joined := mockFP.Join("foo", "bar")
	testutil.AssertEqual(t, "foo/bar", joined)

	cleaned := mockFP.Clean(joined)
	testutil.AssertEqual(t, "foo/bar", cleaned)

	abs, err := mockFP.Abs(cleaned)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, "/home/user/foo/bar", abs)

	base := mockFP.Base(abs)
	testutil.AssertEqual(t, "bar", base)
}
