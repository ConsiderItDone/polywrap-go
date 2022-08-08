package types

import (
	"github.com/consideritdone/polywrap-go/polywrap/msgpack"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
)

func serializeAnotherType(anotherType *AnotherType) []byte {
	ctx := msgpack.NewContext("Serializing (encoding) env-type: AnotherType")
	encoder := msgpack.NewWriteEncoder(ctx)
	writeAnotherType(encoder, anotherType)
	return encoder.Buffer()
}

func writeAnotherType(writer msgpack.Write, anotherType *AnotherType) {
	writer.WriteMapLength(3)

	writer.Context().Push("prop", "string | null", "writing property")
	writer.WriteString("prop")
	writer.WriteOptionalString(anotherType.Prop)
	writer.Context().Pop()

	writer.Context().Push("circular", "Types.CustomType | null", "writing property")
	writer.WriteString("circular")
	if anotherType.Circular.IsSome() {
		CustomTypeWrite(writer, anotherType.Circular.MustGet().(*CustomType))
	} else {
		writer.WriteNil()
	}
	writer.Context().Pop()

	writer.Context().Push("const", "string | null", "writing property")
	writer.WriteString("const")
	writer.WriteOptionalString(anotherType.Const)
	writer.Context().Pop()
}

func deserializeAnotherType(data []byte) *AnotherType {
	ctx := msgpack.NewContext("Deserializing env-type AnotherType")
	reader := msgpack.NewReadDecoder(ctx, data)
	return readAnotherType(reader)
}

func readAnotherType(reader msgpack.Read) *AnotherType {
	var _prop container.Option
	var _circular container.Option
	var _const container.Option

	for i := int32(reader.ReadMapLength()); i > 0; i-- {
		field := reader.ReadString()
		switch field {
		case "prop":
			reader.Context().Push(field, "string | null", "type found, reading property")
			_prop = reader.ReadOptionalString()
			reader.Context().Pop()
		case "circular":
			reader.Context().Push(field, "Types.CustomType | null", "type found, reading property")
			var object *CustomType = nil
			if !reader.IsNil() {
				object = CustomTypeRead(reader)
			}
			if object != nil {
				_circular = container.Some(object)
			} else {
				_circular = container.None()
			}
			reader.Context().Pop()
		case "const":
			reader.Context().Push(field, "string | null", "type found, reading property")
			_const = reader.ReadOptionalString()
			reader.Context().Pop()
		}
	}
	return &AnotherType{_prop, _circular, _const}
}
