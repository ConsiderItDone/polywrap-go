package msgpack

import (
	"math"
	"reflect"
	"testing"

	"github.com/consideritdone/polywrap-go/polywrap/msgpack/big"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
	"github.com/valyala/fastjson"
)

type readcase struct {
	name   string
	bytes  []byte
	format string
	value  any
	fn1    func(reader Read) any
	fn2    func(reader Read) (any, any)
}

func runReadCases(t *testing.T, cases []readcase) {
	for i := range cases {
		tcase := cases[i]
		t.Run(tcase.name, func(t *testing.T) {
			context := NewContext("")
			reader := NewReadDecoder(context, tcase.bytes)

			var v any
			switch cases[i].format {
			case "nil":
				v = reader.IsNil()
			case "bool":
				v = reader.ReadBool()
			case "bool?":
				v = reader.ReadOptionalBool()
			case "int8":
				v = reader.ReadI8()
			case "int8?":
				v = reader.ReadOptionalI8()
			case "int16":
				v = reader.ReadI16()
			case "int16?":
				v = reader.ReadOptionalI16()
			case "int32":
				v = reader.ReadI32()
			case "int32?":
				v = reader.ReadOptionalI32()
			case "int64":
				v = reader.ReadI64()
			case "int64?":
				v = reader.ReadOptionalI64()
			case "uint8":
				v = reader.ReadU8()
			case "uint8?":
				v = reader.ReadOptionalU8()
			case "uint16":
				v = reader.ReadU16()
			case "uint16?":
				v = reader.ReadOptionalU16()
			case "uint32":
				v = reader.ReadU32()
			case "uint32?":
				v = reader.ReadOptionalU32()
			case "uint64":
				v = reader.ReadU64()
			case "uint64?":
				v = reader.ReadOptionalU64()
			case "float32":
				v = reader.ReadF32()
			case "float32?":
				v = reader.ReadOptionalF32()
			case "float64":
				v = reader.ReadF64()
			case "float64?":
				v = reader.ReadOptionalF64()
			case "string":
				v = reader.ReadString()
			case "string?":
				v = reader.ReadOptionalString()
			case "bytes":
				v = reader.ReadBytes()
			case "bytes?":
				v = reader.ReadOptionalBytes()
			case "array":
				v = reader.ReadArray(cases[i].fn1)
			case "array?":
				v = reader.ReadOptionalArray(cases[i].fn1)
			case "map":
				v = reader.ReadMap(cases[i].fn2)
			case "bigint":
				v = reader.ReadBigInt()
			case "json":
				v = reader.ReadJson()
			default:
				t.Fatal("unknown format")
			}
			if (v == nil && tcase.value == nil) && !reflect.DeepEqual(tcase.value, v) {
				t.Errorf("Bad value, got: %v, want: %v", v, tcase.value)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read value",
			bytes:  []byte{192},
			format: "nil",
			value:  true,
		},
	})
}

func TestReadBool(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read false",
			bytes:  []byte{194},
			format: "bool",
			value:  false,
		},
		{
			name:   "can read true",
			bytes:  []byte{195},
			format: "bool",
			value:  true,
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "bool?",
			value:  container.None[bool](),
		},
		{
			name:   "can optional false",
			bytes:  []byte{194},
			format: "bool?",
			value:  container.Some(false),
		},
		{
			name:   "can optional true",
			bytes:  []byte{195},
			format: "bool?",
			value:  container.Some(true),
		},
	})
}

func TestReadI8(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read min i8",
			bytes:  []byte{0xd0, 0x80},
			format: "int8",
			value:  int8(math.MinInt8),
		},
		{
			name:   "can read max i8",
			bytes:  []byte{0x7f},
			format: "int8",
			value:  int8(math.MaxInt8),
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "int8?",
			value:  container.None[int8](),
		},
		{
			name:   "can read optional min i8",
			bytes:  []byte{0xd0, 0x80},
			format: "int8?",
			value:  container.Some(int8(math.MinInt8)),
		},
		{
			name:   "can read optional max i8",
			bytes:  []byte{0x7f},
			format: "int8?",
			value:  container.Some(int8(math.MaxInt8)),
		},
	})
}

