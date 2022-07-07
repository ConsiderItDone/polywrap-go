package msgpack

import (
	"math"

	"github.com/consideritdone/polywrap-go/polywrap/msgpack/format"
)

type ReadDecoder struct {
	context *Context
	view    *DataView
}

func NewReadDecoder(context *Context, data []byte) *ReadDecoder {
	return &ReadDecoder{context: context, view: NewDataViewWithBuf(context, data)}
}

func (rd *ReadDecoder) Context() *Context {
	return rd.context
}

func (rd *ReadDecoder) IsNil() bool {
	return rd.view.ReadFormat() == format.NIL
}

func (rd *ReadDecoder) ReadBool() bool {
	f := rd.view.ReadFormat()
	if f != format.TRUE && f != format.FALSE {
		panic(rd.context.printWithContext("Property must be of type 'bool'. Found ..."))
	}
	return f == format.TRUE
}

func (rd *ReadDecoder) ReadI8() int8 {
	v := rd.ReadI64()
	if math.MinInt8 > v || v > math.MaxInt8 {
		panic(rd.context.printWithContext("int8 overflow"))
	}
	return int8(v)
}

func (rd *ReadDecoder) ReadI16() int16 {
	v := rd.ReadI64()
	if math.MinInt16 > v || v > math.MaxInt16 {
		panic(rd.context.printWithContext("int16 overflow"))
	}
	return int16(v)
}

func (rd *ReadDecoder) ReadI32() int32 {
	v := rd.ReadI64()
	if math.MinInt32 > v || v > math.MaxInt32 {
		panic(rd.context.printWithContext("int32 overflow"))
	}
	return int32(v)
}

func (rd *ReadDecoder) ReadI64() int64 {
	f := rd.view.ReadFormat()
	if isFixedInt(uint8(f)) {
		return int64(f)
	}
	if isNegativeFixedInt(uint8(f)) {
		return int64(f)
	}
	switch f {
	case format.INT8:
		return int64(rd.view.ReadInt8())
	case format.INT16:
		return int64(rd.view.ReadInt16())
	case format.INT32:
		return int64(rd.view.ReadInt32())
	case format.INT64:
		return rd.view.ReadInt64()
	default:
		panic(rd.context.printWithContext("Property must be of type 'int'. Found ..."))
	}
}

func (rd *ReadDecoder) ReadU8() uint8 {
	v := rd.ReadU64()
	if 0 > v || v > math.MaxUint8 {
		panic(rd.context.printWithContext("uint8 overflow"))
	}
	return uint8(v)
}

func (rd *ReadDecoder) ReadU16() uint16 {
	v := rd.ReadU64()
	if 0 > v || v > math.MaxUint16 {
		panic(rd.context.printWithContext("uint16 overflow"))
	}
	return uint16(v)
}

func (rd *ReadDecoder) ReadU32() uint32 {
	v := rd.ReadU64()
	if 0 > v || v > math.MaxUint32 {
		panic(rd.context.printWithContext("uint32 overflow"))
	}
	return uint32(v)
}

func (rd *ReadDecoder) ReadU64() uint64 {
	f := rd.view.ReadFormat()
	if isFixedInt(uint8(f)) {
		return uint64(f)
	}
	if isNegativeFixedInt(uint8(f)) {
		panic(rd.context.printWithContext("Unsigned integer cannot be negative. Found ..."))
	}
	switch f {
	case format.UINT8:
		return uint64(rd.view.ReadUint8())
	case format.UINT16:
		return uint64(rd.view.ReadUint16())
	case format.UINT32:
		return uint64(rd.view.ReadUint32())
	case format.UINT64:
		return rd.view.ReadUint64()
	default:
		panic(rd.context.printWithContext("Property must be of type 'uint'. Found ..."))
	}
}

func (rd *ReadDecoder) ReadF32() float32 {
	if rd.view.ReadFormat() != format.FLOAT32 {
		panic(rd.context.printWithContext("Property must be of type 'float32'. Found ..."))
	}
	return rd.view.ReadFloat32()
}

func (rd *ReadDecoder) ReadF64() float64 {
	if rd.view.ReadFormat() != format.FLOAT64 {
		panic(rd.context.printWithContext("Property must be of type 'float64'. Found ..."))
	}
	return rd.view.ReadFloat64()
}

func (rd *ReadDecoder) ReadBytes() []byte {
	return rd.view.ReadBytes()
}

func (rd *ReadDecoder) ReadString() string {
	return rd.view.ReadString()
}

func isFixedInt(v uint8) bool {
	return v>>7 == 0
}

func isNegativeFixedInt(v uint8) bool {
	return format.Format(v&0xe0) == format.NEGATIVE_FIXINT
}
