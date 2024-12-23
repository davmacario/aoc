package main

import (
	"bufio"
	ds "day23/data_structures"
	"day23/utils"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Me: "This one I'll solve using pointers"

// Given the current element (key of connections), iterate over every pair of
// elements in the associated slice. If the elements in a pair are connected,
// and no element in the pair has been seen already, then add 1 to the count of
// loops.
// Returns the total number of looping triplets for `currElem`
func findLoopTriplet(currElem string, connections map[string][]string, seen map[string]bool) (out int) {
	currArray := connections[currElem]
	for i := 0; i < len(currArray); i++ {
		for j := i; j < len(currArray); j++ {
			if !seen[currArray[i]] && !seen[currArray[j]] && utils.GetIndex(currArray[j], connections[currArray[i]]) != -1 {
				out++
			}
		}
	}
	seen[currElem] = true
	return out
}

// Return the number of looping triplets with 't'
func solve1(connections map[string][]string) int {
	seen := make(map[string]bool) // If set to true, triplets with that
	out := 0
	for k := range connections {
		if k[0] == 't' {
			out += findLoopTriplet(k, connections, seen)
		}
	}
	return out
}

// Bron-Kerbosch algorithm to calculate the maximum complete subgraph
//
// Args:
// - R: current subgraph (complete)
// - P: all nodes that may become part of the subgraph
// - X: all nodes that have been discarded
// - N: map associating each node to all of its neighbors (see input)
func BronKerbosch(R, P, X *ds.Set[string], N map[string]*ds.Set[string]) []*ds.Set[string] {
	if P.Length() == 0 && X.Length() == 0 {
		// R is a maximal clique
		return []*ds.Set[string]{R}
	}

	out := make([]*ds.Set[string], 0)
	for v := range P.S {
		out = append(out, BronKerbosch(ds.SetUnion(R, ds.NewSet([]string{v})), ds.SetIntersect(P, N[v]), ds.SetIntersect(X, N[v]), N)...)

		P.Remove(v)
		X.Add(v)
	}
	return out
}

// Solution for part 2: largest set of computers all connected between each
// other.
// Bron-Kerbosch algorithm
func solve2(connections map[string][]string) string {
	var out string
	// Make connections a map string -> Set (can intersect and unite)
	// In other words: for each key, connSet[k] is N(k) (neighbors)
	connSet := make(map[string]*ds.Set[string])
	P := ds.NewSet([]string{})
	for k, v := range connections {
		connSet[k] = ds.NewSet(v)
		P.Add(k)
	}
	R := ds.NewSet([]string{})
	X := ds.NewSet([]string{})

	allCliques := BronKerbosch(R, P, X, connSet)

	// Find maximum clique
	maximumClique := allCliques[0]
	maxLen := maximumClique.Length()
	for _, k := range allCliques[1:] {
		if k.Length() > maxLen {
			maximumClique = k
			maxLen = k.Length()
		}
	}
	// Get computer names slice and sort alphabetically
	endpoints := make([]string, maxLen)
	count := 0
	for k := range maximumClique.S {
		endpoints[count] = k
		count++
	}

	sort.Strings(endpoints)
	out += endpoints[0]
	for _, e := range endpoints[1:] {
		out += ","
		out += e
	}
	return out
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	conn := make(map[string][]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "-")
		for i, fi := range fields {
			ind_next := utils.Mod(i+1, 2)
			conn[fi] = append(conn[fi], fields[ind_next])
		}
	}
	// for k, v := range conn {
	// 	fmt.Println(k, "->", v)
	// }
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ans1 := solve1(conn)
	fmt.Println("Part 1:", ans1)

	ans2 := solve2(conn)
	fmt.Println("Part 2:", ans2)
}
