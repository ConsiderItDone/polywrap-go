package container

type (
	Result struct {
		value interface{}
		err   error
	}
	ResultResolver func(interface{}) (interface{}, error)
	ResultRejecter func(error) (interface{}, error)
)

func AsResult(value interface{}, err error) Result {
	if err != nil {
		return Err(err)
	}
	return Ok(value)
}

func Ok(value interface{}) Result {
	return Result{
		value: value,
		err:   nil,
	}
}

func Err(err error) Result {
	return Result{
		value: nil,
		err:   err,
	}
}

func (r Result) IsOk() bool {
	return r.err == nil
}

func (r Result) IsError() bool {
	return r.err != nil
}

func (r Result) Error() error {
	return r.err
}

func (r Result) Get() (interface{}, error) {
	return r.value, r.err
}

func (r Result) MustGet() interface{} {
	if r.IsError() {
		panic(r.err)
	}
	return r.value
}

func (r Result) OrElse(fallback interface{}) interface{} {
	if r.IsError() {
		return fallback
	}
	return r.value
}

func (r Result) Match(onSuccess ResultResolver, onError ResultRejecter) Result {
	var (
		v interface{}
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
	return Err(e)
}
