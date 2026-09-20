package main

import (
	"fmt"
	"os"
)

func main() {
	path := "day12/input.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Parte 1: %d\nParte 2: N/A (o dia 12 possui apenas uma parte)\n", Solve(string(data)))
}
