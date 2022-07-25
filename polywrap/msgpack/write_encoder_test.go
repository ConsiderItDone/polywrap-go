package msgpack

import (
	"bytes"
	"math"
	"strings"
	"testing"

	"github.com/consideritdone/polywrap-go/polywrap/msgpack/big"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
	"github.com/valyala/fastjson"
)

type writecase struct {
	name   string
	format string
	value  any
	bytes  []byte
	prefix []byte
	parts  [][]byte
	fn1    func(encoder Write, item any)
	fn2    func(encoder Write, key any, value any)
}

func cast[T any](t *testing.T, value any) T {
	v, ok := value.(T)
	if !ok {
		t.Fatal("can't cast type")
	}
	return v
}

func runWriteCases(t *testing.T, cases []writecase) {
	for i := range cases {
		tcase := cases[i]
		t.Run(tcase.name, func(t *testing.T) {
			context := NewContext("")
			writer := NewWriteEncoder(context)
			switch tcase.format {
			case "nil":
				writer.WriteNil()
			case "bool":
				writer.WriteBool(cast[bool](t, tcase.value))
			case "bool?":
				writer.WriteOptionalBool(cast[container.Option[bool]](t, tcase.value))
			case "int8":
				writer.WriteI8(cast[int8](t, tcase.value))
			case "int8?":
				writer.WriteOptionalI8(cast[container.Option[int8]](t, tcase.value))
			case "int16":
				writer.WriteI16(cast[int16](t, tcase.value))
			case "int16?":
				writer.WriteOptionalI16(cast[container.Option[int16]](t, tcase.value))
			case "int32":
				writer.WriteI32(cast[int32](t, tcase.value))
			case "int32?":
				writer.WriteOptionalI32(cast[container.Option[int32]](t, tcase.value))
			case "int64":
				writer.WriteI64(cast[int64](t, tcase.value))
			case "int64?":
				writer.WriteOptionalI64(cast[container.Option[int64]](t, tcase.value))
			case "uint8":
				writer.WriteU8(cast[uint8](t, tcase.value))
			case "uint8?":
				writer.WriteOptionalU8(cast[container.Option[uint8]](t, tcase.value))
			case "uint16":
				writer.WriteU16(cast[uint16](t, tcase.value))
			case "uint16?":
				writer.WriteOptionalU16(cast[container.Option[uint16]](t, tcase.value))
			case "uint32":
				writer.WriteU32(cast[uint32](t, tcase.value))
			case "uint32?":
				writer.WriteOptionalU32(cast[container.Option[uint32]](t, tcase.value))
			case "uint64":
				writer.WriteU64(cast[uint64](t, tcase.value))
			case "uint64?":
				writer.WriteOptionalU64(cast[container.Option[uint64]](t, tcase.value))
			case "float32":
				writer.WriteFloat32(cast[float32](t, tcase.value))
			case "float32?":
				writer.WriteOptionalFloat32(cast[container.Option[float32]](t, tcase.value))
			case "float64":
				writer.WriteFloat64(cast[float64](t, tcase.value))
			case "float64?":
				writer.WriteOptionalFloat64(cast[container.Option[float64]](t, tcase.value))
			case "string":
				writer.WriteString(cast[string](t, tcase.value))
			case "string?":
				writer.WriteOptionalString(cast[container.Option[string]](t, tcase.value))
			case "bytes":
				writer.WriteBytes(cast[[]byte](t, tcase.value))
			case "bytes?":
				writer.WriteOptionalBytes(cast[container.Option[[]byte]](t, tcase.value))
			case "array":
				if tcase.value == nil {
					writer.WriteArray(nil, tcase.fn1)
				} else {
					writer.WriteArray(cast[[]any](t, tcase.value), tcase.fn1)
				}
			case "array?":
				writer.WriteOptionalArray(cast[container.Option[[]any]](t, tcase.value), tcase.fn1)
			case "bigint":
				if tcase.value == nil {
					writer.WriteBigInt(nil)
				} else {
					writer.WriteBigInt(cast[*big.Int](t, tcase.value))
				}
			case "json":
				if tcase.value == nil {
					writer.WriteJson(nil)
				} else {
					writer.WriteJson(cast[*fastjson.Value](t, tcase.value))
				}
			case "map":
				writer.WriteMap(cast[map[any]any](t, tcase.value), tcase.fn2)
			default:
				t.Fatal("unknown format")
			}
			if tcase.format == "map" {
				tmp := writer.Buffer()
				if !bytes.HasPrefix(tmp, tcase.prefix) {
					t.Errorf("Bad prefix, got: %v, want: %v.", tmp[0:len(tcase.prefix)], tcase.prefix)
				}
				for j := range tcase.parts {
					if !bytes.Contains(tmp, tcase.parts[j]) {
						t.Errorf("Memory doesnt have part of value, got: %v, want: %v.", tmp, tcase.parts[j])
					}
				}
			} else {
				if !bytes.Equal(writer.Buffer(), tcase.bytes) {
					t.Errorf("Bad value, got: %v, want: %v", writer.Buffer(), tcase.bytes)
				}
			}
		})
	}
}

