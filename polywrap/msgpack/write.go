package msgpack

type Write interface {
	Context() *Context
	WriteNil()
	WriteBool(value bool)
	WriteI8(value int8)
	//WriteInt16(value int16)
	WriteStringLength(length uint32)
	WriteString(value string)
	WriteMapLength(length uint32)
}
