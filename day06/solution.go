package main

import (
	"strconv"
	"strings"
)

type span struct{ start, end int }

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
		horizontal := make([]int64, 0, height)
		for row := 0; row < height; row++ {
			var digits strings.Builder
			for col := group.start; col < group.end; col++ {
				if b := charAt(row, col); b >= '0' && b <= '9' {
					digits.WriteByte(b)
				}
			}
			if digits.Len() > 0 {
				value, _ := strconv.ParseInt(digits.String(), 10, 64)
				horizontal = append(horizontal, value)
			}
		}
		vertical := make([]int64, 0, group.end-group.start)
		for col := group.start; col < group.end; col++ {
			var digits strings.Builder
			for row := 0; row < height; row++ {
				if b := charAt(row, col); b >= '0' && b <= '9' {
					digits.WriteByte(b)
				}
			}
			value, _ := strconv.ParseInt(digits.String(), 10, 64)
			vertical = append(vertical, value)
		}
		part1 += evaluate(horizontal, op)
		part2 += evaluate(vertical, op)
	}
	return part1, part2
}
