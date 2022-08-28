package msgpack

import (
	"math"

	"github.com/consideritdone/polywrap-go/polywrap/msgpack/big"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/format"
	"github.com/valyala/fastjson"
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

func (rd ReadDecoder) IsNil() bool {
	return rd.view.PeekFormat() == format.NIL
}

func (rd *ReadDecoder) ReadBool() bool {
	f := rd.view.ReadFormat()
	if f != format.TRUE && f != format.FALSE {
		panic(rd.context.PrintWithContext("Property must be of type 'bool'. Found " + format.ToString(f)))
	}
	return f == format.TRUE
}

func (rd *ReadDecoder) ReadOptionalBool() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadBool())
}

func (rd *ReadDecoder) ReadI8() int8 {
	v := rd.ReadI64()
	if math.MinInt8 > v || v > math.MaxInt8 {
		panic(rd.context.PrintWithContext("int8 overflow"))
	}
	return int8(v)
}

func (rd *ReadDecoder) ReadOptionalI8() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadI8())
}

func (rd *ReadDecoder) ReadI16() int16 {
	v := rd.ReadI64()
	if math.MinInt16 > v || v > math.MaxInt16 {
		panic(rd.context.PrintWithContext("int16 overflow"))
	}
	return int16(v)
}

func (rd *ReadDecoder) ReadOptionalI16() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadI16())
}

func (rd *ReadDecoder) ReadI32() int32 {
	v := rd.ReadI64()
	if math.MinInt32 > v || v > math.MaxInt32 {
		panic(rd.context.PrintWithContext("int32 overflow"))
	}
	return int32(v)
}

func (rd *ReadDecoder) ReadOptionalI32() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadI32())
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
	case format.UINT8:
		return int64(rd.view.ReadUint8())
	case format.UINT16:
		return int64(rd.view.ReadUint16())
	case format.UINT32:
		return int64(rd.view.ReadUint32())
	default:
		panic(rd.context.PrintWithContext("Property must be of type 'int'. Found " + format.ToString(f)))
	}
}

func (rd *ReadDecoder) ReadOptionalI64() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadI64())
}

func (rd *ReadDecoder) ReadU8() uint8 {
	v := rd.ReadU64()
	if 0 > v || v > math.MaxUint8 {
		panic(rd.context.PrintWithContext("uint8 overflow"))
	}
	return uint8(v)
}

func (rd *ReadDecoder) ReadOptionalU8() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadU8())
}

func (rd *ReadDecoder) ReadU16() uint16 {
	v := rd.ReadU64()
	if 0 > v || v > math.MaxUint16 {
		panic(rd.context.PrintWithContext("uint16 overflow"))
	}
	return uint16(v)
}

func (rd *ReadDecoder) ReadOptionalU16() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadU16())
}

func (rd *ReadDecoder) ReadU32() uint32 {
	v := rd.ReadU64()
	if 0 > v || v > math.MaxUint32 {
		panic(rd.context.PrintWithContext("uint32 overflow"))
	}
	return uint32(v)
}

func (rd *ReadDecoder) ReadOptionalU32() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadU32())
}

func (rd *ReadDecoder) ReadU64() uint64 {
	f := rd.view.ReadFormat()
	if isFixedInt(uint8(f)) {
		return uint64(f)
	}
	if isNegativeFixedInt(uint8(f)) {
		panic(rd.context.PrintWithContext("Unsigned integer cannot be negative. Found " + format.ToString(f)))
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
		panic(rd.context.PrintWithContext("Property must be of type 'uint'. Found " + format.ToString(f)))
	}
}

func (rd *ReadDecoder) ReadOptionalU64() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadU64())
}

func (rd *ReadDecoder) ReadF32() float32 {
	f := rd.view.ReadFormat()
	if f != format.FLOAT32 {
		panic(rd.context.PrintWithContext("Property must be of type 'float32'. Found " + format.ToString(f)))
	}
	return rd.view.ReadFloat32()
}

func (rd *ReadDecoder) ReadOptionalF32() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadF32())
}

func (rd *ReadDecoder) ReadF64() float64 {
	f := rd.view.ReadFormat()
	if f != format.FLOAT64 {
		panic(rd.context.PrintWithContext("Property must be of type 'float64'. Found " + format.ToString(f)))
	}
	return rd.view.ReadFloat64()
}

func (rd *ReadDecoder) ReadOptionalF64() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadF64())
}

func (rd *ReadDecoder) ReadBytesLength() uint32 {
	f := rd.view.ReadFormat()
	switch f {
	case format.NIL:
		return 0
	case format.BIN8:
		return uint32(rd.view.ReadUint8())
	case format.BIN16:
		return uint32(rd.view.ReadUint16())
	case format.BIN32:
		return rd.view.ReadUint32()
	}
	panic(rd.context.PrintWithContext("Property must be of type 'binary'. Found " + format.ToString(f)))
}

