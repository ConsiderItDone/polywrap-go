package json

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	var data = "{\"language\": \"Golang\", \"runtime\": \"Polywrap\"}"
	expected := map[string]interface{}{
		"language": "Golang",
		"runtime":  "Polywrap",
	}
	actual, err := Decode(data)
	if err != nil {
		t.Errorf("Decode error: %v", err)
	}

	if !reflect.DeepEqual(*actual, expected) {
		t.Errorf("Decode is incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestEncode(t *testing.T) {
	data := map[string]interface{}{
		"language": "Golang",
		"runtime":  "Polywrap",
	}
	expected := "{\"language\":\"Golang\",\"runtime\":\"Polywrap\"}"
	actual, err := Encode(&data)
	if err != nil {
		t.Errorf("Encode error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Decode is incorrect, got: %v, want: %v.", actual, expected)
	}
}
