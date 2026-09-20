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
		if distance > 0 {
			part2 += int64(distance / 100)
			remainder := distance % 100
			if direction > 0 {
				if position+remainder >= 100 {
					part2++
				}
				position = (position + remainder) % 100
			} else {
				if position != 0 && remainder >= position {
					part2++
				}
				position = (position - remainder + 100) % 100
			}
		}
		if position == 0 {
			part1++
		}
	}
	return part1, part2
}
