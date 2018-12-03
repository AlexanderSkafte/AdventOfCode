package main

import (
	"fmt"
	"sort"
)

func aoc2a(lines []string) string {
	var twos, threes int

	for _, line := range lines {
		set := map[byte]int{}
		for i := range line {
			set[line[i]]++
		}
		var n2, n3 int
		for c := range set {
			if set[c] == 2 {
				n2++
			}
			if set[c] == 3 {
				n3++
			}
		}
		if n2 > 0 {
			twos++
		}
		if n3 > 0 {
			threes++
		}
	}
	return fmt.Sprint(twos * threes)
}

func aoc2b(lines []string) string {
	type Pair struct{ a, b []byte }

	var findDiffering = func(pair Pair) []int {
		differing := []int{}
		for i := 0; i < len(pair.a); i++ {
			if pair.a[i]-pair.b[i] != 0 {
				differing = append(differing, i)
			}
		}
		return differing
	}

	pairs := []Pair{}
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			pair := Pair{
				[]byte(lines[i]),
				[]byte(lines[j]),
			}
			pairs = append(pairs, pair)
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return len(findDiffering(pairs[i])) < len(findDiffering(pairs[j]))
	})

	bestPair := pairs[0]
	diff := findDiffering(bestPair)[0]
	common := append(bestPair.a[:diff], bestPair.a[diff+1:]...)

	return string(common)
}
