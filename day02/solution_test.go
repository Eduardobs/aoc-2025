package main

import "testing"

func TestSolve(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`
	p1, p2 := Solve(input)
	if p1 != 1227775554 || p2 != 4174379265 {
		t.Fatalf("Solve() = (%d, %d), want (1227775554, 4174379265)", p1, p2)
	}
}

func TestRepeated(t *testing.T) {
	for _, tc := range []struct {
		s     string
		twice bool
		any   bool
	}{{"123123", true, true}, {"121212", false, true}, {"123124", false, false}} {
		if got := repeated(tc.s, true); got != tc.twice {
			t.Errorf("repeated(%q, true) = %v, want %v", tc.s, got, tc.twice)
		}
		if got := repeated(tc.s, false); got != tc.any {
			t.Errorf("repeated(%q, false) = %v, want %v", tc.s, got, tc.any)
		}
	}
}
