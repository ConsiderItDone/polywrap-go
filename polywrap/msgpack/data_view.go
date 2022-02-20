package msgpack

import (
	"bytes"
	"encoding/binary"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/format"
)

type DataView struct {
	bufferBytes []byte
	buf         *bytes.Buffer
	context     *Context
}

func NewDataView(context *Context) *DataView {
	return &DataView{
		buf:     new(bytes.Buffer),
		context: context,
	}
}

func (dw *DataView) WriteFormat(value format.Format) {
	err := binary.Write(dw.buf, binary.LittleEndian, value)
	if err != nil {
		panic("WriteUint8 error" + err.Error())
	}
}

func (dw *DataView) WriteUint8(value uint8) {
	err := binary.Write(dw.buf, binary.LittleEndian, value)
	if err != nil {
		panic("WriteUint8 error" + err.Error())
	}
}

func (dw *DataView) WriteString(value string) {
	dw.buf.WriteString(value)
}

func (dw *DataView) Test1() {
	err := binary.Write(dw.buf, binary.LittleEndian, format.NIL)
	if err != nil {
		panic("Test1 error" + err.Error())
	}
}

func (dw *DataView) GetB() *bytes.Buffer {
	return dw.buf
}
