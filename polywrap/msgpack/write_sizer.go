package msgpack

import "math"

type WriteSizer struct {
	length  int32
	context *Context
}

func NewWriteSizer(context *Context) *WriteSizer {
	return &WriteSizer{length: 0, context: context}
}

func (ws *WriteSizer) Context() *Context {
	return ws.context
}

func (ws *WriteSizer) WriteNil() {
	ws.length++
}

func (ws *WriteSizer) WriteBool(_ bool) {
	ws.length++
}
func (ws *WriteSizer) WriteI8(_ int8) {
	ws.length++
}
func (ws *WriteSizer) WriteStringLength(length uint32) {
	if length < 32 {
		ws.length++
	} else if length <= math.MaxUint8 {
		ws.length += 2
	} else if length <= math.MaxUint16 {
		ws.length += 3
	} else {
		ws.length += 5
	}
}
func (ws *WriteSizer) WriteString(value string) {
	ws.WriteStringLength(uint32(len(value)))
	ws.length += int32(len(value))
}

func (ws *WriteSizer) WriteMapLength(length uint32) {
	if length < 16 {
		ws.length++
	} else if length <= math.MaxUint16 {
		ws.length += 3
	} else {
		ws.length += 5
	}
}

func (ws *WriteSizer) Length() int32 {
	return ws.length
}
