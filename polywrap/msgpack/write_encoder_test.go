package msgpack

import (
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/json"
	"github.com/valyala/fastjson"
	"math"
	"reflect"
	"testing"
)

func TestWriteNil(t *testing.T) {
	context := NewContext("")
	writer := NewWriteEncoder(context)
	writer.WriteNil()

	actual := writer.Buffer()
	expected := []byte{192}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("WriteNil is incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestWriteBoolTrue(t *testing.T) {
	context := NewContext("")
	writer := NewWriteEncoder(context)
	writer.WriteBool(true)

	actual := writer.Buffer()
	expected := []byte{195}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("WriteBool(true) is incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestWriteBoolFalse(t *testing.T) {
	context := NewContext("")
	writer := NewWriteEncoder(context)
	writer.WriteBool(false)

	actual := writer.Buffer()
	expected := []byte{194}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("WriteBool(false) is incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestWriteI8(t *testing.T) {
	tests := []struct {
		name  string
		input int8
		want  []byte
	}{
		{name: "zero", input: 0, want: []byte{0}},
		{name: "negative fixed int", input: -1, want: []byte{255}},
		{name: "negative fixed int", input: -31, want: []byte{225}},
		{name: "negative fixed int", input: -32, want: []byte{224}},
		{name: "positive fixed int", input: 1, want: []byte{1}},
		{name: "positive fixed int", input: 127, want: []byte{127}},
		{name: "8-bit signed int", input: -128, want: []byte{208, 128}},
		{name: "8-bit signed int", input: -100, want: []byte{208, 156}},
		{name: "8-bit signed int", input: -33, want: []byte{208, 223}},
	}

	for _, tc := range tests {
		context := NewContext("")
		writer := NewWriteEncoder(context)
		writer.WriteI8(tc.input)

		got := writer.Buffer()
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s (%d): expected: %v, got: %v", tc.name, tc.input, tc.want, got)
		}
	}
}

func TestWriteI16(t *testing.T) {
	tests := []struct {
		name  string
		input int16
		want  []byte
	}{
		{name: "16-bit signed int (negative)", input: -32768, want: []byte{209, 128, 0}},
		{name: "16-bit signed int (negative)", input: -32767, want: []byte{209, 128, 1}},
		{name: "16-bit signed int (negative)", input: -129, want: []byte{209, 255, 127}},
		{name: "16-bit signed int (positive)", input: 128, want: []byte{209, 0, 128}},
		{name: "16-bit signed int (positive)", input: 32767, want: []byte{209, 127, 255}},
	}

	for _, tc := range tests {
		context := NewContext("")
		writer := NewWriteEncoder(context)
		writer.WriteI16(tc.input)

		got := writer.Buffer()
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s (%d): expected: %v, got: %v", tc.name, tc.input, tc.want, got)
		}
	}
}

func TestWriteI32(t *testing.T) {
	tests := []struct {
		name  string
		input int32
		want  []byte
	}{
		{name: "32-bit signed int (negative)", input: -32769, want: []byte{210, 255, 255, 127, 255}},
		{name: "32-bit signed int (negative)", input: -2147483648, want: []byte{210, 128, 0, 0, 0}},
		{name: "32-bit signed int (negative)", input: -2147483647, want: []byte{210, 128, 0, 0, 1}},
		{name: "32-bit signed int (positive)", input: 32768, want: []byte{210, 0, 0, 128, 0}},
		{name: "32-bit signed int (positive)", input: 123456, want: []byte{210, 0, 1, 226, 64}},
		{name: "32-bit signed int (positive)", input: 2147483647, want: []byte{210, 127, 255, 255, 255}},
	}

	for _, tc := range tests {
		context := NewContext("")
		writer := NewWriteEncoder(context)
		writer.WriteI32(tc.input)

		got := writer.Buffer()
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s (%d): expected: %v, got: %v", tc.name, tc.input, tc.want, got)
		}
	}
}

func TestWriteU8(t *testing.T) {
	tests := []struct {
		name  string
		input uint8
		want  []byte
	}{
		{name: "zero", input: 0, want: []byte{0}},
		{name: "positive fixed int", input: 1, want: []byte{1}},
		{name: "positive fixed int", input: 127, want: []byte{127}},
		{name: "8-bit unsigned int", input: 200, want: []byte{204, 200}},
		{name: "8-bit unsigned int", input: 255, want: []byte{204, 255}},
	}

	for _, tc := range tests {
		context := NewContext("")
		writer := NewWriteEncoder(context)
		writer.WriteU8(tc.input)

		got := writer.Buffer()
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s (%d): expected: %v, got: %v", tc.name, tc.input, tc.want, got)
		}
	}
}

