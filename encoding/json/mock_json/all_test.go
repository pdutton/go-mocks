package mock_json

import (
	"errors"
	"testing"

	"github.com/pdutton/go-mocks/internal/testutil"
	"go.uber.org/mock/gomock"
)

// TestMockJSON_Marshal tests basic marshaling.
func TestMockJSON_Marshal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJSON := NewMockJSON(ctrl)

	input := map[string]string{"key": "value"}
	expected := []byte(`{"key":"value"}`)

	mockJSON.EXPECT().Marshal(input).Return(expected, nil)

	result, err := mockJSON.Marshal(input)
	testutil.AssertNil(t, err)
	testutil.AssertBytes(t, expected, result)
}

// TestMockJSON_MarshalError tests marshal error handling.
func TestMockJSON_MarshalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJSON := NewMockJSON(ctrl)

	expectedErr := errors.New("marshal error")
	mockJSON.EXPECT().Marshal(gomock.Any()).Return([]byte(nil), expectedErr)

	result, err := mockJSON.Marshal("invalid")
	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
	testutil.AssertError(t, expectedErr, err)
}

// TestMockJSON_Unmarshal tests basic unmarshaling.
func TestMockJSON_Unmarshal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJSON := NewMockJSON(ctrl)

	data := []byte(`{"key":"value"}`)
	var v map[string]string

	mockJSON.EXPECT().Unmarshal(data, gomock.Any()).Return(nil)

	err := mockJSON.Unmarshal(data, &v)
	testutil.AssertNil(t, err)
}

// TestMockJSON_UnmarshalError tests unmarshal error handling.
func TestMockJSON_UnmarshalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJSON := NewMockJSON(ctrl)

	data := []byte(`invalid json`)
	expectedErr := errors.New("unmarshal error")

	mockJSON.EXPECT().Unmarshal(data, gomock.Any()).Return(expectedErr)

	var v interface{}
	err := mockJSON.Unmarshal(data, &v)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockJSON_Valid tests JSON validity check.
func TestMockJSON_Valid(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJSON := NewMockJSON(ctrl)

	data := []byte(`{"key":"value"}`)
	mockJSON.EXPECT().Valid(data).Return(true)

	valid := mockJSON.Valid(data)
	testutil.AssertEqual(t, true, valid)
}

// TestMockJSON_ValidInvalid tests invalid JSON.
func TestMockJSON_ValidInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJSON := NewMockJSON(ctrl)

	data := []byte(`invalid`)
	mockJSON.EXPECT().Valid(data).Return(false)

	valid := mockJSON.Valid(data)
	testutil.AssertEqual(t, false, valid)
}

// TestMockDecoder_Decode tests single decode operation.
func TestMockDecoder_Decode(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDecoder := NewMockDecoder(ctrl)

	mockDecoder.EXPECT().Decode(gomock.Any()).Return(nil)

	var v map[string]string
	err := mockDecoder.Decode(&v)
	testutil.AssertNil(t, err)
}

// TestMockDecoder_DecodeError tests decode error handling.
func TestMockDecoder_DecodeError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDecoder := NewMockDecoder(ctrl)

	expectedErr := errors.New("decode error")
	mockDecoder.EXPECT().Decode(gomock.Any()).Return(expectedErr)

	var v interface{}
	err := mockDecoder.Decode(&v)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockDecoder_MultipleDecodes tests stream decoding with InOrder.
func TestMockDecoder_MultipleDecodes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDecoder := NewMockDecoder(ctrl)

	gomock.InOrder(
		mockDecoder.EXPECT().Decode(gomock.Any()).Return(nil),
		mockDecoder.EXPECT().Decode(gomock.Any()).Return(nil),
		mockDecoder.EXPECT().Decode(gomock.Any()).Return(errors.New("EOF")),
	)

	var v1, v2, v3 interface{}

	err1 := mockDecoder.Decode(&v1)
	testutil.AssertNil(t, err1)

	err2 := mockDecoder.Decode(&v2)
	testutil.AssertNil(t, err2)

	err3 := mockDecoder.Decode(&v3)
	testutil.AssertNotNil(t, err3)
}

// TestMockDecoder_More tests stream continuation check.
func TestMockDecoder_More(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDecoder := NewMockDecoder(ctrl)

	gomock.InOrder(
		mockDecoder.EXPECT().More().Return(true),
		mockDecoder.EXPECT().More().Return(false),
	)

	hasMore1 := mockDecoder.More()
	testutil.AssertEqual(t, true, hasMore1)

	hasMore2 := mockDecoder.More()
	testutil.AssertEqual(t, false, hasMore2)
}