func TestWriteNil(t *testing.T) {
	runWriteCases(t, []writecase{
		{
			name:   "can write value",
			bytes:  []byte{192},
			format: "nil",
		},
	})
}

func TestWriteBool(t *testing.T) {
	runWriteCases(t, []writecase{
		{
			name:   "can write false",
			format: "bool",
			value:  false,
			bytes:  []byte{194},
		},
		{
			name:   "can write true",
			format: "bool",
			value:  true,
			bytes:  []byte{195},
		},
		{
			name:   "can write optional nil",
			format: "bool?",
			value:  container.None[bool](),
			bytes:  []byte{192},
		},
	})
}

func TestWriteI8(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "zero", format: "int8", value: int8(0), bytes: []byte{0}},
		{name: "negative fixed int", format: "int8", value: int8(-1), bytes: []byte{255}},
		{name: "negative fixed int", format: "int8", value: int8(-31), bytes: []byte{225}},
		{name: "negative fixed int", format: "int8", value: int8(-32), bytes: []byte{224}},
		{name: "positive fixed int", format: "int8", value: int8(1), bytes: []byte{1}},
		{name: "positive fixed int", format: "int8", value: int8(127), bytes: []byte{127}},
		{name: "8-bit signed int", format: "int8", value: int8(-128), bytes: []byte{208, 128}},
		{name: "8-bit signed int", format: "int8", value: int8(-100), bytes: []byte{208, 156}},
		{name: "8-bit signed int", format: "int8", value: int8(-33), bytes: []byte{208, 223}},
		{name: "optional nil", format: "int8?", value: container.None[int8](), bytes: []byte{192}},
		{name: "optional zero", format: "int8?", value: container.Some(int8(0)), bytes: []byte{0}},
		{name: "optional negative fixed int", format: "int8?", value: container.Some(int8(-1)), bytes: []byte{255}},
		{name: "optional negative fixed int", format: "int8?", value: container.Some(int8(-31)), bytes: []byte{225}},
		{name: "optional negative fixed int", format: "int8?", value: container.Some(int8(-32)), bytes: []byte{224}},
		{name: "optional positive fixed int", format: "int8?", value: container.Some(int8(1)), bytes: []byte{1}},
		{name: "optional positive fixed int", format: "int8?", value: container.Some(int8(127)), bytes: []byte{127}},
		{name: "optional 8-bit signed int", format: "int8?", value: container.Some(int8(-128)), bytes: []byte{208, 128}},
		{name: "optional 8-bit signed int", format: "int8?", value: container.Some(int8(-100)), bytes: []byte{208, 156}},
		{name: "optional 8-bit signed int", format: "int8?", value: container.Some(int8(-33)), bytes: []byte{208, 223}},
	})
}

func TestWriteI16(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "16-bit signed int (negative)", format: "int16", value: int16(-32768), bytes: []byte{209, 128, 0}},
		{name: "16-bit signed int (negative)", format: "int16", value: int16(-32767), bytes: []byte{209, 128, 1}},
		{name: "16-bit signed int (negative)", format: "int16", value: int16(-129), bytes: []byte{209, 255, 127}},
		{name: "16-bit signed int (positive)", format: "int16", value: int16(128), bytes: []byte{209, 0, 128}},
		{name: "16-bit signed int (positive)", format: "int16", value: int16(32767), bytes: []byte{209, 127, 255}},
		{name: "optional nil", format: "int16?", value: container.None[int16](), bytes: []byte{192}},
		{name: "optional 16-bit signed int (negative)", format: "int16?", value: container.Some(int16(-32768)), bytes: []byte{209, 128, 0}},
		{name: "optional 16-bit signed int (negative)", format: "int16?", value: container.Some(int16(-32767)), bytes: []byte{209, 128, 1}},
		{name: "optional 16-bit signed int (negative)", format: "int16?", value: container.Some(int16(-129)), bytes: []byte{209, 255, 127}},
		{name: "optional 16-bit signed int (positive)", format: "int16?", value: container.Some(int16(128)), bytes: []byte{209, 0, 128}},
		{name: "optional 16-bit signed int (positive)", format: "int16?", value: container.Some(int16(32767)), bytes: []byte{209, 127, 255}},
	})
}

