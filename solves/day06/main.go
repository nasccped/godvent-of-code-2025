package day06

import (
	"strconv"
	"strings"
)

// Filter the slice based on a given condition.
func filterSlice[T any](slice []T, condition func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		if condition(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map all items of a slice to a new value/type based on a given
// function.
func mapSlice[T any, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = f(item)
	}
	return result
}

func Part1(input string) int {
	rows := strings.Split(input, "\n")
	values := mapSlice(mapSlice(rows[:len(rows)-1], func(row string) []string {
		return filterSlice(strings.Split(row, " "), func(part string) bool {
			return part != " " && part != ""
		})
	}), func(nums []string) []int {
		return mapSlice(nums, func(n string) int {
			v, _ := strconv.Atoi(n)
			return v
		})
	})
	opers := filterSlice(strings.Split(rows[len(rows)-1], " "), func(item string) bool {
		return item != " " && item != ""
	})
	count := 0
	for x := range len(values[0]) {
		current := 0
		if opers[x] == "*" {
			current++
		}
		for y := range len(values) {
			switch opers[x] {
			case "+":
				current += values[y][x]
			default:
				current *= values[y][x]
			}

		}
		count += current
	}
	return count
}

func Part2(input string) int {
	return 0
}
