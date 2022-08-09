package msgpack

import (
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/big"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
	"github.com/valyala/fastjson"
)

type (
	Write interface {
		Context() *Context
		View() *DataView

		WriteNil()

		WriteBool(bool)
		WriteOptionalBool(container.Option[bool])

		WriteInt8(int8)
		WriteOptionalInt8(container.Option[int8])

		WriteInt16(int16)
		WriteOptionalInt16(container.Option[int16])

		WriteInt32(int32)
		WriteOptionalInt32(container.Option[int32])

		WriteInt64(int64)
		WriteOptionalInt64(container.Option[int64])

		WriteUint8(uint8)
		WriteOptionalUint8(container.Option[uint8])

		WriteUint16(uint16)
		WriteOptionalUint16(container.Option[uint16])

		WriteUint32(uint32)
		WriteOptionalUint32(container.Option[uint32])

		WriteUint64(uint64)
		WriteOptionalUint64(container.Option[uint64])

		WriteFloat32(float32)
		WriteOptionalFloat32(container.Option[float32])

		WriteFloat64(float64)
		WriteOptionalFloat64(container.Option[float64])

		WriteBytesLength(uint32)
		WriteBytes([]byte)
		WriteOptionalBytes(container.Option[[]byte])

		WriteStringLength(uint32)
		WriteString(string)
		WriteOptionalString(container.Option[string])
	}

	BigIntWriter         func(Write, *big.Int)
	OptionalBigIntWriter func(Write, container.Option[*big.Int])

	JsonWriter         func(*fastjson.Value)
	OptionalJsonWriter func(container.Option[*fastjson.Value])

	ArrayLengthWriter          func(Write, uint32)
	ArrayItemWriter[T any]     func(Write, T)
	ArrayWriter[T any]         func(Write, []T, ArrayItemWriter[T])
	OptionalArrayWriter[T any] func(Write, container.Option[[]T], func(Write, T))

	MapLengthWriter                     func(Write, uint32)
	MapItemWriter[K Ordered, V any]     func(Write, K, V)
	MapWriter[K Ordered, V any]         func(Write, map[K]V, MapItemWriter[K, V])
	OptionalMapWriter[K Ordered, V any] func(Write, container.Option[map[K]V], func(Write, K, V))
)
