package main

import "strings"

type point struct{ row, col int }

func Solve(input string) (int64, int64) {
	lines := strings.Fields(input)
	if len(lines) == 0 {
		return 0, 0
	}
	height, width := len(lines), len(lines[0])
	present := make([]bool, height*width)
	neighbors := make([]uint8, height*width)
	queued := make([]bool, height*width)
	presentCount := 0
	for r, line := range lines {
		for c := range width {
			if line[c] == '@' {
				present[r*width+c] = true
				presentCount++
			}
		}
	}
	for r := range height {
		for c := range width {
			index := r*width + c
			if !present[index] {
				continue
			}
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					nr, nc := r+dr, c+dc
					if (dr != 0 || dc != 0) && nr >= 0 && nr < height && nc >= 0 && nc < width && present[nr*width+nc] {
						neighbors[index]++
					}
				}
			}
		}
	}
	queue := make([]point, 0, presentCount)
	for r := range height {
		for c := range width {
			index := r*width + c
			if present[index] && neighbors[index] < 4 {
				queue = append(queue, point{r, c})
				queued[index] = true
			}
		}
	}
	part1 := int64(len(queue))
	var part2 int64
	for head := 0; head < len(queue); head++ {
		p := queue[head]
		index := p.row*width + p.col
		if !present[index] {
			continue
		}
		present[index] = false
		part2++
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				nr, nc := p.row+dr, p.col+dc
				if (dr == 0 && dc == 0) || nr < 0 || nr >= height || nc < 0 || nc >= width {
					continue
				}
				neighbor := nr*width + nc
				if !present[neighbor] {
					continue
				}
				neighbors[neighbor]--
				if neighbors[neighbor] < 4 && !queued[neighbor] {
					queued[neighbor] = true
					queue = append(queue, point{nr, nc})
				}
			}
		}
	}
	return part1, part2
}
