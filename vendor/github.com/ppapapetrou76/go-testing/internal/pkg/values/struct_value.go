package values

import "reflect"

// StructValue is a struct that holds a struct value.
type StructValue struct {
	value          interface{}
	ExcludedFields []string
}

// IsEqualTo returns true if the value is equal to the expected value, else false.
func (s StructValue) IsEqualTo(expected interface{}) bool {
	if s.value == nil && expected == nil {
		return true
	}
	actualElement := reflect.ValueOf(s.value)
	expectedElement := reflect.ValueOf(expected)

	if actualElement.Kind() == reflect.Ptr && expectedElement.Kind() == reflect.Ptr {
		actualElement = actualElement.Elem()
		expectedElement = expectedElement.Elem()
	}
	if !(actualElement.Kind() == reflect.Struct && expectedElement.Kind() == reflect.Struct) {
		return false
	}
	if reflect.TypeOf(s.value) != reflect.TypeOf(expected) {
		return false
	}
	for i := 0; i < actualElement.NumField(); i++ {
		actualValue := actualElement.Field(i)
		expectedValue := expectedElement.Field(i)

		if sliceContains(s.ExcludedFields, actualElement.Type().Field(i).Name) {
			continue
		}
		if actualValue.CanInterface() {
			if !reflect.DeepEqual(actualValue.Interface(), expectedValue.Interface()) {
				return false
			}
		} else {
			if !areEqualValues(actualValue, expectedValue) {
				return false
			}
		}
	}
	return true
}

// Value returns the actual value of the structure.
func (s StructValue) Value() interface{} {
	return s.value
}

// NewStructValue creates and returns a StructValue struct initialed with the given value.
func NewStructValue(value interface{}) StructValue {
	return StructValue{value: value}
}

func sliceContains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}
