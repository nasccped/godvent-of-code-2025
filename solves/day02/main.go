package day02

import (
	"strconv"
	"strings"
)

func Part1(input string) int {
	ranges := strings.Split(input, ",")
	invalids := []int{}
	for _, r := range ranges {
		temp := strings.Split(r, "-")
		mini, _ := strconv.Atoi(temp[0])
		maxi, _ := strconv.Atoi(temp[1])
		for i := mini; i <= maxi; i++ {
			current := strconv.Itoa(i)
			if len(current)%2 != 0 {
				continue
			}
			half := len(current) / 2
			if current[:half] == current[half:] {
				invalids = append(invalids, i)
			}
		}
	}
	result := 0
	for _, val := range invalids {
		result += val
	}
	return result
}

func isInvalid(s, part string) bool {
	if part == "" {
		return false
	} else if len(s)%len(part) != 0 {
		return isInvalid(s, part[:len(part)-1])
	}
	for i := len(part); i <= len(s); i += len(part) {
		if s[i-len(part):i] != part {
			return isInvalid(s, part[:len(part)-1])
		}
	}
	return true
}

func Part2(input string) int {
	ranges := strings.Split(input, ",")
	invalids := []int{}
	for _, r := range ranges {
		temp := strings.Split(r, "-")
		mini, _ := strconv.Atoi(temp[0])
		maxi, _ := strconv.Atoi(temp[1])
		for i := mini; i <= maxi; i++ {
			current := strconv.Itoa(i)
			half := current[:len(current)/2]
			if isInvalid(current, half) {
				invalids = append(invalids, i)
			}
		}
	}
	result := 0
	for _, v := range invalids {
		result += v
	}
	return result
}
