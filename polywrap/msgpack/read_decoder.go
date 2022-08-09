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
func (rd *ReadDecoder) View() *DataView {
	return rd.view
}

func (rd ReadDecoder) IsNextNil() bool {
	return rd.view.PeekFormat() == format.NIL
}

func (rd *ReadDecoder) ReadBool() bool {
	f := rd.view.ReadFormat()
	if f != format.TRUE && f != format.FALSE {
		panic(rd.context.PrintWithContext("Property must be of type 'bool'. Found ..."))
	}
	return f == format.TRUE
}

func (rd *ReadDecoder) ReadOptionalBool() container.Option[bool] {
	if rd.IsNextNil() {
		return container.None[bool]()
	}
	return container.Some(rd.ReadBool())
}

func (rd *ReadDecoder) ReadInt8() int8 {
	v := rd.ReadInt64()
	if math.MinInt8 > v || v > math.MaxInt8 {
		panic(rd.context.PrintWithContext("int8 overflow"))
	}
	return int8(v)
}

func (rd *ReadDecoder) ReadOptionalInt8() container.Option[int8] {
	if rd.IsNextNil() {
		return container.None[int8]()
	}
	return container.Some(rd.ReadInt8())
}

func (rd *ReadDecoder) ReadInt16() int16 {
	v := rd.ReadInt64()
	if math.MinInt16 > v || v > math.MaxInt16 {
		panic(rd.context.PrintWithContext("int16 overflow"))
	}
	return int16(v)
}

func (rd *ReadDecoder) ReadOptionalInt16() container.Option[int16] {
	if rd.IsNextNil() {
		return container.None[int16]()
	}
	return container.Some(rd.ReadInt16())
}

func (rd *ReadDecoder) ReadInt32() int32 {
	v := rd.ReadInt64()
	if math.MinInt32 > v || v > math.MaxInt32 {
		panic(rd.context.PrintWithContext("int32 overflow"))
	}
	return int32(v)
}

func (rd *ReadDecoder) ReadOptionalInt32() container.Option[int32] {
	if rd.IsNextNil() {
		return container.None[int32]()
	}
	return container.Some(rd.ReadInt32())
}

func (rd *ReadDecoder) ReadInt64() int64 {
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
		panic(rd.context.PrintWithContext("Property must be of type 'int'. Found ..."))
	}
}

func (rd *ReadDecoder) ReadOptionalI64() container.Option[int64] {
	if rd.IsNextNil() {
		return container.None[int64]()
	}
	return container.Some(rd.ReadInt64())
}

func (rd *ReadDecoder) ReadUint8() uint8 {
	v := rd.ReadUint64()
	if v > math.MaxUint8 {
		panic(rd.context.PrintWithContext("uint8 overflow"))
	}
	return uint8(v)
}

func (rd *ReadDecoder) ReadOptionalUint8() container.Option[uint8] {
	if rd.IsNextNil() {
		return container.None[uint8]()
	}
	return container.Some(rd.ReadUint8())
}

func (rd *ReadDecoder) ReadUint16() uint16 {
	v := rd.ReadUint64()
	if v > math.MaxUint16 {
		panic(rd.context.PrintWithContext("uint16 overflow"))
	}
	return uint16(v)
}

func (rd *ReadDecoder) ReadOptionalU16() container.Option[uint16] {
	if rd.IsNextNil() {
		return container.None[uint16]()
	}
	return container.Some(rd.ReadUint16())
}

func (rd *ReadDecoder) ReadUint32() uint32 {
	v := rd.ReadUint64()
	if v > math.MaxUint32 {
		panic(rd.context.PrintWithContext("uint32 overflow"))
	}
	return uint32(v)
}

func (rd *ReadDecoder) ReadOptionalU32() container.Option[uint32] {
	if rd.IsNextNil() {
		return container.None[uint32]()
	}
	return container.Some(rd.ReadUint32())
}

func (rd *ReadDecoder) ReadUint64() uint64 {
	f := rd.view.ReadFormat()
	if isFixedInt(uint8(f)) {
		return uint64(f)
	}
	if isNegativeFixedInt(uint8(f)) {
		panic(rd.context.PrintWithContext("Unsigned integer cannot be negative. Found ..."))
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
		panic(rd.context.PrintWithContext("Property must be of type 'uint'. Found ..."))
	}
}

func (rd *ReadDecoder) ReadOptionalUint64() container.Option[uint64] {
	if rd.IsNextNil() {
		return container.None[uint64]()
	}
	return container.Some(rd.ReadUint64())
}

func (rd *ReadDecoder) ReadFloat32() float32 {
	if rd.view.ReadFormat() != format.FLOAT32 {
		panic(rd.context.PrintWithContext("Property must be of type 'float32'. Found ..."))
	}
	return rd.view.ReadFloat32()
}

func (rd *ReadDecoder) ReadOptionalF32() container.Option[float32] {
	if rd.IsNextNil() {
		return container.None[float32]()
	}
	return container.Some(rd.ReadFloat32())
}

