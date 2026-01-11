package mock_http

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/pdutton/go-mocks/internal/testutil"
	"go.uber.org/mock/gomock"
)

// TestMockClient_Get tests HTTP GET request.
func TestMockClient_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := NewMockClient(ctrl)
	mockResponse := NewMockResponse(ctrl)

	mockClient.EXPECT().Get("https://example.com").Return(mockResponse, nil)

	resp, err := mockClient.Get("https://example.com")
	testutil.AssertNotNil(t, resp)
	testutil.AssertNil(t, err)
}

// TestMockClient_GetError tests GET error handling.
func TestMockClient_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := NewMockClient(ctrl)

	expectedErr := errors.New("connection refused")
	mockClient.EXPECT().Get("https://invalid.example").Return(nil, expectedErr)

	resp, err := mockClient.Get("https://invalid.example")
	testutil.AssertNil(t, resp)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockClient_Post tests HTTP POST request.
func TestMockClient_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := NewMockClient(ctrl)
	mockResponse := NewMockResponse(ctrl)

	body := strings.NewReader(`{"key":"value"}`)
	mockClient.EXPECT().Post("https://example.com/api", "application/json", body).Return(mockResponse, nil)

	resp, err := mockClient.Post("https://example.com/api", "application/json", body)
	testutil.AssertNotNil(t, resp)
	testutil.AssertNil(t, err)
}

// TestMockClient_PostError tests POST error handling.
func TestMockClient_PostError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := NewMockClient(ctrl)

	expectedErr := errors.New("timeout")
	mockClient.EXPECT().Post(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, expectedErr)

	resp, err := mockClient.Post("https://example.com", "text/plain", nil)
	testutil.AssertNil(t, resp)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockClient_Head tests HTTP HEAD request.
func TestMockClient_Head(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := NewMockClient(ctrl)
	mockResponse := NewMockResponse(ctrl)

	mockClient.EXPECT().Head("https://example.com").Return(mockResponse, nil)

	resp, err := mockClient.Head("https://example.com")
	testutil.AssertNotNil(t, resp)
	testutil.AssertNil(t, err)
}

// TestMockClient_Do tests executing a request.
func TestMockClient_Do(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := NewMockClient(ctrl)
	mockRequest := NewMockRequest(ctrl)
	mockResponse := NewMockResponse(ctrl)

	mockClient.EXPECT().Do(mockRequest).Return(mockResponse, nil)

	resp, err := mockClient.Do(mockRequest)
	testutil.AssertNotNil(t, resp)
	testutil.AssertNil(t, err)
}

// TestMockClient_DoError tests Do error handling.
func TestMockClient_DoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := NewMockClient(ctrl)
	mockRequest := NewMockRequest(ctrl)

	expectedErr := errors.New("network error")
	mockClient.EXPECT().Do(mockRequest).Return(nil, expectedErr)

	resp, err := mockClient.Do(mockRequest)
	testutil.AssertNil(t, resp)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockClient_CloseIdleConnections tests closing idle connections.
func TestMockClient_CloseIdleConnections(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := NewMockClient(ctrl)

	mockClient.EXPECT().CloseIdleConnections()

	mockClient.CloseIdleConnections()
}

// TestMockResponse_StatusCode tests getting status code.
func TestMockResponse_StatusCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockResponse := NewMockResponse(ctrl)

	mockResponse.EXPECT().StatusCode().Return(200)

	code := mockResponse.StatusCode()
	testutil.AssertEqual(t, 200, code)
}

// TestMockResponse_Status tests getting status text.
func TestMockResponse_Status(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockResponse := NewMockResponse(ctrl)

	mockResponse.EXPECT().Status().Return("200 OK")

	status := mockResponse.Status()
	testutil.AssertEqual(t, "200 OK", status)
}

// TestMockResponse_Body tests getting response body.
func TestMockResponse_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockResponse := NewMockResponse(ctrl)

	var expectedBody io.ReadCloser
	mockResponse.EXPECT().Body().Return(expectedBody)

	body := mockResponse.Body()
	testutil.AssertNil(t, body)
}

