package main

import (
	"fmt"
)

type aocFn func() string

func main() {
	aocs := []aocFn{
		aoc1a, aoc1b,
		aoc2a, aoc2b,
	}

	fmt.Println("\n" +
		"   °~.~°~.~°~.~°~.~°~.~°~.~°\n" +
		"      Advent of Code 2018   \n" +
		"   °~.~°~.~°~.~°~.~°~.~°~.~°\n")
	for i := 0; i < len(aocs); i += 2 {
		fmt.Printf("Dec %d:\n", i+1)
		fmt.Printf("  a) %s\n", aocs[i]())
		fmt.Printf("  b) %s\n", aocs[i+1]())
		fmt.Println()
	}
}
