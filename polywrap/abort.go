package polywrap

import "unsafe"

//go:wasm-module w3
//export __w3_abort
func __w3_abort(msgPtr, msgLen, filePtr, fileLen, line, column uint32)

func W3Abort(msg string, file string, line uint32, column uint32) {
	msgPtr := unsafe.Pointer(&msg)
	filePtr := unsafe.Pointer(&file)

	__w3_abort(*(*uint32)(msgPtr), uint32(len(msg)), *(*uint32)(filePtr), uint32(len(file)), line, column)
}
