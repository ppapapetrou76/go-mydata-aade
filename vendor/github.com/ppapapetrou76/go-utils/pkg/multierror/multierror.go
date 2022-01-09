package multierror

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type chain []error

// PrefixedError is a thread-safe, multi-error which will prefix the error output message with the specified prefix.
type PrefixedError struct {
	prefix     string
	errors     chain
	formatFunc FormatFunc

	mu sync.RWMutex
}

// FormatFunc defines a format function which should format a slice of errors into a string.
type FormatFunc func(errs []error) string

// DefaultFormatFunc is the default error formatter that outputs the number of errors
// that occurred along with a bullet point list of the errors.
func DefaultFormatFunc(errs []error) string {
	if len(errs) == 1 {
		return errs[0].Error()
	}

	lines := make([]string, len(errs))
	for i, err := range errs {
		lines[i] = err.Error()
	}

	return fmt.Sprintf(
		"%d errors:[ %s ]",
		len(errs), strings.Join(lines, ", "))
}

// NewPrefixed creates a new pointer of PrefixedError.
func NewPrefixed(prefix string, errs ...error) *PrefixedError {
	return &PrefixedError{prefix: prefix, errors: unwrapErrors(errs...), formatFunc: DefaultFormatFunc}
}

// WithFormatFunc sets the PrefixedError.formatFunc attribute and returns the PrefixedError pointer.
func (p *PrefixedError) WithFormatFunc(formatter FormatFunc) *PrefixedError {
	defer p.mu.RUnlock()
	p.mu.RLock()

	p.formatFunc = formatter

	return p
}

// Append appends a number of errors to the current instance of PrefixedError.
func (p *PrefixedError) Append(errs ...error) *PrefixedError {
	defer p.mu.Unlock()
	p.mu.Lock()

	p.errors = append(p.errors, unwrapErrors(errs...)...)

	return p
}

// ErrorOrNil either returns nil when the type is nil or when there's no errors.
// Otherwise, the type is returned.
func (p *PrefixedError) ErrorOrNil() error {
	defer p.mu.RUnlock()
	p.mu.RLock()

	if len(p.errors) > 0 {
		return p
	}

	return nil
}

// Error returns the stored slice of error formatted using the FormatFunc.
func (p *PrefixedError) Error() string {
	defer p.mu.RUnlock()
	p.mu.RLock()

	if len(p.errors) == 0 {
		return ""
	}

	var prefix string
	if p.prefix != "" {
		prefix = fmt.Sprint(p.prefix, " ")
	}

	return fmt.Sprint(prefix, p.formatFunc(p.errors))
}

// Unwrap returns an error from PrefixedError (or nil if there are no errors).
// This error returned will further support Unwrap to get the next error,
// etc. The order will match the order of errors in the PrefixedError error
// at the time of calling.
//
// The resulting error supports errors.As/Is/Unwrap so you can continue
// to use the stdlib errors package to introspect further.
func (p *PrefixedError) Unwrap() error {
	defer p.mu.RUnlock()
	p.mu.RLock()

	if len(p.errors) <= 1 {
		return nil
	}

	return p.errors[1:]
}

// As implements errors.As by attempting to map to the current value.
func (p *PrefixedError) As(target interface{}) bool {
	defer p.mu.RUnlock()
	p.mu.RLock()

	if len(p.errors) == 0 {
		return false
	}

	return errors.As(p.errors[0], &target)
}

// Is implements errors.Is by comparing the current value directly.
func (p *PrefixedError) Is(target error) bool {
	defer p.mu.RUnlock()
	p.mu.RLock()

	if len(p.errors) == 0 {
		return false
	}

	return errors.Is(p.errors[0], target)
}

func unwrapErrors(errs ...error) []error {
	unWrappedErrors := make([]error, 0, len(errs))
	for _, e := range errs {
		if e == nil {
			continue
		}
		unWrappedErrors = append(unWrappedErrors, e)
	}

	return unWrappedErrors
}

// Error implements the error interface.
func (e chain) Error() string {
	return e[0].Error()
}

// As implements errors.As by attempting to map to the current value.
func (e chain) As(target interface{}) bool {
	return errors.As(e[0], &target)
}

// Is implements errors.Is by comparing the current value directly.
func (e chain) Is(target error) bool {
	return errors.Is(e[0], target)
}
