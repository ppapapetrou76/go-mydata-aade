package assert

import (
	"strings"
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/values"
)

// StringOpt is a configuration option to initialize an AssertableString.
type StringOpt func(*AssertableString)

// AssertableString is the implementation of CommonAssertable for string types.
type AssertableString struct {
	t      *testing.T
	actual values.StringValue
}

// IgnoringCase sets underlying value to lower case.
func IgnoringCase() StringOpt {
	return func(c *AssertableString) {
		c.actual = c.actual.AddDecorator(strings.ToLower)
	}
}

// IgnoringWhiteSpaces removes the whitespaces from the value under assertion.
func IgnoringWhiteSpaces() StringOpt {
	return func(c *AssertableString) {
		c.actual = c.actual.AddDecorator(values.RemoveSpaces)
	}
}

// IgnoringNewLines removes the new lines from the value under assertion.
func IgnoringNewLines() StringOpt {
	return func(c *AssertableString) {
		c.actual = c.actual.AddDecorator(values.RemoveNewLines)
	}
}

// ThatString returns an AssertableString structure initialized with the test reference and the actual value to assert.
func ThatString(t *testing.T, actual string, opts ...StringOpt) AssertableString {
	t.Helper()
	assertable := &AssertableString{
		t:      t,
		actual: values.NewStringValue(actual),
	}
	for _, opt := range opts {
		if opt != nil {
			opt(assertable)
		}
	}
	return *assertable
}

// IsEqualTo asserts if the expected string is equal to the assertable string value
// It errors the tests if the compared values (actual VS expected) are not equal.
func (a AssertableString) IsEqualTo(expected interface{}) AssertableString {
	a.t.Helper()
	if !a.actual.IsEqualTo(expected) {
		a.t.Error(shouldBeEqual(a.actual, expected))
	}
	return a
}

// IsNotEqualTo asserts if the expected string is not equal to the assertable string value
// It errors the tests if the compared values (actual VS expected) are equal.
func (a AssertableString) IsNotEqualTo(expected interface{}) AssertableString {
	a.t.Helper()
	if a.actual.IsEqualTo(expected) {
		a.t.Error(shouldNotBeEqual(a.actual, expected))
	}
	return a
}

// IsEmpty asserts if the expected string is empty
// It errors the tests if the string is not empty.
func (a AssertableString) IsEmpty() AssertableString {
	a.t.Helper()
	if a.actual.IsNotEmpty() {
		a.t.Error(shouldBeEmpty(a.actual))
	}
	return a
}

// IsLowerCase asserts if the expected string is lower case
// It errors the tests if the string is not lower case.
func (a AssertableString) IsLowerCase() AssertableString {
	a.t.Helper()
	if !a.actual.IsLowerCase() {
		a.t.Error(shouldBeLowerCase(a.actual))
	}
	return a
}

// IsUpperCase asserts if the expected string is upper case
// It errors the tests if the string is not upper case.
func (a AssertableString) IsUpperCase() AssertableString {
	a.t.Helper()
	if !a.actual.IsUpperCase() {
		a.t.Error(shouldBeUpperCase(a.actual))
	}
	return a
}

// IsNotEmpty asserts if the expected string is not empty
// It errors the tests if the string is empty.
func (a AssertableString) IsNotEmpty() AssertableString {
	a.t.Helper()
	if a.actual.IsEmpty() {
		a.t.Error(shouldNotBeEmpty(a.actual))
	}
	return a
}

// Contains asserts if the assertable string contains the given element(s)
// It errors the test if it does not contain it.
func (a AssertableString) Contains(substring string) AssertableString {
	a.t.Helper()
	if a.actual.DoesNotContain(substring) {
		a.t.Error(shouldContain(a.actual, substring))
	}
	return a
}

// ContainsIgnoringCase asserts if the assertable string contains the given element(s) case insensitively
// It errors the test if it does not contain it.
func (a AssertableString) ContainsIgnoringCase(substring string) AssertableString {
	a.t.Helper()
	if !a.actual.ContainsIgnoringCase(substring) {
		a.t.Error(shouldContainIgnoringCase(a.actual, substring))
	}
	return a
}

// ContainsOnly asserts if the assertable string only contains the given substring
// It errors the test if it does not contain it.
func (a AssertableString) ContainsOnly(substring string) AssertableString {
	a.t.Helper()
	if !a.actual.ContainsOnly(substring) {
		a.t.Error(shouldContainOnly(a.actual, substring))
	}
	return a
}

// ContainsOnlyOnce asserts if the assertable string contains the given substring only once
// It errors the test if it does not contain it or contains more than once.
func (a AssertableString) ContainsOnlyOnce(substring string) AssertableString {
	a.t.Helper()
	if !a.actual.ContainsOnlyOnce(substring) {
		a.t.Error(shouldContainOnlyOnce(a.actual, substring))
	}
	return a
}

// ContainsWhitespaces asserts if the assertable string contains at least one whitespace
// It errors the test if it does not contain any.
func (a AssertableString) ContainsWhitespaces() AssertableString {
	a.t.Helper()
	if !a.actual.ContainsWhitespaces() {
		a.t.Error(shouldContainWhiteSpace(a.actual))
	}
	return a
}

// DoesNotContainAnyWhitespaces asserts if the assertable string contains no whitespace
// It errors the test if it does contain any.
func (a AssertableString) DoesNotContainAnyWhitespaces() AssertableString {
	a.t.Helper()
	if a.actual.ContainsWhitespaces() {
		a.t.Error(shouldNotContainAnyWhiteSpace(a.actual))
	}
	return a
}

