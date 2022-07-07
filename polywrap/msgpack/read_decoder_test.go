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
			case format.STR16:
				v = reader.ReadString()
			case format.BIN8:
				v = reader.ReadBytes()
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
