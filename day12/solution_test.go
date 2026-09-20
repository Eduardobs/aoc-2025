package main

import "testing"

func TestSolve(t *testing.T) {
	input := `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`
	if got := Solve(input); got != 2 {
		t.Fatalf("Solve() = %d, want 2", got)
	}
}
