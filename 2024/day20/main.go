package main

import (
	"bufio"
	ds "day20/data_structures"
	f "day20/functions"
	"fmt"
	"log"
	"os"
	_ "strings"
)

func findEnd(track []string) ds.Point {
	i, j := f.FindInMatrix(track, 'E')
	return ds.MakePoint(j, i)
}

func findStart(track []string) ds.Point {
	i, j := f.FindInMatrix(track, 'S')
	return ds.MakePoint(j, i)
}

func initVisitedMap(h, w int) map[ds.Point]int {
	out := make(map[ds.Point]int)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			out[ds.MakePoint(j, i)] = h * w // Upper bound
		}
	}
	return out
}

// Assumes only 1 valid path between start and end
func getAllPtsDist(track []string, start, end ds.Point) map[ds.Point]int {
	h := len(track)
	w := len(track[0])
	out := initVisitedMap(h, w)
	currPt := start
	currDist := 0
	out[currPt] = currDist

	for currPt != end {
		// Look around, get neighbors
		for _, d := range ds.Dirs {
			nextPt := currPt.MoveInDir(d)
			if ds.GetCharInPoint(track, nextPt) != "#" && currDist+1 < out[nextPt] {
				currPt = nextPt
				currDist++
				out[nextPt] = currDist
				break
			}
		}
	}
	return out
}

// Returns the total number of cheats allowing to save >= 100 ps
func solve1(track []string) int {
	h := len(track)
	w := len(track[0])
	startPt := findStart(track)
	endPt := findEnd(track)

	// Move from end to start and get, for each point, the distance from the end
	// following the track
	distFromEnd := getAllPtsDist(track, endPt, startPt)

	// Move from start, look for shortcuts
	currPt := startPt
	currDistStart := 0
	mapDistStart := initVisitedMap(h, w)
	mapDistStart[startPt] = currDistStart // Overkill - just used to prevent going back
	var countShortcuts100 int
	for currPt != endPt {
		// Look around, get neighbors
		actualNext := currPt
		for _, d := range ds.Dirs {
			nextPt := currPt.MoveInDir(d)
			nextNext := nextPt.MoveInDir(d)
			if nextPt.InsideBounds(w, h) {
				nextCh := ds.GetCharInPoint(track, nextPt)
				if nextCh != "#" && currDistStart+1 < mapDistStart[nextPt] {
					actualNext = nextPt
					mapDistStart[nextPt] = currDistStart + 1
				} else if nextCh == "#" && nextNext.InsideBounds(w, h) && ds.GetCharInPoint(track, nextNext) != "#" {
					// Exists shortcut - calculate saving
					lenSaving := distFromEnd[currPt] - (distFromEnd[nextNext] + 2)
					if lenSaving >= 100 {
						countShortcuts100++
					}
				}
			}
		}
		if currPt == actualNext {
			panic("Did not move...")
		}
		currPt = actualNext
		currDistStart++
	}
	return countShortcuts100
}

func solve2(track []string) int {
	h := len(track)
	w := len(track[0])
	startPt := findStart(track)
	endPt := findEnd(track)

	// Move from end to start and get, for each point, the distance from the end
	// following the track
	distFromEnd := getAllPtsDist(track, endPt, startPt)

	// Move from start, look for shortcuts
	currPt := startPt
	currDistStart := 0
	mapDistStart := initVisitedMap(h, w)
	mapDistStart[startPt] = currDistStart
	var countShortcuts100 int
	for currPt != endPt {
		// Look around, get neighbors
		actualNext := currPt
		for _, d := range ds.Dirs {
			nextPt := currPt.MoveInDir(d)
			if nextPt.InsideBounds(w, h) {
				nextCh := ds.GetCharInPoint(track, nextPt)
				if nextCh != "#" && currDistStart+1 < mapDistStart[nextPt] {
					actualNext = nextPt
					mapDistStart[nextPt] = currDistStart + 1
				}
			}
		}
		for dst := 2; dst <= 20; dst++ {
			// Get points at distance dst from currPt
			pointsAtD := currPt.GetPointsAtDistBounds(dst, w, h)
			for _, pn := range pointsAtD {
				if ds.GetCharInPoint(track, pn) != "#" {
					// Calculate saving
					lenSaving := distFromEnd[currPt] - (distFromEnd[pn] + dst)
					if lenSaving >= 100 {
						countShortcuts100++
					}
				}
			}
		}
		currPt = actualNext
		currDistStart++
	}
	return countShortcuts100
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	racetrack := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		racetrack = append(racetrack, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ans1 := solve1(racetrack)
	fmt.Println("Part 1:", ans1)
	ans2 := solve2(racetrack)
	fmt.Println("Part 2:", ans2)
}
