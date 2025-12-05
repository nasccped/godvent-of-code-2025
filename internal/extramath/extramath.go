package extramath

import "fmt"

// Returns the absolute value of an integer (stdlib provides only
// `abs` of `float64`).
func AbsInt(value int) int {
	if value < 0 {
		value *= -1
	}
	return value
}

// Converts a numeric byte (`0-9`) into an integer. Can return an
// error if the provided byte isn't contained within the numeric
// range.
func NumByteToInt(b byte) (int, error) {
	var (
		i   int
		err error
	)
	if b >= '0' && b <= '9' {
		i = int(b - '0')
	} else {
		err = fmt.Errorf("the provided byte isn't a valid digit (`%c`)", b)
	}
	return i, err
}
