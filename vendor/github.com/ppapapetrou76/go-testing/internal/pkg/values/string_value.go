package values

import (
	"fmt"
	"strings"
	"unicode"
)

// StringValue value represents a string value.
type StringValue struct {
	value      string
	decorators []StringDecorator
}

// StringDecorator is a function type to decorate a string.
type StringDecorator func(value string) string

// RemoveSpaces removes all spaces from the given string.
func RemoveSpaces(value string) string {
	return strings.ReplaceAll(value, " ", "")
}

// IsEqualTo returns true if the value is equal to the expected value, else false.
func (s StringValue) IsEqualTo(expected interface{}) bool {
	return s.DecoratedValue() == s.decoratedValue(expected)
}

// Value returns the actual value of the structure.
func (s StringValue) Value() interface{} {
	return s.value
}

// IsGreaterThan returns true if the value is greater than the expected value, else false.
func (s StringValue) IsGreaterThan(expected interface{}) bool {
	return s.greaterThan(NewStringValue(s.decoratedValue(expected)))
}

// IsGreaterOrEqualTo returns true if the value is greater than or equal to the expected value, else false.
func (s StringValue) IsGreaterOrEqualTo(expected interface{}) bool {
	return s.greaterOrEqual(NewStringValue(s.decoratedValue(expected)))
}

// IsLessThan returns true if the value is less than the expected value, else false.
func (s StringValue) IsLessThan(expected interface{}) bool {
	return !s.IsGreaterOrEqualTo(s.decoratedValue(expected))
}

// IsLessOrEqualTo returns true if the value is less than or equal to the expected value, else false.
func (s StringValue) IsLessOrEqualTo(expected interface{}) bool {
	return !s.IsGreaterThan(s.decoratedValue(expected))
}

// IsEmpty returns true if the string is empty else false.
func (s StringValue) IsEmpty() bool {
	return s.DecoratedValue() == ""
}

// IsNotEmpty returns true if the string is not empty else false.
func (s StringValue) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Contains returns true if the string contains the given sub-string.
func (s StringValue) Contains(expected interface{}) bool {
	return strings.Contains(s.DecoratedValue(), NewStringValue(s.decoratedValue(expected)).value)
}

// ContainsIgnoringCase returns true if the string contains the given sub-string case insensitively.
func (s StringValue) ContainsIgnoringCase(expected interface{}) bool {
	return strings.Contains(strings.ToLower(s.DecoratedValue()), NewStringValue(strings.ToLower(s.decoratedValue(expected))).value)
}

// DoesNotContain returns true if the string does not contain the given sub-string.
func (s StringValue) DoesNotContain(expected interface{}) bool {
	return !s.Contains(s.decoratedValue(expected))
}

// HasSize returns true if the string has the expected size else false.
func (s StringValue) HasSize(length int) bool {
	return s.Size() == length
}

// HasSizeLessThan returns true if the string has size less than the given value else false.
func (s StringValue) HasSizeLessThan(length int) bool {
	return s.Size() < length
}

// Size returns the string size.
func (s StringValue) Size() int {
	return len(s.value)
}

// StartsWith returns true if the asserted value starts with the given string, else false.
func (s StringValue) StartsWith(substr string) bool {
	return strings.HasPrefix(s.DecoratedValue(), s.decoratedValue(substr))
}

// DoesNotStartWith returns true if the asserted value does not start with the given string, else false.
func (s StringValue) DoesNotStartWith(substr string) bool {
	return !s.StartsWith(substr)
}

// EndsWith returns true if the asserted value ends with the given string, else false.
func (s StringValue) EndsWith(substr string) bool {
	return strings.HasSuffix(s.DecoratedValue(), s.decoratedValue(substr))
}

// DoesNotEndWith returns true if the asserted value does not end with the given string, else false.
func (s StringValue) DoesNotEndWith(substr string) bool {
	return !s.EndsWith(substr)
}

// ContainsOnly returns true if the string contains only the given sub-string
// In other words if performs an equal operation.
func (s StringValue) ContainsOnly(expected interface{}) bool {
	return s.IsEqualTo(s.decoratedValue(expected))
}

// ContainsOnlyOnce returns true if the string contains the given sub-string only once.
func (s StringValue) ContainsOnlyOnce(substr string) bool {
	return strings.Count(s.DecoratedValue(), s.decoratedValue(substr)) == 1
}

// ContainsWhitespaces returns true if the string contains at least one whitespace.
func (s StringValue) ContainsWhitespaces() bool {
	return strings.Count(s.DecoratedValue(), " ") > 0
}

// ContainsOnlyWhitespaces returns true if the string contains only whitespaces.
func (s StringValue) ContainsOnlyWhitespaces() bool {
	return strings.Count(s.DecoratedValue(), " ") == len(s.DecoratedValue())
}

func (s StringValue) greaterThan(expected StringValue) bool {
	return s.DecoratedValue() > expected.value
}

func (s StringValue) greaterOrEqual(expected StringValue) bool {
	return s.DecoratedValue() >= expected.value
}

// HasDigitsOnly returns true if the string has only digits else false.
func (s StringValue) HasDigitsOnly() bool {
	for _, c := range s.DecoratedValue() {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// IsLowerCase returns true if the string is in lower case.
func (s StringValue) IsLowerCase() bool {
	return s.IsEqualTo(strings.ToLower(s.value))
}

// IsUpperCase returns true if the string is in upper case.
func (s StringValue) IsUpperCase() bool {
	return s.IsEqualTo(strings.ToUpper(s.value))
}

// NewStringValue creates and returns a StringValue struct initialed with the given value.
func NewStringValue(value interface{}) StringValue {
	switch v := value.(type) {
	case string:
		return StringValue{value: v}
	default:
		panic(fmt.Sprintf("expected string value type but got %T type", v))
	}
}

// AddDecorator adds a new string decorator to the assertable string value.
func (s StringValue) AddDecorator(decorator StringDecorator) StringValue {
	s.decorators = append(s.decorators, decorator)
	return s
}

func (s StringValue) decoratedValue(value interface{}) string {
	decoratedValue, ok := value.(string)
	if !ok {
		return ""
	}
	for _, decorator := range s.decorators {
		decoratedValue = decorator(decoratedValue)
	}
	return decoratedValue
}

// DecoratedValue returns the asserted string value after applying all the defined decorators.
func (s StringValue) DecoratedValue() string {
	return s.decoratedValue(s.value)
}
