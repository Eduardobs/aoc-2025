package main

import "testing"

func TestPart1Example(t *testing.T) {
	input := `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`
	p1, _ := Solve(input)
	if p1 != 5 {
		t.Fatalf("part 1 = %d, want 5", p1)
	}
}

func TestPart2Example(t *testing.T) {
	input := `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`
	_, p2 := Solve(input)
	if p2 != 2 {
		t.Fatalf("part 2 = %d, want 2", p2)
	}
}
