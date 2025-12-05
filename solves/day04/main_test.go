package day04

import (
	asrt "github.com/nasccped/godvent-of-code-2025/internal/assertion"
	"testing"
)

const input string = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestPart1(t *testing.T) {
	assertion := asrt.NewAssertion(Part1(input), t)
	assertion.ExpectsEq(13).Assert()
}

func TestPart2(t *testing.T) {
	assertion := asrt.NewAssertion(Part2(input), t)
	assertion.ExpectsEq(43).Assert()
}
