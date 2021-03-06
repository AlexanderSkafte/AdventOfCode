package main

import (
	"fmt"
	"time"
)

type aocFn func([]string) string

func main() {
	aocs := []aocFn{
		aoc1a, aoc1b,
		aoc2a, aoc2b,
		aoc3a, aoc3b,
		aoc4a, aoc4b,
		aoc5a, aoc5b,
		aoc6a, aoc6b,
		aoc7a, aoc7b,
	}

	fmt.Println("\n" +
		"   °~.~°~.~°~.~°~.~°~.~°~.~°\n" +
		"      Advent of Code 2018   \n" +
		"   °~.~°~.~°~.~°~.~°~.~°~.~°\n")
	for i := 0; i < len(aocs); i += 2 {
		fmt.Printf("Dec %d:\n", i/2+1)

		lines := readLines(fmt.Sprintf("input%d.txt", i/2+1))

		a, t := measure(aocs[i], lines)
		fmt.Printf(" a) %8s | %s\n", t, a)

		b, t := measure(aocs[i+1], lines)
		fmt.Printf(" b) %8s | %s\n", t, b)

		fmt.Println()
	}
}

func measure(fn aocFn, lines []string) (string, time.Duration) {
	start := time.Now()
	result := fn(lines)
	elapsed := time.Since(start)
	return result, elapsed.Round(time.Millisecond / 10)
}