func TestWriteU16(t *testing.T) {
	tests := []struct {
		name  string
		input uint16
		want  []byte
	}{
		{name: "16-bit unsigned int", input: 256, want: []byte{205, 1, 0}},
		{name: "16-bit unsigned int", input: 32767, want: []byte{205, 127, 255}},
		{name: "16-bit unsigned int", input: 32768, want: []byte{205, 128, 0}},
		{name: "16-bit unsigned int", input: 65535, want: []byte{205, 255, 255}},
	}

	for _, tc := range tests {
		context := NewContext("")
		writer := NewWriteEncoder(context)
		writer.WriteU16(tc.input)

		got := writer.Buffer()
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s (%d): expected: %v, got: %v", tc.name, tc.input, tc.want, got)
		}
	}
}

func TestWriteU32(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  []byte
	}{
		{name: "32-bit unsigned int", input: 65536, want: []byte{206, 0, 1, 0, 0}},
		{name: "32-bit unsigned int", input: 123456, want: []byte{206, 0, 1, 226, 64}},
		{name: "32-bit unsigned int", input: 2147483648, want: []byte{206, 128, 0, 0, 0}},
		{name: "32-bit unsigned int", input: 4294967295, want: []byte{206, 255, 255, 255, 255}},
	}

	for _, tc := range tests {
		context := NewContext("")
		writer := NewWriteEncoder(context)
		writer.WriteU32(tc.input)

		got := writer.Buffer()
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s (%d): expected: %v, got: %v", tc.name, tc.input, tc.want, got)
		}
	}
}

func TestWriteFloat32(t *testing.T) {
	context := NewContext("")
	writer := NewWriteEncoder(context)
	writer.WriteFloat32(0.5)

	actual := writer.Buffer()
	expected := []byte{202, 63, 0, 0, 0}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("WriteFloat32(false) is incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestWriteFloat64(t *testing.T) {
	context := NewContext("")
	writer := NewWriteEncoder(context)
	writer.WriteFloat64(3.141592653589793)

	actual := writer.Buffer()
	expected := []byte{203, 64, 9, 33, 251, 84, 68, 45, 24}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("WriteFloat32(false) is incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestWriteString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []byte
	}{
		{name: "Empty String", input: "", want: []byte{160}},
		{name: "5-char String", input: "hello", want: []byte{165, 104, 101, 108, 108, 111}},
		{name: "11-char String", input: "hello world", want: []byte{171, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}},
		{name: "31-char String", input: "-This string contains 31 chars-", want: []byte{191, 45, 84, 104, 105, 115, 32, 115, 116, 114, 105, 110, 103, 32, 99, 111, 110,
			116, 97, 105, 110, 115, 32, 51, 49, 32, 99, 104, 97, 114, 115, 45}},
		{name: "255-char String", input: "This is a str 8 string of 255 bytes " +
			"AC53LgxLLOKm0hfsPa1V0nfMjXtnmkEttruCPjc51dtEMLRJIEu1YoRGd9" + "oXnM4CxcIiTc9V2DnAidZz22foIzc3kqHBoXgYskevfoJ5RK" +
			"Yp52qvoDPufUebLksFl7astBNEnjPVUX2e3O9O6VKeUpB0iiHQXfzOOjTEK6Xy6ks4zAG2M6jCL01flIJlxplRXCV7 sadsadsadsadasdasaaaaa",
			want: []byte{217, 255, 84, 104, 105, 115, 32, 105, 115, 32, 97, 32, 115, 116, 114, 32, 56, 32, 115,
				116, 114, 105, 110, 103, 32, 111, 102, 32, 50, 53, 53, 32, 98, 121, 116, 101, 115, 32,
				65, 67, 53, 51, 76, 103, 120, 76, 76, 79, 75, 109, 48, 104, 102, 115, 80, 97, 49, 86,
				48, 110, 102, 77, 106, 88, 116, 110, 109, 107, 69, 116, 116, 114, 117, 67, 80, 106, 99,
				53, 49, 100, 116, 69, 77, 76, 82, 74, 73, 69, 117, 49, 89, 111, 82, 71, 100, 57, 111,
				88, 110, 77, 52, 67, 120, 99, 73, 105, 84, 99, 57, 86, 50, 68, 110, 65, 105, 100, 90,
				122, 50, 50, 102, 111, 73, 122, 99, 51, 107, 113, 72, 66, 111, 88, 103, 89, 115, 107,
				101, 118, 102, 111, 74, 53, 82, 75, 89, 112, 53, 50, 113, 118, 111, 68, 80, 117, 102,
				85, 101, 98, 76, 107, 115, 70, 108, 55, 97, 115, 116, 66, 78, 69, 110, 106, 80, 86, 85,
				88, 50, 101, 51, 79, 57, 79, 54, 86, 75, 101, 85, 112, 66, 48, 105, 105, 72, 81, 88,
				102, 122, 79, 79, 106, 84, 69, 75, 54, 88, 121, 54, 107, 115, 52, 122, 65, 71, 50, 77,
				54, 106, 67, 76, 48, 49, 102, 108, 73, 74, 108, 120, 112, 108, 82, 88, 67, 86, 55, 32,
				115, 97, 100, 115, 97, 100, 115, 97, 100, 115, 97, 100, 97, 115, 100, 97, 115, 97, 97,
				97, 97, 97}},
	}

	for _, tc := range tests {
		context := NewContext("")
		writer := NewWriteEncoder(context)
		writer.WriteString(tc.input)

		got := writer.Buffer()
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s (%s): expected: %v, got: %v", tc.name, tc.input, tc.want, got)
		}
	}
}

