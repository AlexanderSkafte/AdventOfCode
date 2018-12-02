package main

import (
	"fmt"
)

func aoc1b() string {
	lines := readLines("input1.txt")

	sumSet := map[int]int{}

	sum := 0
loop:
	for {
		for _, line := range lines {
			sumSet[sum]++
			if sumSet[sum] > 1 {
				break loop
			}
			sum += readInt(line)
		}
	}
	return fmt.Sprint(sum)
}

func aoc1a() string {
	lines := readLines("input1.txt")
	sum := 0
	for _, line := range lines {
		sum += readInt(line)
	}
	return fmt.Sprint(sum)
}
