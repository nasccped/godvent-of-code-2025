package day06

import (
	asrt "github.com/nasccped/godvent-of-code-2025/internal/assertion"
	"testing"
)

const input string = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

func TestPart1(t *testing.T) {
	assertion := asrt.NewAssertion(Part1(input), t)
	assertion.ExpectsEq(4277556).Assert()
}

func TestPart2(t *testing.T) {
	assertion := asrt.NewAssertion(Part2(input), t)
	assertion.ExpectsEq(0).Assert()
}
