package msgpack

import (
	"math"
	"sort"

	"github.com/consideritdone/polywrap-go/polywrap/msgpack/big"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/format"
	"github.com/valyala/fastjson"
)

type WriteEncoder struct {
	context *Context
	view    *DataView
}

func NewWriteEncoder(context *Context) *WriteEncoder {
	return &WriteEncoder{context: context, view: NewDataView(context)}
}

func (we *WriteEncoder) Context() *Context {
	return we.context
}

func (we *WriteEncoder) View() *DataView {
	return we.view
}

func (we *WriteEncoder) Buffer() []byte {
	return we.view.buf.Bytes()
}

func (we *WriteEncoder) WriteNil() {
	we.view.WriteFormat(format.NIL)
}

func (we *WriteEncoder) WriteBool(value bool) {
	if value {
		we.view.WriteFormat(format.TRUE)
	} else {
		we.view.WriteFormat(format.FALSE)
	}
}

func (we *WriteEncoder) WriteOptionalBool(value container.Option[bool]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteBool(value.Unwrap())
}

func (we *WriteEncoder) WriteInt8(value int8) {
	we.WriteInt64(int64(value))
}

func (we *WriteEncoder) WriteOptionalInt8(value container.Option[int8]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteInt8(value.Unwrap())
}

func (we *WriteEncoder) WriteInt16(value int16) {
	we.WriteInt64(int64(value))
}

func (we *WriteEncoder) WriteOptionalInt16(value container.Option[int16]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteInt16(value.Unwrap())
}

func (we *WriteEncoder) WriteInt32(value int32) {
	we.WriteInt64(int64(value))
}

func (we *WriteEncoder) WriteOptionalInt32(value container.Option[int32]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteInt32(value.Unwrap())
}

func (we *WriteEncoder) WriteInt64(value int64) {
	if value >= 0 && value < 1<<7 {
		// positive fixed int
		we.view.WriteInt8(int8(value))
	} else if value < 0 && value >= -(1<<5) {
		// negative fixed int
		we.view.WriteInt8(int8(uint8(value) | uint8(format.NEGATIVE_FIXINT)))
	} else if value <= math.MaxInt8 && value >= math.MinInt8 {
		we.view.WriteFormat(format.INT8)
		we.view.WriteInt8(int8(value))
	} else if value <= math.MaxInt16 && value >= math.MinInt16 {
		we.view.WriteFormat(format.INT16)
		we.view.WriteInt16(int16(value))
	} else if value <= math.MaxInt32 && value >= math.MinInt32 {
		we.view.WriteFormat(format.INT32)
		we.view.WriteInt32(int32(value))
	} else {
		we.view.WriteFormat(format.INT64)
		we.view.WriteInt64(value)
	}
}

func (we *WriteEncoder) WriteOptionalInt64(value container.Option[int64]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteInt64(value.Unwrap())
}

func (we *WriteEncoder) WriteUint8(value uint8) {
	we.WriteUint64(uint64(value))
}

func (we *WriteEncoder) WriteOptionalUint8(value container.Option[uint8]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteUint8(value.Unwrap())
}

func (we *WriteEncoder) WriteUint16(value uint16) {
	we.WriteUint64(uint64(value))
}

func (we *WriteEncoder) WriteOptionalUint16(value container.Option[uint16]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteUint16(value.Unwrap())
}

func (we *WriteEncoder) WriteUint32(value uint32) {
	we.WriteUint64(uint64(value))
}

func (we *WriteEncoder) WriteOptionalUint32(value container.Option[uint32]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteUint32(value.Unwrap())
}

func (we *WriteEncoder) WriteUint64(value uint64) {
	if value < 1<<7 {
		we.view.WriteInt8(int8(value))
	} else if value <= math.MaxUint8 {
		we.view.WriteFormat(format.UINT8)
		we.view.WriteUint8(uint8(value))
	} else if value <= math.MaxUint16 {
		we.view.WriteFormat(format.UINT16)
		we.view.WriteUint16(uint16(value))
	} else if value <= math.MaxUint32 {
		we.view.WriteFormat(format.UINT32)
		we.view.WriteUint32(uint32(value))
	} else {
		we.view.WriteFormat(format.UINT64)
		we.view.WriteUint64(value)
	}
}

func (we *WriteEncoder) WriteOptionalUint64(value container.Option[uint64]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteUint64(value.Unwrap())
}

func (we *WriteEncoder) WriteFloat32(value float32) {
	we.view.WriteFormat(format.FLOAT32)
	we.view.WriteFloat32(value)
}

func (we *WriteEncoder) WriteOptionalFloat32(value container.Option[float32]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteFloat32(value.Unwrap())
}

