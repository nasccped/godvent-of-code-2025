package day03

import (
	asrt "github.com/nasccped/godvent-of-code-2025/internal/assertion"
	"testing"
)

const input string = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPart1(t *testing.T) {
	assertion := asrt.NewAssertion(Part1(input), t)
	assertion.ExpectsEq(357).Assert()
}

func TestPart2(t *testing.T) {
	assertion := asrt.NewAssertion(Part2(input), t)
	assertion.ExpectsEq(3121910778619).Assert()
}
