package main

import "strings"

func Solve(input string) (int64, int64) {
	lines := strings.Fields(input)
	if len(lines) == 0 {
		return 0, 0
	}
	width := len(lines[0])
	start := strings.IndexByte(lines[0], 'S')
	active := make([]bool, width)
	ways := make([]int64, width)
	nextActive := make([]bool, width)
	nextWays := make([]int64, width)
	active[start] = true
	ways[start] = 1
	var part1 int64
	for _, line := range lines[1:] {
		copy(nextActive, active)
		copy(nextWays, ways)
		for col := 0; col < width; col++ {
			if line[col] != '^' {
				continue
			}
			if active[col] {
				part1++
				nextActive[col] = false
				if col > 0 {
					nextActive[col-1] = true
				}
				if col+1 < width {
					nextActive[col+1] = true
				}
			}
			if ways[col] != 0 {
				count := ways[col]
				nextWays[col] -= count
				if col > 0 {
					nextWays[col-1] += count
				}
				if col+1 < width {
					nextWays[col+1] += count
				}
			}
		}
		active, nextActive = nextActive, active
		ways, nextWays = nextWays, ways
	}
	var part2 int64
	for _, count := range ways {
		part2 += count
	}
	return part1, part2
}
