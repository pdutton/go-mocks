package mock_signal

import (
	"context"
	"os"
	"syscall"
	"testing"

	"github.com/pdutton/go-mocks/internal/testutil"
	"go.uber.org/mock/gomock"
)

// TestMockSignal_Notify tests signal notification setup.
func TestMockSignal_Notify(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	c := make(chan os.Signal, 1)
	mockSignal.EXPECT().Notify(c, syscall.SIGINT, syscall.SIGTERM)

	mockSignal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
}

// TestMockSignal_NotifyMultiple tests notifying multiple signals.
func TestMockSignal_NotifyMultiple(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	c := make(chan os.Signal, 1)
	mockSignal.EXPECT().Notify(c, gomock.Any())

	mockSignal.Notify(c, syscall.SIGINT)
}

// TestMockSignal_Stop tests stopping signal notifications.
func TestMockSignal_Stop(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	c := make(chan os.Signal, 1)
	mockSignal.EXPECT().Stop(c)

	mockSignal.Stop(c)
}

// TestMockSignal_NotifyStop tests notify then stop sequence.
func TestMockSignal_NotifyStop(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	c := make(chan os.Signal, 1)

	gomock.InOrder(
		mockSignal.EXPECT().Notify(c, syscall.SIGINT),
		mockSignal.EXPECT().Stop(c),
	)

	mockSignal.Notify(c, syscall.SIGINT)
	mockSignal.Stop(c)
}

// TestMockSignal_Ignore tests ignoring signals.
func TestMockSignal_Ignore(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	mockSignal.EXPECT().Ignore(syscall.SIGHUP)

	mockSignal.Ignore(syscall.SIGHUP)
}

// TestMockSignal_IgnoreMultiple tests ignoring multiple signals.
func TestMockSignal_IgnoreMultiple(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	mockSignal.EXPECT().Ignore(syscall.SIGHUP, syscall.SIGPIPE)

	mockSignal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
}

// TestMockSignal_Ignored tests checking if signal is ignored.
func TestMockSignal_Ignored(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	mockSignal.EXPECT().Ignored(syscall.SIGHUP).Return(true)

	result := mockSignal.Ignored(syscall.SIGHUP)
	testutil.AssertEqual(t, true, result)
}

// TestMockSignal_IgnoredNotIgnored tests signal not ignored.
func TestMockSignal_IgnoredNotIgnored(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	mockSignal.EXPECT().Ignored(syscall.SIGINT).Return(false)

	result := mockSignal.Ignored(syscall.SIGINT)
	testutil.AssertEqual(t, false, result)
}

// TestMockSignal_Reset tests resetting signal handlers.
func TestMockSignal_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	mockSignal.EXPECT().Reset(syscall.SIGINT, syscall.SIGTERM)

	mockSignal.Reset(syscall.SIGINT, syscall.SIGTERM)
}

// TestMockSignal_ResetAll tests resetting all signals.
func TestMockSignal_ResetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	mockSignal.EXPECT().Reset()

	mockSignal.Reset()
}

// TestMockSignal_NotifyContext tests context-based notification.
func TestMockSignal_NotifyContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	parent := context.Background()
	ctx := context.Background()
	cancel := func() {}

	mockSignal.EXPECT().NotifyContext(parent, syscall.SIGINT).Return(ctx, cancel)

	resultCtx, resultCancel := mockSignal.NotifyContext(parent, syscall.SIGINT)
	testutil.AssertNotNil(t, resultCtx)
	testutil.AssertNotNil(t, resultCancel)
}

// TestMockSignal_NotifyContextMultiple tests NotifyContext with multiple signals.
func TestMockSignal_NotifyContextMultiple(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	parent := context.Background()
	ctx := context.Background()
	cancel := func() {}

	mockSignal.EXPECT().NotifyContext(parent, syscall.SIGINT, syscall.SIGTERM).Return(ctx, cancel)

	resultCtx, resultCancel := mockSignal.NotifyContext(parent, syscall.SIGINT, syscall.SIGTERM)
	testutil.AssertNotNil(t, resultCtx)
	testutil.AssertNotNil(t, resultCancel)
}

// TestMockSignal_IgnoreResetSequence tests ignore then reset.
func TestMockSignal_IgnoreResetSequence(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSignal := NewMockSignal(ctrl)

	gomock.InOrder(
		mockSignal.EXPECT().Ignore(syscall.SIGHUP),
		mockSignal.EXPECT().Ignored(syscall.SIGHUP).Return(true),
		mockSignal.EXPECT().Reset(syscall.SIGHUP),
		mockSignal.EXPECT().Ignored(syscall.SIGHUP).Return(false),
	)

	mockSignal.Ignore(syscall.SIGHUP)

	ignored1 := mockSignal.Ignored(syscall.SIGHUP)
	testutil.AssertEqual(t, true, ignored1)

	mockSignal.Reset(syscall.SIGHUP)

	ignored2 := mockSignal.Ignored(syscall.SIGHUP)
	testutil.AssertEqual(t, false, ignored2)
}