func TestWriteI32(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "32-bit signed int (negative)", format: "int32", value: int32(-32769), bytes: []byte{210, 255, 255, 127, 255}},
		{name: "32-bit signed int (negative)", format: "int32", value: int32(-2147483648), bytes: []byte{210, 128, 0, 0, 0}},
		{name: "32-bit signed int (negative)", format: "int32", value: int32(-2147483647), bytes: []byte{210, 128, 0, 0, 1}},
		{name: "32-bit signed int (positive)", format: "int32", value: int32(32768), bytes: []byte{210, 0, 0, 128, 0}},
		{name: "32-bit signed int (positive)", format: "int32", value: int32(123456), bytes: []byte{210, 0, 1, 226, 64}},
		{name: "32-bit signed int (positive)", format: "int32", value: int32(2147483647), bytes: []byte{210, 127, 255, 255, 255}},
		{name: "optional nil", format: "int32?", value: container.None[int32](), bytes: []byte{192}},
		{name: "optional 32-bit signed int (negative)", format: "int32?", value: container.Some(int32(-32769)), bytes: []byte{210, 255, 255, 127, 255}},
		{name: "optional 32-bit signed int (negative)", format: "int32?", value: container.Some(int32(-2147483648)), bytes: []byte{210, 128, 0, 0, 0}},
		{name: "optional 32-bit signed int (negative)", format: "int32?", value: container.Some(int32(-2147483647)), bytes: []byte{210, 128, 0, 0, 1}},
		{name: "optional 32-bit signed int (positive)", format: "int32?", value: container.Some(int32(32768)), bytes: []byte{210, 0, 0, 128, 0}},
		{name: "optional 32-bit signed int (positive)", format: "int32?", value: container.Some(int32(123456)), bytes: []byte{210, 0, 1, 226, 64}},
		{name: "optional 32-bit signed int (positive)", format: "int32?", value: container.Some(int32(2147483647)), bytes: []byte{210, 127, 255, 255, 255}},
	})
}

func TestWriteU8(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "zero", format: "uint8", value: uint8(0), bytes: []byte{0}},
		{name: "positive fixed int", format: "uint8", value: uint8(1), bytes: []byte{1}},
		{name: "positive fixed int", format: "uint8", value: uint8(127), bytes: []byte{127}},
		{name: "8-bit unsigned int", format: "uint8", value: uint8(200), bytes: []byte{204, 200}},
		{name: "8-bit unsigned int", format: "uint8", value: uint8(255), bytes: []byte{204, 255}},
		{name: "optional nil", format: "uint8?", value: container.None[uint8](), bytes: []byte{192}},
		{name: "optional zero", format: "uint8?", value: container.Some(uint8(0)), bytes: []byte{0}},
		{name: "optional positive fixed int", format: "uint8?", value: container.Some(uint8(1)), bytes: []byte{1}},
		{name: "optional positive fixed int", format: "uint8?", value: container.Some(uint8(127)), bytes: []byte{127}},
		{name: "optional 8-bit unsigned int", format: "uint8?", value: container.Some(uint8(200)), bytes: []byte{204, 200}},
		{name: "optional 8-bit unsigned int", format: "uint8?", value: container.Some(uint8(255)), bytes: []byte{204, 255}},
	})
}

func TestWriteU16(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "16-bit unsigned int", format: "uint16", value: uint16(256), bytes: []byte{205, 1, 0}},
		{name: "16-bit unsigned int", format: "uint16", value: uint16(32767), bytes: []byte{205, 127, 255}},
		{name: "16-bit unsigned int", format: "uint16", value: uint16(32768), bytes: []byte{205, 128, 0}},
		{name: "16-bit unsigned int", format: "uint16", value: uint16(65535), bytes: []byte{205, 255, 255}},
		{name: "optional nil", format: "uint16?", value: container.None[uint16](), bytes: []byte{192}},
		{name: "optional 16-bit unsigned int", format: "uint16?", value: container.Some(uint16(256)), bytes: []byte{205, 1, 0}},
		{name: "optional 16-bit unsigned int", format: "uint16?", value: container.Some(uint16(32767)), bytes: []byte{205, 127, 255}},
		{name: "optional 16-bit unsigned int", format: "uint16?", value: container.Some(uint16(32768)), bytes: []byte{205, 128, 0}},
		{name: "optional 16-bit unsigned int", format: "uint16?", value: container.Some(uint16(65535)), bytes: []byte{205, 255, 255}},
	})
}

