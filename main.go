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
	}

	fmt.Println("\n" +
		"   °~.~°~.~°~.~°~.~°~.~°~.~°\n" +
		"      Advent of Code 2018   \n" +
		"   °~.~°~.~°~.~°~.~°~.~°~.~°\n")
	for i := 0; i < len(aocs); i += 2 {
		fmt.Printf("Dec %d:\n", i/2+1)

		lines := readLines(fmt.Sprintf("input%d.txt", i/2+1))

		a, t := measure(aocs[i], lines)
		fmt.Printf(" a) %8s | %s\n", t, a.(string))

		b, t := measure(aocs[i+1], lines)
		fmt.Printf(" b) %8s | %s\n", t, b.(string))

		fmt.Println()
	}
}

func measure(fn aocFn, lines []string) (interface{}, time.Duration) {
	start := time.Now()
	result := fn(lines)
	elapsed := time.Since(start)
	return result, elapsed.Round(time.Millisecond / 10)
}
