// Package runner contém a infraestrutura compartilhada pelos executáveis de cada dia.
package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Run executa uma solução com duas partes.
func Run(day string, solve func(string) (int64, int64)) {
	input := readInput(day)
	part1, part2 := solve(input)
	fmt.Printf("Parte 1: %d\nParte 2: %d\n", part1, part2)
}

// RunSinglePart executa uma solução que possui somente uma parte oficial.
func RunSinglePart(day string, solve func(string) int64) {
	input := readInput(day)
	fmt.Printf(
		"Parte 1: %d\nParte 2: N/A (o dia %s possui apenas uma parte)\n",
		solve(input), strings.TrimPrefix(day, "day"),
	)
}

func readInput(day string) string {
	path := filepath.Join(day, "input.txt")
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return string(data)
}
