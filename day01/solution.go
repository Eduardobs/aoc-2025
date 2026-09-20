package main

import (
	"strconv"
	"strings"
)

// Solve follows the dial rotations. Part 1 counts rotations that finish at
// zero; part 2 counts every click that lands on zero during a rotation.
func Solve(input string) (int64, int64) {
	position := 50
	var part1, part2 int64
	for _, line := range strings.Fields(input) {
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			continue
		}
		direction := 1
		if line[0] == 'L' {
			direction = -1
		}
		for range distance {
			position = (position + direction + 100) % 100
			if position == 0 {
				part2++
			}
		}
		if position == 0 {
			part1++
		}
	}
	return part1, part2
}
