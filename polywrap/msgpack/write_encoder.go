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

func (we *WriteEncoder) WriteOptionalBool(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(bool)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'bool'"))
	}
	we.WriteBool(v)
}

func (we *WriteEncoder) WriteI8(value int8) {
	we.WriteI64(int64(value))
}

func (we *WriteEncoder) WriteOptionalI8(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(int8)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'int8'"))
	}
	we.WriteI8(v)
}

func (we *WriteEncoder) WriteI16(value int16) {
	we.WriteI64(int64(value))
}

func (we *WriteEncoder) WriteOptionalI16(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(int16)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'int16'"))
	}
	we.WriteI16(v)
}

func (we *WriteEncoder) WriteI32(value int32) {
	we.WriteI64(int64(value))
}

func (we *WriteEncoder) WriteOptionalI32(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(int32)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'int32'"))
	}
	we.WriteI32(v)
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

func (we *WriteEncoder) WriteOptionalI64(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(int64)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'int64'"))
	}
	we.WriteI64(v)
}

func (we *WriteEncoder) WriteU8(value uint8) {
	we.WriteU64(uint64(value))
}

func (we *WriteEncoder) WriteOptionalU8(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(uint8)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'uint8'"))
	}
	we.WriteU8(v)
}

func (we *WriteEncoder) WriteU16(value uint16) {
	we.WriteU64(uint64(value))
}

func (we *WriteEncoder) WriteOptionalU16(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(uint16)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'uint16'"))
	}
	we.WriteU16(v)
}

func (we *WriteEncoder) WriteU32(value uint32) {
	we.WriteU64(uint64(value))
}

func (we *WriteEncoder) WriteOptionalU32(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(uint32)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'uint32'"))
	}
	we.WriteU32(v)
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

func (we *WriteEncoder) WriteOptionalU64(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(uint64)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'uint64'"))
	}
	we.WriteU64(v)
}

func (we *WriteEncoder) WriteFloat32(value float32) {
	we.view.WriteFormat(format.FLOAT32)
	we.view.WriteFloat32(value)
}

func (we *WriteEncoder) WriteOptionalFloat32(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(float32)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'float32'"))
	}
	we.WriteFloat32(v)
}

func (we *WriteEncoder) WriteFloat64(value float64) {
	we.view.WriteFormat(format.FLOAT64)
	we.view.WriteFloat64(value)
}

func (we *WriteEncoder) WriteOptionalFloat64(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(float64)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'float64'"))
	}
	we.WriteFloat64(v)
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

func (we *WriteEncoder) WriteOptionalBytes(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().([]byte)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type '[]byte'"))
	}
	we.WriteBytes(v)
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

func (we *WriteEncoder) WriteOptionalString(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(string)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'string'"))
	}
	we.WriteString(v)
}

func (we *WriteEncoder) WriteJson(value *fastjson.Value) {
	if value == nil {
		we.WriteNil()
		return
	}
	we.WriteString(value.String())
}

func (we *WriteEncoder) WriteOptionalJson(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(*fastjson.Value)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type '*fastjson.Value'"))
	}
	we.WriteJson(v)
}

func (we *WriteEncoder) WriteBigInt(value *big.Int) {
	if value == nil {
		we.WriteNil()
		return
	}
	we.WriteString(value.String())
}

func (we *WriteEncoder) WriteOptionalBigInt(value container.Option) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(*big.Int)
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type '*big.Int'"))
	}
	we.WriteBigInt(v)
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

func (we *WriteEncoder) WriteArray(value []interface{}, fn func(encoder Write, item interface{})) {
	if len(value) == 0 {
		we.WriteNil()
		return
	}
	we.WriteArrayLength(uint32(len(value)))
	for i := range value {
		fn(we, value[i])
	}
}

func (we *WriteEncoder) WriteOptionalArray(value container.Option, fn func(encoder Write, item interface{})) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().([]interface{})
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type '[]interface{}'"))
	}
	we.WriteArray(v, fn)
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

func (we *WriteEncoder) WriteMap(value map[interface{}]interface{}, fn func(encoder Write, key interface{}, value interface{})) {
	we.WriteMapLength(uint32(len(value)))
	for key := range value {
		fn(we, key, value[key])
	}
}

func (we *WriteEncoder) WriteOptionalMap(value container.Option, fn func(encoder Write, key interface{}, value interface{})) {
	if value.IsNone() {
		we.WriteNil()
		return
	}
	v, ok := value.MustGet().(map[interface{}]interface{})
	if !ok {
		panic(we.context.PrintWithContext("Argument must be of type 'map[interface{}]interface{}'"))
	}
	we.WriteMap(v, fn)
}
