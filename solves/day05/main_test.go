package day05

import (
	asrt "github.com/nasccped/godvent-of-code-2025/internal/assertion"
	"testing"
)

const input string = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestPart1(t *testing.T) {
	assertion := asrt.NewAssertion(Part1(input), t)
	assertion.ExpectsEq(3).Assert()
}

func TestPart2(t *testing.T) {
	assertion := asrt.NewAssertion(Part2(input), t)
	assertion.ExpectsEq(14).Assert()
}