func TestWriteU32(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "32-bit unsigned int", format: "uint32", value: uint32(65536), bytes: []byte{206, 0, 1, 0, 0}},
		{name: "32-bit unsigned int", format: "uint32", value: uint32(123456), bytes: []byte{206, 0, 1, 226, 64}},
		{name: "32-bit unsigned int", format: "uint32", value: uint32(2147483648), bytes: []byte{206, 128, 0, 0, 0}},
		{name: "32-bit unsigned int", format: "uint32", value: uint32(4294967295), bytes: []byte{206, 255, 255, 255, 255}},
		{name: "optional nil", format: "uint32?", value: container.None[uint32](), bytes: []byte{192}},
		{name: "optional 32-bit unsigned int", format: "uint32?", value: container.Some(uint32(65536)), bytes: []byte{206, 0, 1, 0, 0}},
		{name: "optional 32-bit unsigned int", format: "uint32?", value: container.Some(uint32(123456)), bytes: []byte{206, 0, 1, 226, 64}},
		{name: "optional 32-bit unsigned int", format: "uint32?", value: container.Some(uint32(2147483648)), bytes: []byte{206, 128, 0, 0, 0}},
		{name: "optional 32-bit unsigned int", format: "uint32?", value: container.Some(uint32(4294967295)), bytes: []byte{206, 255, 255, 255, 255}},
	})
}

func TestWriteFloat32(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "can write", format: "float32", value: float32(0.5), bytes: []byte{202, 63, 0, 0, 0}},
		{name: "optional nil", format: "float32?", value: container.None[float32](), bytes: []byte{192}},
		{name: "optional value", format: "float32?", value: container.Some(float32(0.5)), bytes: []byte{202, 63, 0, 0, 0}},
	})
}

func TestWriteFloat64(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "can write", format: "float64", value: float64(3.141592653589793), bytes: []byte{203, 64, 9, 33, 251, 84, 68, 45, 24}},
		{name: "optional nil", format: "float64?", value: container.None[float64](), bytes: []byte{192}},
		{name: "optional value", format: "float64?", value: container.Some(float64(3.141592653589793)), bytes: []byte{203, 64, 9, 33, 251, 84, 68, 45, 24}},
	})
}

func TestWriteString(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "Empty String", format: "string", value: "", bytes: []byte{160}},
		{name: "5-char String", format: "string", value: "hello", bytes: []byte{165, 104, 101, 108, 108, 111}},
		{name: "11-char String", format: "string", value: "hello world", bytes: []byte{171, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}},
		{
			name:   "31-char String",
			format: "string",
			value:  "-This string contains 31 chars-",
			bytes: []byte{
				191, 45, 84, 104, 105, 115, 32, 115, 116, 114, 105, 110, 103, 32, 99, 111, 110,
				116, 97, 105, 110, 115, 32, 51, 49, 32, 99, 104, 97, 114, 115, 45}},
		{
			name:   "255-char String",
			format: "string",
			value: strings.Join([]string{
				"This is a str 8 string of 255 bytes ",
				"AC53LgxLLOKm0hfsPa1V0nfMjXtnmkEttruCPjc51dtEMLRJIEu1YoRGd9",
				"oXnM4CxcIiTc9V2DnAidZz22foIzc3kqHBoXgYskevfoJ5RK",
				"Yp52qvoDPufUebLksFl7astBNEnjPVUX2e3O9O6VKeUpB0iiHQXfzOOjTEK6Xy6ks4zAG2M6jCL01flIJlxplRXCV7 sadsadsadsadasdasaaaaa",
			}, ""),
			bytes: []byte{
				217, 255, 84, 104, 105, 115, 32, 105, 115, 32, 97, 32, 115, 116, 114, 32, 56, 32, 115,
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
				97, 97, 97,
			},
		},
		{name: "optional nil", format: "string?", value: container.None[string](), bytes: []byte{192}},
		{name: "optional empty string", format: "string?", value: container.Some(""), bytes: []byte{160}},
		{name: "optional 5-char String", format: "string?", value: container.Some("hello"), bytes: []byte{165, 104, 101, 108, 108, 111}},
	})
}

