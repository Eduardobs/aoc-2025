package main

import "testing"

func TestSolve(t *testing.T) {
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	p1, p2 := Solve(input)
	if p1 != 50 || p2 != 24 {
		t.Fatalf("Solve() = (%d, %d), want (50, 24)", p1, p2)
	}
}
