package container

import "errors"

var (
	ErrEmptyOption = errors.New("no such element")
)

type (
	Option[T any] struct {
		value *T
	}
	OptionResolver[T any] func(T) (T, bool)
	OptionRejecter[T any] func() (T, bool)
)

func Some[T any](value T) Option[T] {
	return Option[T]{
		value: &value,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		value: nil,
	}
}

func AsOption[T any](value T, ok bool) Option[T] {
	if !ok {
		return None[T]()
	}
	return Some(value)
}

func (o Option[T]) IsSome() bool {
	return o.value != nil
}

func (o Option[T]) IsNone() bool {
	return o.value == nil
}

func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("can't unwrap none value")
	}
	return *o.value
}

func (o Option[T]) Match(onValue OptionResolver[T], onNone OptionRejecter[T]) Option[T] {
	if o.IsNone() {
		return AsOption(onNone())
	}
	return AsOption(onValue(o.Unwrap()))
}
