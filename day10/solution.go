package main

import (
	"strconv"
	"strings"
)

const infinity int64 = 1 << 60

type subset struct {
	presses int64
	effect  []int64
}

func numbers(raw string) []int64 {
	raw = strings.Trim(raw, "(){}")
	if raw == "" {
		return nil
	}
	fields := strings.Split(raw, ",")
	result := make([]int64, 0, len(fields))
	for _, field := range fields {
		value, err := strconv.ParseInt(field, 10, 64)
		if err == nil {
			result = append(result, value)
		}
	}
	return result
}

func parity(values []int64) uint64 {
	var result uint64
	for i, value := range values {
		if value&1 != 0 {
			result |= 1 << uint(i)
		}
	}
	return result
}

func stateKey(values []int64) string {
	var b strings.Builder
	for _, value := range values {
		b.WriteString(strconv.FormatInt(value, 10))
		b.WriteByte(',')
	}
	return b.String()
}

func minPresses(target []int64, options map[uint64][]subset, memo map[string]int64) int64 {
	allZero := true
	for _, value := range target {
		if value < 0 {
			return infinity
		}
		allZero = allZero && value == 0
	}
	if allZero {
		return 0
	}
	key := stateKey(target)
	if cached, ok := memo[key]; ok {
		return cached
	}
	best := infinity
	for _, choice := range options[parity(target)] {
		next := make([]int64, len(target))
		valid := true
		for i := range target {
			if choice.effect[i] > target[i] {
				valid = false
				break
			}
			next[i] = (target[i] - choice.effect[i]) / 2
		}
		if !valid {
			continue
		}
		cost := minPresses(next, options, memo)
		if cost < infinity && choice.presses+2*cost < best {
			best = choice.presses + 2*cost
		}
	}
	memo[key] = best
	return best
}

func Solve(input string) (int64, int64) {
	var part1, part2 int64
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		lightsRaw := strings.Trim(fields[0], "[]")
		var lights uint64
		for i, light := range lightsRaw {
			if light == '#' {
				lights |= 1 << uint(i)
			}
		}
		buttons := make([][]int64, len(fields)-2)
		for i, raw := range fields[1 : len(fields)-1] {
			buttons[i] = numbers(raw)
		}
		target := numbers(fields[len(fields)-1])
		options := make(map[uint64][]subset)
		for mask := 0; mask < 1<<uint(len(buttons)); mask++ {
			effect := make([]int64, len(target))
			var presses int64
			for button, affected := range buttons {
				if mask&(1<<uint(button)) == 0 {
					continue
				}
				presses++
				for _, index := range affected {
					if index >= 0 && index < int64(len(effect)) {
						effect[index]++
					}
				}
			}
			p := parity(effect)
			options[p] = append(options[p], subset{presses, effect})
		}
		bestLights := infinity
		for _, choice := range options[lights] {
			if choice.presses < bestLights {
				bestLights = choice.presses
			}
		}
		part1 += bestLights
		part2 += minPresses(target, options, make(map[string]int64))
	}
	return part1, part2
}