func TestReadI16(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read min i16",
			bytes:  []byte{0xd1, 0x80, 0x0},
			format: "int16",
			value:  int16(math.MinInt16),
		},
		{
			name:   "can read max i16",
			bytes:  []byte{0xd1, 0x7f, 0xff},
			format: "int16",
			value:  int16(math.MaxInt16),
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "int16?",
			value:  container.None[int16](),
		},
		{
			name:   "can read optional min i16",
			bytes:  []byte{0xd1, 0x80, 0x0},
			format: "int16?",
			value:  container.Some(int16(math.MinInt16)),
		},
		{
			name:   "can read optional max i16",
			bytes:  []byte{0xd1, 0x7f, 0xff},
			format: "int16?",
			value:  container.Some(int16(math.MaxInt16)),
		},
	})
}

func TestReadI32(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read min i32",
			bytes:  []byte{0xd2, 0x80, 0x0, 0x0, 0x0},
			format: "int32",
			value:  int32(math.MinInt32),
		},
		{
			name:   "can read max i32",
			bytes:  []byte{0xd2, 0x7f, 0xff, 0xff, 0xff},
			format: "int32",
			value:  int32(math.MaxInt32),
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "int32?",
			value:  container.None[int32](),
		},
		{
			name:   "can read optional min i32",
			bytes:  []byte{0xd2, 0x80, 0x0, 0x0, 0x0},
			format: "int32?",
			value:  container.Some(int32(math.MinInt32)),
		},
		{
			name:   "can read optional max i32",
			bytes:  []byte{0xd2, 0x7f, 0xff, 0xff, 0xff},
			format: "int32?",
			value:  container.Some(int32(math.MaxInt32)),
		},
	})
}

func TestReadI64(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read min i64",
			bytes:  []byte{0xd3, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			format: "int64",
			value:  int64(math.MinInt64),
		},
		{
			name:   "can read max i64",
			bytes:  []byte{0xd3, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			format: "int64",
			value:  int64(math.MaxInt64),
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "int64?",
			value:  container.None[int64](),
		},
		{
			name:   "can read optional min i64",
			bytes:  []byte{0xd3, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			format: "int64?",
			value:  container.Some(int64(math.MinInt64)),
		},
		{
			name:   "can read optional max i64",
			bytes:  []byte{0xd3, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			format: "int64?",
			value:  container.Some(int64(math.MaxInt64)),
		},
	})
}

func TestReadU8(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read min u8",
			bytes:  []byte{0x0},
			format: "uint8",
			value:  uint8(0),
		},
		{
			name:   "can read max u8",
			bytes:  []byte{0xcc, 0xff},
			format: "uint8",
			value:  uint8(math.MaxUint8),
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "uint8?",
			value:  container.None[uint8](),
		},
		{
			name:   "can read optional min u8",
			bytes:  []byte{0x0},
			format: "uint8?",
			value:  container.Some(uint8(0)),
		},
		{
			name:   "can read optional max u8",
			bytes:  []byte{0xcc, 0xff},
			format: "uint8?",
			value:  container.Some(uint8(math.MaxUint8)),
		},
	})
}

func TestReadU16(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read min u16",
			bytes:  []byte{0x0},
			format: "uint16",
			value:  uint16(0),
		},
		{
			name:   "can read max u16",
			bytes:  []byte{0xcd, 0xff, 0xff},
			format: "uint16",
			value:  uint16(math.MaxUint16),
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "uint16?",
			value:  container.None[uint16](),
		},
		{
			name:   "can read optional min u16",
			bytes:  []byte{0x0},
			format: "uint16?",
			value:  container.Some(uint16(0)),
		},
		{
			name:   "can read optional max u16",
			bytes:  []byte{0xcd, 0xff, 0xff},
			format: "uint16?",
			value:  container.Some(uint16(math.MaxUint16)),
		},
	})
}

