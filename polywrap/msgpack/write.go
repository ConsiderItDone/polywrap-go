package msgpack

import (
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/big"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
	"github.com/valyala/fastjson"
)

type Write interface {
	Context() *Context
	WriteNil()

	WriteBool(value bool)
	WriteOptionalBool(value container.Option)

	WriteI8(value int8)
	WriteOptionalI8(value container.Option)

	WriteI16(value int16)
	WriteOptionalI16(value container.Option)

	WriteI32(value int32)
	WriteOptionalI32(value container.Option)

	WriteI64(value int64)
	WriteOptionalI64(value container.Option)

	WriteU8(value uint8)
	WriteOptionalU8(value container.Option)

	WriteU16(value uint16)
	WriteOptionalU16(value container.Option)

	WriteU32(value uint32)
	WriteOptionalU32(value container.Option)

	WriteU64(value uint64)
	WriteOptionalU64(value container.Option)

	WriteFloat32(value float32)
	WriteOptionalFloat32(value container.Option)

	WriteFloat64(value float64)
	WriteOptionalFloat64(value container.Option)

	WriteBytesLength(length uint32)
	WriteBytes(value []byte)
	WriteOptionalBytes(value container.Option)

	WriteStringLength(length uint32)
	WriteString(value string)
	WriteOptionalString(value container.Option)

	WriteJson(data *fastjson.Value)
	WriteOptionalJson(value container.Option)

	WriteBigInt(value *big.Int)
	WriteOptionalBigInt(value container.Option)

	WriteArrayLength(length uint32)
	WriteArray(value []interface{}, fn func(encoder Write, item interface{}))
	WriteOptionalArray(value container.Option, fn func(encoder Write, item interface{}))

	WriteMapLength(length uint32)
	WriteMap(value map[interface{}]interface{}, fn func(encoder Write, key interface{}, value interface{}))
	WriteOptionalMap(value container.Option, fn func(encoder Write, key interface{}, value interface{}))
}
