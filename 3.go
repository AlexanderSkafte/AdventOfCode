package main

import (
	"fmt"
)

func aoc3a(lines []string) string {

	claims := make([]Claim, len(lines))
	for i := range lines {
		claims[i] = parseClaim(lines[i])
	}

	type XY struct{ x, y int }
	nbrClaims := map[XY]int{} // Number of claims on a position (x,y)

	for _, c := range claims {
		for i := c.X; i < c.X+c.W; i++ {
			for j := c.Y; j < c.Y+c.H; j++ {
				nbrClaims[XY{i, j}]++
			}
		}
	}

	count := 0
	for xy := range nbrClaims {
		if nbrClaims[xy] >= 2 {
			count++
		}
	}
	return fmt.Sprint(count)
}

func aoc3b(lines []string) string {

	claims := make([]Claim, len(lines))
	for i := range lines {
		claims[i] = parseClaim(lines[i])
	}

	type XY struct{ x, y int }
	nbrClaims := map[XY]int{} // Number of claims on a position (x,y)

	for _, c := range claims {
		for i := c.X; i < c.X+c.W; i++ {
			for j := c.Y; j < c.Y+c.H; j++ {
				nbrClaims[XY{i, j}]++
			}
		}
	}

	for _, c := range claims {
		this := true
		for i := c.X; i < c.X+c.W; i++ {
			for j := c.Y; j < c.Y+c.H; j++ {
				if nbrClaims[XY{i, j}] != 1 {
					this = false
				}
			}
		}
		if this {
			return fmt.Sprint(c.ID)
		}
	}
	return "No solution."
}
