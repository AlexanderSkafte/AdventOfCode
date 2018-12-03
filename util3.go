package main

import (
	"fmt"
)

// Claim is used for problem 3a.
type Claim struct {
	ID   int
	X, Y int
	W, H int
}

func parseClaim(str string) Claim {
	var id, x, y, w, h int
	n, err := fmt.Sscanf(str, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
	if n != 5 {
		panic("could not read 5 ints")
	}
	if err != nil {
		panic(err)
	}
	return Claim{id, x, y, w, h}
}

func (c Claim) String() string {
	return fmt.Sprintf("#%d @ %d,%d: %dx%d", c.ID, c.X, c.Y, c.W, c.H)
}
