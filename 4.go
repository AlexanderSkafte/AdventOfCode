package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

// https://adventofcode.com/2018/day/4

var exampleLines4a = []string{
	"[1518-11-01 00:00] Guard #10 begins shift",
	"[1518-11-01 00:05] falls asleep",
	"[1518-11-01 00:25] wakes up",
	"[1518-11-01 00:30] falls asleep",
	"[1518-11-01 00:55] wakes up",
	"[1518-11-01 23:58] Guard #99 begins shift",
	"[1518-11-02 00:40] falls asleep",
	"[1518-11-02 00:50] wakes up",
	"[1518-11-03 00:05] Guard #10 begins shift",
	"[1518-11-03 00:24] falls asleep",
	"[1518-11-03 00:29] wakes up",
	"[1518-11-04 00:02] Guard #99 begins shift",
	"[1518-11-04 00:36] falls asleep",
	"[1518-11-04 00:46] wakes up",
	"[1518-11-05 00:03] Guard #99 begins shift",
	"[1518-11-05 00:45] falls asleep",
	"[1518-11-05 00:55] wakes up",
}

var dateFormat = "[%04d-%02d-%02d %02d:%02d]"

// Entry is used for the problems of day 4.
type Entry struct {
	date time.Time
	kind int // id if >= 0, fall asleep if -1, wake up if -2
}

// You can ignore this, only used for debugging
func (entry Entry) String() string {
	date := fmt.Sprintf(dateFormat,
		entry.date.Year(),
		int(entry.date.Month()),
		entry.date.Day(),
		entry.date.Hour(),
		entry.date.Minute(),
	)
	var rest string
	if entry.kind >= 0 {
		rest = "#" + fmt.Sprint(entry.kind)
	} else if entry.kind == -1 {
		rest = "sleep"
	} else if entry.kind == -2 {
		rest = "wake"
	} else {
		panic("invalid kind")
	}
	return date + " " + rest
}

func parseKind(rest string) int {
	rest = strings.TrimSpace(rest)
	tokens := strings.Split(rest, " ")
	switch tokens[0] {
	case "Guard":
		var id int
		fmt.Sscanf(tokens[1][1:], "%d", &id)
		return id
	case "falls":
		return -1
	case "wakes":
		return -2
	default:
		fmt.Println("unhandled token " + tokens[0])
		panic("!")
	}
}

func aoc4a(lines []string) string {

	entries := make([]Entry, len(lines))

	var y, m, d, hr, mn int
	for i := range lines {
		fmt.Sscanf(lines[i], dateFormat, &y, &m, &d, &hr, &mn)
		entries[i] = Entry{
			time.Date(y, time.Month(m), d, hr, mn, 0, 0, time.UTC),
			parseKind(lines[i][18:]),
		}
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].date.Before(entries[j].date)
	})

	type (
		GuardID = int
		Minute  = int
	)

	// GuardID -> (Minute -> Days slept during Minute)
	data := map[GuardID]map[Minute]int{}

	// Read GuardID and all the guard's sleep wake times.
outer:
	for i := 0; i < len(entries); {
		id := entries[i].kind

		if _, ok := data[id]; !ok {
			data[id] = map[Minute]int{}
		}

		i++

		for {
			ms := entries[i].date.Minute()   // Minute guard fell asleep
			mw := entries[i+1].date.Minute() // Minute guard woke up

			for m := ms; m < mw; m++ {
				data[id][m]++
			}

			i += 2

			if i >= len(entries) {
				break outer
			} else if entries[i].kind >= 0 {
				break
			}
		}
	}

	// Map GuardID -> Total minutes sleeping
	totalFor := map[GuardID]int{}
	for id := range data {
		for _, daysSleptForMinute := range data[id] {
			totalFor[id] += daysSleptForMinute
		}
	}

	// Find which guard sleeps the most (highest total)
	laziest := -1
	highestTotal := -1
	for id, mins := range totalFor {
		if mins > highestTotal {
			highestTotal = mins
			laziest = id
		}
	}

	// Find the minute during which the laziest guard slept the most
	highestN := -1
	var best Minute
	for minute, n := range data[laziest] {
		if n > highestN {
			highestN = n
			best = minute
		}
	}

	return fmt.Sprint(laziest * best)
}

func aoc4b(lines []string) string {
	return "not done"
}

// Ignore this
func writeToFile(entries []Entry, path string) {
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	for _, entry := range entries {
		fmt.Fprintf(file, entry.String()+"\n")
	}
}
