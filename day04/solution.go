package main

import "strings"

type point struct{ row, col int }

func Solve(input string) (int64, int64) {
	lines := strings.Fields(input)
	if len(lines) == 0 {
		return 0, 0
	}
	height, width := len(lines), len(lines[0])
	present := make([][]bool, height)
	neighbors := make([][]int, height)
	queued := make([][]bool, height)
	for r := range height {
		present[r] = make([]bool, width)
		neighbors[r] = make([]int, width)
		queued[r] = make([]bool, width)
		for c := range width {
			present[r][c] = lines[r][c] == '@'
		}
	}
	for r := range height {
		for c := range width {
			if !present[r][c] {
				continue
			}
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					nr, nc := r+dr, c+dc
					if (dr != 0 || dc != 0) && nr >= 0 && nr < height && nc >= 0 && nc < width && present[nr][nc] {
						neighbors[r][c]++
					}
				}
			}
		}
	}
	queue := make([]point, 0)
	for r := range height {
		for c := range width {
			if present[r][c] && neighbors[r][c] < 4 {
				queue = append(queue, point{r, c})
				queued[r][c] = true
			}
		}
	}
	part1 := int64(len(queue))
	var part2 int64
	for head := 0; head < len(queue); head++ {
		p := queue[head]
		if !present[p.row][p.col] {
			continue
		}
		present[p.row][p.col] = false
		part2++
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				nr, nc := p.row+dr, p.col+dc
				if (dr == 0 && dc == 0) || nr < 0 || nr >= height || nc < 0 || nc >= width || !present[nr][nc] {
					continue
				}
				neighbors[nr][nc]--
				if neighbors[nr][nc] < 4 && !queued[nr][nc] {
					queued[nr][nc] = true
					queue = append(queue, point{nr, nc})
				}
			}
		}
	}
	return part1, part2
}
