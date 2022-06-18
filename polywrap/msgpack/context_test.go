package msgpack

import "testing"

func TestContextPushAndPop(t *testing.T) {
	c := NewContext("some description")
	if !c.IsEmpty() {
		t.Errorf("Context is not empty")
	}

	if c.Length() != 0 {
		t.Errorf("Length is incorrect, got: %d, want: %d.", c.Length(), 0)
	}

	c.Push("property", "string", "")
	c.Push("property", "i32", "")
	c.Push("property", "bool", "")

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

func TestPrints(t *testing.T) {
	c := NewContext("Deserializing MyObject")
	c.Push("propertyOne", "unknown", "searching for property type")

	actual := c.toString()
	expected := "Context: Deserializing MyObject\n  at propertyOne: unknown >> searching for property type"

	if actual != expected {
		t.Errorf("toString() is incorrect: \ngot \n%s \nwant \n%s", actual, expected)
	}

	actual = c.printWithContext("Invalid length")
	expected = "Invalid length\n  Context: Deserializing MyObject\n    at propertyOne: unknown >> searching for property type"
	if actual != expected {
		t.Errorf("printWithContext() is incorrect: \ngot \n%s \nwant \n%s", actual, expected)
	}

	c.Push("propertyOne", "i32", "type found, reading property")
	actual = c.printWithContext("Invalid length")
	expected = "Invalid length\n  Context: Deserializing MyObject\n    at propertyOne: i32 >> type found, reading property\n      at propertyOne: unknown >> searching for property type"
	if actual != expected {
		t.Errorf("printWithContext() is incorrect: \ngot \n%s \nwant \n%s", actual, expected)
	}
}