func TestWriteBytes(t *testing.T) {
	context := NewContext("")
	writer := NewWriteEncoder(context)
	writer.WriteBytes([]byte{1})

	actual := writer.Buffer()
	expected := []byte{196, 1, 1}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("WriteBytes(false) is incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestWriteArray(t *testing.T) {
	cases := []struct {
		name     string
		data     []any
		expected []byte
		fn       func(encoder Write, item any)
	}{
		{
			name:     "nil",
			data:     nil,
			expected: []byte{192},
			fn:       nil,
		},
		{
			name:     "[]int8",
			data:     []any{int8(math.MinInt8), int8(math.MaxInt8)},
			expected: []byte{146, 208, 128, 127},
			fn: func(encoder Write, item any) {
				encoder.WriteI8(item.(int8))
			},
		},
		{
			name:     "[]int16",
			data:     []any{int16(math.MinInt16), int16(math.MaxInt16)},
			expected: []byte{146, 209, 128, 0, 209, 127, 255},
			fn: func(encoder Write, item any) {
				encoder.WriteI16(item.(int16))
			},
		},
		{
			name:     "[]int32",
			data:     []any{int32(math.MinInt32), int32(math.MaxInt32)},
			expected: []byte{146, 210, 128, 0, 0, 0, 210, 127, 255, 255, 255},
			fn: func(encoder Write, item any) {
				encoder.WriteI32(item.(int32))
			},
		},
		{
			name:     "[]int64",
			data:     []any{int64(math.MinInt64), int64(math.MaxInt64)},
			expected: []byte{146, 211, 128, 0, 0, 0, 0, 0, 0, 0, 211, 127, 255, 255, 255, 255, 255, 255, 255},
			fn: func(encoder Write, item any) {
				encoder.WriteI64(item.(int64))
			},
		},
		{
			name:     "[]uint8",
			data:     []any{uint8(0), uint8(math.MaxUint8)},
			expected: []byte{146, 0, 204, 255},
			fn: func(encoder Write, item any) {
				encoder.WriteU8(item.(uint8))
			},
		},
		{
			name:     "[]uint16",
			data:     []any{uint16(0), uint16(math.MaxUint16)},
			expected: []byte{146, 0, 205, 255, 255},
			fn: func(encoder Write, item any) {
				encoder.WriteU16(item.(uint16))
			},
		},
		{
			name:     "[]uint32",
			data:     []any{uint32(0), uint32(math.MaxUint32)},
			expected: []byte{146, 0, 206, 255, 255, 255, 255},
			fn: func(encoder Write, item any) {
				encoder.WriteU32(item.(uint32))
			},
		},
		{
			name:     "[]uint64",
			data:     []any{uint64(0), uint64(math.MaxUint64)},
			expected: []byte{146, 0, 207, 255, 255, 255, 255, 255, 255, 255, 255},
			fn: func(encoder Write, item any) {
				encoder.WriteU64(item.(uint64))
			},
		},
		{
			name:     "[]float32",
			data:     []any{float32(0.6046603), float32(0.9405091)},
			expected: []byte{146, 202, 63, 26, 203, 4, 202, 63, 112, 197, 52},
			fn: func(encoder Write, item any) {
				encoder.WriteFloat32(item.(float32))
			},
		},
		{
			name:     "[]float64",
			data:     []any{float64(0.6645600532184904), float64(0.4377141871869802)},
			expected: []byte{146, 203, 63, 229, 68, 19, 113, 217, 165, 93, 203, 63, 220, 3, 130, 93, 189, 166, 190},
			fn: func(encoder Write, item any) {
				encoder.WriteFloat64(item.(float64))
			},
		},
	}

	for i := range cases {
		tcase := cases[i]
		t.Run(tcase.name, func(t *testing.T) {
			context := NewContext("")
			writer := NewWriteEncoder(context)
			writer.WriteArray(tcase.data, tcase.fn)

			actual := writer.Buffer()
			if !reflect.DeepEqual(actual, tcase.expected) {
				t.Logf("%#+v", tcase)
				t.Errorf("TestWriteArray(%s) is incorrect, got: %v, want: %v.", tcase.name, actual, tcase.expected)
			}
		})
	}
}
