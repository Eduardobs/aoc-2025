package main

import (
	"strings"
)

type span struct{ start, end int }

const maxInt64 = int64(^uint64(0) >> 1)

func appendDigit(value int64, digit byte) int64 {
	d := int64(digit - '0')
	if value > (maxInt64-d)/10 {
		return maxInt64
	}
	return value*10 + d
}

func evaluate(values []int64, operator byte) int64 {
	if operator == '+' {
		var total int64
		for _, value := range values {
			total += value
		}
		return total
	}
	result := int64(1)
	for _, value := range values {
		result *= value
	}
	return result
}

func Solve(input string) (int64, int64) {
	input = strings.TrimSuffix(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	lines := strings.Split(input, "\n")
	if len(lines) < 2 {
		return 0, 0
	}
	height := len(lines) - 1
	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}
	charAt := func(row, col int) byte {
		if col < 0 || col >= len(lines[row]) {
			return ' '
		}
		return lines[row][col]
	}
	groups := make([]span, 0)
	for col := 0; col < width; {
		hasDigit := false
		for row := 0; row < height; row++ {
			if b := charAt(row, col); b >= '0' && b <= '9' {
				hasDigit = true
				break
			}
		}
		if !hasDigit {
			col++
			continue
		}
		start := col
		for col < width {
			occupied := false
			for row := 0; row < height; row++ {
				if b := charAt(row, col); b >= '0' && b <= '9' {
					occupied = true
					break
				}
			}
			if !occupied {
				break
			}
			col++
		}
		groups = append(groups, span{start, col})
	}
	operators := strings.Fields(lines[height])
	var part1, part2 int64
	for i, group := range groups {
		if i >= len(operators) {
			break
		}
		op := operators[i][0]
		horizontal := int64(1)
		if op == '+' {
			horizontal = 0
		}
		for row := 0; row < height; row++ {
			var value int64
			hasDigit := false
			for col := group.start; col < group.end; col++ {
				if b := charAt(row, col); b >= '0' && b <= '9' {
					value = appendDigit(value, b)
					hasDigit = true
				}
			}
			if hasDigit {
				if op == '+' {
					horizontal += value
				} else {
					horizontal *= value
				}
			}
		}
		vertical := int64(1)
		if op == '+' {
			vertical = 0
		}
		for col := group.start; col < group.end; col++ {
			var value int64
			for row := 0; row < height; row++ {
				if b := charAt(row, col); b >= '0' && b <= '9' {
					value = appendDigit(value, b)
				}
			}
			if op == '+' {
				vertical += value
			} else {
				vertical *= value
			}
		}
		part1 += horizontal
		part2 += vertical
	}
	return part1, part2
}