func TestWriteBytes(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "can write", format: "bytes", value: []byte{1}, bytes: []byte{196, 1, 1}},
		{name: "optional nil", format: "bytes?", value: container.None[[]byte](), bytes: []byte{192}},
		{name: "optional value", format: "bytes?", value: container.Some([]byte{1}), bytes: []byte{196, 1, 1}},
	})
}

func TestWriteArray(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "nil", format: "array", value: nil, bytes: []byte{192}},
		{
			name:   "[]int8",
			format: "array",
			value:  []any{int8(math.MinInt8), int8(math.MaxInt8)},
			bytes:  []byte{146, 208, 128, 127},
			fn1: func(encoder Write, item any) {
				encoder.WriteI8(item.(int8))
			},
		},
		{
			name:   "[]int16",
			format: "array",
			value:  []any{int16(math.MinInt16), int16(math.MaxInt16)},
			bytes:  []byte{146, 209, 128, 0, 209, 127, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteI16(item.(int16))
			},
		},
		{
			name:   "[]int32",
			format: "array",
			value:  []any{int32(math.MinInt32), int32(math.MaxInt32)},
			bytes:  []byte{146, 210, 128, 0, 0, 0, 210, 127, 255, 255, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteI32(item.(int32))
			},
		},
		{
			name:   "[]int64",
			format: "array",
			value:  []any{int64(math.MinInt64), int64(math.MaxInt64)},
			bytes:  []byte{146, 211, 128, 0, 0, 0, 0, 0, 0, 0, 211, 127, 255, 255, 255, 255, 255, 255, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteI64(item.(int64))
			},
		},
		{
			name:   "[]uint8",
			format: "array",
			value:  []any{uint8(0), uint8(math.MaxUint8)},
			bytes:  []byte{146, 0, 204, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteU8(item.(uint8))
			},
		},
		{
			name:   "[]uint16",
			format: "array",
			value:  []any{uint16(0), uint16(math.MaxUint16)},
			bytes:  []byte{146, 0, 205, 255, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteU16(item.(uint16))
			},
		},
		{
			name:   "[]uint32",
			format: "array",
			value:  []any{uint32(0), uint32(math.MaxUint32)},
			bytes:  []byte{146, 0, 206, 255, 255, 255, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteU32(item.(uint32))
			},
		},
		{
			name:   "[]uint64",
			format: "array",
			value:  []any{uint64(0), uint64(math.MaxUint64)},
			bytes:  []byte{146, 0, 207, 255, 255, 255, 255, 255, 255, 255, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteU64(item.(uint64))
			},
		},
		{
			name:   "[]float32",
			format: "array",
			value:  []any{float32(0.6046603), float32(0.9405091)},
			bytes:  []byte{146, 202, 63, 26, 203, 4, 202, 63, 112, 197, 52},
			fn1: func(encoder Write, item any) {
				encoder.WriteFloat32(item.(float32))
			},
		},
		{
			name:   "[]float64",
			format: "array",
			value:  []any{float64(0.6645600532184904), float64(0.4377141871869802)},
			bytes:  []byte{146, 203, 63, 229, 68, 19, 113, 217, 165, 93, 203, 63, 220, 3, 130, 93, 189, 166, 190},
			fn1: func(encoder Write, item any) {
				encoder.WriteFloat64(item.(float64))
			},
		},
		{name: "optional nil", format: "array?", value: container.None[[]any](), bytes: []byte{192}},
		{
			name:   "optional []int8",
			format: "array?",
			value:  container.Some([]any{int8(math.MinInt8), int8(math.MaxInt8)}),
			bytes:  []byte{146, 208, 128, 127},
			fn1: func(encoder Write, item any) {
				encoder.WriteI8(item.(int8))
			},
		},
		{
			name:   "optional []int16",
			format: "array?",
			value:  container.Some([]any{int16(math.MinInt16), int16(math.MaxInt16)}),
			bytes:  []byte{146, 209, 128, 0, 209, 127, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteI16(item.(int16))
			},
		},
		{
			name:   "optional []int32",
			format: "array?",
			value:  container.Some([]any{int32(math.MinInt32), int32(math.MaxInt32)}),
			bytes:  []byte{146, 210, 128, 0, 0, 0, 210, 127, 255, 255, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteI32(item.(int32))
			},
		},
		{
			name:   "optional []int64",
			format: "array?",
			value:  container.Some([]any{int64(math.MinInt64), int64(math.MaxInt64)}),
			bytes:  []byte{146, 211, 128, 0, 0, 0, 0, 0, 0, 0, 211, 127, 255, 255, 255, 255, 255, 255, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteI64(item.(int64))
			},
		},
		{
			name:   "optional []uint8",
			format: "array?",
			value:  container.Some([]any{uint8(0), uint8(math.MaxUint8)}),
			bytes:  []byte{146, 0, 204, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteU8(item.(uint8))
			},
		},
		{
			name:   "optional []uint16",
			format: "array?",
			value:  container.Some([]any{uint16(0), uint16(math.MaxUint16)}),
			bytes:  []byte{146, 0, 205, 255, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteU16(item.(uint16))
			},
		},
		{
			name:   "optional []uint32",
			format: "array?",
			value:  container.Some([]any{uint32(0), uint32(math.MaxUint32)}),
			bytes:  []byte{146, 0, 206, 255, 255, 255, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteU32(item.(uint32))
			},
		},
		{
			name:   "optional []uint64",
			format: "array?",
			value:  container.Some([]any{uint64(0), uint64(math.MaxUint64)}),
			bytes:  []byte{146, 0, 207, 255, 255, 255, 255, 255, 255, 255, 255},
			fn1: func(encoder Write, item any) {
				encoder.WriteU64(item.(uint64))
			},
		},
		{
			name:   "optional []float32",
			format: "array?",
			value:  container.Some([]any{float32(0.6046603), float32(0.9405091)}),
			bytes:  []byte{146, 202, 63, 26, 203, 4, 202, 63, 112, 197, 52},
			fn1: func(encoder Write, item any) {
				encoder.WriteFloat32(item.(float32))
			},
		},
		{
			name:   "optional []float64",
			format: "array?",
			value:  container.Some([]any{float64(0.6645600532184904), float64(0.4377141871869802)}),
			bytes:  []byte{146, 203, 63, 229, 68, 19, 113, 217, 165, 93, 203, 63, 220, 3, 130, 93, 189, 166, 190},
			fn1: func(encoder Write, item any) {
				encoder.WriteFloat64(item.(float64))
			},
		},
	})
}

