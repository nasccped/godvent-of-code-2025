package assertion

import (
	"bytes"
	"fmt"
	c "github.com/nasccped/godvent-of-code-2025/internal/color"
	"github.com/nasccped/godvent-of-code-2025/internal/pointer"
	"testing"
)

const (
	// Assert if the provided value is equals to the expected
	// one.
	assertEq int = 1
	// Assert if the provided value ISN'T equals to the other
	// one.
	assertNeq int = 2
)

// Common panic function for the Assertion.Assert method.
func assertionAssertPanic(message string) {
	panic(fmt.Sprintf("Assertion.Assert panic: %s", message))
}

// Type alias for the testing.
type Assertion[C comparable] struct {
	input      C
	other      *C
	assertMode *int
	message    *string
	test       *testing.T
}

// Generates a new assertion based on a input value and a testing
// struct pointer.
func NewAssertion[C comparable](input C, test *testing.T) *Assertion[C] {
	return &Assertion[C]{input: input, test: test}
}

// Set the expected value + the assertion mode (assertEq).
func (a *Assertion[C]) ExpectsEq(other C) *Assertion[C] {
	a.assertMode = pointer.GetPointer(assertEq)
	a.other = &other
	return a
}

// Set the non-expected value (expecting different from) + the
// assertion mode (assertNeq).
func (a *Assertion[C]) ExpectsNeq(other C) *Assertion[C] {
	a.assertMode = pointer.GetPointer(assertNeq)
	a.other = &other
	return a
}

// Set the message if test fails.
func (a *Assertion[C]) WithMessage(message string) *Assertion[C] {
	a.message = &message
	return a
}

// Runs the assertion.
func (a *Assertion[T]) Assert() {
	if a.test == nil {
		assertionAssertPanic("the `test` field can't be nil.")
	} else if a.assertMode == nil {
		assertionAssertPanic(
			"the `assertionMode` field wasn't set.\n" +
				"You can set it using the `ExpectsEq` or `ExpectsNeq` methods.\n" +
				"note: use `ExpectsEq(true)` for condition assert.",
		)
	} else if a.other == nil {
		assertionAssertPanic("the `other` value must be a non nil pointer.")
	}
	message := a.getMessage()
	switch *a.assertMode {
	case assertEq:
		if a.input != *a.other {
			a.test.Errorf(message)
		}
	case assertNeq:
		if a.input == *a.other {
			a.test.Errorf(message)
		}
	default:
		panic("undefined variant")
	}
}

func (a *Assertion[C]) getMessage() string {
	var buf bytes.Buffer
	var temp string
	var intoRed func(string) string = func(title string) string {
		return c.Colored(title).Red()
	}
	switch *a.assertMode {
	case assertEq:
		buf.WriteString(intoRed("AssertEq error: "))
		temp = fmt.Sprintf("left value is `%+v` but right value is `%+v`.", a.input, *a.other)
	case assertNeq:
		buf.WriteString(intoRed("AssertNeq error: "))
		temp = fmt.Sprintf("values are equals (`%+v`)", a.input)
	default:
		panic("unexpected variant")
	}
	if a.message != nil && *a.message != "" {
		buf.WriteString(*a.message)
	} else {
		buf.WriteString(temp)
	}
	return buf.String()
}
