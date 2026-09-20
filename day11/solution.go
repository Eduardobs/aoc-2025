package main

import "strings"

type state struct {
	node    string
	seenDAC bool
	seenFFT bool
}

func parseGraph(input string) map[string][]string {
	graph := make(map[string][]string)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		from := strings.TrimSuffix(parts[0], ":")
		graph[from] = append([]string(nil), parts[1:]...)
	}
	return graph
}

func countPaths(graph map[string][]string, current string, seenDAC, seenFFT bool, memo map[state]int64) int64 {
	seenDAC = seenDAC || current == "dac"
	seenFFT = seenFFT || current == "fft"
	if current == "out" {
		if seenDAC && seenFFT {
			return 1
		}
		return 0
	}
	key := state{current, seenDAC, seenFFT}
	if cached, ok := memo[key]; ok {
		return cached
	}
	var total int64
	for _, next := range graph[current] {
		total += countPaths(graph, next, seenDAC, seenFFT, memo)
	}
	memo[key] = total
	return total
}

func Solve(input string) (int64, int64) {
	graph := parseGraph(input)
	part1 := countPaths(graph, "you", true, true, make(map[state]int64))
	part2 := countPaths(graph, "svr", false, false, make(map[state]int64))
	return part1, part2
}
