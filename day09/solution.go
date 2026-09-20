package main

import (
	"slices"
	"strconv"
	"strings"
)

type point struct{ x, y int64 }
type rectangle struct{ x1, y1, x2, y2 int64 }

func makeRectangle(a, b point) rectangle {
	r := rectangle{a.x, a.y, b.x, b.y}
	if r.x1 > r.x2 {
		r.x1, r.x2 = r.x2, r.x1
	}
	if r.y1 > r.y2 {
		r.y1, r.y2 = r.y2, r.y1
	}
	return r
}

func (r rectangle) area() int64 { return (r.x2 - r.x1 + 1) * (r.y2 - r.y1 + 1) }

// crossesInterior reports whether a polygon edge enters the open interior of r.
// Touching or following the rectangle boundary is allowed.
func (r rectangle) crossesInterior(edge rectangle) bool {
	return edge.x1 < r.x2 && edge.x2 > r.x1 && edge.y1 < r.y2 && edge.y2 > r.y1
}

func insideOrBoundary(x, y float64, polygon []point) bool {
	inside := false
	for i, a := range polygon {
		b := polygon[(i+1)%len(polygon)]
		cross := (float64(b.x-a.x))*(y-float64(a.y)) - (float64(b.y-a.y))*(x-float64(a.x))
		if cross == 0 && x >= float64(min(a.x, b.x)) && x <= float64(max(a.x, b.x)) &&
			y >= float64(min(a.y, b.y)) && y <= float64(max(a.y, b.y)) {
			return true
		}
		ay, by := float64(a.y), float64(b.y)
		if (ay > y) != (by > y) {
			intersectionX := float64(b.x-a.x)*(y-ay)/(by-ay) + float64(a.x)
			if x < intersectionX {
				inside = !inside
			}
		}
	}
	return inside
}

func parsePoints(input string) []point {
	result := make([]point, 0)
	for _, line := range strings.Fields(input) {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, err1 := strconv.ParseInt(parts[0], 10, 64)
		y, err2 := strconv.ParseInt(parts[1], 10, 64)
		if err1 == nil && err2 == nil {
			result = append(result, point{x, y})
		}
	}
	return result
}

func Solve(input string) (int64, int64) {
	vertices := parsePoints(input)
	if len(vertices) < 2 {
		return 0, 0
	}
	pairs := make([]rectangle, 0, len(vertices)*(len(vertices)-1)/2)
	for i := range vertices {
		for j := i + 1; j < len(vertices); j++ {
			pairs = append(pairs, makeRectangle(vertices[i], vertices[j]))
		}
	}
	slices.SortFunc(pairs, func(a, b rectangle) int {
		areaA, areaB := a.area(), b.area()
		if areaA > areaB {
			return -1
		}
		if areaA < areaB {
			return 1
		}
		return 0
	})
	part1 := pairs[0].area()
	edges := make([]rectangle, len(vertices))
	for i := range vertices {
		edges[i] = makeRectangle(vertices[i], vertices[(i+1)%len(vertices)])
	}
	var part2 int64
	for _, candidate := range pairs {
		valid := true
		for _, edge := range edges {
			if candidate.crossesInterior(edge) {
				valid = false
				break
			}
		}
		if valid && insideOrBoundary(float64(candidate.x1+candidate.x2)/2, float64(candidate.y1+candidate.y2)/2, vertices) {
			part2 = candidate.area()
			break
		}
	}
	return part1, part2
}
