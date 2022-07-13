package wrap

import "github.com/consideritdone/polywrap-go/polywrap/msgpack"

type Args_sampleMethod struct {
	arg string
}

func deserializesampleMethodArgs(argsBuf []byte) *Args_sampleMethod {
	c := msgpack.NewContext("Deserializing module-type: sampleMethod")

	return &Args_sampleMethod{}
}
