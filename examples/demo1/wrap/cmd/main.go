package main

import (
	"github.com/consideritdone/polywrap-go/examples/demo1/wrap/module"
	"github.com/consideritdone/polywrap-go/polywrap"
)

//export _w3_invoke
func _wrap_invoke(methodSize, argsSize, envSize uint32) bool {
	args := polywrap.WrapInvokeArgs(methodSize, argsSize)

	if args.Method == "sampleMethod" {
		return polywrap.WrapInvoke(args, envSize, module.SampleMethodWrapped)
	} else {
		return polywrap.WrapInvoke(args, envSize, nil)
	}
}

func main() {
}
