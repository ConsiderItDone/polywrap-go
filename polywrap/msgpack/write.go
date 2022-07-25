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
	WriteI8(value int8)
	WriteI16(value int16)
	WriteI32(value int32)
	WriteI64(value int64)
	WriteU8(value uint8)
	WriteU16(value uint16)
	WriteU32(value uint32)
	WriteU64(value uint64)
	WriteFloat32(value float32)
	WriteFloat64(value float64)
	WriteStringLength(length uint32)
	WriteString(value string)
	WriteBytesLength(length uint32)
	WriteBytes(value []byte)
	WriteMapLength(length uint32)
	WriteMap(value map[any]any, fn func(encoder Write, key any, value any))
	WriteArrayLength(length uint32)
	WriteArray(value []any, fn func(encoder Write, item any))
	WriteJson(data *fastjson.Value)
	WriteBigInt(value *big.Int)
	WriteOptionalBool(value container.Option[bool])
	WriteOptionalI8(value container.Option[int8])
	WriteOptionalI16(value container.Option[int16])
	WriteOptionalI32(value container.Option[int32])
	WriteOptionalI64(value container.Option[int64])
	WriteOptionalU8(value container.Option[uint8])
	WriteOptionalU16(value container.Option[uint16])
	WriteOptionalU32(value container.Option[uint32])
	WriteOptionalU64(value container.Option[uint64])
	WriteOptionalFloat32(value container.Option[float32])
	WriteOptionalFloat64(value container.Option[float64])
	WriteOptionalString(value container.Option[string])
	WriteOptionalBytes(value container.Option[[]byte])
	WriteOptionalArray(value container.Option[[]any], fn func(encoder Write, item any))
}
