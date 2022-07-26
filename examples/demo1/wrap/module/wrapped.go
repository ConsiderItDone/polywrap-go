package module

import (
	"github.com/consideritdone/polywrap-go/examples/demo1/wrap/moduleTypes"
	"github.com/consideritdone/polywrap-go/examples/demo1/wrap/sampleResult"
)

type SampleMethodFn func(args *moduleTypes.ArgsSampleMethod) sampleResult.SampleResult

var SampleMethod SampleMethodFn

func SampleMethodWrapped(argsBuf []byte, envSize uint32) []byte {
	args := deserializeSampleMethodArgs(argsBuf)

	result := SampleMethod(args)

	return serializeSampleMethodResult(result)
}
