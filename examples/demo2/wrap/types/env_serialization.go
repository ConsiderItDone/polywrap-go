package types

import (
	"github.com/consideritdone/polywrap-go/polywrap/msgpack"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
)

func serializeEnv(env *Env) []byte {
	ctx := msgpack.NewContext("Serializing (encoding) env-type: Env")
	encoder := msgpack.NewWriteEncoder(ctx)
	writeEnv(encoder, env)
	return encoder.Buffer()
}

func writeEnv(writer msgpack.Write, env *Env) {
	writer.WriteMapLength(3)

	writer.Context().Push("prop", "string", "writing property")
	writer.WriteString("prop")
	writer.WriteString(env.Prop)
	writer.Context().Pop()

	writer.Context().Push("optProp", "string | null", "writing property")
	writer.WriteString("optProp")
	writer.WriteOptionalString(env.OptProp)
	writer.Context().Pop()

	writer.Context().Push("optMap", "Map<string, Option<i32>> | null", "writing property")
	writer.WriteString("optMap")
	writer.WriteOptionalMap(env.OptMap, func(encoder msgpack.Write, key, value interface{}) {
		writer.WriteString(key.(string))
		writer.WriteI32(value.(int32))
	})
	writer.Context().Pop()
}

func deserializeEnv(data []byte) *Env {
	ctx := msgpack.NewContext("Deserializing env-type Env")
	reader := msgpack.NewReadDecoder(ctx, data)
	return readEnv(reader)
}

func readEnv(reader msgpack.Read) *Env {
	var _prop string
	var _propSet bool
	var _optProp container.Option
	var _optMap container.Option

	for i := int32(reader.ReadMapLength()); i > 0; i-- {
		field := reader.ReadString()
		switch field {
		case "prop":
			reader.Context().Push(field, "string", "type found, reading property")
			_prop = reader.ReadString()
			_propSet = true
			reader.Context().Pop()
		case "optProp":
			reader.Context().Push(field, "string | null", "type found, reading property")
			_optProp = reader.ReadOptionalString()
			reader.Context().Pop()
		case "optMap":
			reader.Context().Push(field, "Map<string, Option<i32>> | null", "type found, reading property")
			_optMap = reader.ReadOptionalMap(func(reader msgpack.Read) (interface{}, interface{}) {
				return reader.ReadString(), reader.ReadI32()
			})
			reader.Context().Pop()
		}

		if !_propSet {
			panic(reader.Context().PrintWithContext("Missing required property: 'prop: String'"))
		}
	}

	return &Env{_prop, _optProp, _optMap}
}