func TestReadU32(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read min u32",
			bytes:  []byte{0x0},
			format: "uint32",
			value:  uint32(0),
		},
		{
			name:   "can read max u32",
			bytes:  []byte{0xce, 0xff, 0xff, 0xff, 0xff},
			format: "uint32",
			value:  uint32(math.MaxUint32),
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "uint32?",
			value:  container.None[uint32](),
		},
		{
			name:   "can read optional min u32",
			bytes:  []byte{0x0},
			format: "uint32?",
			value:  container.Some(uint32(0)),
		},
		{
			name:   "can read optional max u32",
			bytes:  []byte{0xce, 0xff, 0xff, 0xff, 0xff},
			format: "uint32?",
			value:  container.Some(uint32(math.MaxUint32)),
		},
	})
}

func TestReadU64(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read min u64",
			bytes:  []byte{0x0},
			format: "uint64",
			value:  uint64(0),
		},
		{
			name:   "can read max u64",
			bytes:  []byte{0xcf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			format: "uint64",
			value:  uint64(math.MaxUint64),
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "uint64?",
			value:  container.None[uint64](),
		},
		{
			name:   "can read optional min u64",
			bytes:  []byte{0x0},
			format: "uint64?",
			value:  container.Some(uint64(0)),
		},
		{
			name:   "can read optional max u64",
			bytes:  []byte{0xcf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			format: "uint64?",
			value:  container.Some(uint64(math.MaxUint64)),
		},
	})
}

func TestReadF32(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read negative f32",
			bytes:  []byte{0xca, 0xbf, 0x0, 0x0, 0x0},
			format: "float32",
			value:  float32(-0.5),
		},
		{
			name:   "can read zero f32",
			bytes:  []byte{0xca, 0x0, 0x0, 0x0, 0x0},
			format: "float32",
			value:  float32(0),
		},
		{
			name:   "can read f32",
			bytes:  []byte{0xca, 0x3d, 0xe3, 0x8e, 0x39},
			format: "float32",
			value:  float32(0.1111111111111),
		},
		{
			name:   "can read max f32",
			bytes:  []byte{0xca, 0x7f, 0x7f, 0xff, 0xff},
			format: "float32",
			value:  float32(math.MaxFloat32),
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "float32?",
			value:  container.None[float32](),
		},
		{
			name:   "can read optional negative f32",
			bytes:  []byte{0xca, 0xbf, 0x0, 0x0, 0x0},
			format: "float32?",
			value:  container.Some(float32(-0.5)),
		},
		{
			name:   "can read optional zero f32",
			bytes:  []byte{0xca, 0x0, 0x0, 0x0, 0x0},
			format: "float32?",
			value:  container.Some(float32(0)),
		},
		{
			name:   "can read optional f32",
			bytes:  []byte{0xca, 0x3d, 0xe3, 0x8e, 0x39},
			format: "float32?",
			value:  container.Some(float32(0.1111111111111)),
		},
		{
			name:   "can read optional max f32",
			bytes:  []byte{0xca, 0x7f, 0x7f, 0xff, 0xff},
			format: "float32?",
			value:  container.Some(float32(math.MaxFloat32)),
		},
	})
}

