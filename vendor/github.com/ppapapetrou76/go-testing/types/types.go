package types

// Assertable is the basic interface for all assertable values.
type Assertable interface {
	Value() interface{}
}

// Comparable is the interface for basic comparable operations.
type Comparable interface {
	IsEqualTo(expected interface{}) bool
	Assertable
}

// ExtendedComparable is the interface for advanced comparable operations.
type ExtendedComparable interface {
	IsGreaterOrEqualTo(expected interface{}) bool
	IsGreaterThan(expected interface{}) bool
	IsLessThan(expected interface{}) bool
	IsLessOrEqualTo(expected interface{}) bool
	Comparable
}

// Sizeable is the interface for operations related to sizeable values.
type Sizeable interface {
	IsEmpty() bool
	IsNotEmpty() bool
	HasSize(length int) bool
	Size() int
	Comparable
}

// Containable is the interface for operations related to containable values such as string or slice.
type Containable interface {
	Contains(elements interface{}) bool
	ContainsOnly(elements interface{}) bool
	DoesNotContain(elements interface{}) bool
	HasUniqueElements() bool
	IsSorted(descending bool) bool
	Sizeable
}

// Nullable is the interface for operations related to nullable values such as pointers or slices.
type Nullable interface {
	IsNil() bool
	IsNotNil() bool
	Assertable
}

// Map is the interface for operations related to map values.
type Map interface {
	HasKey(key interface{}) bool
	HasValue(value interface{}) bool
	HasEntry(entry MapEntry) bool
	Sizeable
}

// MapEntry is the struct to represent a key-value pair used in maps.
type MapEntry struct {
	key, value interface{}
}

// Key returns the key of the map entry.
func (entry MapEntry) Key() interface{} {
	return entry.key
}

// Value returns the value of the map entry.
func (entry MapEntry) Value() interface{} {
	return entry.value
}

// NewMapEntry creates and returns a new map entry with the given values of key and value fields.
func NewMapEntry(key, value interface{}) MapEntry {
	return MapEntry{
		key:   key,
		value: value,
	}
}
