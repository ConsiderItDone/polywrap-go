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
		panic("WriteUint8 error " + err.Error())
	}
}

func (dw *DataView) PeekFormat() format.Format {
	f := dw.ReadUint8()
	dw.buf.UnreadByte()
	return format.Format(f)
}

func (dw *DataView) ReadFormat() format.Format {
	return format.Format(dw.ReadUint8())
}

func (dw *DataView) WriteUint8(value uint8) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteUint8 error " + err.Error())
	}
}

func (dw *DataView) ReadUint8() uint8 {
	var result uint8
	err := binary.Read(dw.buf, binary.BigEndian, &result)
	if err != nil {
		panic("ReadUint8 error " + err.Error())
	}
	return result
}

func (dw *DataView) WriteUint16(value uint16) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteUint16 error " + err.Error())
	}
}

func (dw *DataView) ReadUint16() uint16 {
	var result uint16
	err := binary.Read(dw.buf, binary.BigEndian, &result)
	if err != nil {
		panic("ReadUint16 error " + err.Error())
	}
	return result
}

func (dw *DataView) WriteUint32(value uint32) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteUint32 error " + err.Error())
	}
}

func (dw *DataView) ReadUint32() uint32 {
	var result uint32
	err := binary.Read(dw.buf, binary.BigEndian, &result)
	if err != nil {
		panic("ReadUint8 error " + err.Error())
	}
	return result
}

func (dw *DataView) WriteUint64(value uint64) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteUint64 error " + err.Error())
	}
}

func (dw *DataView) ReadUint64() uint64 {
	var result uint64
	err := binary.Read(dw.buf, binary.BigEndian, &result)
	if err != nil {
		panic("ReadUint8 error " + err.Error())
	}
	return result
}

func (dw *DataView) WriteInt8(value int8) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteInt8 error " + err.Error())
	}
}

func (dw *DataView) ReadInt8() int8 {
	var result int8
	err := binary.Read(dw.buf, binary.BigEndian, &result)
	if err != nil {
		panic("ReadInt8 error " + err.Error())
	}
	return result
}

func (dw *DataView) WriteInt16(value int16) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteInt16 error " + err.Error())
	}
}

func (dw *DataView) ReadInt16() int16 {
	var result int16
	err := binary.Read(dw.buf, binary.BigEndian, &result)
	if err != nil {
		panic("ReadInt16 error " + err.Error())
	}
	return result
}

func (dw *DataView) WriteInt32(value int32) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteInt32 error " + err.Error())
	}
}

func (dw *DataView) ReadInt32() int32 {
	var result int32
	err := binary.Read(dw.buf, binary.BigEndian, &result)
	if err != nil {
		panic("ReadInt8 error " + err.Error())
	}
	return result
}

func (dw *DataView) WriteInt64(value int64) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteInt64 error " + err.Error())
	}
}

func (dw *DataView) ReadInt64() int64 {
	var result int64
	err := binary.Read(dw.buf, binary.BigEndian, &result)
	if err != nil {
		panic("ReadInt8 error " + err.Error())
	}
	return result
}

func (dw *DataView) WriteFloat32(value float32) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteFloat32 error " + err.Error())
	}
}

func (dw *DataView) ReadFloat32() float32 {
	var result float32
	err := binary.Read(dw.buf, binary.BigEndian, &result)
	if err != nil {
		panic("ReadFloat32 error " + err.Error())
	}
	return result
}

func (dw *DataView) WriteFloat64(value float64) {
	err := binary.Write(dw.buf, binary.BigEndian, value)
	if err != nil {
		panic("WriteFloat64 error " + err.Error())
	}
}

func (dw *DataView) ReadFloat64() float64 {
	var result float64
	err := binary.Read(dw.buf, binary.BigEndian, &result)
	if err != nil {
		panic("ReadFloat64 error " + err.Error())
	}
	return result
}

func (dw *DataView) WriteString(value string) {
	dw.buf.WriteString(value)
}

func (dw *DataView) ReadString() string {
	return string(dw.buf.Bytes())
}

func (dw *DataView) WriteBytes(value []byte) {
	dw.buf.Write(value)
}

func (dw *DataView) ReadBytes(ln uint32) []byte {
	tmp := make([]byte, ln)
	err := binary.Read(dw.buf, binary.BigEndian, tmp)
	if err != nil {
		panic("ReadFloat64 error " + err.Error())
	}
	return tmp
}
