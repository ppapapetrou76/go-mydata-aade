package types

// NewBool creates a new bool pointer from a boolean.
func NewBool(value bool) *bool { return &value }

// NewString creates a new string pointer from a string.
func NewString(value string) *string { return &value }

// NewUint32 creates a new uint32 pointer from an uint32.
func NewUint32(value uint32) *uint32 { return &value }

// NewUint64 creates a new uint64 pointer from an uint64.
func NewUint64(value uint64) *uint64 { return &value }

// NewUint creates a new uint pointer from an uint.
func NewUint(value uint) *uint { return &value }

// NewInt32 creates a new int32 pointer from an int32.
func NewInt32(value int32) *int32 { return &value }

// NewInt64 creates a new int64 pointer from an int64.
func NewInt64(value int64) *int64 { return &value }

// NewInt creates a new int pointer from an int.
func NewInt(value int) *int { return &value }
