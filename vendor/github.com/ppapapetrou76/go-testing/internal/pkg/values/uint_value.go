package values

import (
	"fmt"
)

// UIntValue is a struct that holds a uint value.
type UIntValue struct {
	value uint
}

// IsEqualTo returns true if the value is equal to the expected value, else false.
func (i UIntValue) IsEqualTo(expected interface{}) bool {
	return i.equals(NewUIntValue(expected))
}

// IsGreaterThan returns true if the value is greater than the expected value, else false.
func (i UIntValue) IsGreaterThan(expected interface{}) bool {
	return i.greaterThan(NewUIntValue(expected))
}

// IsGreaterOrEqualTo returns true if the value is greater than or equal to the expected value, else false.
func (i UIntValue) IsGreaterOrEqualTo(expected interface{}) bool {
	return i.greaterOrEqual(NewUIntValue(expected))
}

// IsLessThan returns true if the value is less than the expected value, else false.
func (i UIntValue) IsLessThan(expected interface{}) bool {
	return !i.IsGreaterOrEqualTo(expected)
}

// IsLessOrEqualTo returns true if the value is less than or equal to the expected value, else false.
func (i UIntValue) IsLessOrEqualTo(expected interface{}) bool {
	return !i.IsGreaterThan(expected)
}

// Value returns the actual value of the structure.
func (i UIntValue) Value() interface{} {
	return i.value
}

func (i UIntValue) greaterThan(expected UIntValue) bool {
	return i.value > expected.value
}

func (i UIntValue) greaterOrEqual(expected UIntValue) bool {
	return i.value >= expected.value
}

func (i UIntValue) equals(expected UIntValue) bool {
	return i.value == expected.value
}

// NewUIntValue creates and returns a UIntValue struct initialed with the given value.
func NewUIntValue(value interface{}) UIntValue {
	switch v := value.(type) {
	case uint:
		return UIntValue{value: value.(uint)}
	case uint8:
		return UIntValue{value: uint(value.(uint8))}
	case uint16:
		return UIntValue{value: uint(value.(uint16))}
	case uint32:
		return UIntValue{value: uint(value.(uint32))}
	case uint64:
		return UIntValue{value: uint(value.(uint64))}
	default:
		panic(fmt.Sprintf("expected uint value type but got %T type", v))
	}
}
