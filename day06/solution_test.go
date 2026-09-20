package main

import "testing"

func TestSolve(t *testing.T) {
	input := "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  \n"
	p1, p2 := Solve(input)
	if p1 != 4277556 || p2 != 3263827 {
		t.Fatalf("Solve() = (%d, %d), want (4277556, 3263827)", p1, p2)
	}
}