func (rd *ReadDecoder) ReadFloat64() float64 {
	if rd.view.ReadFormat() != format.FLOAT64 {
		panic(rd.context.PrintWithContext("Property must be of type 'float64'. Found ..."))
	}
	return rd.view.ReadFloat64()
}

func (rd *ReadDecoder) ReadOptionalF64() container.Option[float64] {
	if rd.IsNextNil() {
		return container.None[float64]()
	}
	return container.Some(rd.ReadFloat64())
}

func (rd *ReadDecoder) ReadBytesLength() uint32 {
	switch rd.view.ReadFormat() {
	case format.NIL:
		return 0
	case format.BIN8:
		return uint32(rd.view.ReadUint8())
	case format.BIN16:
		return uint32(rd.view.ReadUint16())
	case format.BIN32:
		return uint32(rd.view.ReadUint32())
	}
	panic(rd.context.PrintWithContext("Property must be of type 'binary'. Found ..."))
}

func (rd *ReadDecoder) ReadBytes() []byte {
	if rd.IsNextNil() {
		return nil
	}
	return rd.view.ReadBytes(rd.ReadBytesLength())
}

func (rd *ReadDecoder) ReadOptionalBytes() container.Option[[]byte] {
	if rd.IsNextNil() {
		return container.None[[]byte]()
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
	panic(rd.context.PrintWithContext("Property must be of type 'string'. Found ..."))
}

func (rd *ReadDecoder) ReadString() string {
	ln := rd.ReadStringLength()
	if ln == 0 {
		return ""
	}
	return string(rd.view.ReadBytes(ln))
}

func (rd *ReadDecoder) ReadOptionalString() container.Option[string] {
	if rd.IsNextNil() {
		return container.None[string]()
	}
	return container.Some(rd.ReadString())
}

func ReadBigInt(decoder Read) *big.Int {
	tmp := decoder.ReadString()
	if tmp == "" {
		return nil
	}
	val, ok := new(big.Int).SetString(tmp, 10)
	if !ok {
		panic(decoder.Context().PrintWithContext("Property must be of type 'BigInt'. Found ..."))
	}
	return val
}

func ReadOptionalBigInt(decoder Read) container.Option[*big.Int] {
	if decoder.IsNextNil() {
		return container.None[*big.Int]()
	}
	return container.Some(ReadBigInt(decoder))
}

func ReadJson(decoder Read) *fastjson.Value {
	tmp := decoder.ReadString()
	if tmp == "" {
		return nil
	}
	return fastjson.MustParse(tmp)
}

func ReadOptionalJson(decoder Read) container.Option[*fastjson.Value] {
	if decoder.IsNextNil() {
		return container.None[*fastjson.Value]()
	}
	return container.Some(ReadJson(decoder))
}

func ReadArrayLength(decoder Read) uint32 {
	f := decoder.View().ReadFormat()
	if f == format.NIL {
		return 0
	}
	if isFixedArray(uint8(f)) {
		return uint32(f & format.FOUR_LEAST_SIG_BITS_IN_BYTE)
	}
	switch f {
	case format.ARRAY16:
		return uint32(decoder.View().ReadUint16())
	case format.ARRAY32:
		return decoder.View().ReadUint32()
	case format.NIL:
		return 0
	}
	panic(decoder.Context().PrintWithContext("Property must be of type 'array'. Found ..."))
}

func ReadArray[T any](decoder Read, fn ArrayItemReader[T]) []T {
	size := ReadArrayLength(decoder)
	data := make([]T, size)
	for i := uint32(0); i < size; i++ {
		data[i] = fn(decoder)
	}
	return data
}

func ReadOptionalArray[T any](decoder Read, fn ArrayItemReader[T]) container.Option[[]T] {
	if decoder.IsNextNil() {
		return container.None[[]T]()
	}
	return container.Some(ReadArray(decoder, fn))
}

func ReadMapLength(decoder Read) uint32 {
	f := decoder.View().ReadFormat()
	if f == format.NIL {
		return 0
	}
	if isFixedMap(uint8(f)) {
		return uint32(f & format.FOUR_LEAST_SIG_BITS_IN_BYTE)
	}
	switch f {
	case format.MAP16:
		return uint32(decoder.View().ReadUint16())
	case format.MAP32:
		return decoder.View().ReadUint32()
	case format.NIL:
		return 0
	}
	panic(decoder.Context().PrintWithContext("Property must be of type 'map'. Found ..."))
}

func ReadMap[K Ordered, V any](decoder Read, fn MapItemReader[K, V]) map[K]V {
	size := ReadMapLength(decoder)
	data := make(map[K]V)
	for i := uint32(0); i < size; i++ {
		k, v := fn(decoder)
		data[k] = v
	}
	return data
}

func ReadOptionalMap[K Ordered, V any](decoder Read, fn MapItemReader[K, V]) container.Option[map[K]V] {
	if decoder.IsNextNil() {
		return container.None[map[K]V]()
	}
	return container.Some(ReadMap(decoder, fn))
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
