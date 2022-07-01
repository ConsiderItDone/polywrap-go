package msgpack

import (
	"bytes"
	"encoding/binary"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/format"
)

type DataView struct {
	buf     *bytes.Buffer
	context *Context
}

func NewDataView(context *Context) *DataView {
	return &DataView{
		buf:     new(bytes.Buffer),
		context: context,
	}
}

func NewDataViewWithBuf(context *Context, data []byte) *DataView {
	return &DataView{
		buf:     bytes.NewBuffer(data),
		context: context,
	}
}

func (dw *DataView) WriteFormat(value format.Format) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteUint8 error" + err.Error())
	}
}

func (dw *DataView) WriteUint8(value uint8) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteUint8 error" + err.Error())
	}
}

func (dw *DataView) WriteUint16(value uint16) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteUint16 error" + err.Error())
	}
}

func (dw *DataView) WriteUint32(value uint32) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteUint32 error" + err.Error())
	}
}

func (dw *DataView) WriteUint64(value uint64) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteUint64 error" + err.Error())
	}
}

func (dw *DataView) WriteInt8(value int8) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteInt8 error" + err.Error())
	}
}

func (dw *DataView) WriteInt16(value int16) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteInt16 error" + err.Error())
	}
}

func (dw *DataView) WriteInt32(value int32) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteInt32 error" + err.Error())
	}
}

func (dw *DataView) WriteInt64(value int64) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteInt64 error" + err.Error())
	}
}

func (dw *DataView) WriteFloat32(value float32) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteFloat32 error" + err.Error())
	}
}

func (dw *DataView) WriteFloat64(value float64) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteFloat64 error" + err.Error())
	}
}

func (dw *DataView) WriteString(value string) {
	dw.buf.WriteString(value)
}

func (dw *DataView) WriteBytes(value []byte) {
	dw.buf.Write(value)
}

func (dw *DataView) ReadUint8() uint8 {
	var result uint8
	err := binary.Read(dw.buf, binary.BigEndian, &result)
	if err != nil {
		panic("ReadUint8 error" + err.Error())
	}

	return result
}
