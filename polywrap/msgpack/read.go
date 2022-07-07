package msgpack

type Read interface {
	Context() *Context
	IsNil()
	ReadBool(value bool)
	ReadI8(value int8)
	ReadI16(value int16)
	ReadI32(value int32)
	ReadI64(value int64)
	ReadU8(value int8)
	ReadU16(value int16)
	ReadU32(value int32)
	ReadU64(value int64)
	ReadString(value string)
	ReadBytes(value []byte)
}
