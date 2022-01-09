package values

import (
	"fmt"
	"time"
)

// DurationValue is a struct that holds a duration value.
type DurationValue struct {
	value time.Duration
}

// IsEqualTo returns true if the value is the same as the expected value, else false.
func (d DurationValue) IsEqualTo(expected interface{}) bool {
	return d.value == expected
}

// IsNotEqualTo returns true if the value is not the same as the expected value, else false.
func (d DurationValue) IsNotEqualTo(expected interface{}) bool {
	return d.value != expected
}

// IsLongerThan returns true if the value is longer than the expected value, else false.
func (d DurationValue) IsLongerThan(expected interface{}) bool {
	return d.isLonger(NewDurationValue(expected))
}

// IsShorterThan returns true if the value is shorter than the expected value, else false.
func (d DurationValue) IsShorterThan(expected interface{}) bool {
	return d.isShorter(NewDurationValue(expected))
}

// Value returns the actual value of the structure.
func (d DurationValue) Value() interface{} {
	return d.value
}

func (d DurationValue) isLonger(expected DurationValue) bool {
	return d.value > expected.value
}

func (d DurationValue) isShorter(expected DurationValue) bool {
	return d.value < expected.value
}

// NewDurationValue creates and returns an DurationValue struct initialed with the given value.
func NewDurationValue(value interface{}) DurationValue {
	switch v := value.(type) {
	case time.Duration:
		return DurationValue{value: value.(time.Duration)}
	default:
		panic(fmt.Sprintf("expected time.Duration value type but got %T type", v))
	}
}
