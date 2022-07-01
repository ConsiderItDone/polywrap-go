package msgpack

import "github.com/consideritdone/polywrap-go/polywrap/msgpack/format"

type ReadDecoder struct {
	context *Context
	view    *DataView
}

func NewReadDecoder(context *Context, data []byte) *ReadDecoder {
	return &ReadDecoder{context: context, view: NewDataViewWithBuf(context, data)}
}

func (rd *ReadDecoder) ReadBool() bool {
	value := rd.view.ReadUint8()
	if value == uint8(format.TRUE) {
		return true
	} else if value == uint8(format.FALSE) {
		return false
	}

	panic(rd.context.printWithContext("Property must be of type 'bool'. Found ..."))
}
