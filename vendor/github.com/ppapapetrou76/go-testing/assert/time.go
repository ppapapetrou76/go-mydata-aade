package assert

import (
	"testing"
	"time"

	"github.com/ppapapetrou76/go-testing/internal/pkg/types"
)

// AssertableTime is the assertable structure for time.Time values.
type AssertableTime struct {
	t      *testing.T
	actual types.TimeValue
}

// ThatTime returns an AssertableTime structure initialized with the test reference and the actual value to assert.
func ThatTime(t *testing.T, actual time.Time) AssertableTime {
	t.Helper()
	return AssertableTime{
		t:      t,
		actual: types.NewTimeValue(actual),
	}
}

// IsSameAs asserts if the expected time.Time is equal to the assertable time.Time value
// It errors the tests if the compared values (actual VS expected) are not equal.
func (a AssertableTime) IsSameAs(expected time.Time) AssertableTime {
	a.t.Helper()
	if a.actual.IsNotSameAs(expected) {
		a.t.Error(shouldBeEqual(a.actual, expected))
	}
	return a
}

// IsAlmostSameAs asserts if the expected time.Time is almost equal to the assertable time.Time value
// It errors the tests if the compared values (actual VS expected) are not equal.
func (a AssertableTime) IsAlmostSameAs(expected time.Time) AssertableTime {
	a.t.Helper()
	if !a.actual.IsAlmostSameAs(expected) {
		a.t.Error(shouldBeAlmostSame(a.actual, expected))
	}
	return a
}

// IsNotTheSameAs asserts if the expected time.Time is not equal to the assertable time.Time value
// It errors the tests if the compared values (actual VS expected) are equal.
func (a AssertableTime) IsNotTheSameAs(expected time.Time) AssertableTime {
	a.t.Helper()
	if a.actual.IsSameAs(expected) {
		a.t.Error(shouldNotBeEqual(a.actual, expected))
	}
	return a
}

// IsBefore asserts if the assertable time.Time value is before the expected value
// It errors the tests if is not greater.
func (a AssertableTime) IsBefore(expected time.Time) AssertableTime {
	a.t.Helper()
	if !a.actual.IsBefore(expected) {
		a.t.Error(shouldBeGreater(a.actual, expected))
	}
	return a
}

// IsAfter asserts if the assertable time.Time value is after the expected value
// It errors the tests if is not later.
func (a AssertableTime) IsAfter(expected time.Time) AssertableTime {
	a.t.Helper()
	if !a.actual.IsAfter(expected) {
		a.t.Error(shouldBeGreaterOrEqual(a.actual, expected))
	}
	return a
}

// IsDefined asserts if the expected time.Time is defined.
// It errors the tests if the value is not defined.
func (a AssertableTime) IsDefined() AssertableTime {
	a.t.Helper()
	if a.actual.IsNotDefined() {
		a.t.Error(shouldBeDefined(a.actual))
	}
	return a
}

// IsNotDefined asserts if the expected time.Time is not defined.
// It errors the tests if the value is defined.
func (a AssertableTime) IsNotDefined() AssertableTime {
	a.t.Helper()
	if a.actual.IsDefined() {
		a.t.Error(shouldNotBeDefined(a.actual))
	}
	return a
}
