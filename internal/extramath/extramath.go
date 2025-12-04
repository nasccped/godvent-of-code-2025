package extramath

// Returns the absolute value of an integer (stdlib provides only
// `abs` of `float64`).
func AbsInt(value int) int {
	if value < 0 {
		value *= -1
	}
	return value
}
