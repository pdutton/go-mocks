package mock_http

import (
	"context"
	"errors"
	"testing"

	"github.com/pdutton/go-mocks/internal/testutil"
	"go.uber.org/mock/gomock"
)

// TestMockServer_ListenAndServe tests starting HTTP server.
func TestMockServer_ListenAndServe(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockServer := NewMockServer(ctrl)

	mockServer.EXPECT().ListenAndServe().Return(nil)

	err := mockServer.ListenAndServe()
	testutil.AssertNil(t, err)
}

// TestMockServer_ListenAndServeError tests server start error.
func TestMockServer_ListenAndServeError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockServer := NewMockServer(ctrl)

	expectedErr := errors.New("address already in use")
	mockServer.EXPECT().ListenAndServe().Return(expectedErr)

	err := mockServer.ListenAndServe()
	testutil.AssertError(t, expectedErr, err)
}

// TestMockServer_ListenAndServeTLS tests starting HTTPS server.
func TestMockServer_ListenAndServeTLS(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockServer := NewMockServer(ctrl)

	mockServer.EXPECT().ListenAndServeTLS("cert.pem", "key.pem").Return(nil)

	err := mockServer.ListenAndServeTLS("cert.pem", "key.pem")
	testutil.AssertNil(t, err)
}

// TestMockServer_Close tests closing server.
func TestMockServer_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockServer := NewMockServer(ctrl)

	mockServer.EXPECT().Close().Return(nil)

	err := mockServer.Close()
	testutil.AssertNil(t, err)
}

// TestMockServer_Shutdown tests graceful shutdown.
func TestMockServer_Shutdown(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockServer := NewMockServer(ctrl)

	ctx := context.Background()
	mockServer.EXPECT().Shutdown(ctx).Return(nil)

	err := mockServer.Shutdown(ctx)
	testutil.AssertNil(t, err)
}

// TestMockServer_ShutdownError tests shutdown error handling.
func TestMockServer_ShutdownError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockServer := NewMockServer(ctrl)

	expectedErr := errors.New("shutdown timeout")
	mockServer.EXPECT().Shutdown(gomock.Any()).Return(expectedErr)

	err := mockServer.Shutdown(context.Background())
	testutil.AssertError(t, expectedErr, err)
}

// TestMockServer_SetKeepAlivesEnabled tests enabling keep-alives.
func TestMockServer_SetKeepAlivesEnabled(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockServer := NewMockServer(ctrl)

	mockServer.EXPECT().SetKeepAlivesEnabled(true)

	mockServer.SetKeepAlivesEnabled(true)
}

// TestMockServer_RegisterOnShutdown tests registering shutdown callback.
func TestMockServer_RegisterOnShutdown(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockServer := NewMockServer(ctrl)

	callback := func() {}
	mockServer.EXPECT().RegisterOnShutdown(gomock.Any())

	mockServer.RegisterOnShutdown(callback)
}

// TestMockServer_Lifecycle tests server lifecycle.
func TestMockServer_Lifecycle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockServer := NewMockServer(ctrl)

	gomock.InOrder(
		mockServer.EXPECT().SetKeepAlivesEnabled(true),
		mockServer.EXPECT().RegisterOnShutdown(gomock.Any()),
		mockServer.EXPECT().ListenAndServe().Return(nil),
		mockServer.EXPECT().Shutdown(gomock.Any()).Return(nil),
	)

	mockServer.SetKeepAlivesEnabled(true)
	mockServer.RegisterOnShutdown(func() {})
	err1 := mockServer.ListenAndServe()
	testutil.AssertNil(t, err1)

	err2 := mockServer.Shutdown(context.Background())
	testutil.AssertNil(t, err2)
}

// TestMockHTTP_ListenAndServe tests HTTP ListenAndServe.
func TestMockHTTP_ListenAndServe(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)

	mockHTTP.EXPECT().ListenAndServe(":8080", nil).Return(nil)

	err := mockHTTP.ListenAndServe(":8080", nil)
	testutil.AssertNil(t, err)
}

// TestMockHTTP_ListenAndServeTLS tests HTTPS ListenAndServeTLS.
func TestMockHTTP_ListenAndServeTLS(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)

	mockHTTP.EXPECT().ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil).Return(nil)

	err := mockHTTP.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	testutil.AssertNil(t, err)
}

// TestMockHTTP_Handle tests registering handler.
func TestMockHTTP_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)

	mockHTTP.EXPECT().Handle("/", nil)

	mockHTTP.Handle("/", nil)
}

// TestMockHTTP_HandleFunc tests registering handler function.
func TestMockHTTP_HandleFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)

	mockHTTP.EXPECT().HandleFunc("/api", gomock.Any())

	mockHTTP.HandleFunc("/api", nil)
}

// TestMockHTTP_CanonicalHeaderKey tests canonicalizing header key.
func TestMockHTTP_CanonicalHeaderKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)

	mockHTTP.EXPECT().CanonicalHeaderKey("content-type").Return("Content-Type")

	result := mockHTTP.CanonicalHeaderKey("content-type")
	testutil.AssertEqual(t, "Content-Type", result)
}

// TestMockHTTP_DetectContentType tests detecting content type.
func TestMockHTTP_DetectContentType(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)

	data := []byte("hello")
	mockHTTP.EXPECT().DetectContentType(data).Return("text/plain")

	contentType := mockHTTP.DetectContentType(data)
	testutil.AssertEqual(t, "text/plain", contentType)
}

// TestMockHTTP_Error tests sending error response.
func TestMockHTTP_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)

	mockHTTP.EXPECT().Error(nil, "Not Found", 404)

	mockHTTP.Error(nil, "Not Found", 404)
}

// TestMockHTTP_NotFound tests sending 404 response.
func TestMockHTTP_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)

	mockHTTP.EXPECT().NotFound(nil, nil)

	mockHTTP.NotFound(nil, nil)
}

// TestMockHTTP_Redirect tests redirecting request.
func TestMockHTTP_Redirect(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)

	mockHTTP.EXPECT().Redirect(nil, nil, "https://example.com", 302)

	mockHTTP.Redirect(nil, nil, "https://example.com", 302)
}

// TestMockHTTP_ServeFile tests serving file.
func TestMockHTTP_ServeFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)

	mockHTTP.EXPECT().ServeFile(nil, nil, "/tmp/file.txt")

	mockHTTP.ServeFile(nil, nil, "/tmp/file.txt")
}

// TestMockHTTP_ParseHTTPVersion tests parsing HTTP version.
func TestMockHTTP_ParseHTTPVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)

	mockHTTP.EXPECT().ParseHTTPVersion("HTTP/1.1").Return(1, 1, true)

	major, minor, ok := mockHTTP.ParseHTTPVersion("HTTP/1.1")
	testutil.AssertEqual(t, 1, major)
	testutil.AssertEqual(t, 1, minor)
	testutil.AssertEqual(t, true, ok)
}
