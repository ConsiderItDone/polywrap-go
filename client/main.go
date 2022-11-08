package main

import (
	"fmt"
	wasmtime "github.com/bytecodealliance/wasmtime-go"
	"io/ioutil"
	"net/http"
	"strconv"
)

//func add(a int, b int) int {
//	return a + b
//}
//
//func sleep()

func main() {
	store := wasmtime.NewStore(wasmtime.NewEngine())

	// double_and_add(2, 3) = 2*2 + 3 =
	wasm, err := wasmtime.Wat2Wasm(`
      (module
		  (import "" "hello1" (func $hello1))
		  (import "" "hello2" (func $hello2))
		  (import "wrap" "external_double" (func $double (param i32) (result i32)))
		  (func (export "double_and_add") (param i32 i32) (result i32)
			call $hello1
			local.get 0
			call $double
			call $hello2
			local.get 1
			i32.add
		  )
		)
    `)
	check(err)

	//wasm2, err := wasmtime.Wat2Wasm(`
	//  (module
	//	  (func (export "external_double") (param i32) (result i32)
	//		local.get 0
	//		i32.mul
	//	  )
	//	)
	//`)
	//check(err)
	//
	//go add(1, 2)

	module, err := wasmtime.NewModule(store.Engine, wasm)
	check(err)

	linker := wasmtime.NewLinker(store.Engine)
	linker.FuncWrap("wrap", "external_double", func(a int32) int32 {
		resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:3000/double?value=%d", a))
		check(err)
		defer resp.Body.Close()
		resBody, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("client: response body: %s\n", resBody)

		double, err := strconv.Atoi(string(resBody))
		check(err)

		return int32(double)
	})
	linker.FuncWrap("", "hello1", func() {
		fmt.Println("hello 1")
	})
	linker.FuncWrap("", "hello2", func() {
		fmt.Println("hello 2")
	})

	instance, err := linker.Instantiate(store, module)
	check(err)

	run := instance.GetFunc(store, "double_and_add")
	if run == nil {
		panic("not a function")
	}
	// 6*6 + 10 = 36 + 10 = 46
	result, err := run.Call(store, 5, 10)
	check(err)
	fmt.Println(result)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//func main() {
//	wasmBytes, err := ioutil.ReadFile("simple.wasm")
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	engine := wasmer.NewEngine()
//	store := wasmer.NewStore(engine)
//
//	// Compiles the module
//	module, err := wasmer.NewModule(store, wasmBytes)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	// Instantiates the module
//	importObject := wasmer.NewImportObject()
//	instance, err := wasmer.NewInstance(module, importObject)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	// Gets the `sum` exported function from the WebAssembly instance.
//	sum, err := instance.Exports.GetFunction("sum")
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	// Calls that exported function with Go standard values. The WebAssembly
//	// types are inferred and values are casted automatically.
//	result, err := sum(5, 37)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	fmt.Println(result) // 42!
//}
