package functions

import (
	"day4/data_structures"
)

// Returns the number of accessible rolls in a grid
func CountAccessibleRolls(grid [][]string) int {
	h := len(grid)
	w := len(grid[0])

	res := 0
	for i := range h {
		for j := range w {
			currPoint := datastructures.Point{X: j, Y: i}
			currChar, err := datastructures.GetCharInPoint(grid, currPoint)
			if err != nil {
				// Should not be here!
				panic(err)
			}
			if currChar == "@" {
				// Count how many "free" spaces are around
				count := 0
				for _, d := range datastructures.DirsAll {
					neighbor := currPoint.MoveInDir(d)
					neighborChar, err := datastructures.GetCharInPoint(grid, neighbor)
					if err != nil {
						continue
					}
					if neighborChar == "@" {
						count++
					}
				}
				if count < 4 {
					res += 1
				}
			}
		}
	}

	return res
}

// Returns the number of accessible rolls in a grid and the updated grid after
// removing all accessible rolls
func CountAccessibleRollsAndRemove(grid [][]string) (int, [][]string) {
	h := len(grid)
	w := len(grid[0])

	upd := make([][]string, h)

	res := 0
	for i := range h {
		upd_row := make([]string, w)
		for j := range w {
			currPoint := datastructures.Point{X: j, Y: i}
			currChar, err := datastructures.GetCharInPoint(grid, currPoint)
			if err != nil {
				// Should not be here!
				panic(err)
			}
			if currChar == "@" {
				// Count how many "free" spaces are around
				count := 0
				for _, d := range datastructures.DirsAll {
					neighbor := currPoint.MoveInDir(d)
					neighborChar, err := datastructures.GetCharInPoint(grid, neighbor)
					if err != nil {
						continue
					}
					if neighborChar == "@" {
						count++
					}
				}
				if count < 4 {
					res += 1
					upd_row[j] = "."
					continue
				}
			}
			upd_row[j] = currChar
		}
		upd[i] = upd_row
	}

	return res, upd
}
