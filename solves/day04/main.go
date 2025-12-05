package day04

import (
	"bytes"
	"strings"
)

func Part1(input string) int {
	count := 0
	grid := strings.Split(input, "\n")
	for y, row := range grid {
		for x, cel := range row {
			if cel == '.' {
				continue
			}
			paperCount := 0
			for ny := y - 1; ny <= y+1; ny++ {
				if ny < 0 || ny >= len(grid) {
					continue
				}
				for nx := x - 1; nx <= x+1; nx++ {
					if nx < 0 || nx >= len(grid[ny]) || (ny == y && nx == x) {
						continue
					}
					if grid[ny][nx] == '@' {
						paperCount++
					}
				}
			}
			if paperCount < 4 {
				count++
			}
		}
	}
	return count
}

func Part2(input string) int {
	count := 0
	rows := strings.Split(input, "\n")
	grid := make([][]byte, len(rows))
	for i, r := range rows {
		grid[i] = []byte(r)
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != '@' {
				continue
			}
			paperCount := 0
			for ny := y - 1; ny <= y+1; ny++ {
				if ny < 0 || ny >= len(grid) {
					continue
				}
				for nx := x - 1; nx <= x+1; nx++ {
					if nx < 0 || nx >= len(grid[ny]) || (ny == y && nx == x) {
						continue
					}
					if grid[ny][nx] == '@' {
						paperCount++
					}
				}
			}
			if paperCount < 4 {
				grid[y][x] = 'x'
				count++
			}
		}
	}
	rec := 0
	if count > 0 {
		var newInput bytes.Buffer
		for i, row := range grid {
			newInput.Write(row)
			if i == len(grid)-1 {
				break
			}
			newInput.WriteByte('\n')
		}
		rec = Part2(newInput.String())
	}
	return count + rec
}
