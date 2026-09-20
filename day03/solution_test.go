package main

import "testing"

func TestSolve(t *testing.T) {
	input := "987654321111111\n811111111111119\n234234234234278\n818181911112111\n"
	p1, p2 := Solve(input)
	if p1 != 357 || p2 != 3121910778619 {
		t.Fatalf("Solve() = (%d, %d), want (357, 3121910778619)", p1, p2)
	}
}

func TestMaxJoltageKeepsOrder(t *testing.T) {
	if got := maxJoltage("12345", 2); got != 45 {
		t.Fatalf("maxJoltage() = %d, want 45", got)
	}
}
