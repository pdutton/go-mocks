// Package testutil provides shared testing utilities for mock tests.
package testutil

import (
	"bytes"
	"crypto/rand"
	"errors"
	"testing"
)

// AssertPanic asserts that a function panics with an optional message check.
func AssertPanic(t *testing.T, f func(), msg string) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic%s, but function did not panic", formatMsg(msg))
		}
	}()
	f()
}

// AssertError asserts that the actual error matches the expected error.
func AssertError(t *testing.T, expected, actual error) {
	t.Helper()
	if expected == nil && actual == nil {
		return
	}
	if expected == nil && actual != nil {
		t.Errorf("expected no error, got %v", actual)
		return
	}
	if expected != nil && actual == nil {
		t.Errorf("expected error %v, got nil", expected)
		return
	}
	if !errors.Is(actual, expected) && expected.Error() != actual.Error() {
		t.Errorf("expected error %v, got %v", expected, actual)
	}
}

// AssertBytes compares byte slices and reports differences.
func AssertBytes(t *testing.T, expected, actual []byte) {
	t.Helper()
	if !bytes.Equal(expected, actual) {
		t.Errorf("byte slices differ:\nexpected: %v\nactual:   %v", expected, actual)
	}
}

// AssertEqual compares two values and reports if they're not equal.
func AssertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if expected != actual {
		t.Errorf("values differ:\nexpected: %v\nactual:   %v", expected, actual)
	}
}

// AssertNotEqual ensures two values are different.
func AssertNotEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if expected == actual {
		t.Errorf("expected values to differ, but both are: %v", expected)
	}
}

// AssertNil checks that a value is nil.
func AssertNil(t *testing.T, value any) {
	t.Helper()
	if value != nil {
		t.Errorf("expected nil, got %v", value)
	}
}

// AssertNotNil checks that a value is not nil.
func AssertNotNil(t *testing.T, value any) {
	t.Helper()
	if value == nil {
		t.Error("expected non-nil value, got nil")
	}
}

// RandomBytes generates a random byte slice of the specified length.
func RandomBytes(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic("failed to generate random bytes: " + err.Error())
	}
	return b
}

// formatMsg formats an optional message for error output.
func formatMsg(msg string) string {
	if msg == "" {
		return ""
	}
	return " (" + msg + ")"
}
