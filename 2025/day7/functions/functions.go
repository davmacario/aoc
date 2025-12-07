package functions

import (
	ds "day7/data_structures"
	"log"
)

const (
	Start    = "S"
	Empty    = "."
	Splitter = "^"
	Beam     = "|"
)

func FindStart(matrix [][]string) ds.Point {
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == Start {
				return ds.Point{X: j, Y: i}
			}
		}
	}
	log.Fatal("No start found")
	return ds.Point{}
}

// Recursive function for following beams using DFS.
//
// Steps:
//  1. Increase Y by 1 until point is not `Empty`
//  2. If out of bounds, we reached the end. Return 0
//  3. If in bounds, cases are either `Beam` (i.e., already seen), so return,
//     or `Splitter`, i.e., return (1 + recur left + recur right)
//
// NOTE: assuming no splitters can be right next to each other
func FollowBeamsRecur(currPoint ds.Point, matrix [][]string) int {
	for true {
		currChar, err := ds.GetCharInPoint(matrix, currPoint)
		if err != nil {
			// Out of bounds
			return 0
		}
		if currChar == Beam {
			// Already been here, counted somewhere else
			return 0
		}

		if currChar == Splitter {
			break
		}

		ds.SetCharInPoint(matrix, currPoint, Beam)
		currPoint = currPoint.MoveInDir(ds.Down)
	}

	return 1 + FollowBeamsRecur(currPoint.MoveInDir(ds.Left), matrix) + FollowBeamsRecur(currPoint.MoveInDir(ds.Right), matrix)
}

// Solve part 1, returning the number of "used" splitters
func FollowBeams(matrix [][]string) int {
	startPoint := FindStart(matrix)

	return FollowBeamsRecur(startPoint.MoveInDir(ds.Down), matrix)
}

// Part 2
func FollowQuantumBeamsRecur(currPoint ds.Point, matrix [][]string, cache map[ds.Point]int) int {
	for true {
		currChar, err := ds.GetCharInPoint(matrix, currPoint)
		if err != nil {
			// Out of bounds
			break
		}
		if currChar == Splitter {
			val, ok := cache[currPoint]
			if ok {
				return val
			}

			newVal := FollowQuantumBeamsRecur(currPoint.MoveInDir(ds.Left), matrix, cache) +
				FollowQuantumBeamsRecur(currPoint.MoveInDir(ds.Right), matrix, cache)

			cache[currPoint] = newVal
			return newVal
		}
		currPoint = currPoint.MoveInDir(ds.Down)
	}

	return 1
}

func FollowQuantumBeams(matrix [][]string) int {
	startPoint := FindStart(matrix)
	// Cache mapping, for each splitter, the number of possible paths to end
	cache := make(map[ds.Point]int)
	return FollowQuantumBeamsRecur(startPoint.MoveInDir(ds.Down), matrix, cache)
}
