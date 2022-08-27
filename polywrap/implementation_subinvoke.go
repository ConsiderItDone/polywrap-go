package polywrap

import (
	"errors"
	"unsafe"
)

// Implementation Subinvoke Interface

//go:wasm-module wrap
//export __wrap_subinvokeImplementation
func __wrap_subinvokeImplementation(interfaceUriPtr, interfaceUriLen, implUriPtr, implUriLen, methodPtr, methodLen, argsPtr, argsLen uint32) bool

// Implementation Subinvoke Result

//go:wasm-module wrap
//export __wrap_subinvokeImplementation_result_len
func __wrap_subinvokeImplementation_result_len() uint32

//go:wasm-module wrap
//export __wrap_subinvokeImplementation_result
func __wrap_subinvokeImplementation_result(ptr uint32)

// Subinvoke Error

//go:wasm-module wrap
//export __wrap_subinvokeImplementation_error_len
func __wrap_subinvokeImplementation_error_len() uint32

//go:wasm-module wrap
//export __wrap_subinvokeImplementation_error
func __wrap_subinvokeImplementation_error(ptr uint32)

func WrapSubinvokeImplementation(interfaceUri, implUri, method string, args []byte) ([]byte, error) {
	interfaceUriPtr := unsafe.Pointer(&interfaceUri)
	implUriPtr := unsafe.Pointer(&implUri)
	methodPtr := unsafe.Pointer(&method)
	argsPtr := unsafe.Pointer(&args)

	result := __wrap_subinvokeImplementation(*(*uint32)(interfaceUriPtr), uint32(len(interfaceUri)),
		*(*uint32)(implUriPtr), uint32(len(implUri)), *(*uint32)(methodPtr), uint32(len(method)),
		*(*uint32)(argsPtr), uint32(len(args)))

	if !result {
		errorLen := __wrap_subinvokeImplementation_error_len()
		errorBuf := make([]byte, errorLen)
		errorPtr := unsafe.Pointer(&errorBuf)

		__wrap_subinvokeImplementation_error(*(*uint32)(errorPtr))
		return nil, errors.New(string(errorBuf))
	}

	resultLen := __wrap_subinvokeImplementation_result_len()
	resultBuf := make([]byte, resultLen)
	resultPtr := unsafe.Pointer(&resultBuf)

	__wrap_subinvokeImplementation_result(*(*uint32)(resultPtr))
	return resultBuf, nil
}
