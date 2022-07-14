package sampleResult

import "github.com/consideritdone/polywrap-go/polywrap/msgpack"

func serializeSampleResult(args SampleResult) []byte {
	context := msgpack.NewContext("Serializing (encoding) object-type: SampleResult")
	encoder := msgpack.NewWriteEncoder(context)
	writeSampleResult(encoder, args)

	return encoder.Buffer()
}

func writeSampleResult(writer msgpack.Write, args SampleResult) {
	writer.WriteMapLength(1)
	writer.Context().Push("value", "string", "writing property")
	writer.WriteString("value")
	writer.WriteString(args.Value)
	writer.Context().Pop()
}
