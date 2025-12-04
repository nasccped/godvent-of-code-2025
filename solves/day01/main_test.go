package day01

import (
	asrt "github.com/nasccped/godvent-of-code-2025/internal/assertion"
	"testing"
)

const input string = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestPart1(t *testing.T) {
	assertion := asrt.NewAssertion(Part1(INPUT), t)
	assertion.ExpectsEq(3).Assert()
}

func TestPart2(t *testing.T) {
	assertion := asrt.NewAssertion(Part2(INPUT), t)
	assertion.ExpectsEq(6).Assert()
}
