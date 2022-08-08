package types

import (
	"github.com/consideritdone/polywrap-go/polywrap/msgpack"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
)

type AnotherType struct {
	Prop     container.Option
	Circular container.Option
	Const    container.Option
}

func AnotherTypeToBuffer(anotherType *AnotherType) []byte {
	return serializeAnotherType(anotherType)
}

func AnotherTypeFromBuffer(data []byte) *AnotherType {
	return deserializeAnotherType(data)
}

func AnotherTypeWrite(writer msgpack.Write, anotherType *AnotherType) {
	writeAnotherType(writer, anotherType)
}

func AnotherTypeRead(reader msgpack.Read) *AnotherType {
	return readAnotherType(reader)
}
