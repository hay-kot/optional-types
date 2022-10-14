package tests

import (
	"encoding/json"
	"testing"

	"github.com/hay-kot/optional-types"
)

var err error
var body = []byte(`{"Name":"Bob","Age":1}`)

type Struct struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Age         int    `json:"Age"`
}

type OptionalStruct struct {
	Name        optional.Optional[string] `json:"Name"`
	Description optional.Optional[string] `json:"Description"`
	Age         optional.Optional[int]    `json:"Age"`
}

type PointerStruct struct {
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Age         *int    `json:"Age"`
}

type SetterFunc = func(*Struct)

func SetValues(setters ...SetterFunc) {
	var s Struct
	for _, setter := range setters {
		if setter != nil {
			setter(&s)
		}
	}
}

func BenchmarkStructMarshal(b *testing.B) {
	var v Struct
	setters := make([]SetterFunc, 3)

	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(body, &v)

		if v.Name != "" {
			setters[0] = func(s *Struct) {
				s.Name = v.Name
			}
		}

		if v.Description != "" {
			setters[1] = func(s *Struct) {
				s.Description = v.Description
			}
		}

		if v.Age != 0 {
			setters[2] = func(s *Struct) {
				s.Age = v.Age
			}
		}

		SetValues(setters...)
	}
}

func BenchmarkMapMarshal(b *testing.B) {
	var s map[string]interface{}
	setters := make([]SetterFunc, 3)

	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(body, &s)

		if v, ok := s["Name"]; ok {
			setters[0] = func(s *Struct) {
				s.Name = v.(string)
			}
		}

		if v, ok := s["Description"]; ok {
			setters[1] = func(s *Struct) {
				s.Description = v.(string)
			}
		}

		if v, ok := s["Age"]; ok {
			setters[2] = func(s *Struct) {
				s.Age = int(v.(float64))
			}
		}

		SetValues(setters...)
	}
}

func BenchmarkOptionalMarshal(b *testing.B) {
	var o OptionalStruct
	setters := make([]SetterFunc, 3)

	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(body, &o)

		if o.Name.IsPresent() {
			setters[0] = func(s *Struct) {
				s.Name = o.Name.Unwrap()
			}
		}

		if o.Description.IsPresent() {
			setters[1] = func(s *Struct) {
				s.Description = o.Description.Unwrap()
			}
		}

		if o.Age.IsPresent() {
			setters[2] = func(s *Struct) {
				s.Age = o.Age.Unwrap()
			}
		}

		SetValues(setters...)
	}
}

func BenchmarkPointerMarshal(b *testing.B) {
	var p PointerStruct
	setters := make([]SetterFunc, 3)

	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(body, &p)

		if p.Name != nil {
			setters[0] = func(s *Struct) {
				s.Name = *p.Name
			}
		}

		if p.Description != nil {
			setters[1] = func(s *Struct) {
				s.Description = *p.Description
			}
		}

		if p.Age != nil {
			setters[2] = func(s *Struct) {
				s.Age = *p.Age
			}
		}

	}
}
