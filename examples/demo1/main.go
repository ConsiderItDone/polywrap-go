package demo1

import (
	"github.com/consideritdone/polywrap-go/examples/demo1/wrap/moduleTypes"
	"github.com/consideritdone/polywrap-go/examples/demo1/wrap/sampleResult"
)

func SampleMethod(args *moduleTypes.ArgsSampleMethod) sampleResult.SampleResult {
	return sampleResult.SampleResult{Value: args.Arg}
}
