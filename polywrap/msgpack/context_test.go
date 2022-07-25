package msgpack

import (
	"runtime"
	"testing"
)

func TestContextPushAndPop(t *testing.T) {
	c := NewContext("some description")
	if !c.IsEmpty() {
		t.Errorf("Context is not empty")
	}

	if c.Length() != 0 {
		t.Errorf("Length is incorrect, got: %d, want: %d.", c.Length(), 0)
	}

	c.Push("property", "string", "some string")
	c.Push("property", "i32", "100500")
	c.Push("property", "bool", "true")

	if c.IsEmpty() {
		t.Errorf("Context is empty")
	}

	if c.Length() != 3 {
		t.Errorf("Length is incorrect, got: %d, want: %d.", c.Length(), 3)
	}

	c.Pop()
	c.Pop()
	c.Pop()

	if c.Length() != 0 {
		t.Errorf("Length is incorrect, got: %d, want: %d.", c.Length(), 0)
	}
}

func TestEmptyPop(t *testing.T) {
	if runtime.Compiler == "tinygo" {
		t.Log("Skipping due tinygo limitations")
		return
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	c := NewContext("some description")

	c.Pop()
}

func TestPrints(t *testing.T) {
	c := NewContext("Deserializing MyObject")
	// check empty context printing
	actual := c.toString()
	expected := "Context: Deserializing MyObject\n  context stack is empty"
	if actual != expected {
		t.Errorf("toString() is incorrect: \ngot \n%s \nwant \n%s", actual, expected)
	}

	c.Push("propertyOne", "unknown", "searching for property type")

	actual = c.toString()
	expected = "Context: Deserializing MyObject\n  at propertyOne: unknown >> searching for property type"

	if actual != expected {
		t.Errorf("toString() is incorrect: \ngot \n%s \nwant \n%s", actual, expected)
	}

	actual = c.PrintWithContext("Invalid length")
	expected = "Invalid length\n  Context: Deserializing MyObject\n    at propertyOne: unknown >> searching for property type"
	if actual != expected {
		t.Errorf("PrintWithContext() is incorrect: \ngot \n%s \nwant \n%s", actual, expected)
	}

	c.Push("propertyOne", "i32", "type found, reading property")
	actual = c.PrintWithContext("Invalid length")
	expected = "Invalid length\n  Context: Deserializing MyObject\n    at propertyOne: i32 >> type found, reading property\n      at propertyOne: unknown >> searching for property type"
	if actual != expected {
		t.Errorf("PrintWithContext() is incorrect: \ngot \n%s \nwant \n%s", actual, expected)
	}
}
