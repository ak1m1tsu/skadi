package skadi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Encode_Success(t *testing.T) {
	testCases := []struct {
		name     string
		data     interface{}
		expected []byte
	}{
		{
			name:     "encode request",
			data:     &Request{CommandCode: Message, Body: []byte{'r', 'e', 'q', 'u', 'e', 's', 't'}},
			expected: []byte{0x0, 0x1, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74},
		},
		{
			name:     "encode response",
			data:     &Response{Body: []byte{'r', 'e', 's', 'p', 'o', 'n', 's', 'e'}},
			expected: []byte{0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Encode(tc.data)
			assert.Nil(t, err)
			assert.NotNil(t, res)
			assert.NotEmpty(t, res)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func Test_Encode_Fail(t *testing.T) {
	testCases := []struct {
		name        string
		data        interface{}
		expectedErr string
	}{
		{
			name:        "unsupported type string",
			data:        "hello world",
			expectedErr: "unsupported data type: string",
		},
		{
			name:        "unsupported type int",
			data:        int(10),
			expectedErr: "unsupported data type: int",
		},
		{
			name:        "unsupported type float64",
			data:        float64(10),
			expectedErr: "unsupported data type: float64",
		},
		{
			name:        "unsupported type byte",
			data:        byte(10),
			expectedErr: "unsupported data type: uint8",
		},
		{
			name:        "unsupported type uint",
			data:        uint(10),
			expectedErr: "unsupported data type: uint",
		},
		{
			name:        "unsupported type struct",
			data:        struct{ field string }{field: "hello world"},
			expectedErr: "unsupported data type: struct { field string }",
		},
		{
			name:        "unsupported type interface",
			data:        interface{}(nil),
			expectedErr: "unsupported data type: <nil>",
		},
		{
			name:        "unsupported type slice",
			data:        []interface{}{nil},
			expectedErr: "unsupported data type: []interface {}",
		},
		{
			name:        "unsupported type interface",
			data:        map[string]interface{}{"key": nil},
			expectedErr: "unsupported data type: map[string]interface {}",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Encode(tc.data)
			assert.Error(t, err)
			assert.Equal(t, tc.expectedErr, err.Error())
		})
	}
}

func Test_Decode_Success(t *testing.T) {
	testCases := []struct {
		name     string
		data     []byte
		dest     interface{}
		expected interface{}
	}{
		{
			name:     "decode request",
			data:     []byte{0x0, 0x1, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74},
			dest:     &Request{},
			expected: &Request{CommandCode: Message, Body: []byte{'r', 'e', 'q', 'u', 'e', 's', 't'}},
		},
		{
			name:     "decode response",
			data:     []byte{0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65},
			dest:     &Response{},
			expected: &Response{Body: []byte{'r', 'e', 's', 'p', 'o', 'n', 's', 'e'}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Decode(tc.data, tc.dest)
			assert.Nil(t, err)
			assert.NotNil(t, tc.dest)
			assert.NotEmpty(t, tc.dest)
			assert.Equal(t, tc.expected, tc.dest)
		})
	}
}

func Test_Decode_Fail(t *testing.T) {
	// testCases := []struct{}{}

}
