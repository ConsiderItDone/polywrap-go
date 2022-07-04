package msgpack

import (
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
		{name: "zero", input: 0, want: []byte{208, 0}},
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
		data     []interface{}
		expected []byte
	}{
		{
			name:     "nil",
			data:     nil,
			expected: []byte{192},
		},
		{
			name:     "[]int8",
			data:     []interface{}{int8(math.MinInt8), int8(math.MaxInt8)},
			expected: []byte{146, 208, 128, 127},
		},
		{
			name:     "[]int16",
			data:     []interface{}{int16(math.MinInt16), int16(math.MaxInt16)},
			expected: []byte{146, 209, 128, 0, 209, 127, 255},
		},
		{
			name:     "[]int32",
			data:     []interface{}{int32(math.MinInt32), int32(math.MaxInt32)},
			expected: []byte{146, 210, 128, 0, 0, 0, 210, 127, 255, 255, 255},
		},
		{
			name:     "[]int64",
			data:     []interface{}{int64(math.MinInt64), int64(math.MaxInt64)},
			expected: []byte{146, 211, 128, 0, 0, 0, 0, 0, 0, 0, 211, 127, 255, 255, 255, 255, 255, 255, 255},
		},
		{
			name:     "[]uint8",
			data:     []interface{}{uint8(0), uint8(math.MaxUint8)},
			expected: []byte{146, 0, 204, 255},
		},
		{
			name:     "[]uint16",
			data:     []interface{}{uint16(0), uint16(math.MaxUint16)},
			expected: []byte{146, 0, 205, 255, 255},
		},
		{
			name:     "[]uint32",
			data:     []interface{}{uint32(0), uint32(math.MaxUint32)},
			expected: []byte{146, 0, 206, 255, 255, 255, 255},
		},
		{
			name:     "[]uint64",
			data:     []interface{}{uint64(0), uint64(math.MaxUint64)},
			expected: []byte{146, 0, 207, 255, 255, 255, 255, 255, 255, 255, 255},
		},
		{
			name:     "[]float32",
			data:     []interface{}{float32(0.6046603), float32(0.9405091)},
			expected: []byte{146, 202, 63, 26, 203, 4, 202, 63, 112, 197, 52},
		},
		{
			name:     "[]float64",
			data:     []interface{}{float64(0.6645600532184904), float64(0.4377141871869802)},
			expected: []byte{146, 203, 63, 229, 68, 19, 113, 217, 165, 93, 203, 63, 220, 3, 130, 93, 189, 166, 190},
		},
	}

	for i := range cases {
		tcase := cases[i]
		t.Run(tcase.name, func(t *testing.T) {
			context := NewContext("")
			writer := NewWriteEncoder(context)
			writer.WriteArray(tcase.data)

			actual := writer.Buffer()
			if !reflect.DeepEqual(actual, tcase.expected) {
				t.Logf("%#+v", tcase)
				t.Errorf("TestWriteArray(%s) is incorrect, got: %v, want: %v.", tcase.name, actual, tcase.expected)
			}
		})
	}
}
