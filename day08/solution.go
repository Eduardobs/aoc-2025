package main

import (
	"sort"
	"strconv"
	"strings"
)

type box struct{ x, y, z int64 }
type edge struct {
	distance int64
	a, b     int
}

type disjointSet struct {
	parent, size []int
	components   int
}

func newDisjointSet(n int) *disjointSet {
	d := &disjointSet{parent: make([]int, n), size: make([]int, n), components: n}
	for i := range n {
		d.parent[i], d.size[i] = i, 1
	}
	return d
}

func (d *disjointSet) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *disjointSet) unite(a, b int) bool {
	a, b = d.find(a), d.find(b)
	if a == b {
		return false
	}
	if d.size[a] < d.size[b] {
		a, b = b, a
	}
	d.parent[b] = a
	d.size[a] += d.size[b]
	d.components--
	return true
}

func parseBoxes(input string) []box {
	result := make([]box, 0)
	for _, line := range strings.Fields(input) {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		values := [3]int64{}
		ok := true
		for i := range 3 {
			var err error
			values[i], err = strconv.ParseInt(parts[i], 10, 64)
			ok = ok && err == nil
		}
		if ok {
			result = append(result, box{values[0], values[1], values[2]})
		}
	}
	return result
}

func solveWithLimit(input string, limit int) (int64, int64) {
	boxes := parseBoxes(input)
	edges := make([]edge, 0, len(boxes)*(len(boxes)-1)/2)
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			dx, dy, dz := boxes[i].x-boxes[j].x, boxes[i].y-boxes[j].y, boxes[i].z-boxes[j].z
			edges = append(edges, edge{dx*dx + dy*dy + dz*dz, i, j})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].distance != edges[j].distance {
			return edges[i].distance < edges[j].distance
		}
		if edges[i].a != edges[j].a {
			return edges[i].a < edges[j].a
		}
		return edges[i].b < edges[j].b
	})
	dsu := newDisjointSet(len(boxes))
	var part1, part2 int64
	for i, connection := range edges {
		merged := dsu.unite(connection.a, connection.b)
		if i+1 == limit {
			sizes := make([]int, 0)
			for node := range boxes {
				if dsu.find(node) == node {
					sizes = append(sizes, dsu.size[node])
				}
			}
			sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
			part1 = 1
			for j := 0; j < 3 && j < len(sizes); j++ {
				part1 *= int64(sizes[j])
			}
		}
		if merged && dsu.components == 1 {
			part2 = boxes[connection.a].x * boxes[connection.b].x
			break
		}
	}
	return part1, part2
}

func Solve(input string) (int64, int64) { return solveWithLimit(input, 1000) }
