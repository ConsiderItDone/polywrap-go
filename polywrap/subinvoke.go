package polywrap

import (
	"errors"
	"unsafe"
)

//go:wasm-module wrap
//export __wrap_subinvoke
func __wrap_subinvoke(uriPtr, uriLen, methodPtr, methodLen, argsPtr, argsLen uint32) bool

// Subinvoke Result

//go:wasm-module wrap
//export __wrap_subinvoke_result_len
func __wrap_subinvoke_result_len() uint32

//go:wasm-module wrap
//export __wrap_subinvoke_result
func __wrap_subinvoke_result(ptr uint32)

// Subinvoke Error

//go:wasm-module wrap
//export __wrap_subinvoke_error_len
func __wrap_subinvoke_error_len() uint32

//go:wasm-module wrap
//export __wrap_subinvoke_error
func __wrap_subinvoke_error(ptr uint32)

func WrapSubinvoke(uri, method string, args []byte) ([]byte, error) {
	uriPtr := unsafe.Pointer(&uri)
	methodPtr := unsafe.Pointer(&method)
	argsPtr := unsafe.Pointer(&args)

	result := __wrap_subinvoke(*(*uint32)(uriPtr), uint32(len(uri)), *(*uint32)(methodPtr), uint32(len(method)),
		*(*uint32)(argsPtr), uint32(len(args)))

	if !result {
		errorLen := __wrap_subinvoke_error_len()
		errorBuf := make([]byte, errorLen)
		errorPtr := unsafe.Pointer(&errorBuf)

		__wrap_subinvoke_error(*(*uint32)(errorPtr))
		return nil, errors.New(string(errorBuf))
	}

	resultLen := __wrap_subinvoke_result_len()
	resultBuf := make([]byte, resultLen)
	resultPtr := unsafe.Pointer(&resultBuf)

	__wrap_subinvoke_result(*(*uint32)(resultPtr))
	return resultBuf, nil
}
