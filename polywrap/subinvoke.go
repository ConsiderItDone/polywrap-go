package polywrap

import "unsafe"

//go:wasm-module w3
//export __w3_subinvoke
func __w3_subinvoke(uriPtr, uriLen, modulePtr, moduleLen, methodPtr, methodLen, inputPtr, inputLen uint32)

func W3Subinvoke(uri, module, method string, input []byte) {
	uriPtr := unsafe.Pointer(&uri)
	modulePtr := unsafe.Pointer(&module)
	methodPtr := unsafe.Pointer(&method)
	inputPtr := unsafe.Pointer(&input)

	__w3_subinvoke(*(*uint32)(uriPtr), uint32(len(uri)), *(*uint32)(modulePtr), uint32(len(module)),
		*(*uint32)(methodPtr), uint32(len(method)), *(*uint32)(inputPtr), uint32(len(input)))
}
