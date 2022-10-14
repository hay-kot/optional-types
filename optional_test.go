// Path: optional_test.go
package optional

import (
	"encoding/json"
	"testing"
)

type TestStruct struct {
	String Optional[string] `json:"String"`
	Int    Optional[int]    `json:"Int"`
}

func TestOption_String_NotNull(t *testing.T) {
	// json
	byts := []byte(`{"String":"test","Int":1}`)

	// unmarshal
	var testStruct TestStruct
	err := json.Unmarshal(byts, &testStruct)

	// assert
	if err != nil {
		t.Error(err)
	}

	if !testStruct.String.IsPresent() {
		t.Error("String is not present")
	}

	if testStruct.String.Unwrap() != "test" {
		t.Error("String is not equal")
	}

	v, ok := testStruct.String.Get()
	if !ok {
		t.Error("String is not present")
	}

	if v != "test" {
		t.Error("String is not equal")
	}
}

func TestOption_String_Null(t *testing.T) {
	// json
	byts := []byte(`{"String":null,"Int":1}`)

	// unmarshal
	var testStruct TestStruct
	err := json.Unmarshal(byts, &testStruct)

	// assert
	if err != nil {
		t.Error(err)
	}

	if testStruct.String.IsPresent() {
		t.Error("String is present")
	}

	v, ok := testStruct.String.Get()
	if ok {
		t.Error("String is present")
	}

	if v != "" {
		t.Error("String is not equal")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("String is not nil")
		}
	}()

	testStruct.String.Unwrap()

}