func TestReadF64(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read negative f64",
			bytes:  []byte{0xcb, 0xbf, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			format: "float64",
			value:  float64(-0.5),
		},
		{
			name:   "can read zero f64",
			bytes:  []byte{0xcb, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			format: "float64",
			value:  float64(0),
		},
		{
			name:   "can read f64",
			bytes:  []byte{0xcb, 0x3f, 0xbc, 0x71, 0xc7, 0x1c, 0x71, 0xc3, 0xfc},
			format: "float64",
			value:  float64(0.1111111111111),
		},
		{
			name:   "can read max f64",
			bytes:  []byte{0xcb, 0x7f, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			format: "float64",
			value:  float64(math.MaxFloat64),
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "float64?",
			value:  container.None[float64](),
		},
		{
			name:   "can read optional negative f64",
			bytes:  []byte{0xcb, 0xbf, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			format: "float64?",
			value:  container.Some(float64(-0.5)),
		},
		{
			name:   "can read optional zero f64",
			bytes:  []byte{0xcb, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			format: "float64?",
			value:  container.Some(float64(0)),
		},
		{
			name:   "can read optional f64",
			bytes:  []byte{0xcb, 0x3f, 0xbc, 0x71, 0xc7, 0x1c, 0x71, 0xc3, 0xfc},
			format: "float64?",
			value:  container.Some(float64(0.1111111111111)),
		},
		{
			name:   "can read optional max f64",
			bytes:  []byte{0xcb, 0x7f, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			format: "float64?",
			value:  container.Some(float64(math.MaxFloat64)),
		},
	})
}

func TestReadBytes(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can read nil",
			bytes:  []byte{0xc0},
			format: "bytes",
			value:  []byte{},
		},
		{
			name:   "can read bytes",
			bytes:  []byte{0xc4, 0x1, 0x1},
			format: "bytes",
			value:  []byte{1},
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "bytes?",
			value:  container.None[[]byte](),
		},
		{
			name:   "can read optional bytes",
			bytes:  []byte{0xc4, 0x1, 0x1},
			format: "bytes?",
			value:  container.Some([]byte{1}),
		},
	})
}

func TestReadString(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "can empty string",
			bytes:  []byte{0xa0},
			format: "string",
			value:  "",
		},
		{
			name:   "can read string",
			bytes:  []byte{0xab, 0x73, 0x6f, 0x6d, 0x65, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67},
			format: "string",
			value:  "some string",
		},
		{
			name:   "can optional nil",
			bytes:  []byte{192},
			format: "string?",
			value:  container.None[string](),
		},
		{
			name:   "can read optional empty string",
			bytes:  []byte{0xa0},
			format: "string?",
			value:  container.Some(""),
		},
		{
			name:   "can read optional string",
			bytes:  []byte{0xab, 0x73, 0x6f, 0x6d, 0x65, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67},
			format: "string?",
			value:  container.Some("some string"),
		},
	})
}

