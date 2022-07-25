package module

import "github.com/consideritdone/polywrap-go/examples/demo1"

func SampleMethodWrapped(argsBuf []byte, envSize uint32) []byte {
	args := deserializeSampleMethodArgs(argsBuf)

	result := demo1.SampleMethod(args)

	return serializeSampleMethodResult(result)
}
