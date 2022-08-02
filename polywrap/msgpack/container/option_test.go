package container

import (
	"reflect"
	"testing"
)

func compare(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Bad value, got: %v, want: %v", actual, expected)
	}
}

func TestOptionSome(t *testing.T) {
	compare(t, Some(42), Option{isValue: true, value: 42})
}

func TestOptionNone(t *testing.T) {
	compare(t, None(), Option{isValue: false, value: nil})
}

func TestOptionIsSome(t *testing.T) {
	compare(t, Some(42).IsSome(), true)
	compare(t, None().IsSome(), false)
}

func TestOptionIsNone(t *testing.T) {
	compare(t, Some(42).IsNone(), false)
	compare(t, None().IsNone(), true)
}

func TestOptionGet(t *testing.T) {
	v1, ok1 := Some(42).Get()
	v2, ok2 := None().Get()
	compare(t, v1, 42)
	compare(t, ok1, true)
	compare(t, v2, nil)
	compare(t, ok2, false)
}

func TestOptionOrElse(t *testing.T) {
	compare(t, Some(42).OrElse(21), 42)
	compare(t, None().OrElse(21), 21)
}

func TestOptionMatch(t *testing.T) {
	onValue1 := func(i interface{}) (interface{}, bool) {
		return i.(int) * 2, true
	}
	onValue2 := func(i interface{}) (interface{}, bool) {
		return i.(int) * 2, false
	}
	onNone1 := func() (interface{}, bool) {
		return 11, true
	}
	onNone2 := func() (interface{}, bool) {
		return 11, false
	}

	compare(t, Some(21).Match(onValue1, onNone1), Some(42))
	compare(t, Some(21).Match(onValue2, onNone2), None())
	compare(t, None().Match(onValue1, onNone1), Some(11))
	compare(t, None().Match(onValue1, onNone2), None())
}