func TestReadArray(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "nil",
			bytes:  []byte{192},
			format: "array",
			value:  []any{},
			fn1:    nil,
		},
		{
			name:   "[]int8",
			bytes:  []byte{146, 208, 128, 127},
			format: "array",
			value:  []any{int8(math.MinInt8), int8(math.MaxInt8)},
			fn1: func(reader Read) any {
				return reader.ReadI8()
			},
		},
		{
			name:   "[]int16",
			bytes:  []byte{146, 209, 128, 0, 209, 127, 255},
			format: "array",
			value:  []any{int16(math.MinInt16), int16(math.MaxInt16)},
			fn1: func(reader Read) any {
				return reader.ReadI16()
			},
		},
		{
			name:   "[]int32",
			bytes:  []byte{146, 210, 128, 0, 0, 0, 210, 127, 255, 255, 255},
			format: "array",
			value:  []any{int32(math.MinInt32), int32(math.MaxInt32)},
			fn1: func(reader Read) any {
				return reader.ReadI32()
			},
		},
		{
			name:   "[]int64",
			bytes:  []byte{146, 211, 128, 0, 0, 0, 0, 0, 0, 0, 211, 127, 255, 255, 255, 255, 255, 255, 255},
			format: "array",
			value:  []any{int64(math.MinInt64), int64(math.MaxInt64)},
			fn1: func(reader Read) any {
				return reader.ReadI64()
			},
		},
		{
			name:   "[]uint8",
			bytes:  []byte{146, 0, 204, 255},
			format: "array",
			value:  []any{uint8(0), uint8(math.MaxUint8)},
			fn1: func(reader Read) any {
				return reader.ReadU8()
			},
		},
		{
			name:   "[]uint16",
			bytes:  []byte{146, 0, 205, 255, 255},
			format: "array",
			value:  []any{uint16(0), uint16(math.MaxUint16)},
			fn1: func(reader Read) any {
				return reader.ReadU16()
			},
		},
		{
			name:   "[]uint32",
			bytes:  []byte{146, 0, 206, 255, 255, 255, 255},
			format: "array",
			value:  []any{uint32(0), uint32(math.MaxUint32)},
			fn1: func(reader Read) any {
				return reader.ReadU32()
			},
		},
		{
			name:   "[]uint64",
			bytes:  []byte{146, 0, 207, 255, 255, 255, 255, 255, 255, 255, 255},
			format: "array",
			value:  []any{uint64(0), uint64(math.MaxUint64)},
			fn1: func(reader Read) any {
				return reader.ReadU64()
			},
		},
		{
			name:   "[]float32",
			bytes:  []byte{146, 202, 63, 26, 203, 4, 202, 63, 112, 197, 52},
			format: "array",
			value:  []any{float32(0.6046603), float32(0.9405091)},
			fn1: func(reader Read) any {
				return reader.ReadF32()
			},
		},
		{
			name:   "[]float64",
			bytes:  []byte{146, 203, 63, 229, 68, 19, 113, 217, 165, 93, 203, 63, 220, 3, 130, 93, 189, 166, 190},
			format: "array",
			value:  []any{float64(0.6645600532184904), float64(0.4377141871869802)},
			fn1: func(reader Read) any {
				return reader.ReadF64()
			},
		},
		{
			name:   "optional nil",
			bytes:  []byte{192},
			format: "array?",
			value:  container.None[[]any](),
		},
		{
			name:   "optional []int8",
			bytes:  []byte{146, 208, 128, 127},
			format: "array?",
			value:  container.Some([]any{int8(math.MinInt8), int8(math.MaxInt8)}),
			fn1: func(reader Read) any {
				return reader.ReadI8()
			},
		},
		{
			name:   "optional []int16",
			bytes:  []byte{146, 209, 128, 0, 209, 127, 255},
			format: "array?",
			value:  container.Some([]any{int16(math.MinInt16), int16(math.MaxInt16)}),
			fn1: func(reader Read) any {
				return reader.ReadI16()
			},
		},
		{
			name:   "optional []int32",
			bytes:  []byte{146, 210, 128, 0, 0, 0, 210, 127, 255, 255, 255},
			format: "array?",
			value:  container.Some([]any{int32(math.MinInt32), int32(math.MaxInt32)}),
			fn1: func(reader Read) any {
				return reader.ReadI32()
			},
		},
		{
			name:   "optional []int64",
			bytes:  []byte{146, 211, 128, 0, 0, 0, 0, 0, 0, 0, 211, 127, 255, 255, 255, 255, 255, 255, 255},
			format: "array?",
			value:  container.Some([]any{int64(math.MinInt64), int64(math.MaxInt64)}),
			fn1: func(reader Read) any {
				return reader.ReadI64()
			},
		},
		{
			name:   "optional []uint8",
			bytes:  []byte{146, 0, 204, 255},
			format: "array?",
			value:  container.Some([]any{uint8(0), uint8(math.MaxUint8)}),
			fn1: func(reader Read) any {
				return reader.ReadU8()
			},
		},
		{
			name:   "optional []uint16",
			bytes:  []byte{146, 0, 205, 255, 255},
			format: "array?",
			value:  container.Some([]any{uint16(0), uint16(math.MaxUint16)}),
			fn1: func(reader Read) any {
				return reader.ReadU16()
			},
		},
		{
			name:   "optional []uint32",
			bytes:  []byte{146, 0, 206, 255, 255, 255, 255},
			format: "array?",
			value:  container.Some([]any{uint32(0), uint32(math.MaxUint32)}),
			fn1: func(reader Read) any {
				return reader.ReadU32()
			},
		},
		{
			name:   "optional []uint64",
			bytes:  []byte{146, 0, 207, 255, 255, 255, 255, 255, 255, 255, 255},
			format: "array?",
			value:  container.Some([]any{uint64(0), uint64(math.MaxUint64)}),
			fn1: func(reader Read) any {
				return reader.ReadU64()
			},
		},
		{
			name:   "optional []float32",
			bytes:  []byte{146, 202, 63, 26, 203, 4, 202, 63, 112, 197, 52},
			format: "array?",
			value:  container.Some([]any{float32(0.6046603), float32(0.9405091)}),
			fn1: func(reader Read) any {
				return reader.ReadF32()
			},
		},
		{
			name:   "optional []float64",
			bytes:  []byte{146, 203, 63, 229, 68, 19, 113, 217, 165, 93, 203, 63, 220, 3, 130, 93, 189, 166, 190},
			format: "array?",
			value:  container.Some([]any{float64(0.6645600532184904), float64(0.4377141871869802)}),
			fn1: func(reader Read) any {
				return reader.ReadF64()
			},
		},
	})
}

