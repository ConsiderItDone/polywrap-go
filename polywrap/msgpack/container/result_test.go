package container

import (
	"errors"
	"testing"
)

var (
	someResInt = 42
	errSomeRes = errors.New("some test err")
)

func compareResults[T bool | int | string](t *testing.T, actual Result[T], expected Result[T]) {
	if actual.IsErr() != expected.IsErr() {
		t.Errorf("Bad isNone, got: %v, want: %v", actual, expected)
	}
	if actual.IsOk() != expected.IsOk() {
		t.Errorf("Bad isSome, got: %v, want: %v", actual, expected)
	}
	if actual.IsOk() && expected.IsOk() && actual.Unwrap() != expected.Unwrap() {
		t.Errorf("Bad value, got: %v, want: %v", actual.Unwrap(), expected.Unwrap())
	}
	if actual.IsErr() && expected.IsErr() && actual.Error() != expected.Error() {
		t.Errorf("Bad value, got: %v, want: %v", actual.Unwrap(), expected.Unwrap())
	}
}

func TestResultOk(t *testing.T) {
	compareResults(t, Ok(42), Result[int]{value: &someResInt, err: nil})
}

func TestResultErr(t *testing.T) {
	compareResults(t, Err[int](errSomeRes), Result[int]{value: nil, err: errSomeRes})
}

func TestResultAsResult(t *testing.T) {
	compareResults(t, AsResult(someResInt, nil), Ok(someResInt))
	compareResults(t, AsResult(someResInt, errSomeRes), Err[int](errSomeRes))
}

func TestResultMatch(t *testing.T) {
	onValue1 := ResultResolver[int](func(i int) (int, error) {
		return i * 2, nil
	})
	onValue2 := ResultResolver[int](func(i int) (int, error) {
		return 0, errSomeRes
	})
	onNone1 := ResultRejecter[int](func(e error) (int, error) {
		return someResInt, nil
	})
	onNone2 := ResultRejecter[int](func(e error) (int, error) {
		return someResInt, errSomeRes
	})

	compareResults(t, Ok(someResInt/2).Match(onValue1, onNone1), Ok(someResInt))
	compareResults(t, Ok(someResInt/2).Match(onValue2, onNone1), Err[int](errSomeRes))
	compareResults(t, Err[int](errSomeRes).Match(onValue1, onNone1), Ok(42))
	compareResults(t, Err[int](errSomeRes).Match(onValue1, onNone2), Err[int](errSomeRes))
}
