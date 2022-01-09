package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/values"
	"github.com/ppapapetrou76/go-testing/types"
)

// SliceOpt is a configuration option to initialize an AssertableAny Slice.
type SliceOpt func(*AssertableSlice)

// AssertableSlice is the implementation of AssertableSlice for string slices.
type AssertableSlice struct {
	t             *testing.T
	actual        types.Containable
	customMessage string
}

// WithCustomMessage provides a custom message to be added before the assertion error message.
func WithCustomMessage(customMessage string) SliceOpt {
	return func(c *AssertableSlice) {
		c.customMessage = customMessage
	}
}

// ThatSlice returns a proper assertable structure based on the slice type.
func ThatSlice(t *testing.T, actual interface{}, opts ...SliceOpt) AssertableSlice {
	t.Helper()
	assertable := &AssertableSlice{
		t:      t,
		actual: values.NewSliceValue(actual),
	}
	for _, opt := range opts {
		opt(assertable)
	}
	return *assertable
}

// IsEqualTo asserts if the expected slice is equal to the assertable slice value
// It errors the tests if the compared values (actual VS expected) are not equal.
func (a AssertableSlice) IsEqualTo(expected interface{}) AssertableSlice {
	a.t.Helper()
	if !a.actual.IsEqualTo(expected) {
		a.t.Error(shouldBeEqual(a.actual, expected))
	}
	return a
}

// IsNotEqualTo asserts if the expected slice is not equal to the assertable slice value
// It errors the tests if the compared values (actual VS expected) are equal.
func (a AssertableSlice) IsNotEqualTo(expected interface{}) AssertableSlice {
	a.t.Helper()
	if a.actual.IsEqualTo(expected) {
		a.t.Error(shouldNotBeEqual(a.actual, expected))
	}
	return a
}

// HasSize asserts if the assertable string slice has the expected length size
// It errors the test if it doesn't have the expected size.
func (a AssertableSlice) HasSize(size int) AssertableSlice {
	a.t.Helper()
	if !a.actual.HasSize(size) {
		a.t.Error(shouldHaveSize(a.actual, size))
	}
	return a
}

// IsEmpty asserts if the assertable string slice is empty or not.
func (a AssertableSlice) IsEmpty() AssertableSlice {
	a.t.Helper()
	if a.actual.IsNotEmpty() {
		a.t.Error(shouldBeEmpty(a.actual))
	}
	return a
}

// IsNotEmpty asserts if the assertable string slice is not empty.
func (a AssertableSlice) IsNotEmpty() AssertableSlice {
	a.t.Helper()
	if a.actual.IsEmpty() {
		a.t.Error(shouldNotBeEmpty(a.actual))
	}
	return a
}

// Contains asserts if the assertable string slice contains the given element(s)
// It errors the test if it does not contain it/them.
func (a AssertableSlice) Contains(elements interface{}) AssertableSlice {
	a.t.Helper()
	if a.actual.DoesNotContain(elements) {
		a.t.Error(shouldContain(a.actual, elements))
	}
	return a
}

// ContainsOnly asserts if the assertable string slice contains only the given element(s)
// It errors the test if it does not contain it/them.
func (a AssertableSlice) ContainsOnly(elements interface{}) AssertableSlice {
	a.t.Helper()
	if !a.actual.ContainsOnly(elements) {
		a.t.Error(shouldContainOnly(a.actual, elements))
	}
	return a
}

// DoesNotContain asserts if the assertable string slice does not contain the given element
// It errors the test if it contains it/them.
func (a AssertableSlice) DoesNotContain(elements interface{}) AssertableSlice {
	a.t.Helper()
	if a.actual.Contains(elements) {
		a.t.Error(shouldNotContain(a.actual, elements))
	}
	return a
}

// HasUniqueElements asserts if the assertable string slice does not contain the given element
// It errors the test if it contains it/them.
func (a AssertableSlice) HasUniqueElements() AssertableSlice {
	if !(a.actual.HasUniqueElements()) {
		a.t.Error(shouldContainUniqueElements(a.actual))
	}
	return a
}

// IsSortedAscending asserts if the assertable slice is sorted on ascending order.
// It supports the following slices : []string, []int, []int32, []int64, []float64 and any slice that implements
// sort.Interface
// If a non-supported type is given then the assertion will fail.
func (a AssertableSlice) IsSortedAscending() AssertableSlice {
	if !(a.actual.IsSorted(false)) {
		a.t.Error(shouldBeSorted(a.actual, "ascending"))
	}
	return a
}

// IsSortedDescending asserts if the assertable slice is sorted on descending order.
// It supports the following slices : []string, []int, []int32, []int64, []float64 and any slice that implements
// sort.Interface
// If a non-supported type is given then the assertion will fail.
func (a AssertableSlice) IsSortedDescending() AssertableSlice {
	if !(a.actual.IsSorted(true)) {
		a.t.Error(shouldBeSorted(a.actual, "descending"))
	}
	return a
}