func (we *WriteEncoder) WriteFloat64(value float64) {
	we.view.WriteFormat(format.FLOAT64)
	we.view.WriteFloat64(value)
}

func (we *WriteEncoder) WriteOptionalFloat64(value container.Option[float64]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteFloat64(value.Unwrap())
}

func (we *WriteEncoder) WriteBytesLength(length uint32) {
	if length < math.MaxUint8 {
		we.view.WriteFormat(format.BIN8)
		we.view.WriteUint8(uint8(length))
	} else if length < math.MaxUint16 {
		we.view.WriteFormat(format.BIN16)
		we.view.WriteUint16(uint16(length))
	} else {
		we.view.WriteFormat(format.BIN32)
		we.view.WriteUint32(length)
	}
}

func (we *WriteEncoder) WriteBytes(value []byte) {
	if len(value) == 0 {
		we.WriteNil()
		return
	}
	we.WriteBytesLength(uint32(len(value)))
	we.view.WriteBytes(value)
}

func (we *WriteEncoder) WriteOptionalBytes(value container.Option[[]byte]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteBytes(value.Unwrap())
}

func (we *WriteEncoder) WriteStringLength(length uint32) {
	if length < 32 {
		we.view.WriteUint8(uint8(length) | uint8(format.FIXSTR))
	} else if length <= math.MaxUint8 {
		we.view.WriteUint8(uint8(format.STR8))
		we.view.WriteUint8(uint8(length))
	} else if length <= math.MaxUint16 {
		we.view.WriteUint8(uint8(format.STR16))
		we.view.WriteUint16(uint16(length))
	} else {
		we.view.WriteUint8(uint8(format.STR32))
		we.view.WriteUint32(length)
	}
}

func (we *WriteEncoder) WriteString(value string) {
	we.WriteStringLength(uint32(len(value)))
	we.view.WriteString(value)
}

func (we *WriteEncoder) WriteOptionalString(value container.Option[string]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteString(value.Unwrap())
}

func WriteBigInt(encoder Write, value *big.Int) {
	if value == nil {
		encoder.WriteNil()
		return
	}
	encoder.WriteString(value.String())
}

func WriteOptionalBigInt(encoder Write, value container.Option[*big.Int]) {
	if value.IsNone() {
		encoder.WriteNil()
		return
	}
	WriteBigInt(encoder, value.Unwrap())
}

func WriteJson(encoder Write, value *fastjson.Value) {
	if value == nil {
		encoder.WriteNil()
		return
	}
	encoder.WriteString(value.String())
}

func WriteOptionalJson(encoder Write, value container.Option[*fastjson.Value]) {
	if value.IsNone() {
		encoder.WriteNil()
		return
	}
	WriteJson(encoder, value.Unwrap())
}

func WriteArrayLength(encoder Write, length uint32) {
	if length < 16 {
		encoder.View().WriteUint8(uint8(length) | uint8(format.FIXARRAY))
	} else if length <= math.MaxUint16 {
		encoder.View().WriteFormat(format.ARRAY16)
		encoder.View().WriteUint16(uint16(length))
	} else {
		encoder.View().WriteFormat(format.ARRAY32)
		encoder.View().WriteUint32(length)
	}
}

func WriteArray[T any](encoder Write, value []T, fn ArrayItemWriter[T]) {
	if len(value) == 0 {
		encoder.WriteNil()
		return
	}
	WriteArrayLength(encoder, uint32(len(value)))
	for i := range value {
		fn(encoder, value[i])
	}
}

func WriteOptionalArray[T any](encoder Write, value container.Option[[]T], fn ArrayItemWriter[T]) {
	if value.IsNone() {
		encoder.WriteNil()
		return
	}
	WriteArray(encoder, value.Unwrap(), fn)
}

func WriteMapLength(encoder Write, length uint32) {
	if length < 16 {
		encoder.View().WriteUint8(uint8(length) | uint8(format.FIXMAP))
	} else if length <= math.MaxUint16 {
		encoder.View().WriteUint8(uint8(format.MAP16))
		encoder.View().WriteUint16(uint16(length))
	} else {
		encoder.View().WriteUint8(uint8(format.MAP32))
		encoder.View().WriteUint32(length)
	}
}

func WriteMap[K Ordered, V any](encoder Write, value map[K]V, fn MapItemWriter[K, V]) {
	i := 0
	ln := len(value)
	keys := make([]K, ln)
	for k := range value {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	WriteMapLength(encoder, uint32(ln))
	for i := range keys {
		fn(encoder, keys[i], value[keys[i]])
	}
}

func WriteOptionalMap[K Ordered, V any](encoder Write, value container.Option[map[K]V], fn MapItemWriter[K, V]) {
	if value.IsNone() {
		encoder.WriteNil()
		return
	}
	WriteMap(encoder, value.Unwrap(), fn)
}
