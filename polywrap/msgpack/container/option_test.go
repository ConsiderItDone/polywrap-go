package container

import (
	"reflect"
	"testing"
)

func compare(t *testing.T, actual any, expected any) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Bad value, got: %v, want: %v", actual, expected)
	}
}

func TestOptionSome(t *testing.T) {
	compare(t, Some(42), Option[int]{isValue: true, value: 42})
}

func TestOptionNone(t *testing.T) {
	compare(t, None[int](), Option[int]{isValue: false})
}

func TestOptionIsSome(t *testing.T) {
	compare(t, Some(42).IsSome(), true)
	compare(t, None[int]().IsSome(), false)
}

func TestOptionIsNone(t *testing.T) {
	compare(t, Some(42).IsNone(), false)
	compare(t, None[int]().IsNone(), true)
}

func TestOptionGet(t *testing.T) {
	v1, ok1 := Some(42).Get()
	v2, ok2 := None[int]().Get()
	compare(t, v1, 42)
	compare(t, ok1, true)
	compare(t, v2, 0)
	compare(t, ok2, false)
}

func TestOptionMustGet(t *testing.T) {
	compare(t, Some(42).MustGet(), 42)

	defer func() {
		compare(t, recover(), ErrEmptyOption)
	}()
	compare(t, None[int]().MustGet(), 0)
}

func TestOptionOrElse(t *testing.T) {
	compare(t, Some(42).OrElse(21), 42)
	compare(t, None[int]().OrElse(21), 21)
}

func TestOptionMatch(t *testing.T) {
	onValue1 := func(i int) (int, bool) {
		return i * 2, true
	}
	onValue2 := func(i int) (int, bool) {
		return i * 2, false
	}
	onNone1 := func() (int, bool) {
		return 11, true
	}
	onNone2 := func() (int, bool) {
		return 11, false
	}

	compare(t, Some(21).Match(onValue1, onNone1), Some(42))
	compare(t, Some(21).Match(onValue2, onNone2), None[int]())
	compare(t, None[int]().Match(onValue1, onNone1), Some(11))
	compare(t, None[int]().Match(onValue1, onNone2), None[int]())
}
