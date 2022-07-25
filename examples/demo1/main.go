package demo1

import (
	"github.com/consideritdone/polywrap-go/examples/demo1/wrap/moduleTypes"
	"github.com/consideritdone/polywrap-go/examples/demo1/wrap/sampleResult"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/big"
)

func SampleMethod(args *moduleTypes.ArgsSampleMethod) sampleResult.SampleResult {
	result := "0"
	if num, ok := new(big.Int).SetString(args.Arg, 10); ok {
		result = num.Add(num, big.NewInt(1)).String()
	}
	return sampleResult.SampleResult{Value: result}
}
