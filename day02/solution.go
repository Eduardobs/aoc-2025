package main

import (
	"strconv"
	"strings"
)

func repeated(s string, exactlyTwice bool) bool {
	for size := 1; size*2 <= len(s); size++ {
		if len(s)%size != 0 {
			continue
		}
		copies := len(s) / size
		if exactlyTwice && copies != 2 {
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

// Solve sums invalid product IDs in the comma-separated inclusive ranges.
func Solve(input string) (int64, int64) {
	var part1, part2 int64
	for _, raw := range strings.Split(strings.TrimSpace(input), ",") {
		bounds := strings.Split(strings.TrimSpace(raw), "-")
		if len(bounds) != 2 {
			continue
		}
		start, err1 := strconv.ParseInt(bounds[0], 10, 64)
		end, err2 := strconv.ParseInt(bounds[1], 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}
		for id := start; id <= end; id++ {
			s := strconv.FormatInt(id, 10)
			if repeated(s, true) {
				part1 += id
			}
			if repeated(s, false) {
				part2 += id
			}
		}
	}
	return part1, part2
}
