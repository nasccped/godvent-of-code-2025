package day03

import (
	"bytes"
	"github.com/nasccped/godvent-of-code-2025/internal/extramath"
	"strconv"
	"strings"
)

// Function alias for the `NumByteToInt` function.
func nbti(b byte) int {
	res, _ := extramath.NumByteToInt(b)
	return res
}

// Searches for the greatest byte within a string input and returns
// the byte itself + its index. This function panics if the string
// input is empty.
func maxByte(s string) (byte, int) {
	if s == "" {
		panic("empty string isn't allowed")
	}
	ind := 0
	b := s[ind]
	for i := 1; i < len(s); i++ {
		if temp := s[i]; temp > b {
			b, ind = temp, i
		}
	}
	return b, ind
}

func Part1(input string) int {
	count := 0
	for _, row := range strings.Split(input, "\n") {
		ten, tenInd := maxByte(row[:len(row)-1])
		unit, _ := maxByte(row[tenInd+1:])
		count += nbti(ten)*10 + nbti(unit)
	}
	return count
}

func Part2(input string) int {
	count := 0
	for _, row := range strings.Split(input, "\n") {
		seq := make([]int, 12)
		for i := range 12 {
			seq[11-i] = len(row) - 1 - i
		}
		for i := range len(seq) {
			minimum := 0
			if i != 0 {
				minimum = seq[i-1] + 1
			}
			for j := seq[i]; j >= minimum; j-- {
				if row[j] >= row[seq[i]] {
					seq[i] = j
				}
			}
		}
		var buf bytes.Buffer
		for _, i := range seq {
			buf.WriteByte(row[i])
		}
		temp, _ := strconv.Atoi(buf.String())
		count += temp
	}
	return count
}
