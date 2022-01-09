package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/values"
)

// AssertableError is the assertable structure for error values.
type AssertableError struct {
	t      *testing.T
	actual values.ErrorValue
}

// ThatError returns an AssertableError structure initialized with the test reference and the actual value to assert.
func ThatError(t *testing.T, actual error) AssertableError {
	t.Helper()
	return AssertableError{
		t:      t,
		actual: values.NewErrorValue(actual),
	}
}

// IsNil asserts if the expected error is nil.
func (a AssertableError) IsNil() AssertableError {
	a.t.Helper()
	errAnyValue := values.NewAnyValue(a.actual.Value())
	if errAnyValue.IsNotNil() {
		a.t.Error(shouldBeNil(errAnyValue))
	}
	return a
}

// IsNotNil asserts if the expected error is nil.
func (a AssertableError) IsNotNil() AssertableError {
	a.t.Helper()
	errAnyValue := values.NewAnyValue(a.actual.Value())
	if errAnyValue.IsNil() {
		a.t.Error(shouldNotBeNil(errAnyValue))
	}
	return a
}

// HasExactMessage asserts if the expected error contains exactly the given message.
func (a AssertableError) HasExactMessage(expectedMessage string) AssertableError {
	a.t.Helper()
	errAnyValue := values.NewAnyValue(a.actual.Value())
	if errAnyValue.IsNil() {
		a.t.Error(shouldContain(errAnyValue, expectedMessage))
		return a
	}

	errStringValue := values.NewStringValue(a.actual.Error().Error())
	if !errStringValue.ContainsOnly(expectedMessage) {
		a.t.Error(shouldContain(errAnyValue, expectedMessage))
	}
	return a
}

// IsSameAs asserts if the expected error is the same with the given error.
func (a AssertableError) IsSameAs(err error) AssertableError {
	a.t.Helper()
	actualAnyValue := values.NewAnyValue(a.actual.Value())
	expectedAnyValue := values.NewAnyValue(err)

	if actualAnyValue.IsNil() != expectedAnyValue.IsNil() {
		a.t.Error(shouldBeEqual(a.actual, expectedAnyValue))
		return a
	}
	if actualAnyValue.IsNil() && expectedAnyValue.IsNil() {
		return a
	}

	actualStringValue := values.NewStringValue(a.actual.Error().Error())
	if !actualStringValue.IsEqualTo(err.Error()) {
		a.t.Error(shouldBeEqual(a.actual, err.Error()))
	}
	return a
}
