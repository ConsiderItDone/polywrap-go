package main

import (
	"errors"

	"github.com/consideritdone/polywrap-go/polywrap/msgpack/container"
)

func wrap(num int, err error) container.Result[int] {
	if err != nil {
		return container.Err[int](err)
	}
	return container.Ok(num)
}

func main() {
	wrap(0, errors.New("some error"))
	wrap(0, nil)
}
