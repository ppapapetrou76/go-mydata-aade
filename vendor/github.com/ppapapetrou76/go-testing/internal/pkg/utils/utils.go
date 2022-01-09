package utils

import (
	"reflect"
	"unicode"
	"unicode/utf8"
)

// HasUnexportedFields returns true if the given value is a struct and has at least one unexported field.
func HasUnexportedFields(value reflect.Value) bool {
	if value.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < value.NumField(); i++ {
		if !IsFieldExported(value.Type().Field(i).Name) {
			return true
		}
	}
	return false
}

// IsFieldExported returns true if the given name is uppercase.
func IsFieldExported(name string) bool {
	ch, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(ch)
}
