package sampleResult

import "github.com/consideritdone/polywrap-go/polywrap/msgpack"

type SampleResult struct {
	Value string
}

func toBuffer(args SampleResult) []byte {
	return serializeSampleResult(args)
}

func Write(writer msgpack.Write, args SampleResult) {
	writeSampleResult(writer, args)
}