func TestWriteBigInt(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "nil", format: "bigint", value: nil, bytes: []byte{192}},
		{name: "zero", format: "bigint", value: big.NewInt(0), bytes: []byte{161, 48}},
		{name: "maxInt64", format: "bigint", value: big.NewInt(math.MaxInt64), bytes: []byte{179, 57, 50, 50, 51, 51, 55, 50, 48, 51, 54, 56, 53, 52, 55, 55, 53, 56, 48, 55}},
	})
}

func TestWriteJSON(t *testing.T) {
	runWriteCases(t, []writecase{
		{name: "nil", format: "json", value: nil, bytes: []byte{192}},
		{name: "obj", format: "json", value: fastjson.MustParse(`{"key1":1,"key2":"string","key3":true}`), bytes: []byte{217, 38, 123, 34, 107, 101, 121, 49, 34, 58, 49, 44, 34, 107, 101, 121, 50, 34, 58, 34, 115, 116, 114, 105, 110, 103, 34, 44, 34, 107, 101, 121, 51, 34, 58, 116, 114, 117, 101, 125}},
	})
}

func TestWriteMap(t *testing.T) {
	runWriteCases(t, []writecase{
		{
			name:   "map[int8]int64",
			format: "map",
			value: map[any]any{
				int8(1): int64(1),
				int8(2): int64(2),
				int8(3): int64(3),
			},
			prefix: []byte{131},
			parts:  [][]byte{{1, 1}, {2, 2}, {3, 3}},
			fn2: func(encoder Write, key any, value any) {
				k := key.(int8)
				encoder.WriteI8(k)
				v := value.(int64)
				encoder.WriteI64(v)
			},
		},
		{
			name:   "map[string]string",
			format: "map",
			value: map[any]any{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			prefix: []byte{131},
			parts: [][]byte{
				{164, 107, 101, 121, 49, 166, 118, 97, 108, 117, 101, 49},
				{164, 107, 101, 121, 50, 166, 118, 97, 108, 117, 101, 50},
				{164, 107, 101, 121, 51, 166, 118, 97, 108, 117, 101, 51},
			},
			fn2: func(encoder Write, key any, value any) {
				k := key.(string)
				encoder.WriteString(k)
				v := value.(string)
				encoder.WriteString(v)
			},
		},
	})
}
