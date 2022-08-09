package msgpack

import (
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/big"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
	"github.com/valyala/fastjson"
)

type (
	Read interface {
		Context() *Context
		View() *DataView

		IsNextNil() bool

		ReadBool() bool
		ReadOptionalBool() container.Option[bool]

		ReadInt8() int8
		ReadOptionalInt8() container.Option[int8]

		ReadInt16() int16
		ReadOptionalInt16() container.Option[int16]

		ReadInt32() int32
		ReadOptionalInt32() container.Option[int32]

		ReadInt64() int64
		ReadOptionalInt64() container.Option[int64]

		ReadUint8() uint8
		ReadOptionalUint8() container.Option[uint8]

		ReadUint16() uint16
		ReadOptionalUint16() container.Option[uint16]

		ReadUint32() uint32
		ReadOptionalUint32() container.Option[uint32]

		ReadUint64() uint64
		ReadOptionalUint64() container.Option[uint64]

		ReadFloat32() float32
		ReadOptionalFloat32() container.Option[float32]

		ReadFloat64() float64
		ReadOptionalFloat64() container.Option[float64]

		ReadBytesLength() uint32
		ReadBytes() []byte
		ReadOptionalBytes() container.Option[[]byte]

		ReadStringLength() uint32
		ReadString() string
		ReadOptionalString() container.Option[string]
	}

	BigIntReader         func(Read) *big.Int
	OptionalBigIntReader func(Read) container.Option[*big.Int]

	JsonReader         func(Read) *fastjson.Value
	OptionalJsonReader func(Read) container.Option[*fastjson.Value]

	ArrayLengthReader          func(Read) uint32
	ArrayItemReader[T any]     func(Read) T
	ArrayReader[T any]         func(Read, ArrayItemReader[T]) []T
	OptionalArrayReader[T any] func(Read, ArrayItemReader[T]) container.Option[[]T]

	MapLengthReader                     func(Read) uint32
	MapItemReader[K Ordered, V any]     func(Read) (K, V)
	MapReader[K Ordered, V any]         func(Read, MapItemWriter[K, V]) map[K]V
	OptionalMapReader[K Ordered, V any] func(Read, MapItemWriter[K, V]) container.Option[map[K]V]
)
