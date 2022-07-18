package container

import (
	"errors"
	"testing"
)

var err = errors.New("some test err")

func TestResultOk(t *testing.T) {
	compare(t, Ok(42), Result[int]{isErr: false, value: 42})
}

func TestResultErr(t *testing.T) {
	compare(t, Err[int](err), Result[int]{isErr: true, err: err})
}

func TestResultAsResult(t *testing.T) {
	compare(t, AsResult(42, nil), Ok(42))
	compare(t, AsResult(0, err), Err[int](err))
}

func TestResultIsOk(t *testing.T) {
	compare(t, Ok(42).IsOk(), true)
	compare(t, Err[int](err).IsOk(), false)
}

func TestResultIsErr(t *testing.T) {
	compare(t, Ok(42).IsError(), false)
	compare(t, Err[int](err).IsError(), true)
}

func TestResultGet(t *testing.T) {
	v1, err1 := Ok(42).Get()
	v2, err2 := Err[int](err).Get()
	compare(t, v1, 42)
	compare(t, err1, nil)
	compare(t, v2, 0)
	compare(t, err2, err)
}

func TestResultMustGet(t *testing.T) {
	compare(t, Ok(42).MustGet(), 42)

	defer func() {
		compare(t, recover(), err)
	}()
	compare(t, Err[int](err).MustGet(), 0)
}

func TestResultOrElse(t *testing.T) {
	compare(t, Ok(42).OrElse(21), 42)
	compare(t, Err[int](err).OrElse(21), 21)
}

func TestResultMatch(t *testing.T) {
	onValue1 := func(i int) (int, error) {
		return i * 2, nil
	}
	onValue2 := func(i int) (int, error) {
		return 0, err
	}
	onNone1 := func(e error) (int, error) {
		return 11, nil
	}
	onNone2 := func(e error) (int, error) {
		return 11, err
	}

	compare(t, Ok(21).Match(onValue1, onNone1), Ok(42))
	compare(t, Ok(21).Match(onValue2, onNone1), Err[int](err))
	compare(t, Err[int](err).Match(onValue1, onNone1), Ok(11))
	compare(t, Err[int](err).Match(onValue1, onNone2), Err[int](err))
}
