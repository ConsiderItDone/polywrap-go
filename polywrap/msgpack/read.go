package msgpack

import (
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/big"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
	"github.com/valyala/fastjson"
)

type Read interface {
	Context() *Context
	IsNil() bool
	ReadBool() bool
	ReadI8() int8
	ReadI16() int16
	ReadI32() int32
	ReadI64() int64
	ReadU8() uint8
	ReadU16() uint16
	ReadU32() uint32
	ReadU64() uint64
	ReadF32() float32
	ReadF64() float64
	ReadBytesLength() uint32
	ReadBytes() []byte
	ReadMapLength() uint32
	ReadMap(fn func(reader Read) (any, any)) map[any]any
	ReadStringLength() uint32
	ReadString() string
	ReadArrayLength() uint32
	ReadArray(fn func(reader Read) any) []any
	ReadJson() *fastjson.Value
	ReadBigInt() *big.Int
	ReadOptionalBool() container.Option[bool]
	ReadOptionalI8() container.Option[int8]
	ReadOptionalI16() container.Option[int16]
	ReadOptionalI32() container.Option[int32]
	ReadOptionalI64() container.Option[int64]
	ReadOptionalU8() container.Option[uint8]
	ReadOptionalU16() container.Option[uint16]
	ReadOptionalU32() container.Option[uint32]
	ReadOptionalU64() container.Option[uint64]
	ReadOptionalF32() container.Option[float32]
	ReadOptionalF64() container.Option[float64]
	ReadOptionalBytes() container.Option[[]byte]
	ReadOptionalString() container.Option[string]
	ReadOptionalArray(fn func(reader Read) any) container.Option[[]any]
}
