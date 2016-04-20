package util

import "time"

// OptionalUInt8 Returns a pointer to a value
func OptionalUInt8(value uint8) *uint8 {
	return &value
}

// OptionalUInt16 Returns a pointer to a value
func OptionalUInt16(value uint16) *uint16 {
	return &value
}

// OptionalUInt Returns a pointer to a value
func OptionalUInt(value uint) *uint {
	return &value
}

// OptionalUInt64 Returns a pointer to a value
func OptionalUInt64(value uint64) *uint64 {
	return &value
}

// OptionalInt16 Returns a pointer to a value
func OptionalInt16(value int16) *int16 {
	return &value
}

// OptionalInt Returns a pointer to a value
func OptionalInt(value int) *int {
	return &value
}

// OptionalInt64 Returns a pointer to a value
func OptionalInt64(value int64) *int64 {
	return &value
}

// OptionalFloat32 Returns a pointer to a value
func OptionalFloat32(value float32) *float32 {
	return &value
}

// OptionalFloat64 Returns a pointer to a value
func OptionalFloat64(value float64) *float64 {
	return &value
}

// OptionalString Returns a pointer to a value
func OptionalString(value string) *string {
	return &value
}

// OptionalBool Returns a pointer to a value
func OptionalBool(value bool) *bool {
	return &value
}

// OptionalTime Returns a pointer to a value
func OptionalTime(value time.Time) *time.Time {
	return &value
}
