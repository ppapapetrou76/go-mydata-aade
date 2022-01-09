package types

import (
	"fmt"
	"time"
)

// TimeValue is a struct that holds a time value.
type TimeValue struct {
	value time.Time
}

// IsSameAs returns true if the value is the same as the expected value, else false.
func (t TimeValue) IsSameAs(expected interface{}) bool {
	return t.value == expected
}

// IsAlmostSameAs returns true if the value is the almost the same as the expected value, else false.
func (t TimeValue) IsAlmostSameAs(expected interface{}) bool {
	return NewTimeValue(t.value.Add(time.Millisecond*100)).IsAfter(expected) && NewTimeValue(t.value.Add(-time.Millisecond*100)).IsBefore(expected)
}

// IsNotDefined returns true if the time value is not defined (has no value) else false.
func (t TimeValue) IsNotDefined() bool {
	return t.value.Nanosecond() == 0
}

// IsDefined returns true if the time value is defined ( has some value ) else false.
func (t TimeValue) IsDefined() bool {
	return !t.IsNotDefined()
}

// IsNotSameAs returns true if the value is not the same as the expected value, else false.
func (t TimeValue) IsNotSameAs(expected interface{}) bool {
	return t.value != expected
}

// IsAfter returns true if the value is after than the expected value, else false.
func (t TimeValue) IsAfter(expected interface{}) bool {
	return t.isAfter(NewTimeValue(expected))
}

// IsBefore returns true if the value is before the expected value, else false.
func (t TimeValue) IsBefore(expected interface{}) bool {
	return t.isBefore(NewTimeValue(expected))
}

// Value returns the actual value of the structure.
func (t TimeValue) Value() interface{} {
	return t.value
}

func (t TimeValue) isAfter(expected TimeValue) bool {
	return t.value.After(expected.value)
}

func (t TimeValue) isBefore(expected TimeValue) bool {
	return t.value.Before(expected.value)
}

// NewTimeValue creates and returns an TimeValue struct initialed with the given value.
func NewTimeValue(value interface{}) TimeValue {
	switch v := value.(type) {
	case time.Time:
		return TimeValue{value: value.(time.Time)}
	default:
		panic(fmt.Sprintf("expected time.Time value type but got %T type", v))
	}
}
