package polywrap

import "unsafe"

//go:wasm-module w3
//export __w3_invoke_args
func __w3_invoke_args(methodPtr, argsPtr uint32)

//go:wasm-module w3
//export __w3_invoke_result
func __w3_invoke_result(ptr, len uint32)

//go:wasm-module w3
//export __w3_invoke_error
func __w3_invoke_error(ptr, len uint32)

type invokeFunction func(argsBuf []byte) []byte

type InvokeArgs struct {
	Method string
	Args   []byte
}

func W3InvokeArgs(methodSize, argsSize uint32) InvokeArgs {
	methodBuf := make([]byte, methodSize)
	methodPtr := unsafe.Pointer(&methodBuf)

	argsBuf := make([]byte, argsSize)
	argsPtr := unsafe.Pointer(&argsBuf)

	__w3_invoke_args(*(*uint32)(methodPtr), *(*uint32)(argsPtr))

	method := string(methodBuf)

	return InvokeArgs{
		Method: method,
		Args:   argsBuf,
	}
}

func W3Invoke(args InvokeArgs, fn invokeFunction) bool {
	if fn != nil {
		result := fn(args.Args)
		resultPtr := unsafe.Pointer(&result)

		__w3_invoke_result(*(*uint32)(resultPtr), uint32(len(result)))

		return true
	} else {
		message := "Could not find invoke function \"" + args.Method + "\""
		messagePtr := unsafe.Pointer(&message)

		__w3_invoke_error(*(*uint32)(messagePtr), uint32(len(message)))

		return false
	}
}
