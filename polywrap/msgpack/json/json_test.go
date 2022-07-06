package json

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	var data = `{"language": "Golang", "version": 100500}`
	v, err := Decode(data)
	if err != nil {
		t.Errorf("Decode error: %v", err)
	}

	expected := "Golang"
	actual := string(v.GetStringBytes("language"))

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Decode is incorrect, got: %s, want: %s.", actual, expected)
	}

	expectedInt := 100500
	actualInt := v.GetInt("version")

	if !reflect.DeepEqual(actualInt, expectedInt) {
		t.Errorf("Decode is incorrect, got: %d, want: %d.", actualInt, expectedInt)
	}
}

func TestEncode(t *testing.T) {
	json := NewJSON()
	obj := json.NewObject()
	obj.Set("language", json.NewString("Golang"))
	obj.Set("version", json.NewNumberInt(100500))

	expected := `{"language":"Golang","version":100500}`
	actual := obj.String()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Decode is incorrect, got: %v, want: %v.", actual, expected)
	}
}
