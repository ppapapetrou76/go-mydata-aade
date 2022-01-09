package values

import "reflect"

func areEqualValues(actualValue, expectedValue reflect.Value) bool {
	switch actualValue.Kind() {
	case reflect.String:
		return NewStringValue(actualValue.String()).IsEqualTo(expectedValue.String())
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return NewIntValue(actualValue.Int()).IsEqualTo(expectedValue.Int())
	case reflect.Bool:
		return NewBoolValue(actualValue.Bool()).IsEqualTo(expectedValue.Bool())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return NewUIntValue(actualValue.Uint()).IsEqualTo(expectedValue.Uint())
	case reflect.Float32, reflect.Float64:
		// This might panic - we need to implement a NewFloatValue
		return actualValue.Float() == expectedValue.Float()
	case reflect.Array, reflect.Slice:
		return areSlicesEqual(actualValue, expectedValue)
	case reflect.Map:
		return areMapsEqual(actualValue, expectedValue)
	case reflect.Struct:
		if actualValue.CanInterface() && expectedValue.CanInterface() {
			return NewStructValue(actualValue.Interface()).IsEqualTo(expectedValue.Interface())
		}
		return NewStructValue(actualValue).IsEqualTo(expectedValue)
	case reflect.Interface:
		return NewAnyValue(actualValue.Interface()).IsEqualTo(expectedValue.Interface())
	case reflect.Chan, reflect.Complex64, reflect.Complex128, reflect.Func, reflect.Invalid, reflect.Ptr, reflect.Uintptr, reflect.UnsafePointer:
		// not supported yet
		return true
	default:
		return true
	}
}

func areMapsEqual(actualValue, expectedValue reflect.Value) bool {
	if actualValue.Len() != expectedValue.Len() {
		return false
	}
	if actualValue.Len() > 0 && expectedValue.Len() > 0 {
		for _, k := range actualValue.MapKeys() {
			if !expectedValue.MapIndex(k).IsValid() {
				return false
			}
			if !areEqualValues(actualValue.MapIndex(k), expectedValue.MapIndex(k)) {
				return false
			}
		}
	}
	return true
}

func areSlicesEqual(actualValue, expectedValue reflect.Value) bool {
	if actualValue.Len() != expectedValue.Len() {
		return false
	}
	if actualValue.Len() > 0 && expectedValue.Len() > 0 {
		for i := 0; i < actualValue.Len(); i++ {
			if !areEqualValues(actualValue.Index(i), expectedValue.Index(i)) {
				return false
			}
		}
	}
	return true
}
