package main

import (
	"slices"
	"strconv"
	"strings"
)

type box struct{ x, y, z int64 }
type edge struct {
	distance int64
	a, b     int32
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
	for d.parent[x] != x {
		d.parent[x] = d.parent[d.parent[x]]
		x = d.parent[x]
	}
	return x
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

func squaredDistance(a, b box) int64 {
	dx, dy, dz := a.x-b.x, a.y-b.y, a.z-b.z
	return dx*dx + dy*dy + dz*dz
}

func edgeLess(a, b edge) bool {
	return a.distance < b.distance ||
		(a.distance == b.distance && (a.a < b.a || (a.a == b.a && a.b < b.b)))
}

// retainShortest keeps the limit globally shortest edges in a max-heap.
func retainShortest(edges []edge, candidate edge, limit int) []edge {
	if limit <= 0 {
		return edges
	}
	if len(edges) < limit {
		edges = append(edges, candidate)
		for child := len(edges) - 1; child > 0; {
			parent := (child - 1) / 2
			if !edgeLess(edges[parent], edges[child]) {
				break
			}
			edges[parent], edges[child] = edges[child], edges[parent]
			child = parent
		}
		return edges
	}
	if !edgeLess(candidate, edges[0]) {
		return edges
	}
	edges[0] = candidate
	for parent := 0; ; {
		child := parent*2 + 1
		if child >= len(edges) {
			break
		}
		if right := child + 1; right < len(edges) && edgeLess(edges[child], edges[right]) {
			child = right
		}
		if !edgeLess(edges[parent], edges[child]) {
			break
		}
		edges[parent], edges[child] = edges[child], edges[parent]
		parent = child
	}
	return edges
}

// connectionThreshold finds the lowest distance at which the complete graph
// becomes connected. A dense Prim pass avoids materializing all O(n²) edges.
func connectionThreshold(boxes []box) int64 {
	if len(boxes) < 2 {
		return 0
	}
	used := make([]bool, len(boxes))
	distances := make([]int64, len(boxes))
	known := make([]bool, len(boxes))
	known[0] = true
	var threshold int64
	thresholdSet := false
	for step := 0; step < len(boxes); step++ {
		next := -1
		for node := range boxes {
			if !used[node] && known[node] && (next < 0 || distances[node] < distances[next]) {
				next = node
			}
		}
		used[next] = true
		if step > 0 && (!thresholdSet || distances[next] > threshold) {
			threshold = distances[next]
			thresholdSet = true
		}
		for node := range boxes {
			if used[node] {
				continue
			}
			distance := squaredDistance(boxes[next], boxes[node])
			if !known[node] || distance < distances[node] {
				distances[node], known[node] = distance, true
			}
		}
	}
	return threshold
}

func solveWithLimit(input string, limit int) (int64, int64) {
	boxes := parseBoxes(input)
	totalEdges := len(boxes) * (len(boxes) - 1) / 2
	capacity := min(max(limit, 0), totalEdges)
	edges := make([]edge, 0, capacity)
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			candidate := edge{squaredDistance(boxes[i], boxes[j]), int32(i), int32(j)}
			edges = retainShortest(edges, candidate, limit)
		}
	}
	slices.SortFunc(edges, func(a, b edge) int {
		if edgeLess(a, b) {
			return -1
		}
		if a == b {
			return 0
		}
		return 1
	})
	dsu := newDisjointSet(len(boxes))
	var part1 int64
	for i, connection := range edges {
		a, b := int(connection.a), int(connection.b)
		merged := dsu.unite(a, b)
		if i+1 == limit {
			largest := [3]int{}
			for node := range boxes {
				if dsu.find(node) == node {
					size := dsu.size[node]
					for j := range largest {
						if size > largest[j] {
							size, largest[j] = largest[j], size
						}
					}
				}
			}
			part1 = 1
			for _, size := range largest {
				if size != 0 {
					part1 *= int64(size)
				}
			}
		}
		if merged && dsu.components == 1 {
			break
		}
	}

	if len(boxes) < 2 {
		return part1, 0
	}
	threshold := connectionThreshold(boxes)
	dsu = newDisjointSet(len(boxes))
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			if squaredDistance(boxes[i], boxes[j]) < threshold {
				dsu.unite(i, j)
			}
		}
	}
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			if squaredDistance(boxes[i], boxes[j]) == threshold && dsu.unite(i, j) && dsu.components == 1 {
				return part1, boxes[i].x * boxes[j].x
			}
		}
	}
	return part1, 0
}

func Solve(input string) (int64, int64) { return solveWithLimit(input, 1000) }
