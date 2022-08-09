package container

type (
	Result[T any] struct {
		value *T
		err   error
	}
	ResultResolver[T any] func(T) (T, error)
	ResultRejecter[T any] func(error) (T, error)
)

func AsResult[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(value)
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: &value,
		err:   nil,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		value: nil,
		err:   err,
	}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) Error() error {
	return r.err
}

func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic("can't unwrap err value")
	}
	return *r.value
}

func (r Result[T]) Match(onValue ResultResolver[T], onError ResultRejecter[T]) Result[T] {
	if r.IsErr() {
		return AsResult(onError(r.err))
	}
	return AsResult(onValue(r.Unwrap()))
}
