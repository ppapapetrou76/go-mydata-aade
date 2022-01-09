package values

import (
	"reflect"

	"github.com/ppapapetrou76/go-testing/types"
)

// MapValue is a struct that holds a string map value.
type MapValue struct {
	value interface{}
}

// Value returns the actual value of the map.
func (s MapValue) Value() interface{} {
	return s.value
}

// IsEqualTo returns true if the value is equal to the expected value, else false.
func (s MapValue) IsEqualTo(expected interface{}) bool {
	if !IsMap(expected) {
		return false
	}

	if reflect.TypeOf(s.Value()).Elem() != reflect.TypeOf(expected).Elem() {
		return false
	}

	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(s.Value())
	return areMapsEqual(actualValue, expectedValue)
}

// IsEmpty returns true if the map is empty else false.
func (s MapValue) IsEmpty() bool {
	return s.HasSize(0)
}

// IsNotEmpty returns true if the map is not empty else false.
func (s MapValue) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// HasSize returns true if the map has the expected size else false.
func (s MapValue) HasSize(length int) bool {
	return s.Size() == length
}

// Size returns the map size.
func (s MapValue) Size() int {
	if !IsMap(s.Value()) {
		return 0
	}
	return reflect.ValueOf(s.Value()).Len()
}

func (s MapValue) hasKeyValue(key, value interface{}) bool {
	actualValue := reflect.ValueOf(s.Value())
	keyValue := actualValue.MapIndex(reflect.ValueOf(key))

	return areEqualValues(keyValue, reflect.ValueOf(value))
}

// HasEntry returns true if the map has the given key,value pair (map entry).
func (s MapValue) HasEntry(entry types.MapEntry) bool {
	return s.HasKey(entry.Key()) && s.hasKeyValue(entry.Key(), entry.Value())
}

// HasKey returns true if the map has the given key.
func (s MapValue) HasKey(key interface{}) bool {
	if !reflect.TypeOf(key).Comparable() {
		return false
	}

	actualValue := reflect.ValueOf(s.Value())
	return actualValue.MapIndex(reflect.ValueOf(key)).IsValid()
}

// HasValue returns true if the map has the given value.
func (s MapValue) HasValue(value interface{}) bool {
	actualValue := reflect.ValueOf(s.Value())
	keys := reflect.ValueOf(s.Value()).MapKeys()
	for _, k := range keys {
		if areEqualValues(actualValue.MapIndex(k), reflect.ValueOf(value)) {
			return true
		}
	}
	return false
}

// NewKeyStringMap creates and returns a MapValue struct initialed with the given value.
func NewKeyStringMap(value interface{}) MapValue {
	return MapValue{value: value}
}

// IsMap returns true if the given value is a map, else false.
func IsMap(value interface{}) bool {
	return reflect.ValueOf(value).Kind() == reflect.Map
}
