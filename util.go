package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) []string {
	lines := []string{}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func readInt(line string) int {
	var number int
	n, err := fmt.Sscanf(line, "%d", &number)
	if n != 1 {
		panic("could not read 1 integer from line")
	}
	if err != nil {
		panic(err)
	}
	return number
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
