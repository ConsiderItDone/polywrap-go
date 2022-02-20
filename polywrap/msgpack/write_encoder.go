package msgpack

import "github.com/consideritdone/polywrap-go/polywrap/msgpack/format"

type WriteEncoder struct {
	context *Context
	view    *DataView
}

func NewWriteEncoder(context *Context, view *DataView) *WriteEncoder {
	return &WriteEncoder{context: context, view: view}
}

func (we *WriteEncoder) Context() *Context {
	return we.context
}

func (we *WriteEncoder) WriteNil() {
	we.view.WriteFormat(format.NIL)
}

func (we *WriteEncoder) WriteBool(value bool) {
	if value {
		we.view.WriteFormat(format.TRUE)
	} else {
		we.view.WriteFormat(format.FALSE)
	}
}

func (we *WriteEncoder) WriteI8(value int8) {
	we.WriteI32(int32(value))
}

func (we *WriteEncoder) WriteI32(value int32) {
	if value > 0 && value < 1<<7 {
		we.view.WriteUint8(uint8(value))
	}
}

func (we *WriteEncoder) WriteStringLength(length uint32) {
	if length < 32 {
		we.view.WriteUint8(uint8(length) | uint8(format.FIXSTR))
	}
}

func (we *WriteEncoder) WriteString(value string) {
	we.WriteStringLength(uint32(len(value)))
	we.view.WriteString(value)
}

func (we *WriteEncoder) WriteMapLength(length uint32) {
	if length < 16 {
		we.view.WriteUint8(uint8(length) | uint8(format.FIXMAP))
	}
}
