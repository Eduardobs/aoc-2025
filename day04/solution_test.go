package main

import "testing"

func TestSolve(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	p1, p2 := Solve(input)
	if p1 != 13 || p2 != 43 {
		t.Fatalf("Solve() = (%d, %d), want (13, 43)", p1, p2)
	}
}

func TestEmptyGrid(t *testing.T) {
	p1, p2 := Solve("...\n...\n")
	if p1 != 0 || p2 != 0 {
		t.Fatalf("Solve() = (%d, %d), want (0, 0)", p1, p2)
	}
}