// TestMockEncoder_Encode tests single encode operation.
func TestMockEncoder_Encode(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockEncoder := NewMockEncoder(ctrl)

	value := map[string]string{"key": "value"}
	mockEncoder.EXPECT().Encode(value).Return(nil)

	err := mockEncoder.Encode(value)
	testutil.AssertNil(t, err)
}

// TestMockEncoder_EncodeError tests encode error handling.
func TestMockEncoder_EncodeError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockEncoder := NewMockEncoder(ctrl)

	expectedErr := errors.New("encode error")
	mockEncoder.EXPECT().Encode(gomock.Any()).Return(expectedErr)

	err := mockEncoder.Encode("invalid")
	testutil.AssertError(t, expectedErr, err)
}

// TestMockEncoder_MultipleEncodes tests stream encoding.
func TestMockEncoder_MultipleEncodes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockEncoder := NewMockEncoder(ctrl)

	val1 := map[string]string{"a": "1"}
	val2 := map[string]string{"b": "2"}
	val3 := map[string]string{"c": "3"}

	gomock.InOrder(
		mockEncoder.EXPECT().Encode(val1).Return(nil),
		mockEncoder.EXPECT().Encode(val2).Return(nil),
		mockEncoder.EXPECT().Encode(val3).Return(nil),
	)

	err1 := mockEncoder.Encode(val1)
	testutil.AssertNil(t, err1)

	err2 := mockEncoder.Encode(val2)
	testutil.AssertNil(t, err2)

	err3 := mockEncoder.Encode(val3)
	testutil.AssertNil(t, err3)
}

// TestMockEncoder_SetIndent tests indentation configuration.
func TestMockEncoder_SetIndent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockEncoder := NewMockEncoder(ctrl)

	mockEncoder.EXPECT().SetIndent("", "  ")

	mockEncoder.SetIndent("", "  ")
}

// TestMockEncoder_SetEscapeHTML tests HTML escaping configuration.
func TestMockEncoder_SetEscapeHTML(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockEncoder := NewMockEncoder(ctrl)

	mockEncoder.EXPECT().SetEscapeHTML(false)

	mockEncoder.SetEscapeHTML(false)
}

// TestMockDecoder_InputOffset tests input offset tracking.
func TestMockDecoder_InputOffset(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDecoder := NewMockDecoder(ctrl)

	mockDecoder.EXPECT().InputOffset().Return(int64(42))

	offset := mockDecoder.InputOffset()
	testutil.AssertEqual(t, int64(42), offset)
}

// TestMockDecoder_Token tests token reading.
func TestMockDecoder_Token(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDecoder := NewMockDecoder(ctrl)

	mockDecoder.EXPECT().Token().Return(gomock.Any(), nil)

	token, err := mockDecoder.Token()
	testutil.AssertNotNil(t, token)
	testutil.AssertNil(t, err)
}

// TestMockDecoder_TokenError tests token read error.
func TestMockDecoder_TokenError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDecoder := NewMockDecoder(ctrl)

	expectedErr := errors.New("token error")
	mockDecoder.EXPECT().Token().Return(nil, expectedErr)

	token, err := mockDecoder.Token()
	testutil.AssertNil(t, token)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockJSON_MarshalIndent tests marshaling with indentation.
func TestMockJSON_MarshalIndent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJSON := NewMockJSON(ctrl)

	input := map[string]string{"key": "value"}
	expected := []byte(`{
  "key": "value"
}`)

	mockJSON.EXPECT().MarshalIndent(input, "", "  ").Return(expected, nil)

	result, err := mockJSON.MarshalIndent(input, "", "  ")
	testutil.AssertNil(t, err)
	testutil.AssertBytes(t, expected, result)
}

// TestMockJSON_MarshalIndentError tests MarshalIndent error handling.
func TestMockJSON_MarshalIndentError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockJSON := NewMockJSON(ctrl)

	expectedErr := errors.New("marshal indent error")
	mockJSON.EXPECT().MarshalIndent(gomock.Any(), "", "  ").Return([]byte(nil), expectedErr)

	result, err := mockJSON.MarshalIndent("invalid", "", "  ")
	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
	testutil.AssertError(t, expectedErr, err)
}
