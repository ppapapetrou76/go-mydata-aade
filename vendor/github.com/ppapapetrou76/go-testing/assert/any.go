package assert

import (
	"reflect"
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/values"
)

// AssertableAny is the assertable structure for interface{} values.
type AssertableAny struct {
	t      *testing.T
	actual values.AnyValue
}

// That returns an AssertableAny structure initialized with the test reference and the actual value to assert.
func That(t *testing.T, actual interface{}) AssertableAny {
	t.Helper()
	return AssertableAny{
		t:      t,
		actual: values.NewAnyValue(actual),
	}
}

// IsEqualTo asserts if the expected interface is equal to the assertable value
// It errors the tests if the compared values (actual VS expected) are not equal.
func (a AssertableAny) IsEqualTo(expected interface{}) AssertableAny {
	a.t.Helper()
	if !a.actual.IsEqualTo(expected) {
		a.t.Error(shouldBeEqual(a.actual, expected))
	}
	return a
}

// IsNotEqualTo asserts if the expected interface is not qual to the assertable value
// It errors the tests if the compared values (actual VS expected) are equal.
func (a AssertableAny) IsNotEqualTo(expected interface{}) AssertableAny {
	a.t.Helper()
	if a.actual.IsEqualTo(expected) {
		a.t.Error(shouldNotBeEqual(a.actual, expected))
	}
	return a
}

// IsNil asserts if the expected value is nil.
func (a AssertableAny) IsNil() AssertableAny {
	a.t.Helper()
	if !a.actual.IsNil() {
		a.t.Error(shouldBeNil(a.actual))
	}
	return a
}

// IsNotNil asserts if the expected value is not nil.
func (a AssertableAny) IsNotNil() AssertableAny {
	a.t.Helper()
	if !a.actual.IsNotNil() {
		a.t.Error(shouldNotBeNil(a.actual))
	}
	return a
}

// IsTrue asserts if the expected value is true.
func (a AssertableAny) IsTrue() AssertableAny {
	a.t.Helper()
	a.IsEqualTo(true)
	return a
}

// IsFalse asserts if the expected value is false.
func (a AssertableAny) IsFalse() AssertableAny {
	a.t.Helper()
	a.IsEqualTo(false)
	return a
}

// HasTypeOf asserts if the expected value has the type of a given value.
func (a AssertableAny) HasTypeOf(t reflect.Type) AssertableAny {
	a.t.Helper()
	if !a.actual.HasTypeOf(t) {
		a.t.Error(shouldHaveType(a.actual, t))
	}
	return a
}
