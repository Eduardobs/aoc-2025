package main

import (
	"math/bits"
	"strconv"
	"strings"
)

const infinity int64 = 1 << 60

type optionSet struct {
	dimensions int
	presses    []int64
	effects    []int64
	byParity   map[uint64][]int
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
	b.Grow(8 * len(values))
	for _, value := range values {
		for shift := 0; shift < 64; shift += 8 {
			b.WriteByte(byte(uint64(value) >> shift))
		}
	}
	return b.String()
}

func minPresses(target []int64, options *optionSet, memo map[string]int64) int64 {
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
	next := make([]int64, len(target))
	for _, choice := range options.byParity[parity(target)] {
		effect := options.effects[choice*options.dimensions : (choice+1)*options.dimensions]
		valid := true
		for i := range target {
			if effect[i] > target[i] {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		for i := range target {
			next[i] = (target[i] - effect[i]) / 2
		}
		cost := minPresses(next, options, memo)
		if cost < infinity && options.presses[choice]+2*cost < best {
			best = options.presses[choice] + 2*cost
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
		totalOptions := 1 << uint(len(buttons))
		options := optionSet{
			dimensions: len(target),
			presses:    make([]int64, totalOptions),
			effects:    make([]int64, totalOptions*len(target)),
			byParity:   make(map[uint64][]int),
		}
		buttonParity := make([]uint64, len(buttons))
		for button, affected := range buttons {
			for _, index := range affected {
				if index >= 0 && index < int64(len(target)) {
					buttonParity[button] ^= 1 << uint(index)
				}
			}
		}
		parities := make([]uint64, totalOptions)
		options.byParity[0] = append(options.byParity[0], 0)
		for mask := 1; mask < totalOptions; mask++ {
			button := bits.TrailingZeros(uint(mask))
			previous := mask & (mask - 1)
			options.presses[mask] = options.presses[previous] + 1
			currentEffect := options.effects[mask*len(target) : (mask+1)*len(target)]
			copy(currentEffect, options.effects[previous*len(target):(previous+1)*len(target)])
			for _, index := range buttons[button] {
				if index >= 0 && index < int64(len(currentEffect)) {
					currentEffect[index]++
				}
			}
			parities[mask] = parities[previous] ^ buttonParity[button]
			options.byParity[parities[mask]] = append(options.byParity[parities[mask]], mask)
		}
		bestLights := infinity
		for _, choice := range options.byParity[lights] {
			if options.presses[choice] < bestLights {
				bestLights = options.presses[choice]
			}
		}
		part1 += bestLights
		part2 += minPresses(target, &options, make(map[string]int64))
	}
	return part1, part2
}
