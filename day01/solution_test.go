package main

import "testing"

func TestSolve(t *testing.T) {
	input := "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82\n"
	p1, p2 := Solve(input)
	if p1 != 3 || p2 != 6 {
		t.Fatalf("Solve() = (%d, %d), want (3, 6)", p1, p2)
	}
}

func TestCrossesZeroMoreThanOnce(t *testing.T) {
	p1, p2 := Solve("R250\n")
	if p1 != 1 || p2 != 3 {
		t.Fatalf("Solve() = (%d, %d), want (1, 3)", p1, p2)
	}
}
