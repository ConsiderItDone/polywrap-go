package msgpack

import (
	"testing"
)

func TestReadTrue(t *testing.T) {
	context := NewContext("")
	reader := NewReadDecoder(context, []byte{195})

	actual := reader.ReadBool()

	if !actual {
		t.Errorf("TestReadTrue is incorrect, got: %v, want: true.", actual)
	}
}

func TestReadFalse(t *testing.T) {
	context := NewContext("")
	reader := NewReadDecoder(context, []byte{194})

	actual := reader.ReadBool()

	if actual {
		t.Errorf("TestReadFalse is incorrect, got: %v, want: false.", actual)
	}
}
