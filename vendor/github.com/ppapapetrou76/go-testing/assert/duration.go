package assert

import (
	"testing"
	"time"

	"github.com/ppapapetrou76/go-testing/internal/pkg/values"
)

// AssertableDuration is the assertable structure for time.Duration values.
type AssertableDuration struct {
	t      *testing.T
	actual values.DurationValue
}

// ThatDuration returns an AssertableDuration structure initialized with the test reference and the actual value to assert.
func ThatDuration(t *testing.T, actual time.Duration) AssertableDuration {
	t.Helper()
	return AssertableDuration{
		t:      t,
		actual: values.NewDurationValue(actual),
	}
}

// IsEqualTo asserts if the expected time.Duration is equal to the assertable time.Duration value
// It errors the tests if the compared values (actual VS expected) are not equal.
func (a AssertableDuration) IsEqualTo(expected time.Duration) AssertableDuration {
	a.t.Helper()
	if a.actual.IsNotEqualTo(expected) {
		a.t.Error(shouldBeEqual(a.actual, expected))
	}
	return a
}

// IsNotEqualTo asserts if the expected time.Duration is not equal to the assertable time.Duration value
// It errors the tests if the compared values (actual VS expected) are equal.
func (a AssertableDuration) IsNotEqualTo(expected time.Duration) AssertableDuration {
	a.t.Helper()
	if a.actual.IsEqualTo(expected) {
		a.t.Error(shouldNotBeEqual(a.actual, expected))
	}
	return a
}

// IsShorterThan asserts if the assertable time.Duration value is shorter than the expected value
// It errors the tests if is not shorter.
func (a AssertableDuration) IsShorterThan(expected time.Duration) AssertableDuration {
	a.t.Helper()
	if !a.actual.IsShorterThan(expected) {
		a.t.Error(shouldBeShorter(a.actual, expected))
	}
	return a
}

// IsLongerThan asserts if the assertable time.v value is longer than the expected value
// It errors the tests if is not longer.
func (a AssertableDuration) IsLongerThan(expected time.Duration) AssertableDuration {
	a.t.Helper()
	if !a.actual.IsLongerThan(expected) {
		a.t.Error(shouldBeLonger(a.actual, expected))
	}
	return a
}