// TestMockResponse_ContentLength tests getting content length.
func TestMockResponse_ContentLength(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockResponse := NewMockResponse(ctrl)

	mockResponse.EXPECT().ContentLength().Return(int64(1024))

	length := mockResponse.ContentLength()
	testutil.AssertEqual(t, int64(1024), length)
}

// TestMockResponse_Close tests closing response.
func TestMockResponse_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockResponse := NewMockResponse(ctrl)

	mockResponse.EXPECT().Close().Return(true)

	closed := mockResponse.Close()
	testutil.AssertEqual(t, true, closed)
}

// TestMockRequest_Write tests writing request.
func TestMockRequest_Write(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRequest := NewMockRequest(ctrl)

	mockRequest.EXPECT().Write(gomock.Any()).Return(nil)

	err := mockRequest.Write(nil)
	testutil.AssertNil(t, err)
}

// TestMockRequest_WriteProxy tests writing proxy request.
func TestMockRequest_WriteProxy(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRequest := NewMockRequest(ctrl)

	mockRequest.EXPECT().WriteProxy(gomock.Any()).Return(nil)

	err := mockRequest.WriteProxy(nil)
	testutil.AssertNil(t, err)
}

// TestMockHTTP_NewRequest tests creating a new request.
func TestMockHTTP_NewRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)
	mockRequest := NewMockRequest(ctrl)

	mockHTTP.EXPECT().NewRequest("GET", "https://example.com", nil).Return(mockRequest, nil)

	req, err := mockHTTP.NewRequest("GET", "https://example.com", nil)
	testutil.AssertNotNil(t, req)
	testutil.AssertNil(t, err)
}

// TestMockHTTP_NewRequestError tests NewRequest error handling.
func TestMockHTTP_NewRequestError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(ctrl)

	expectedErr := errors.New("invalid URL")
	mockHTTP.EXPECT().NewRequest("GET", "://invalid", nil).Return(nil, expectedErr)

	req, err := mockHTTP.NewRequest("GET", "://invalid", nil)
	testutil.AssertNil(t, req)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockClient_GetPostSequence tests GET then POST sequence.
func TestMockClient_GetPostSequence(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := NewMockClient(ctrl)
	mockResponse1 := NewMockResponse(ctrl)
	mockResponse2 := NewMockResponse(ctrl)

	gomock.InOrder(
		mockClient.EXPECT().Get("https://example.com").Return(mockResponse1, nil),
		mockResponse1.EXPECT().StatusCode().Return(200),
		mockClient.EXPECT().Post("https://example.com/api", gomock.Any(), gomock.Any()).Return(mockResponse2, nil),
		mockResponse2.EXPECT().StatusCode().Return(201),
	)

	resp1, err1 := mockClient.Get("https://example.com")
	testutil.AssertNotNil(t, resp1)
	testutil.AssertNil(t, err1)

	code1 := resp1.StatusCode()
	testutil.AssertEqual(t, 200, code1)

	resp2, err2 := mockClient.Post("https://example.com/api", "application/json", nil)
	testutil.AssertNotNil(t, resp2)
	testutil.AssertNil(t, err2)

	code2 := resp2.StatusCode()
	testutil.AssertEqual(t, 201, code2)
}

// TestMockResponse_StatusCodes tests different status codes.
func TestMockResponse_StatusCodes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockResponse := NewMockResponse(ctrl)

	testCases := []struct {
		code   int
		status string
	}{
		{200, "200 OK"},
		{201, "201 Created"},
		{400, "400 Bad Request"},
		{404, "404 Not Found"},
		{500, "500 Internal Server Error"},
	}

	for _, tc := range testCases {
		mockResponse.EXPECT().StatusCode().Return(tc.code)
		mockResponse.EXPECT().Status().Return(tc.status)

		code := mockResponse.StatusCode()
		testutil.AssertEqual(t, tc.code, code)

		status := mockResponse.Status()
		testutil.AssertEqual(t, tc.status, status)
	}
}
