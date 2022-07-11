package msgpack

import (
	"math"
	"reflect"
	"testing"

	"github.com/consideritdone/polywrap-go/polywrap/msgpack/format"
)

type readcase struct {
	name   string
	bytes  []byte
	format format.Format
	value  interface{}
	fn     func(reader Read) any
}

func runcases(t *testing.T, cases []readcase) {
	for i := range cases {
		tcase := cases[i]
		t.Run(tcase.name, func(t *testing.T) {
			context := NewContext("")
			reader := NewReadDecoder(context, tcase.bytes)

			var v interface{}
			switch cases[i].format {
			case format.NIL:
				v = reader.IsNil()
			case format.TRUE, format.FALSE:
				v = reader.ReadBool()
			case format.INT8:
				v = reader.ReadI8()
			case format.INT16:
				v = reader.ReadI16()
			case format.INT32:
				v = reader.ReadI32()
			case format.INT64:
				v = reader.ReadI64()
			case format.UINT8:
				v = reader.ReadU8()
			case format.UINT16:
				v = reader.ReadU16()
			case format.UINT32:
				v = reader.ReadU32()
			case format.UINT64:
				v = reader.ReadU64()
			case format.FLOAT32:
				v = reader.ReadF32()
			case format.FLOAT64:
				v = reader.ReadF64()
			case format.STR8, format.STR16, format.STR32:
				v = reader.ReadString()
			case format.BIN8, format.BIN16, format.BIN32:
				v = reader.ReadBytes()
			case format.ARRAY16, format.ARRAY32:
				v = reader.ReadArray(cases[i].fn)
			default:
				t.Fatal("unknown format")
			}
			if !reflect.DeepEqual(tcase.value, v) {
				t.Errorf("Bad value, got: %v, want: %v", v, tcase.value)
			}
		})

	}
}

func TestIsNil(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can read value",
			bytes:  []byte{192},
			format: format.NIL,
			value:  true,
		},
	})
}

func TestReadBool(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can read false",
			bytes:  []byte{194},
			format: format.FALSE,
			value:  false,
		},
		{
			name:   "can read true",
			bytes:  []byte{195},
			format: format.TRUE,
			value:  true,
		},
	})
}

func TestReadI8(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can read min i8",
			bytes:  []byte{0xd0, 0x80},
			format: format.INT8,
			value:  int8(math.MinInt8),
		},
		{
			name:   "can read max i8",
			bytes:  []byte{0x7f},
			format: format.INT8,
			value:  int8(math.MaxInt8),
		},
	})
}

func TestReadI16(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can read min i16",
			bytes:  []byte{0xd1, 0x80, 0x0},
			format: format.INT16,
			value:  int16(math.MinInt16),
		},
		{
			name:   "can read max i16",
			bytes:  []byte{0xd1, 0x7f, 0xff},
			format: format.INT16,
			value:  int16(math.MaxInt16),
		},
	})
}

func TestReadI32(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can read min i32",
			bytes:  []byte{0xd2, 0x80, 0x0, 0x0, 0x0},
			format: format.INT32,
			value:  int32(math.MinInt32),
		},
		{
			name:   "can read max i32",
			bytes:  []byte{0xd2, 0x7f, 0xff, 0xff, 0xff},
			format: format.INT32,
			value:  int32(math.MaxInt32),
		},
	})
}

func TestReadI64(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can read min i64",
			bytes:  []byte{0xd3, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			format: format.INT64,
			value:  int64(math.MinInt64),
		},
		{
			name:   "can read max i64",
			bytes:  []byte{0xd3, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			format: format.INT64,
			value:  int64(math.MaxInt64),
		},
	})
}

func TestReadU8(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can read min u8",
			bytes:  []byte{0x0},
			format: format.UINT8,
			value:  uint8(0),
		},
		{
			name:   "can read max u8",
			bytes:  []byte{0xcc, 0xff},
			format: format.UINT8,
			value:  uint8(math.MaxUint8),
		},
	})
}

func TestReadU16(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can read min u16",
			bytes:  []byte{0x0},
			format: format.UINT16,
			value:  uint16(0),
		},
		{
			name:   "can read max u16",
			bytes:  []byte{0xcd, 0xff, 0xff},
			format: format.UINT16,
			value:  uint16(math.MaxUint16),
		},
	})
}

func TestReadU32(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can read min u32",
			bytes:  []byte{0x0},
			format: format.UINT32,
			value:  uint32(0),
		},
		{
			name:   "can read max u32",
			bytes:  []byte{0xce, 0xff, 0xff, 0xff, 0xff},
			format: format.UINT32,
			value:  uint32(math.MaxUint32),
		},
	})
}

func TestReadU64(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can read min u64",
			bytes:  []byte{0x0},
			format: format.UINT64,
			value:  uint64(0),
		},
		{
			name:   "can read max u64",
			bytes:  []byte{0xcf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			format: format.UINT64,
			value:  uint64(math.MaxUint64),
		},
	})
}

func TestReadF32(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can read negative f32",
			bytes:  []byte{0xca, 0xbf, 0x0, 0x0, 0x0},
			format: format.FLOAT32,
			value:  float32(-0.5),
		},
		{
			name:   "can read zero f32",
			bytes:  []byte{0xca, 0x0, 0x0, 0x0, 0x0},
			format: format.FLOAT32,
			value:  float32(0),
		},
		{
			name:   "can read f32",
			bytes:  []byte{0xca, 0x3d, 0xe3, 0x8e, 0x39},
			format: format.FLOAT32,
			value:  float32(0.1111111111111),
		},
		{
			name:   "can read max f32",
			bytes:  []byte{0xca, 0x7f, 0x7f, 0xff, 0xff},
			format: format.FLOAT32,
			value:  float32(math.MaxFloat32),
		},
	})
}

