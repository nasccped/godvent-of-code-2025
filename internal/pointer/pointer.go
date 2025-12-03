package pointer

// Returns a pointer of the provided value (since I can't get
// pointers of constants).
func GetPointer[T any](value T) *T {
	return &value
}