func TestReadMap(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "nil",
			bytes:  []byte{192},
			format: "map",
			value:  map[any]any{},
			fn2:    nil,
		},
		{
			name:   "map[string]int64",
			bytes:  []byte{131, 164, 107, 101, 121, 51, 3, 164, 107, 101, 121, 49, 1, 164, 107, 101, 121, 50, 2},
			format: "map",
			value: map[any]any{
				"key1": int64(1),
				"key2": int64(2),
				"key3": int64(3),
			},
			fn2: func(reader Read) (any, any) {
				key := reader.ReadString()
				val := reader.ReadI64()
				return key, val
			},
		},
		{
			name:   "map[string]string",
			bytes:  []byte{131, 164, 107, 101, 121, 49, 164, 118, 97, 108, 49, 164, 107, 101, 121, 50, 164, 118, 97, 108, 50, 164, 107, 101, 121, 51, 164, 118, 97, 108, 51},
			format: "map",
			value: map[any]any{
				"key1": "val1",
				"key2": "val2",
				"key3": "val3",
			},
			fn2: func(reader Read) (any, any) {
				key := reader.ReadString()
				val := reader.ReadString()
				return key, val
			},
		},
	})
}

func TestReadBigInt(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "nil",
			bytes:  []byte{192},
			format: "bigint",
			value:  nil,
		},
		{
			name:   "zero",
			bytes:  []byte{161, 48},
			format: "bigint",
			value:  big.NewInt(0),
		},
		{
			name:   "maxInt64",
			bytes:  []byte{179, 57, 50, 50, 51, 51, 55, 50, 48, 51, 54, 56, 53, 52, 55, 55, 53, 56, 48, 55},
			format: "bigint",
			value:  big.NewInt(math.MaxInt64),
		},
	})
}

func TestReadJson(t *testing.T) {
	runReadCases(t, []readcase{
		{
			name:   "nil",
			bytes:  []byte{192},
			format: "json",
			value:  nil,
		},
		{
			name:   "obj",
			bytes:  []byte{217, 38, 123, 34, 107, 101, 121, 49, 34, 58, 49, 44, 34, 107, 101, 121, 50, 34, 58, 34, 115, 116, 114, 105, 110, 103, 34, 44, 34, 107, 101, 121, 51, 34, 58, 116, 114, 117, 101, 125},
			format: "json",
			value:  fastjson.MustParse(`{"key1":1,"key2":"string","key3":true}`),
		},
	})
}