func TestReadF64(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can read negative f64",
			bytes:  []byte{0xcb, 0xbf, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			format: format.FLOAT64,
			value:  float64(-0.5),
		},
		{
			name:   "can read zero f64",
			bytes:  []byte{0xcb, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			format: format.FLOAT64,
			value:  float64(0),
		},
		{
			name:   "can read f64",
			bytes:  []byte{0xcb, 0x3f, 0xbc, 0x71, 0xc7, 0x1c, 0x71, 0xc3, 0xfc},
			format: format.FLOAT64,
			value:  float64(0.1111111111111),
		},
		{
			name:   "can read max f64",
			bytes:  []byte{0xcb, 0x7f, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			format: format.FLOAT64,
			value:  float64(math.MaxFloat64),
		},
	})
}

func TestReadBytes(t *testing.T) {
	ctx := NewContext("")
	wri := NewWriteEncoder(ctx)
	wri.WriteBytes(nil)
	t.Logf("%#+v", wri.Buffer())
	runcases(t, []readcase{
		{
			name:   "can read nil",
			bytes:  []byte{0xc0},
			format: format.BIN8,
			value:  []byte{},
		},
		{
			name:   "can read bytes",
			bytes:  []byte{0xc4, 0x1, 0x1},
			format: format.BIN8,
			value:  []byte{1},
		},
	})
}

func TestReadString(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "can empty string",
			bytes:  []byte{0xa0},
			format: format.STR8,
			value:  "",
		},
		{
			name:   "can read string",
			bytes:  []byte{0xab, 0x73, 0x6f, 0x6d, 0x65, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67},
			format: format.STR8,
			value:  "some string",
		},
	})
}

func TestReadArray(t *testing.T) {
	runcases(t, []readcase{
		{
			name:   "nil",
			bytes:  []byte{192},
			format: format.ARRAY16,
			value:  []any{},
			fn:     nil,
		},
		{
			name:   "[]int8",
			bytes:  []byte{146, 208, 128, 127},
			format: format.ARRAY16,
			value:  []any{int8(math.MinInt8), int8(math.MaxInt8)},
			fn: func(reader Read) any {
				return reader.ReadI8()
			},
		},
		{
			name:   "[]int16",
			bytes:  []byte{146, 209, 128, 0, 209, 127, 255},
			format: format.ARRAY16,
			value:  []any{int16(math.MinInt16), int16(math.MaxInt16)},
			fn: func(reader Read) any {
				return reader.ReadI16()
			},
		},
		{
			name:   "[]int32",
			bytes:  []byte{146, 210, 128, 0, 0, 0, 210, 127, 255, 255, 255},
			format: format.ARRAY16,
			value:  []any{int32(math.MinInt32), int32(math.MaxInt32)},
			fn: func(reader Read) any {
				return reader.ReadI32()
			},
		},
		{
			name:   "[]int64",
			bytes:  []byte{146, 211, 128, 0, 0, 0, 0, 0, 0, 0, 211, 127, 255, 255, 255, 255, 255, 255, 255},
			format: format.ARRAY16,
			value:  []any{int64(math.MinInt64), int64(math.MaxInt64)},
			fn: func(reader Read) any {
				return reader.ReadI64()
			},
		},

		{
			name:   "[]uint8",
			bytes:  []byte{146, 0, 204, 255},
			format: format.ARRAY16,
			value:  []any{uint8(0), uint8(math.MaxUint8)},
			fn: func(reader Read) any {
				return reader.ReadU8()
			},
		},
		{
			name:   "[]uint16",
			bytes:  []byte{146, 0, 205, 255, 255},
			format: format.ARRAY16,
			value:  []any{uint16(0), uint16(math.MaxUint16)},
			fn: func(reader Read) any {
				return reader.ReadU16()
			},
		},
		{
			name:   "[]uint32",
			bytes:  []byte{146, 0, 206, 255, 255, 255, 255},
			format: format.ARRAY16,
			value:  []any{uint32(0), uint32(math.MaxUint32)},
			fn: func(reader Read) any {
				return reader.ReadU32()
			},
		},
		{
			name:   "[]uint64",
			bytes:  []byte{146, 0, 207, 255, 255, 255, 255, 255, 255, 255, 255},
			format: format.ARRAY16,
			value:  []any{uint64(0), uint64(math.MaxUint64)},
			fn: func(reader Read) any {
				return reader.ReadU64()
			},
		},
		{
			name:   "[]float32",
			bytes:  []byte{146, 202, 63, 26, 203, 4, 202, 63, 112, 197, 52},
			format: format.ARRAY16,
			value:  []any{float32(0.6046603), float32(0.9405091)},
			fn: func(reader Read) any {
				return reader.ReadF32()
			},
		},
		{
			name:   "[]float64",
			bytes:  []byte{146, 203, 63, 229, 68, 19, 113, 217, 165, 93, 203, 63, 220, 3, 130, 93, 189, 166, 190},
			format: format.ARRAY16,
			value:  []any{float64(0.6645600532184904), float64(0.4377141871869802)},
			fn: func(reader Read) any {
				return reader.ReadF64()
			},
		},
	})
}
