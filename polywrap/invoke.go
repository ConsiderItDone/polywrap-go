package polywrap

import "unsafe"

//go:wasm-module wrap
//export __wrap_invoke_args
func __wrap_invoke_args(methodPtr, argsPtr uint32)

//go:wasm-module wrap
//export __wrap_invoke_result
func __wrap_invoke_result(ptr, len uint32)

//go:wasm-module wrap
//export __wrap_invoke_error
func __wrap_invoke_error(ptr, len uint32)

type invokeFunction func(argsBuf []byte) []byte

type InvokeArgs struct {
	Method string
	Args   []byte
}

func WrapInvokeArgs(methodSize, argsSize uint32) InvokeArgs {
	methodBuf := make([]byte, methodSize)
	methodPtr := unsafe.Pointer(&methodBuf)

	argsBuf := make([]byte, argsSize)
	argsPtr := unsafe.Pointer(&argsBuf)

	__wrap_invoke_args(*(*uint32)(methodPtr), *(*uint32)(argsPtr))

	method := string(methodBuf)

	return InvokeArgs{
		Method: method,
		Args:   argsBuf,
	}
}

func WrapInvoke(args InvokeArgs, fn invokeFunction) bool {
	if fn != nil {
		result := fn(args.Args)
		resultPtr := unsafe.Pointer(&result)

		__wrap_invoke_result(*(*uint32)(resultPtr), uint32(len(result)))

		return true
	} else {
		message := "Could not find invoke function \"" + args.Method + "\""
		messagePtr := unsafe.Pointer(&message)

		__wrap_invoke_error(*(*uint32)(messagePtr), uint32(len(message)))

		return false
	}
}
