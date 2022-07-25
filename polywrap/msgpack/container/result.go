package container

type (
	Result[T any] struct {
		value T
		err   error
		isErr bool
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
		value: value,
		isErr: false,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		err:   err,
		isErr: true,
	}
}

func (r Result[T]) IsOk() bool {
	return !r.isErr
}

func (r Result[T]) IsError() bool {
	return r.isErr
}

func (r Result[T]) Error() error {
	return r.err
}

func (r Result[T]) Get() (T, error) {
	if r.isErr {
		return empty[T](), r.err
	}
	return r.value, nil
}

func (r Result[T]) MustGet() T {
	if r.isErr {
		panic(r.err)
	}
	return r.value
}

func (r Result[T]) OrElse(fallback T) T {
	if r.isErr {
		return fallback
	}
	return r.value
}

func (r Result[T]) Match(onSuccess ResultResolver[T], onError ResultRejecter[T]) Result[T] {
	var (
		v T
		e error
	)
	if r.IsOk() {
		v, e = onSuccess(r.value)
	} else {
		v, e = onError(r.err)
	}
	if e == nil {
		return Ok(v)
	}
	return Err[T](e)
}
