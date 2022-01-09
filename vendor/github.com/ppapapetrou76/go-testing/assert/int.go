package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/values"
)

// AssertableInt is the assertable structure for int values.
type AssertableInt struct {
	t      *testing.T
	actual values.IntValue
}

// ThatInt returns an AssertableInt structure initialized with the test reference and the actual value to assert.
func ThatInt(t *testing.T, actual int) AssertableInt {
	t.Helper()
	return AssertableInt{
		t:      t,
		actual: values.NewIntValue(actual),
	}
}

// IsEqualTo asserts if the expected int is equal to the assertable int value
// It errors the tests if the compared values (actual VS expected) are not equal.
func (a AssertableInt) IsEqualTo(expected int) AssertableInt {
	a.t.Helper()
	if !a.actual.IsEqualTo(expected) {
		a.t.Error(shouldBeEqual(a.actual, expected))
	}
	return a
}

// IsNotEqualTo asserts if the expected int is not equal to the assertable int value
// It errors the tests if the compared values (actual VS expected) are equal.
func (a AssertableInt) IsNotEqualTo(expected int) AssertableInt {
	a.t.Helper()
	if a.actual.IsEqualTo(expected) {
		a.t.Error(shouldNotBeEqual(a.actual, expected))
	}
	return a
}

// IsGreaterThan asserts if the assertable int value is greater than the expected value
// It errors the tests if is not greater.
func (a AssertableInt) IsGreaterThan(expected int) AssertableInt {
	a.t.Helper()
	if !a.actual.IsGreaterThan(expected) {
		a.t.Error(shouldBeGreater(a.actual, expected))
	}
	return a
}

// IsGreaterThanOrEqualTo asserts if the assertable int value is greater than or equal to the expected value
// It errors the tests if is not greater.
func (a AssertableInt) IsGreaterThanOrEqualTo(expected int) AssertableInt {
	a.t.Helper()
	if !a.actual.IsGreaterOrEqualTo(expected) {
		a.t.Error(shouldBeGreaterOrEqual(a.actual, expected))
	}
	return a
}

// IsLessThan asserts if the assertable int value is less than the expected value
// It errors the tests if is not greater.
func (a AssertableInt) IsLessThan(expected int) AssertableInt {
	a.t.Helper()
	if !a.actual.IsLessThan(expected) {
		a.t.Error(shouldBeLessThan(a.actual, expected))
	}
	return a
}

// IsLessThanOrEqualTo asserts if the assertable int value is less than or equal to the expected value
// It errors the tests if is not greater.
func (a AssertableInt) IsLessThanOrEqualTo(expected int) AssertableInt {
	a.t.Helper()
	if !a.actual.IsLessOrEqualTo(expected) {
		a.t.Error(shouldBeLessOrEqual(a.actual, expected))
	}

	return a
}
