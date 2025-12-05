package day05

import (
	"fmt"
	"math/big"
	"slices"
	"sort"
	"strings"
)

// Type abstraction for range input.
type rangeFields struct {
	minimum, maximum *big.Int
}

// Generate a new rangeFields struct from the given string.
func rangeFieldsFromString(str string) (rangeFields, error) {
	var (
		rf  rangeFields
		err error
	)
	parts := strings.Split(str, "-")
	bigMin, bigMax := new(big.Int), new(big.Int)
	if _, ok := bigMin.SetString(parts[0], 10); !ok {
		err = fmt.Errorf("couldn't parse the string to big int: %s", parts[0])
	} else if _, ok := bigMax.SetString(parts[1], 10); !ok {
		err = fmt.Errorf("couldn't parse the string to big int: %s", parts[1])
	} else {
		rf = rangeFields{minimum: bigMin, maximum: bigMax}
	}
	return rf, err
}

// Checks if a value is contained with the self range struct.
func (r *rangeFields) isContained(value *big.Int) bool {
	return value.Cmp(r.minimum) >= 0 && value.Cmp(r.maximum) <= 0
}

func Part1(input string) int {
	var (
		rf         rangeFields
		err        error
		count      int           = 0
		ranges     []rangeFields = []rangeFields{}
		inputParts []string      = strings.Split(input, "\n\n")
		row        string
	)
	for _, row = range strings.Split(inputParts[0], "\n") {
		if rf, err = rangeFieldsFromString(row); err != nil {
			fmt.Println("error:", err)
			continue
		}
		ranges = append(ranges, rf)
	}
	for _, row = range strings.Split(inputParts[1], "\n") {
		currentBig := new(big.Int)
		if _, ok := currentBig.SetString(row, 10); !ok {
			fmt.Println("error:", fmt.Errorf("couldn't parse the big int (%s)", row))
			continue
		}
		if slices.ContainsFunc(ranges, func(rf rangeFields) bool {
			return rf.isContained(currentBig)
		}) {
			count++
		}
	}
	return count
}

func Part2(input string) int {
	rows := strings.Split(strings.Split(input, "\n\n")[0], "\n")
	ranges := make([]rangeFields, 0, len(rows))
	for _, r := range rows {
		rf, err := rangeFieldsFromString(r)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		ranges = append(ranges, rf)
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].minimum.Cmp(ranges[j].minimum) < 0
	})
	merged := []rangeFields{}
	for _, curr := range ranges {
		n := len(merged)
		if n == 0 {
			merged = append(merged, curr)
			continue
		}
		last := &merged[n-1]

		if curr.minimum.Cmp(new(big.Int).Add(last.maximum, big.NewInt(1))) <= 0 {
			if curr.maximum.Cmp(last.maximum) > 0 {
				last.maximum = curr.maximum
			}
		} else {
			merged = append(merged, curr)
		}
	}
	total := big.NewInt(0)
	one := big.NewInt(1)
	for _, r := range merged {
		length := new(big.Int).Sub(r.maximum, r.minimum)
		length.Add(length, one)
		total.Add(total, length)
	}
	return int(total.Int64())
}
