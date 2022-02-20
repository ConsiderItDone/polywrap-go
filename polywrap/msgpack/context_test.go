package msgpack

import "testing"

func TestContextPush(t *testing.T) {
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
}
