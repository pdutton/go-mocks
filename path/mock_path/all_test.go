package mock_path

import (
	"errors"
	"testing"

	"github.com/pdutton/go-mocks/internal/testutil"
	"go.uber.org/mock/gomock"
)

// TestMockPath_Base tests base name extraction.
func TestMockPath_Base(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Base("/foo/bar/file.txt").Return("file.txt")

	result := mockPath.Base("/foo/bar/file.txt")
	testutil.AssertEqual(t, "file.txt", result)
}

// TestMockPath_BaseRoot tests base of root path.
func TestMockPath_BaseRoot(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Base("/").Return("/")

	result := mockPath.Base("/")
	testutil.AssertEqual(t, "/", result)
}

// TestMockPath_Clean tests path cleaning.
func TestMockPath_Clean(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Clean("/foo//bar/../baz").Return("/foo/baz")

	result := mockPath.Clean("/foo//bar/../baz")
	testutil.AssertEqual(t, "/foo/baz", result)
}

// TestMockPath_CleanDot tests cleaning current directory.
func TestMockPath_CleanDot(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Clean("./foo").Return("foo")

	result := mockPath.Clean("./foo")
	testutil.AssertEqual(t, "foo", result)
}

// TestMockPath_Dir tests directory extraction.
func TestMockPath_Dir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Dir("/foo/bar/file.txt").Return("/foo/bar")

	result := mockPath.Dir("/foo/bar/file.txt")
	testutil.AssertEqual(t, "/foo/bar", result)
}

// TestMockPath_DirRoot tests directory of root.
func TestMockPath_DirRoot(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Dir("/").Return("/")

	result := mockPath.Dir("/")
	testutil.AssertEqual(t, "/", result)
}

// TestMockPath_Ext tests extension extraction.
func TestMockPath_Ext(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Ext("file.txt").Return(".txt")

	result := mockPath.Ext("file.txt")
	testutil.AssertEqual(t, ".txt", result)
}

// TestMockPath_ExtNone tests file with no extension.
func TestMockPath_ExtNone(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Ext("file").Return("")

	result := mockPath.Ext("file")
	testutil.AssertEqual(t, "", result)
}

// TestMockPath_IsAbs tests absolute path check.
func TestMockPath_IsAbs(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().IsAbs("/foo/bar").Return(true)

	result := mockPath.IsAbs("/foo/bar")
	testutil.AssertEqual(t, true, result)
}

// TestMockPath_IsAbsRelative tests relative path check.
func TestMockPath_IsAbsRelative(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().IsAbs("foo/bar").Return(false)

	result := mockPath.IsAbs("foo/bar")
	testutil.AssertEqual(t, false, result)
}

// TestMockPath_Join tests path joining.
func TestMockPath_Join(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Join("foo", "bar", "baz").Return("foo/bar/baz")

	result := mockPath.Join("foo", "bar", "baz")
	testutil.AssertEqual(t, "foo/bar/baz", result)
}

// TestMockPath_JoinEmpty tests joining empty paths.
func TestMockPath_JoinEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Join().Return("")

	result := mockPath.Join()
	testutil.AssertEqual(t, "", result)
}

// TestMockPath_JoinSingle tests joining single element.
func TestMockPath_JoinSingle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Join("foo").Return("foo")

	result := mockPath.Join("foo")
	testutil.AssertEqual(t, "foo", result)
}

// TestMockPath_Match tests pattern matching.
func TestMockPath_Match(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Match("*.txt", "file.txt").Return(true, nil)

	matched, err := mockPath.Match("*.txt", "file.txt")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, true, matched)
}

// TestMockPath_MatchNoMatch tests non-matching pattern.
func TestMockPath_MatchNoMatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Match("*.txt", "file.md").Return(false, nil)

	matched, err := mockPath.Match("*.txt", "file.md")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, false, matched)
}

// TestMockPath_MatchError tests pattern match error.
func TestMockPath_MatchError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	expectedErr := errors.New("bad pattern")
	mockPath.EXPECT().Match("[", "file.txt").Return(false, expectedErr)

	matched, err := mockPath.Match("[", "file.txt")
	testutil.AssertEqual(t, false, matched)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockPath_Split tests path splitting.
func TestMockPath_Split(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Split("/foo/bar/file.txt").Return("/foo/bar/", "file.txt")

	dir, file := mockPath.Split("/foo/bar/file.txt")
	testutil.AssertEqual(t, "/foo/bar/", dir)
	testutil.AssertEqual(t, "file.txt", file)
}

// TestMockPath_SplitRoot tests splitting root path.
func TestMockPath_SplitRoot(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	mockPath.EXPECT().Split("/").Return("/", "")

	dir, file := mockPath.Split("/")
	testutil.AssertEqual(t, "/", dir)
	testutil.AssertEqual(t, "", file)
}

// TestMockPath_MultipleOperations tests sequence of path operations.
func TestMockPath_MultipleOperations(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPath := NewMockPath(ctrl)

	gomock.InOrder(
		mockPath.EXPECT().Join("foo", "bar").Return("foo/bar"),
		mockPath.EXPECT().Clean("foo/bar").Return("foo/bar"),
		mockPath.EXPECT().Base("foo/bar").Return("bar"),
		mockPath.EXPECT().Dir("foo/bar").Return("foo"),
		mockPath.EXPECT().Ext("bar").Return(""),
	)

	joined := mockPath.Join("foo", "bar")
	testutil.AssertEqual(t, "foo/bar", joined)

	cleaned := mockPath.Clean(joined)
	testutil.AssertEqual(t, "foo/bar", cleaned)

	base := mockPath.Base(cleaned)
	testutil.AssertEqual(t, "bar", base)

	dir := mockPath.Dir(cleaned)
	testutil.AssertEqual(t, "foo", dir)

	ext := mockPath.Ext(base)
	testutil.AssertEqual(t, "", ext)
}
