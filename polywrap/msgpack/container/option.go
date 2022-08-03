package container

import "errors"

var (
	ErrEmptyOption = errors.New("no such element")
)

type (
	Option struct {
		isValue bool
		value   interface{}
	}
	OptionResolver func(interface{}) (interface{}, bool)
	OptionRejecter func() (interface{}, bool)
)

func Some(value interface{}) Option {
	return Option{
		isValue: true,
		value:   value,
	}
}

func None() Option {
	return Option{
		isValue: false,
		value:   nil,
	}
}

func (o Option) IsSome() bool {
	return o.isValue
}

func (o Option) IsNone() bool {
	return !o.isValue
}

func (o Option) Get() (interface{}, bool) {
	if o.IsNone() {
		return nil, false
	}
	return o.value, true
}

func (o Option) MustGet() interface{} {
	if o.IsNone() {
		panic(ErrEmptyOption)
	}
	return o.value
}

func (o Option) OrElse(fallback interface{}) interface{} {
	if o.IsNone() {
		return fallback
	}
	return o.value
}

func (o Option) Match(onValue OptionResolver, onNone OptionRejecter) Option {
	var (
		v  interface{}
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
	return None()
}
