package main

import (
	"fmt"
	"sort"
)

func aoc2b() string {
	lines := readLines("input2.txt")

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
	comm := append(bestPair.a[:diff], bestPair.a[diff+1:]...)

	return string(comm)
}

func aoc2a() string {
	lines := readLines("input2.txt")

	var isOK = func(set map[byte]int) (bool, bool) {
		var twos, threes int
		for c := range set {
			if set[c] == 2 {
				twos++
			}
			if set[c] == 3 {
				threes++
			}
		}
		return twos > 0, threes > 0
	}

	var twos, threes int
	for _, line := range lines {
		set := map[byte]int{}
		for i := range line {
			b := line[i]
			set[b]++
		}
		ok2, ok3 := isOK(set)
		if ok2 {
			twos++
		}
		if ok3 {
			threes++
		}
	}
	return fmt.Sprint(twos * threes)
}
