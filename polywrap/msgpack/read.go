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
	ReadOptionalBool() container.Option

	ReadI8() int8
	ReadOptionalI8() container.Option

	ReadI16() int16
	ReadOptionalI16() container.Option

	ReadI32() int32
	ReadOptionalI32() container.Option

	ReadI64() int64
	ReadOptionalI64() container.Option

	ReadU8() uint8
	ReadOptionalU8() container.Option

	ReadU16() uint16
	ReadOptionalU16() container.Option

	ReadU32() uint32
	ReadOptionalU32() container.Option

	ReadU64() uint64
	ReadOptionalU64() container.Option

	ReadF32() float32
	ReadOptionalF32() container.Option

	ReadF64() float64
	ReadOptionalF64() container.Option

	ReadBytesLength() uint32
	ReadBytes() []byte
	ReadOptionalBytes() container.Option

	ReadStringLength() uint32
	ReadString() string
	ReadOptionalString() container.Option

	ReadJson() *fastjson.Value
	ReadOptionalJson() container.Option

	ReadBigInt() *big.Int
	ReadOptionalBigInt() container.Option

	ReadArrayLength() uint32
	ReadArray(fn func(reader Read) interface{}) []interface{}
	ReadOptionalArray(fn func(reader Read) interface{}) container.Option

	ReadMapLength() uint32
	ReadMap(fn func(reader Read) (interface{}, interface{})) map[interface{}]interface{}
	ReadOptionalMap(fn func(reader Read) (interface{}, interface{})) container.Option
}
