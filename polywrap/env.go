package polywrap

import "unsafe"

//go:wasm-module wrap
//export __wrap_load_env
func __wrap_load_env(envPtr uint32)

func WrapLoadEnv(envSize uint32) []byte {
	envBuf := make([]byte, envSize)
	envPtr := unsafe.Pointer(&envBuf)

	__wrap_load_env(*(*uint32)(envPtr))

	return envBuf
}
