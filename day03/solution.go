package main

import "strings"

func maxJoltage(bank string, digits int) int64 {
	start := 0
	var result int64
	for remaining := digits; remaining > 0; remaining-- {
		end := len(bank) - remaining
		best := start
		for i := start + 1; i <= end; i++ {
			if bank[i] > bank[best] {
				best = i
			}
		}
		result = result*10 + int64(bank[best]-'0')
		start = best + 1
	}
	return result
}

func Solve(input string) (int64, int64) {
	var part1, part2 int64
	for _, bank := range strings.Fields(input) {
		part1 += maxJoltage(bank, 2)
		part2 += maxJoltage(bank, 12)
	}
	return part1, part2
}
