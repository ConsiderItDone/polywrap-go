package types

import (
	"github.com/consideritdone/polywrap-go/polywrap/msgpack"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
)

type Env struct {
	Prop    string
	OptProp container.Option
	OptMap  container.Option
}

func ToBuffer(env *Env) []byte {
	return serializeEnv(env)
}

func FromBuffer(data []byte) *Env {
	return deserializeEnv(data)
}

func Write(writer msgpack.Write, env *Env) {
	writeEnv(writer, env)
}

func Read(reader msgpack.Read) *Env {
	return readEnv(reader)
}
