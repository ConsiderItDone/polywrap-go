package msgpack

type Write interface {
	Context() *Context
	WriteNil()
	WriteBool(value bool)
	WriteI8(value int8)
	WriteI16(value int16)
	//WriteI64(value int64)
	WriteU8(value uint8)
	WriteU16(value uint16)
	WriteU32(value uint32)
	WriteStringLength(length uint32)
	WriteString(value string)
	WriteMapLength(length uint32)
}
