package optional

import (
	"encoding/json"
)

type Optional[T any] struct {
	value *T
}

func (o Optional[T]) IsPresent() bool {
	return o.value != nil
}

func (o Optional[T]) Get() (v T, ok bool) {
	if !o.IsPresent() {
		var t T
		return t, false
	}

	return *o.value, true
}

func (o Optional[T]) Unwrap() T {
	if o.value == nil {
		panic("unwrap on nil value")
	}

	return *o.value
}

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.value)
}

func (o Optional[T]) UnwrapOr(defaultValue T) T {
	if o.value == nil {
		return defaultValue
	}

	return *o.value
}
