package msgpack

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
	WriteArrayLength(length uint32)
	WriteArray(value []any, fn func(encoder Write, item any))
}
