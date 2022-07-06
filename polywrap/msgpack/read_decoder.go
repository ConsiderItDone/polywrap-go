package msgpack

import "github.com/consideritdone/polywrap-go/polywrap/msgpack/format"

type ReadDecoder struct {
	context *Context
	view    *DataView
}

func NewReadDecoder(context *Context, data []byte) *ReadDecoder {
	return &ReadDecoder{context: context, view: NewDataViewWithBuf(context, data)}
}

func (rd *ReadDecoder) Context() *Context {
	return rd.context
}

func (rd *ReadDecoder) IsNil() bool {
	return rd.view.ReadFormat() == format.NIL
}

func (rd *ReadDecoder) ReadBool() bool {
	f := rd.view.ReadFormat()
	if f != format.TRUE && f != format.FALSE {
		panic(rd.context.printWithContext("Property must be of type 'bool'. Found ..."))
	}
	return f == format.TRUE
}

func (rd *ReadDecoder) ReadI8() int8 {
	return rd.view.ReadInt8()
}

func (rd *ReadDecoder) ReadI16() int16 {
	return rd.view.ReadInt16()
}

func (rd *ReadDecoder) ReadI32() int32 {
	return rd.view.ReadInt32()
}

func (rd *ReadDecoder) ReadI64() int64 {
	return rd.view.ReadInt64()
}

func (rd *ReadDecoder) ReadU8() uint8 {
	return rd.view.ReadUint8()
}

func (rd *ReadDecoder) ReadU16() uint16 {
	return rd.view.ReadUint16()
}

func (rd *ReadDecoder) ReadU32() uint32 {
	return rd.view.ReadUint32()
}

func (rd *ReadDecoder) ReadU64() uint64 {
	return rd.view.ReadUint64()
}

func (rd *ReadDecoder) ReadF32() float32 {
	return rd.view.ReadFloat32()
}

func (rd *ReadDecoder) ReadF64() float64 {
	return rd.view.ReadFloat64()
}

func (rd *ReadDecoder) ReadBytes() []byte {
	return rd.view.ReadBytes()
}

func (rd *ReadDecoder) ReadString() string {
	return rd.view.ReadString()
}
