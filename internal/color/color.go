package color

import "fmt"

const (
	RESET  string = "\x1b[0m"
	RED    string = "\x1b[91m"
	GREEN  string = "\x1b[92m"
	YELLOW string = "\x1b[93m"
	CYAN   string = "\x1b[96m"
)

// Type abstraction for colored texting.
type Colored string

// Turns the string into a red string.
func (c Colored) Red() string {
	return fmt.Sprintf("%s%s%s", RED, c, RESET)
}

// Turns the string into a green string.
func (c Colored) Green() string {
	return fmt.Sprintf("%s%s%s", GREEN, c, RESET)
}

// Turns the string into a yellow string.
func (c Colored) Yellow() string {
	return fmt.Sprintf("%s%s%s", YELLOW, c, RESET)
}

// Turns the string into a cyan string.
func (c Colored) Cyan() string {
	return fmt.Sprintf("%s%s%s", CYAN, c, RESET)
}
