package values

import (
	"fmt"
)

// IntValue is a struct that holds an int value.
type IntValue struct {
	value int
}

// IsEqualTo returns true if the value is equal to the expected value, else false.
func (i IntValue) IsEqualTo(expected interface{}) bool {
	return i.equals(NewIntValue(expected))
}

// IsGreaterThan returns true if the value is greater than the expected value, else false.
func (i IntValue) IsGreaterThan(expected interface{}) bool {
	return i.greaterThan(NewIntValue(expected))
}

// IsGreaterOrEqualTo returns true if the value is greater than or equal to the expected value, else false.
func (i IntValue) IsGreaterOrEqualTo(expected interface{}) bool {
	return i.greaterOrEqual(NewIntValue(expected))
}

// IsLessThan returns true if the value is less than the expected value, else false.
func (i IntValue) IsLessThan(expected interface{}) bool {
	return !i.IsGreaterOrEqualTo(expected)
}

// IsLessOrEqualTo returns true if the value is less than or equal to the expected value, else false.
func (i IntValue) IsLessOrEqualTo(expected interface{}) bool {
	return !i.IsGreaterThan(expected)
}

// Value returns the actual value of the structure.
func (i IntValue) Value() interface{} {
	return i.value
}

func (i IntValue) greaterThan(expected IntValue) bool {
	return i.value > expected.value
}

func (i IntValue) greaterOrEqual(expected IntValue) bool {
	return i.value >= expected.value
}

func (i IntValue) equals(expected IntValue) bool {
	return i.value == expected.value
}

// NewIntValue creates and returns an IntValue struct initialed with the given value.
func NewIntValue(value interface{}) IntValue {
	switch v := value.(type) {
	case int:
		return IntValue{value: value.(int)}
	case int8:
		return IntValue{value: int(value.(int8))}
	case int16:
		return IntValue{value: int(value.(int16))}
	case int32:
		return IntValue{value: int(value.(int32))}
	case int64:
		return IntValue{value: int(value.(int64))}
	default:
		panic(fmt.Sprintf("expected int value type but got %T type", v))
	}
}
