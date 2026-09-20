package main

import (
	"strconv"
	"strings"
)

func repeated(s string, exactlyTwice bool) bool {
	if exactlyTwice {
		half := len(s) / 2
		return len(s) > 0 && len(s)%2 == 0 && s[:half] == s[half:]
	}
	for size := 1; size*2 <= len(s); size++ {
		if len(s)%size != 0 {
			continue
		}
		ok := true
		for i := size; i < len(s); i += size {
			if s[i:i+size] != s[:size] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}

func repeatedSums(start, end int64) (int64, int64) {
	if start > end || end < 11 {
		return 0, 0
	}
	if start < 1 {
		start = 1
	}

	var powers [20]uint64
	powers[0] = 1
	for i := 1; i < len(powers); i++ {
		powers[i] = powers[i-1] * 10
	}

	low, high := uint64(start), uint64(end)
	all := make(map[int64]struct{})
	var part1 int64
	for digits := 2; digits <= 19; digits++ {
		for copies := 2; copies <= digits; copies++ {
			if digits%copies != 0 {
				continue
			}
			patternDigits := digits / copies
			var multiplier uint64
			for offset := 0; offset < digits; offset += patternDigits {
				multiplier += powers[offset]
			}
			first := powers[patternDigits-1]
			last := powers[patternDigits] - 1
			minimum := low / multiplier
			if low%multiplier != 0 {
				minimum++
			}
			minimum = max(minimum, first)
			maximum := min(high/multiplier, last)
			for pattern := minimum; pattern <= maximum; pattern++ {
				id := int64(pattern * multiplier)
				if copies == 2 {
					part1 += id
				}
				all[id] = struct{}{}
			}
		}
	}
	var part2 int64
	for id := range all {
		part2 += id
	}
	return part1, part2
}

// Solve sums invalid product IDs in the comma-separated inclusive ranges.
func Solve(input string) (int64, int64) {
	var part1, part2 int64
	for _, raw := range strings.Split(strings.TrimSpace(input), ",") {
		lower, upper, ok := strings.Cut(strings.TrimSpace(raw), "-")
		if !ok {
			continue
		}
		start, err1 := strconv.ParseInt(lower, 10, 64)
		end, err2 := strconv.ParseInt(upper, 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}
		p1, p2 := repeatedSums(start, end)
		part1 += p1
		part2 += p2
	}
	return part1, part2
}
