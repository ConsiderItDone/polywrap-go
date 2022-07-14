package module

import (
	"github.com/consideritdone/polywrap-go/examples/demo1/wrap/moduleTypes"
	"github.com/consideritdone/polywrap-go/examples/demo1/wrap/sampleResult"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack"
)

func deserializeSampleMethodArgs(argsBuf []byte) *moduleTypes.ArgsSampleMethod {
	context := msgpack.NewContext("Deserializing module-type: sampleMethod")
	reader := msgpack.NewReadDecoder(context, argsBuf)

	numFields := reader.ReadMapLength()

	var _arg string = ""
	var _argSet bool = false

	for i := numFields; i > 0; i-- {
		field := reader.ReadString()

		reader.Context().Push(field, "unknown", "searching for property type")
		if field == "arg" {
			reader.Context().Push(field, "string", "type found, reading property")
			_arg = reader.ReadString()
			_argSet = true
			reader.Context().Pop()
		}
		reader.Context().Pop()
	}

	if !_argSet {
		panic(reader.Context().PrintWithContext("Missing required argument: 'arg: String'"))
	}

	return &moduleTypes.ArgsSampleMethod{
		Arg: _arg,
	}
}

func serializeSampleMethodResult(result sampleResult.SampleResult) []byte {
	context := msgpack.NewContext("Serializing (encoding) module-type: sampleMethod")
	encoder := msgpack.NewWriteEncoder(context)
	writeSampleMethodResult(encoder, result)

	return encoder.Buffer()
}

func writeSampleMethodResult(writer msgpack.Write, result sampleResult.SampleResult) {
	writer.Context().Push("sampleMethod", "Types.SampleResult", "writing property")
	sampleResult.Write(writer, result)
	writer.Context().Pop()
}
