package container

import (
	"testing"
)

var (
	someOptInt = 42
	someOptStr = "text"
)

func compareOptions[T bool | int | string](t *testing.T, actual Option[T], expected Option[T]) {
	if actual.IsNone() != expected.IsNone() {
		t.Errorf("Bad isNone, got: %v, want: %v", actual, expected)
	}
	if actual.IsSome() != expected.IsSome() {
		t.Errorf("Bad isSome, got: %v, want: %v", actual, expected)
	}
	if actual.IsSome() && expected.IsSome() && actual.Unwrap() != expected.Unwrap() {
		t.Errorf("Bad value, got: %v, want: %v", actual.Unwrap(), expected.Unwrap())
	}
}

func TestOptionSome(t *testing.T) {
	compareOptions(t, Some(someOptInt), Option[int]{value: &someOptInt})
	compareOptions(t, Some(someOptStr), Option[string]{value: &someOptStr})
}

func TestOptionNone(t *testing.T) {
	compareOptions(t, None[int](), Option[int]{value: nil})
	compareOptions(t, None[string](), Option[string]{value: nil})
}

func TestAsOption(t *testing.T) {
	compareOptions(t, AsOption(42, true), Some(someOptInt))
	compareOptions(t, AsOption("text", true), Some(someOptStr))
	compareOptions(t, AsOption(someOptInt, false), None[int]())
	compareOptions(t, AsOption(someOptStr, false), None[string]())
}

func TestOptionMatch(t *testing.T) {
	onValue1 := OptionResolver[int](func(i int) (int, bool) {
		return i * 2, true
	})
	onValue2 := OptionResolver[int](func(i int) (int, bool) {
		return i * 2, false
	})
	onNone1 := OptionRejecter[int](func() (int, bool) {
		return someOptInt, true
	})
	onNone2 := OptionRejecter[int](func() (int, bool) {
		return someOptInt, false
	})

	compareOptions(t, Some(someOptInt/2).Match(onValue1, onNone1), Some(someOptInt))
	compareOptions(t, Some(someOptInt/2).Match(onValue2, onNone2), None[int]())
	compareOptions(t, None[int]().Match(onValue1, onNone1), Some(someOptInt))
	compareOptions(t, None[int]().Match(onValue1, onNone2), None[int]())
}
