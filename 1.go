package main

import (
	"fmt"
)

func aoc1a(lines []string) string {
	sum := 0
	for _, line := range lines {
		sum += readInt(line)
	}
	return fmt.Sprint(sum)
}

func aoc1b(lines []string) string {
	sumSet := map[int]int{}

	ints := make([]int, len(lines))
	for i := range lines {
		ints[i] = readInt(lines[i])
	}

	sum := 0
	for {
		for _, num := range ints {
			sumSet[sum]++
			if sumSet[sum] > 1 {
				return fmt.Sprint(sum)
			}
			sum += num
		}
	}
}
