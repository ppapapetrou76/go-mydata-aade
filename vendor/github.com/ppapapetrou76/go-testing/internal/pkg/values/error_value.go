package values

import "fmt"

// ErrorValue is a struct that holds an error value.
type ErrorValue struct {
	value error
}

func (v ErrorValue) Error() error {
	return v.value
}

// Value returns the error value as an interface object.
func (v ErrorValue) Value() interface{} {
	return v.value
}

// NewErrorValue creates and returns a ErrorValue struct initialed with the given value.
func NewErrorValue(value interface{}) ErrorValue {
	switch cValue := value.(type) {
	case nil:
		return ErrorValue{}
	case error:
		return ErrorValue{value: cValue}
	default:
		panic(fmt.Sprintf("expected error value type but got %T type", cValue))
	}
}
