package values

import (
	"fmt"
	"strconv"
)

// BoolValue is a struct that holds a bool value.
type BoolValue struct {
	value bool
}

// IsEqualTo returns true if the value is equal to the expected value, else false.
func (s BoolValue) IsEqualTo(expected interface{}) bool {
	return s.value == expected
}

// Value returns the actual value of the structure.
func (s BoolValue) Value() interface{} {
	return s.value
}

// NewBoolValue creates and returns an BoolValue struct initialed with the given value.
func NewBoolValue(value interface{}) BoolValue {
	switch v := value.(type) {
	case bool:
		return BoolValue{value: v}
	default:
		panic(fmt.Sprintf("expected bool value type but got %T type", v))
	}
}

func (s BoolValue) String() string {
	return strconv.FormatBool(s.value)
}
