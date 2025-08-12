package utils

// ToPtr converts a value to a pointer
func ToPtr[T any](v T) *T {
	return &v
}
