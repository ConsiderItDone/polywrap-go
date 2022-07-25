package msgpack

import (
	"math"

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

func (we *WriteEncoder) WriteI8(value int8) {
	we.WriteI64(int64(value))
}

func (we *WriteEncoder) WriteI16(value int16) {
	we.WriteI64(int64(value))
}

func (we *WriteEncoder) WriteI32(value int32) {
	we.WriteI64(int64(value))
}

func (we *WriteEncoder) WriteI64(value int64) {
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

func (we *WriteEncoder) WriteU8(value uint8) {
	we.WriteU64(uint64(value))
}

func (we *WriteEncoder) WriteU16(value uint16) {
	we.WriteU64(uint64(value))
}

func (we *WriteEncoder) WriteU32(value uint32) {
	we.WriteU64(uint64(value))
}

func (we *WriteEncoder) WriteU64(value uint64) {
	if value < 1<<7 {
		// positive fixed int
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

func (we *WriteEncoder) WriteFloat32(value float32) {
	we.view.WriteFormat(format.FLOAT32)
	we.view.WriteFloat32(value)
}

func (we *WriteEncoder) WriteFloat64(value float64) {
	we.view.WriteFormat(format.FLOAT64)
	we.view.WriteFloat64(value)
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

func (we *WriteEncoder) WriteMapLength(length uint32) {
	if length < 16 {
		we.view.WriteUint8(uint8(length) | uint8(format.FIXMAP))
	} else if length <= math.MaxUint16 {
		we.view.WriteUint8(uint8(format.MAP16))
		we.view.WriteUint16(uint16(length))
	} else {
		we.view.WriteUint8(uint8(format.MAP32))
		we.view.WriteUint32(length)
	}
}

func (we *WriteEncoder) WriteMap(value map[any]any, fn func(encoder Write, key any, value any)) {
	we.WriteMapLength(uint32(len(value)))
	for key := range value {
		fn(we, key, value[key])
	}
}

func (we *WriteEncoder) WriteArrayLength(length uint32) {
	if length < 16 {
		we.view.WriteUint8(uint8(length) | uint8(format.FIXARRAY))
	} else if length <= math.MaxUint16 {
		we.view.WriteFormat(format.ARRAY16)
		we.view.WriteUint16(uint16(length))
	} else {
		we.view.WriteFormat(format.ARRAY32)
		we.view.WriteUint32(length)
	}
}

func (we *WriteEncoder) WriteArray(value []any, fn func(encoder Write, item any)) {
	if len(value) == 0 {
		we.WriteNil()
		return
	}
	we.WriteArrayLength(uint32(len(value)))
	for i := range value {
		fn(we, value[i])
	}
}

func (we *WriteEncoder) WriteJson(value *fastjson.Value) {
	if value == nil {
		we.WriteNil()
		return
	}
	we.WriteString(value.String())
}

func (we *WriteEncoder) WriteBigInt(value *big.Int) {
	if value == nil {
		we.WriteNil()
		return
	}
	we.WriteString(value.String())
}

func (we *WriteEncoder) WriteOptionalBool(value container.Option[bool]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteBool(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalI8(value container.Option[int8]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteI8(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalI16(value container.Option[int16]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteI16(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalI32(value container.Option[int32]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteI32(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalI64(value container.Option[int64]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteI64(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalU8(value container.Option[uint8]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteU8(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalU16(value container.Option[uint16]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteU16(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalU32(value container.Option[uint32]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteU32(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalU64(value container.Option[uint64]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteU64(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalFloat32(value container.Option[float32]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteFloat32(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalFloat64(value container.Option[float64]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteFloat64(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalString(value container.Option[string]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteString(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalBytes(value container.Option[[]byte]) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteBytes(value.MustGet())
}

func (we *WriteEncoder) WriteOptionalArray(value container.Option[[]any], fn func(encoder Write, item any)) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	we.WriteArray(value.MustGet(), fn)
}
