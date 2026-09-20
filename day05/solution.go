package main

import (
	"sort"
	"strconv"
	"strings"
)

type interval struct{ start, end int64 }

func Solve(input string) (int64, int64) {
	sections := strings.SplitN(strings.TrimSpace(input), "\n\n", 2)
	if len(sections) != 2 {
		return 0, 0
	}
	ranges := make([]interval, 0)
	for _, line := range strings.Fields(sections[0]) {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}
		start, err1 := strconv.ParseInt(parts[0], 10, 64)
		end, err2 := strconv.ParseInt(parts[1], 10, 64)
		if err1 == nil && err2 == nil {
			ranges = append(ranges, interval{start, end})
		}
	}
	if len(ranges) == 0 {
		return 0, 0
	}
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].start == ranges[j].start {
			return ranges[i].end < ranges[j].end
		}
		return ranges[i].start < ranges[j].start
	})
	merged := []interval{ranges[0]}
	for _, current := range ranges[1:] {
		last := &merged[len(merged)-1]
		if current.start <= last.end+1 {
			if current.end > last.end {
				last.end = current.end
			}
		} else {
			merged = append(merged, current)
		}
	}
	var part1 int64
	for _, field := range strings.Fields(sections[1]) {
		id, err := strconv.ParseInt(field, 10, 64)
		if err != nil {
			continue
		}
		i := sort.Search(len(merged), func(i int) bool { return merged[i].start > id })
		if i > 0 && id <= merged[i-1].end {
			part1++
		}
	}
	var part2 int64
	for _, r := range merged {
		part2 += r.end - r.start + 1
	}
	return part1, part2
}
