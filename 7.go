package main

import (
	"fmt"
	"sort"
	"strings"
)

// type Node struct {
// 	data  byte
// 	left  []Node
// 	right []Node
// }

func aoc7a(lines []string) string {

	// lines = []string{
	// 	"Step C must be finished before step A can begin.",
	// 	"Step C must be finished before step F can begin.",
	// 	"Step A must be finished before step B can begin.",
	// 	"Step A must be finished before step D can begin.",
	// 	"Step B must be finished before step E can begin.",
	// 	"Step D must be finished before step E can begin.",
	// 	"Step F must be finished before step E can begin.",
	// }

	// LVDIJFACEHZPGNRQWSBKUTMXOY
	// LVFDJIHECAZPRNGWQSUKBTMXOY
	// LDEVCIPFHGNQZJBKSUATRMWXY

	type Pair struct{ x, y byte }

	pairs := make([]Pair, len(lines))
	nodes := make(map[byte]struct{})

	for i, line := range lines {
		x, y := line[5], line[36]
		pairs[i] = Pair{x, y}
		nodes[x] = struct{}{}
		nodes[y] = struct{}{}
	}

	deps := map[byte][]byte{}
	revs := map[byte][]byte{}

	for _, p := range pairs {
		if _, ok := deps[p.y]; !ok {
			deps[p.y] = []byte{}
		}
		if _, ok := revs[p.x]; !ok {
			revs[p.x] = []byte{}
		}
		deps[p.y] = append(deps[p.y], p.x)
		revs[p.x] = append(revs[p.x], p.y)
	}

	for _, xs := range deps {
		sort.Slice(xs, func(i, j int) bool { return xs[i] < xs[j] })
	}
	for _, xs := range revs {
		sort.Slice(xs, func(i, j int) bool { return xs[i] < xs[j] })
	}

	var first, last byte
	for node := range nodes {
		if _, ok := deps[node]; !ok {
			first = node
		}
		if _, ok := revs[node]; !ok {
			last = node
		}
	}
	fmt.Printf("First, last = %c, %c\n", first, last)

	var reverse = func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	res := kahn(deps, revs, last)
	// res := postorder(deps, first)
	// res := bfs(deps, last)
	var sb strings.Builder
	sb.Write(res)
	fmt.Println("\n" + reverse(sb.String()))

	// fmt.Println()
	// for k := range deps {
	// 	fmt.Printf("%c:", k)
	// 	for _, dep := range deps[k] {
	// 		fmt.Printf(" %c", dep)
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()
	// for k := range revs {
	// 	fmt.Printf("%c:", k)
	// 	for _, dep := range revs[k] {
	// 		fmt.Printf(" %c", dep)
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	return "not done"
}

type ByteGraph map[byte][]byte
type ByteSet map[byte]struct{}

// https://en.wikipedia.org/wiki/Topological_sorting
func kahn(deps, revs ByteGraph, root byte) []byte {
	L := []byte{}
	s := make(chan byte, 300)

	s <- root

	for len(s) > 0 {
		n := <-s         // Remove a node n from s
		L = append(L, n) // Add n to tail of L

		// fmt.Println("n =", string(n))

		for len(deps[n]) > 0 {
			m := deps[n][len(deps[n])-1]
			deps[n] = deps[n][:len(deps[n])-1]

			fmt.Print("m = ", string(m)+". Incoming are [")
			for _, x := range revs[m] {
				fmt.Print(" " + string(x))
			}
			fmt.Println(" ]")

			for i := range revs[m] {
				if revs[m][i] == n {
					fmt.Println("         Removed", string(n))
					revs[m] = append(revs[m][:i], revs[m][i+1:]...)
					break
				}
			}

			// fmt.Print("m = ", string(m)+". Incoming to m are [")
			// for _, x := range revs[m] {
			// 	fmt.Print(" " + string(x))
			// }
			// fmt.Println(" ]")

			if len(revs[m]) == 0 {
				s <- m // Insert m into s
			}
		}

		// for i, m := range deps[n] {
		// 	fmt.Println(i, string(n), string(m), deps[n])
		// 	// Remove edge from graph
		// 	fmt.Println(string(deps[n][0]))
		// 	deps[n] = deps[n][1:]
		// 	// if i >= len(deps[n]) {
		// 	// 	deps[n] = deps[n][:i]
		// 	// } else {
		// 	// 	deps[n] = append(deps[n][:i], deps[n][i+1:]...)
		// 	// }
		// 	// If m has no incoming edges then
		// 	if _, ok := revs[m]; !ok {
		// 		s <- m // Insert m into s
		// 		fmt.Println("asdasd")
		// 	}
		// }
	}

	return L
}

func bfs(graph ByteGraph, root byte) []byte {
	r := []byte{}
	visited := ByteSet{}
	q := make(chan byte, len(graph))
	q <- root
	for len(q) > 0 {
		v := <-q
		fmt.Println(string(v))
		for _, w := range graph[v] {
			if _, in := visited[w]; !in {
				q <- w
				visited[w] = struct{}{}
			}
		}
	}
	return r
}

func postorder(graph ByteGraph, root byte) []byte {
	r := []byte{}
	visited := ByteSet{}
	rec(root, graph, visited, &r)
	return r
}

func rec(n byte, graph ByteGraph, visited ByteSet, r *[]byte) {
	visited[n] = struct{}{}
	for m := range graph {
		if _, in := visited[m]; !in {
			rec(m, graph, visited, r)
		}
	}
	*r = append(*r, n)

}

func dfs(nodes map[byte]struct{}, revs map[byte][]byte) []byte {
	r := []byte{}
	marked := map[byte]struct{}{}
	for n := range nodes {
		// ??????
		// if len(unmarked) == 0 {
		// 	break
		// }
		visit(n, marked, revs, &r)
	}
	return r
}

func visit(n byte, marked map[byte]struct{}, revs map[byte][]byte, r *[]byte) {
	if _, ok := marked[n]; ok {
		return
	}
	for _, m := range revs[n] {
		visit(m, marked, revs, r)
	}
	marked[n] = struct{}{}
	*r = append(*r, n)
}

func stuff(current, first, last byte, level int, deps map[byte][]byte, r *[]byte) {
	xs := deps[current]
	if current == first {
		return
	}
	*r = append(*r, current)
	for i := len(xs) - 1; i >= 0; i-- {
		stuff(xs[i], first, last, level+1, deps, r)
	}
}

func buildResult(current, last byte, revs map[byte][]byte, result *[]byte) {

	if current == last {
		return
	}
	*result = append(*result, current)
	for _, other := range revs[current] {
		buildResult(other, last, revs, result)
	}
}

// 5, 36

func aoc7b(lines []string) string {
	return ""
}
