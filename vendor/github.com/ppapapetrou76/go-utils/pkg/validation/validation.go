package validation

import (
	"fmt"
	"reflect"

	"github.com/ppapapetrou76/go-utils/pkg/multierror"
)

// IsRequired returns a formatted error message if the given element is nil or empty, else it returns nil.
// For now it supports the following : reflect.Interface, reflect.Ptr, reflect.String, reflect.Array,
// reflect.Slice and reflect.Map.
func IsRequired(element interface{}, name string) error {
	value := reflect.ValueOf(element)
	kind := value.Kind()

	switch kind { //nolint:exhaustive // no reason to check for the other reflection types here.
	case reflect.Invalid:
		if element == nil {
			return fmt.Errorf("%s is nil", name)
		}
	case reflect.Interface, reflect.Ptr:
		if element == nil || value.IsNil() {
			return fmt.Errorf("%s is nil", name)
		}
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		if value.Len() == 0 {
			return fmt.Errorf("%s is empty", name)
		}
	}

	return nil
}

// HasNoNilElements accepts a Slice or an Array and validates that all elements of the container are not nil.
// If one or more elements are nil then it returns a multierror.PrefixedError error with information about which elements
// are nil.
//
// If there are no nil elements or the target is not an Array or Slice then it returns nil.
func HasNoNilElements(target interface{}) error {
	value := reflect.ValueOf(target)
	kind := value.Kind()

	switch kind { //nolint:exhaustive //no reason to check for the other reflection types here.
	case reflect.Array, reflect.Slice:
		errs := multierror.NewPrefixed("")
		for i := 0; i < value.Len(); i++ {
			errs = errs.Append(IsRequired(value.Index(i).Interface(), fmt.Sprintf("index %d", i)))
		}

		return errs.ErrorOrNil() //nolint:wrapcheck // it's already using the multierror
	default:
		return nil
	}
}
