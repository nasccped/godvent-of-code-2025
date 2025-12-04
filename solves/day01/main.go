package day01

import (
	"github.com/nasccped/godvent-of-code-2025/internal/extramath"
	"strconv"
	"strings"
)

func Part1(input string) int {
	current := 50
	count := 0
	for _, r := range strings.Split(input, "\n") {
		direction := 1
		if r[0] == 'L' {
			direction = -1
		}
		sure, _ := strconv.Atoi(r[1:])
		amount := direction * sure
		current = (current + amount) % 100
		if current == 0 {
			count++
		}
	}
	return count
}

func Part2(input string) int {
	current := 50
	count := 0
	for _, r := range strings.Split(input, "\n") {
		sure, _ := strconv.Atoi(r[1:])
		direction := 1
		amount := direction * sure
		count += extramath.AbsInt(amount / 100)
		amount %= 100
		var temp int
		if r[0] == 'L' {
			temp = current - amount
			if current != 0 && temp <= 0 {
				count++
			}
			current = (100 + temp) % 100
		} else {
			temp = current + amount
			if current != 0 && temp >= 100 {
				count++
			}
			current = temp % 100
		}
	}
	return count
}
