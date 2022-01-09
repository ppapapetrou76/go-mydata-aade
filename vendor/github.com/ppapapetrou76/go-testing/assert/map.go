package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/values"
	"github.com/ppapapetrou76/go-testing/types"
)

// AssertableMap is the structure to assert maps.
type AssertableMap struct {
	t      *testing.T
	actual types.Map
}

// ThatMap returns a proper assertable structure based on the map key type.
func ThatMap(t *testing.T, actual interface{}) AssertableMap {
	t.Helper()
	return AssertableMap{
		t:      t,
		actual: values.NewKeyStringMap(actual),
	}
}

// IsEqualTo asserts if the expected map is equal to the assertable map value
// It errors the tests if the compared values (actual VS expected) are not equal.
func (a AssertableMap) IsEqualTo(expected interface{}) AssertableMap {
	a.t.Helper()
	if !values.IsMap(a.actual.Value()) {
		a.t.Error(shouldBeMap(a.actual))
		return a
	}

	if !a.actual.IsEqualTo(expected) {
		a.t.Error(shouldBeEqual(a.actual, expected))
	}
	return a
}

// IsNotEqualTo asserts if the expected map is not equal to the assertable map value
// It errors the tests if the compared values (actual VS expected) are equal.
func (a AssertableMap) IsNotEqualTo(expected interface{}) AssertableMap {
	a.t.Helper()
	if !values.IsMap(a.actual.Value()) {
		a.t.Error(shouldBeMap(a.actual))
		return a
	}
	if a.actual.IsEqualTo(expected) {
		a.t.Error(shouldNotBeEqual(a.actual, expected))
	}
	return a
}

// HasSize asserts if the assertable string map has the expected length size
// It errors the test if it doesn't have the expected size.
func (a AssertableMap) HasSize(size int) AssertableMap {
	a.t.Helper()
	if !values.IsMap(a.actual.Value()) {
		a.t.Error(shouldBeMap(a.actual))
		return a
	}
	if !a.actual.HasSize(size) {
		a.t.Error(shouldHaveSize(a.actual, size))
	}
	return a
}

// IsEmpty asserts if the assertable string map is empty or not.
func (a AssertableMap) IsEmpty() AssertableMap {
	a.t.Helper()
	if !values.IsMap(a.actual.Value()) {
		a.t.Error(shouldBeMap(a.actual))
		return a
	}
	if a.actual.IsNotEmpty() {
		a.t.Error(shouldBeEmpty(a.actual))
	}
	return a
}

// IsNotEmpty asserts if the assertable string map is not empty.
func (a AssertableMap) IsNotEmpty() AssertableMap {
	a.t.Helper()
	if !values.IsMap(a.actual.Value()) {
		a.t.Error(shouldBeMap(a.actual))
		return a
	}
	if a.actual.IsEmpty() {
		a.t.Error(shouldNotBeEmpty(a.actual))
	}
	return a
}

// HasKey asserts if the assertable map has the given key
// It errors the test if
// * they key can't be found
// * the key is not comparable
// * the asserted type is not a map.
func (a AssertableMap) HasKey(elements interface{}) AssertableMap {
	a.t.Helper()
	if !values.IsMap(a.actual.Value()) {
		a.t.Error(shouldBeMap(a.actual))
		return a
	}
	if !a.actual.HasKey(elements) {
		a.t.Error(shouldHaveKey(a.actual, elements))
	}
	return a
}

// HasValue asserts if the assertable map has the given value
// It errors the test if
// * they key can't be found
// * the key is not comparable
// * the asserted type is not a map.
func (a AssertableMap) HasValue(elements interface{}) AssertableMap {
	a.t.Helper()
	if !values.IsMap(a.actual.Value()) {
		a.t.Error(shouldBeMap(a.actual))
		return a
	}
	if !a.actual.HasValue(elements) {
		a.t.Error(shouldHaveValue(a.actual, elements))
	}
	return a
}

// HasEntry asserts if the assertable map has the given entry
// It errors the test if
// * the entry can't be found
// * the key is not comparable
// * the asserted type is not a map.
func (a AssertableMap) HasEntry(value types.MapEntry) AssertableMap {
	a.t.Helper()
	if !values.IsMap(a.actual.Value()) {
		a.t.Error(shouldBeMap(a.actual))
		return a
	}
	if !a.actual.HasEntry(value) {
		a.t.Error(shouldHaveEntry(a.actual, value))
	}
	return a
}

// HasNotKey asserts if the assertable map has not the given key
// It errors the test if
// * they key is found
// * the key is not comparable
// * the asserted type is not a map.
func (a AssertableMap) HasNotKey(elements interface{}) AssertableMap {
	a.t.Helper()
	if !values.IsMap(a.actual.Value()) {
		a.t.Error(shouldBeMap(a.actual))
		return a
	}
	if a.actual.HasKey(elements) {
		a.t.Error(shouldNotHaveKey(a.actual, elements))
	}
	return a
}

// HasNotValue asserts if the assertable map has not the given value
// It errors the test if
// * they key can be found
// * the key is not comparable
// * the asserted type is not a map.
func (a AssertableMap) HasNotValue(elements interface{}) AssertableMap {
	a.t.Helper()
	if !values.IsMap(a.actual.Value()) {
		a.t.Error(shouldBeMap(a.actual))
		return a
	}
	if a.actual.HasValue(elements) {
		a.t.Error(shouldNotHaveValue(a.actual, elements))
	}
	return a
}

// HasNotEntry asserts if the assertable map has not the given entry
// It errors the test if
// * the entry can be found
// * the key is not comparable
// * the asserted type is not a map.
func (a AssertableMap) HasNotEntry(value types.MapEntry) AssertableMap {
	a.t.Helper()
	if !values.IsMap(a.actual.Value()) {
		a.t.Error(shouldBeMap(a.actual))
		return a
	}
	if a.actual.HasEntry(value) {
		a.t.Error(shouldNotHaveEntry(a.actual, value))
	}
	return a
}
