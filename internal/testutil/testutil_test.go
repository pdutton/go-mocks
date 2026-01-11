package testutil

import (
	"bytes"
	"errors"
	"testing"
)

func TestAssertPanic_Panics(t *testing.T) {
	// This test expects the function to panic
	AssertPanic(t, func() {
		panic("expected panic")
	}, "")
}

func TestAssertError_BothNil(t *testing.T) {
	AssertError(t, nil, nil)
}

func TestAssertError_ExpectedNil(t *testing.T) {
	// Create a mock testing.T to capture the error
	mockT := &testing.T{}
	AssertError(mockT, nil, errors.New("unexpected"))
	if !mockT.Failed() {
		t.Error("expected test to fail when actual error is not nil")
	}
}

func TestAssertError_ActualNil(t *testing.T) {
	mockT := &testing.T{}
	AssertError(mockT, errors.New("expected"), nil)
	if !mockT.Failed() {
		t.Error("expected test to fail when actual error is nil")
	}
}

func TestAssertError_SameError(t *testing.T) {
	err := errors.New("test error")
	AssertError(t, err, err)
}

func TestAssertError_SameMessage(t *testing.T) {
	err1 := errors.New("test error")
	err2 := errors.New("test error")
	AssertError(t, err1, err2)
}

func TestAssertBytes_Equal(t *testing.T) {
	b1 := []byte{1, 2, 3}
	b2 := []byte{1, 2, 3}
	AssertBytes(t, b1, b2)
}

func TestAssertBytes_NotEqual(t *testing.T) {
	mockT := &testing.T{}
	b1 := []byte{1, 2, 3}
	b2 := []byte{4, 5, 6}
	AssertBytes(mockT, b1, b2)
	if !mockT.Failed() {
		t.Error("expected test to fail when byte slices differ")
	}
}

func TestAssertEqual_Equal(t *testing.T) {
	AssertEqual(t, 42, 42)
	AssertEqual(t, "test", "test")
	AssertEqual(t, true, true)
}

func TestAssertEqual_NotEqual(t *testing.T) {
	mockT := &testing.T{}
	AssertEqual(mockT, 42, 43)
	if !mockT.Failed() {
		t.Error("expected test to fail when values are not equal")
	}
}

func TestAssertNotEqual_NotEqual(t *testing.T) {
	AssertNotEqual(t, 42, 43)
	AssertNotEqual(t, "foo", "bar")
}

func TestAssertNotEqual_Equal(t *testing.T) {
	mockT := &testing.T{}
	AssertNotEqual(mockT, 42, 42)
	if !mockT.Failed() {
		t.Error("expected test to fail when values are equal")
	}
}

func TestAssertNil_Nil(t *testing.T) {
	var err error
	AssertNil(t, err)
	AssertNil(t, nil)
}

func TestAssertNil_NotNil(t *testing.T) {
	mockT := &testing.T{}
	AssertNil(mockT, errors.New("not nil"))
	if !mockT.Failed() {
		t.Error("expected test to fail when value is not nil")
	}
}

func TestAssertNotNil_NotNil(t *testing.T) {
	AssertNotNil(t, errors.New("not nil"))
	AssertNotNil(t, 42)
}

func TestAssertNotNil_Nil(t *testing.T) {
	mockT := &testing.T{}
	var err error
	AssertNotNil(mockT, err)
	if !mockT.Failed() {
		t.Error("expected test to fail when value is nil")
	}
}

func TestRandomBytes_Length(t *testing.T) {
	sizes := []int{0, 1, 10, 100, 1024}
	for _, size := range sizes {
		b := RandomBytes(size)
		if len(b) != size {
			t.Errorf("expected %d bytes, got %d", size, len(b))
		}
	}
}

func TestRandomBytes_Randomness(t *testing.T) {
	// Generate two random byte slices and ensure they're different
	// (extremely unlikely to be the same for any reasonable size)
	b1 := RandomBytes(32)
	b2 := RandomBytes(32)
	if bytes.Equal(b1, b2) {
		t.Error("expected random bytes to be different, but they're identical")
	}
}

func TestFormatMsg_Empty(t *testing.T) {
	result := formatMsg("")
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestFormatMsg_NonEmpty(t *testing.T) {
	result := formatMsg("test message")
	expected := " (test message)"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
