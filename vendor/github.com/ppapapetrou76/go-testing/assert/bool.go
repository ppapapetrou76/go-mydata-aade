package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/values"
)

// AssertableBool is the assertable structure for bool values.
type AssertableBool struct {
	t      *testing.T
	actual values.BoolValue
}

// ThatBool returns an AssertableBool structure initialized with the test reference and the actual bool value to assert.
func ThatBool(t *testing.T, actual bool) AssertableBool {
	t.Helper()
	return AssertableBool{
		t:      t,
		actual: values.NewBoolValue(actual),
	}
}

// IsEqualTo asserts if the expected bool is equal to the assertable bool value
// It errors the tests if the compared values (actual VS expected) are not equal.
func (a AssertableBool) IsEqualTo(expected interface{}) AssertableBool {
	a.t.Helper()
	if !a.actual.IsEqualTo(expected) {
		a.t.Error(shouldBeEqual(a.actual, expected))
	}
	return a
}

// IsNotEqualTo asserts if the expected bool is not equal to the assertable bool value
// It errors the tests if the compared values (actual VS expected) are equal.
func (a AssertableBool) IsNotEqualTo(expected interface{}) AssertableBool {
	a.t.Helper()
	if a.actual.IsEqualTo(expected) {
		a.t.Error(shouldNotBeEqual(a.actual, expected))
	}
	return a
}

// IsTrue asserts if the expected bool value is true.
func (a AssertableBool) IsTrue() AssertableBool {
	a.t.Helper()
	return a.IsEqualTo(true)
}

// IsFalse asserts if the expected bool value is false.
func (a AssertableBool) IsFalse() AssertableBool {
	a.t.Helper()
	return a.IsEqualTo(false)
}