// ContainsOnlyWhitespaces asserts if the assertable string contains only whitespaces
// It errors the test if it contains any other character.
func (a AssertableString) ContainsOnlyWhitespaces() AssertableString {
	a.t.Helper()
	if !a.actual.ContainsOnlyWhitespaces() {
		a.t.Error(shouldContainOnlyWhiteSpaces(a.actual))
	}
	return a
}

// DoesNotContainOnlyWhitespaces asserts if the assertable string does not contain only whitespaces
// It errors the test if it contains only whitespaces.
func (a AssertableString) DoesNotContainOnlyWhitespaces() AssertableString {
	a.t.Helper()
	if a.actual.ContainsOnlyWhitespaces() {
		a.t.Error(shouldNotContainOnlyWhiteSpaces(a.actual))
	}
	return a
}

// DoesNotContain asserts if the assertable string does not contain the given substring
// It errors the test if it contains it.
func (a AssertableString) DoesNotContain(substring string) AssertableString {
	a.t.Helper()
	if a.actual.Contains(substring) {
		a.t.Error(shouldNotContain(a.actual, substring))
	}
	return a
}

// IsSubstringOf asserts if the assertable string is a substring of the given element
// It errors the test if it does not contain it.
func (a AssertableString) IsSubstringOf(someString string) AssertableString {
	a.t.Helper()
	if !a.actual.IsSubstringOf(someString) {
		a.t.Error(shouldBeSubstringOf(a.actual, someString))
	}
	return a
}

// StartsWith asserts if the assertable string starts with the given substring
// It errors the test if it doesn't start with the given substring.
func (a AssertableString) StartsWith(substring string) AssertableString {
	a.t.Helper()
	if !a.actual.StartsWith(substring) {
		a.t.Error(shouldStartWith(a.actual, substring))
	}
	return a
}

// DoesNotStartWith asserts if the assertable string doesn't start with the given substring
// It errors the test if it starts with the given substring.
func (a AssertableString) DoesNotStartWith(substring string) AssertableString {
	a.t.Helper()
	if a.actual.StartsWith(substring) {
		a.t.Error(shouldNotStartWith(a.actual, substring))
	}
	return a
}

// EndsWith asserts if the assertable string ends with the given substring
// It errors the test if it doesn't end with the given substring.
func (a AssertableString) EndsWith(substring string) AssertableString {
	a.t.Helper()
	if !a.actual.EndsWith(substring) {
		a.t.Error(shouldEndWith(a.actual, substring))
	}
	return a
}

// DoesNotEndWith asserts if the assertable string doesn't end with the given substring
// It errors the test if it end with the given substring.
func (a AssertableString) DoesNotEndWith(substring string) AssertableString {
	a.t.Helper()
	if a.actual.EndsWith(substring) {
		a.t.Error(shouldNotEndWith(a.actual, substring))
	}
	return a
}

// HasSameSizeAs asserts if the assertable string has the same size with the given string
// It errors the test if they don't have the same size.
func (a AssertableString) HasSameSizeAs(substring string) AssertableString {
	a.t.Helper()
	if !(a.actual.HasSize(len(substring))) {
		a.t.Error(shouldHaveSameSizeAs(a.actual, substring))
	}
	return a
}

// HasSizeBetween asserts if the assertable string has bigger size of the first given string and less size than the second given string
// It errors the test if the assertable string has the same or less size than the first string or greater than or the same size to the second string.
func (a AssertableString) HasSizeBetween(shortString, longString string) AssertableString {
	a.t.Helper()
	if a.actual.HasSizeLessThanOrEqual(len(shortString)) || a.actual.HasSizeGreaterThanOrEqual(len(longString)) {
		a.t.Error(shouldHaveSizeBetween(a.actual, shortString, longString))
	}
	return a
}

// HasSizeGreaterThan asserts if the assertable string has bigger size of the given string
// It errors the test if the assertable string has the less or equal size to the given one.
func (a AssertableString) HasSizeGreaterThan(substring string) AssertableString {
	a.t.Helper()
	if !(a.actual.HasSizeGreaterThan(len(substring))) {
		a.t.Error(shouldHaveGreaterSizeThan(a.actual, substring))
	}
	return a
}

// HasSizeGreaterThanOrEqualTo asserts if the assertable string has bigger os the same size of the given string
// It errors the test if the assertable string has the less size to the given one.
func (a AssertableString) HasSizeGreaterThanOrEqualTo(substring string) AssertableString {
	a.t.Helper()
	if !(a.actual.HasSizeGreaterThanOrEqual(len(substring))) {
		a.t.Error(shouldHaveGreaterSizeThanOrEqual(a.actual, substring))
	}
	return a
}

// HasSizeLessThan asserts if the assertable string's length is less than the size of the given string
// It errors the test if they don't have the same size.
func (a AssertableString) HasSizeLessThan(substring string) AssertableString {
	a.t.Helper()
	if !(a.actual.HasSizeLessThan(len(substring))) {
		a.t.Error(shouldHaveLessSizeThan(a.actual, substring))
	}
	return a
}

// HasSizeLessThanOrEqualTo asserts if the assertable string's length is less than or equal to the size of the given string
// It errors the test if they don't have the same size.
func (a AssertableString) HasSizeLessThanOrEqualTo(substring string) AssertableString {
	a.t.Helper()
	if !(a.actual.HasSizeLessThanOrEqual(len(substring))) {
		a.t.Error(shouldHaveLessSizeThanOrEqual(a.actual, substring))
	}
	return a
}

// ContainsOnlyDigits asserts if the expected string contains only digits
// It errors the tests if the string has other characters than digits.
func (a AssertableString) ContainsOnlyDigits() AssertableString {
	a.t.Helper()
	if !(a.actual.HasDigitsOnly()) {
		a.t.Error(shouldContainOnlyDigits(a.actual))
	}
	return a
}
