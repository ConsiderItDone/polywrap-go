package msgpack

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
	ReadStringLength() uint32
	ReadString() string
	ReadArrayLength() uint32
	ReadArray(fn func(reader Read) any) []any
}
