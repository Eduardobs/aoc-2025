package main

import "testing"

func TestSolve(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	p1, p2 := Solve(input)
	if p1 != 3 || p2 != 14 {
		t.Fatalf("Solve() = (%d, %d), want (3, 14)", p1, p2)
	}
}

func TestContainedAndAdjacentRanges(t *testing.T) {
	p1, p2 := Solve("1-10\n3-5\n11-12\n\n1\n12\n13\n")
	if p1 != 2 || p2 != 12 {
		t.Fatalf("Solve() = (%d, %d), want (2, 12)", p1, p2)
	}
}
