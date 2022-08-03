package container

import (
	"errors"
	"testing"
)

var err = errors.New("some test err")

func TestResultOk(t *testing.T) {
	compare(t, Ok(42), Result{value: 42, err: nil})
}

func TestResultErr(t *testing.T) {
	compare(t, Err(err), Result{value: nil, err: err})
}

func TestResultAsResult(t *testing.T) {
	compare(t, AsResult(42, nil), Ok(42))
	compare(t, AsResult(nil, err), Err(err))
}

func TestResultIsOk(t *testing.T) {
	compare(t, Ok(42).IsOk(), true)
	compare(t, Err(err).IsOk(), false)
}

func TestResultIsErr(t *testing.T) {
	compare(t, Ok(42).IsError(), false)
	compare(t, Err(err).IsError(), true)
}

func TestResultGet(t *testing.T) {
	v1, err1 := Ok(42).Get()
	v2, err2 := Err(err).Get()
	compare(t, v1, 42)
	compare(t, err1, nil)
	compare(t, v2, nil)
	compare(t, err2, err)
}

func TestResultOrElse(t *testing.T) {
	compare(t, Ok(42).OrElse(21), 42)
	compare(t, Err(err).OrElse(21), 21)
}

func TestResultMatch(t *testing.T) {
	onValue1 := func(i interface{}) (interface{}, error) {
		return i.(int) * 2, nil
	}
	onValue2 := func(i interface{}) (interface{}, error) {
		return 0, err
	}
	onNone1 := func(e error) (interface{}, error) {
		return 11, nil
	}
	onNone2 := func(e error) (interface{}, error) {
		return 11, err
	}

	compare(t, Ok(21).Match(onValue1, onNone1), Ok(42))
	compare(t, Ok(21).Match(onValue2, onNone1), Err(err))
	compare(t, Err(err).Match(onValue1, onNone1), Ok(11))
	compare(t, Err(err).Match(onValue1, onNone2), Err(err))
}
