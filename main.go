package main

import (
	"github.com/consideritdone/polywrap-go/polywrap"
	"github.com/consideritdone/polywrap-go/polywrap/msgpack"
)

//import "unsafe"

////go:wasm-module w3
////export __w3_abort
//func __w3_abort(msgPtr, msgLen, filePtr, fileLen, line, column uint32)

func main() {
}

//export multiply
func multiply(x, y int) int {
	return x * y
}

//export testAbort
func testAbort() {
	const message = "some text message"
	const file = "/some/path/to/file.go"
	polywrap.W3Abort(message, file, 10, 20)
}

//export testStr
func testStr() string {
	var s string
	s = "some long string"
	return s
}

//export _w3_invoke
func _w3_invoke(methodSize, argsSize uint32) bool {
	args := polywrap.W3InvokeArgs(methodSize, argsSize)

	if args.Method == "testHttp" {
		polywrap.W3Invoke(args, testHttpWrapped)
	}
	return true
}

//go:generate ...
func testHttpWrapped(argsBuf []byte) []byte {
	testHttp()

	return serializeResult()
}

func serializeResult() []byte {
	context := msgpack.NewContext("Serializing (sizing) imported query-type: get")
	sizer := msgpack.NewWriteSizer(context)
	writeTestHttpResult(sizer)

	context = msgpack.NewContext("Serializing (encoding) imported query-type: get")
	encoder := msgpack.NewWriteEncoder(context)
	writeTestHttpResult(encoder)

	return encoder.Buffer()
}

//export testHttp
func testHttp() {
	context := msgpack.NewContext("Serializing (sizing) imported query-type: get")
	sizer := msgpack.NewWriteSizer(context)
	writeArgs(sizer)

	context = msgpack.NewContext("Serializing (encoding) imported query-type: post")
	encoder := msgpack.NewWriteEncoder(context)
	writeArgs(encoder)

	polywrap.W3Subinvoke("w3://ens/http.web3api.eth", "query", "get", dataView.GetB().Bytes())

	//dw := msgpack.NewDataView(context)
	//dw.Test1()

	//buf := dataView.GetB().Bytes()
	//bufPtr := unsafe.Pointer(&buf)
	//return *(*int32)(bufPtr)

	//return sizer.Length()

}

func writeArgs(w msgpack.Write) {
	w.WriteMapLength(2)
	w.Context().Push("url", "string", "writing property")
	w.WriteString("url")
	//w.WriteString("https://google.com")
	w.WriteString("http://localhost:4040/ens")
	w.Context().Pop()
	w.Context().Push("request", "Types.HTTP_Request | null", "writing property")
	w.WriteString("request")
	w.WriteNil()
	w.Context().Pop()
}

func writeTestHttpResult(w msgpack.Write) {
	w.Context().Push("testHttp", "string", "writing property")
	w.WriteNil()
	w.Context().Pop()
}

////export testStr2
//func testStr2() *string {
//	var s = "another string"
//	//bytes := []byte{104, 101, 108, 108, 111}
//	p := unsafe.Pointer(&s)
//
//	__w3_abort(*(*uint32)(p), uint32(len(s)), 200, 20, 2, 3)
//
//	return &s
//}
