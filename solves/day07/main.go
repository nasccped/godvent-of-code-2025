package day07

import (
	"strings"
)

func Part1(input string) int {
	singleRow := make([]bool, strings.Index(input, "\n"))
	start := strings.Index(input, "S")
	singleRow[start] = true
	counter := 0
	for _, row := range strings.Split(input, "\n")[1:] {
		for i, item := range singleRow {
			if item && row[i] == '^' {
				if i > 0 {
					singleRow[i-1] = true
				}
				if i < len(row)-1 {
					singleRow[i+1] = true
				}
				singleRow[i] = false
				counter++
			}
		}
	}
	return counter
}

func Part2(input string) int {
	coords := map[int]int{}
	start := strings.Index(input, "S")
	coords[start] = 1
	for _, row := range strings.Split(input, "\n")[1:] {
		localCoords := map[int]int{}
		for k, v := range coords {
			if row[k] == '^' {
				if k > 0 {
					localCoords[k-1] += v
				}
				if k < len(row)-1 {
					localCoords[k+1] += v
				}
			} else {
				localCoords[k] += v
			}
		}
		coords = localCoords
	}
	accum := 0
	for _, v := range coords {
		accum += v
	}
	return accum
}
