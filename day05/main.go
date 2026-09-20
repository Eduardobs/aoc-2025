package main

import (
	"fmt"
	"os"
)

func main() {
	path := "day05/input.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	p1, p2 := Solve(string(data))
	fmt.Printf("Parte 1: %d\nParte 2: %d\n", p1, p2)
}
