package assert

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/values"
)

// AssertableStruct is the implementation of AssertableAny for structs.
type AssertableStruct struct {
	actual values.StructValue
	t      *testing.T
}

// ThatStruct returns a proper assertable structure based on the slice type.
func ThatStruct(t *testing.T, actual interface{}) AssertableStruct {
	t.Helper()
	return AssertableStruct{
		actual: values.NewStructValue(actual),
		t:      t,
	}
}

// IsEqualTo asserts if the expected structure is equal to the assertable structure value
// It errors the tests if the compared values (actual VS expected) are not equal.
func (s AssertableStruct) IsEqualTo(expected interface{}) AssertableStruct {
	s.t.Helper()
	if !s.actual.IsEqualTo(expected) {
		s.t.Error(shouldBeEqual(s.actual, expected))
	}
	return s
}

// IsNotEqualTo asserts if the expected structure is not equal to the assertable structure value
// It errors the tests if the compared values (actual VS expected) are equal.
func (s AssertableStruct) IsNotEqualTo(expected interface{}) AssertableStruct {
	s.t.Helper()
	if s.actual.IsEqualTo(expected) {
		s.t.Error(shouldNotBeEqual(s.actual, expected))
	}
	return s
}

// ExcludingFields excludes the given fields from struct comparisons.
func (s AssertableStruct) ExcludingFields(fields ...string) AssertableStruct {
	s.actual.ExcludedFields = fields
	return s
}
