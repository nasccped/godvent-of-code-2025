package day02

import (
	asrt "github.com/nasccped/godvent-of-code-2025/internal/assertion"
	"testing"
)

const input string = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestPart1(t *testing.T) {
	assertion := asrt.NewAssertion(Part1(INPUT), t)
	assertion.ExpectsEq(1227775554).Assert()
}

func TestPart2(t *testing.T) {
	assertion := asrt.NewAssertion(Part2(INPUT), t)
	assertion.ExpectsEq(4174379265).Assert()
}
