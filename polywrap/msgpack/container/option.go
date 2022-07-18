package container

import "errors"

var (
	ErrEmptyOption = errors.New("no such element")
)

type Option[T any] struct {
	isValue bool
	value   T
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		isValue: true,
		value:   value,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		isValue: false,
	}
}

func (o Option[T]) IsSome() bool {
	return o.isValue
}

func (o Option[T]) IsNone() bool {
	return !o.isValue
}

func (o Option[T]) Get() (T, bool) {
	if o.IsNone() {
		return empty[T](), false
	}
	return o.value, true
}

func (o Option[T]) MustGet() T {
	if o.IsNone() {
		panic(ErrEmptyOption)
	}
	return o.value
}

func (o Option[T]) OrElse(fallback T) T {
	if o.IsNone() {
		return fallback
	}
	return o.value
}

func (o Option[T]) Match(onValue func(value T) (T, bool), onNone func() (T, bool)) Option[T] {
	var (
		v  T
		ok bool
	)
	if o.IsSome() {
		v, ok = onValue(o.value)
	} else {
		v, ok = onNone()
	}
	if ok {
		return Some(v)
	}
	return None[T]()
}