func (rd *ReadDecoder) ReadBytes() []byte {
	if rd.IsNil() {
		return nil
	}
	return rd.view.ReadBytes(rd.ReadBytesLength())
}

func (rd *ReadDecoder) ReadOptionalBytes() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadBytes())
}

func (rd *ReadDecoder) ReadStringLength() uint32 {
	f := rd.view.ReadFormat()
	if isFixedString(uint8(f)) {
		return uint32(uint8(f) & 0x1f)
	}
	if isFixedArray(uint8(f)) {
		return uint32(f & format.FOUR_LEAST_SIG_BITS_IN_BYTE)
	}
	switch f {
	case format.NIL:
		return 0
	case format.STR8:
		return uint32(rd.view.ReadUint8())
	case format.STR16:
		return uint32(rd.view.ReadUint16())
	case format.STR32:
		return rd.view.ReadUint32()
	}
	panic(rd.context.PrintWithContext("Property must be of type 'string'. Found " + format.ToString(f)))
}

func (rd *ReadDecoder) ReadString() string {
	ln := rd.ReadStringLength()
	if ln == 0 {
		return ""
	}
	return string(rd.view.ReadBytes(ln))
}

func (rd *ReadDecoder) ReadOptionalString() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadString())
}

func (rd *ReadDecoder) ReadJson() *fastjson.Value {
	tmp := rd.ReadString()
	if tmp == "" {
		return nil
	}
	return fastjson.MustParse(tmp)
}

func (rd *ReadDecoder) ReadOptionalJson() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadJson())
}

func (rd *ReadDecoder) ReadBigInt() *big.Int {
	tmp := rd.ReadString()
	if tmp == "" {
		return nil
	}
	val, ok := new(big.Int).SetString(tmp, 10)
	if !ok {
		panic(rd.context.PrintWithContext("Property must be of type 'BigInt'"))
	}
	return val
}

func (rd *ReadDecoder) ReadOptionalBigInt() container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadBigInt())
}

func (rd *ReadDecoder) ReadArrayLength() uint32 {
	f := rd.view.ReadFormat()
	if f == format.NIL {
		return 0
	}
	if isFixedArray(uint8(f)) {
		return uint32(f & format.FOUR_LEAST_SIG_BITS_IN_BYTE)
	}
	switch f {
	case format.ARRAY16:
		return uint32(rd.view.ReadUint16())
	case format.ARRAY32:
		return rd.view.ReadUint32()
	case format.NIL:
		return 0
	}
	panic(rd.context.PrintWithContext("Property must be of type 'array'. Found " + format.ToString(f)))
}

func (rd *ReadDecoder) ReadArray(fn func(reader Read) interface{}) []interface{} {
	size := rd.ReadArrayLength()
	data := make([]interface{}, size)
	for i := uint32(0); i < size; i++ {
		data[i] = fn(rd)
	}
	return data
}

func (rd *ReadDecoder) ReadOptionalArray(fn func(reader Read) interface{}) container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadArray(fn))
}

func (rd *ReadDecoder) ReadMapLength() uint32 {
	f := rd.view.ReadFormat()
	if f == format.NIL {
		return 0
	}
	if isFixedMap(uint8(f)) {
		return uint32(f & format.FOUR_LEAST_SIG_BITS_IN_BYTE)
	}
	switch f {
	case format.MAP16:
		return uint32(rd.view.ReadUint16())
	case format.MAP32:
		return rd.view.ReadUint32()
	case format.NIL:
		return 0
	}
	panic(rd.context.PrintWithContext("Property must be of type 'map'. Found " + format.ToString(f)))
}

func (rd *ReadDecoder) ReadMap(fn func(reader Read) (interface{}, interface{})) map[interface{}]interface{} {
	size := rd.ReadMapLength()
	data := make(map[interface{}]interface{})
	for i := uint32(0); i < size; i++ {
		k, v := fn(rd)
		data[k] = v
	}
	return data
}

func (rd *ReadDecoder) ReadOptionalMap(fn func(reader Read) (interface{}, interface{})) container.Option {
	if rd.IsNil() {
		return container.None()
	}
	return container.Some(rd.ReadMap(fn))
}

func isFixedInt(v uint8) bool {
	return v>>7 == 0
}

func isFixedMap(v uint8) bool {
	return format.Format(v&0xf0) == format.FIXMAP
}

func isFixedString(v uint8) bool {
	return format.Format(v&0xe0) == format.FIXSTR
}

func isFixedArray(v uint8) bool {
	return format.Format(v&0xf0) == format.FIXARRAY
}

func isNegativeFixedInt(v uint8) bool {
	return format.Format(v&0xe0) == format.NEGATIVE_FIXINT
}
