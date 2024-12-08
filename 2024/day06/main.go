package main

import (
	"bufio"
	"day06/functions"
	"fmt"
	"log"
	"os"
	// "slices"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Dir struct {
	x, y int
}

var up = Dir{x: 0, y: -1}
var right = Dir{x: 1, y: 0}
var down = Dir{x: 0, y: 1}

var left = Dir{x: -1, y: 0}
var dirs = [4]Dir{up, right, down, left}
var dirChar = [4]string{"U", "R", "D", "L"}

type Pos struct {
	x, y int
}

// Returns true if, by moving along the specified direction, we will encounter
// an obstacle
func encounterObstacle(lab []string, start []int, dirInd int) bool {
	h := len(lab)
	w := len(lab[0])
	currPos := start
	for functions.InsideOfBounds(currPos, h, w) {
		if string(lab[currPos[0]][currPos[1]]) == "#" {
			return true
		}
		currPos = []int{currPos[0] + dirs[dirInd].y, currPos[1] + dirs[dirInd].x}
	}
	return false
}

// Returns true if, traversing the labyrinth from start (with specific dir),
// we end up in a position we already visited and with the same orientation
func findLoop(labyrinth []string, start []int, startDirInd int) bool {
	currDirInd := startDirInd
    currPos := Pos{x: start[1], y: start[0]}
	h := len(labyrinth)
	w := len(labyrinth[0])
    // Map each position to the directions they have been traversed through (ind)
    posDirMap := make(map[Pos][]int)
	for {
        dirsCurrPos, ok := posDirMap[currPos]
        if ok {
            // Been here before
            if functions.IntInSlice(currDirInd, dirsCurrPos) {
                return true
            } else {
                posDirMap[currPos] = append(posDirMap[currPos], currDirInd)
            }
        } else {
            // Never been here
            posDirMap[currPos] = []int{currDirInd}
        }
        nextPos := Pos{y: currPos.y + dirs[currDirInd].y, x: currPos.x + dirs[currDirInd].x}
		nextPosValid := false
		for !nextPosValid {
			if functions.InsideOfBounds([]int{nextPos.x, nextPos.y}, h, w) {
				if string(labyrinth[nextPos.y][nextPos.x]) == "#" || string(labyrinth[nextPos.y][nextPos.x]) == "O" {
					// Change dir
					currDirInd = (currDirInd + 1) % len(dirs)
					// Calculate new next pos
                    nextPos = Pos{y: currPos.y + dirs[currDirInd].y, x: currPos.x + dirs[currDirInd].x}
				} else {
					// Next pos is valid
					currPos = nextPos
					nextPosValid = true
				}
			} else {
				// Out of bounds
				return false
			}
		}
	}
}

// Count distinct positions
// Mark visited positions with "X"
func solveMaze(labyrinth []string, start []int) (int, int) {
	var coundIndPos int
	var count2 int
	inLastPos := false
	currDirInd := 0 // Index of `dirs`
	currPos := start
	h := len(labyrinth)
	w := len(labyrinth[0])
	labyrinth[start[0]] = functions.ReplaceInStr(labyrinth[start[0]], ".", start[1])
	labBak := make([]string, len(labyrinth))
	_ = copy(labBak, labyrinth)
	indObstacles := make([][]int, 0) // Track the positions of the O's
	for !inLastPos {
		// Mark current cell
		if string(labyrinth[currPos[0]][currPos[1]]) != "X" {
			coundIndPos++
			currLine := labyrinth[currPos[0]]
			labyrinth[currPos[0]] = functions.ReplaceInStr(currLine, "X", currPos[1])
		}

		// Calculate next pos
		nextPos := []int{currPos[0] + dirs[currDirInd].y, currPos[1] + dirs[currDirInd].x}
		nextPosValid := false
		for !nextPosValid {
			if functions.InsideOfBounds(nextPos, h, w) {
				if string(labyrinth[nextPos[0]][nextPos[1]]) == "#" {
					// Change dir
					currDirInd = (currDirInd + 1) % len(dirs)
					// Calculate new next pos
					nextPos = []int{currPos[0] + dirs[currDirInd].y, currPos[1] + dirs[currDirInd].x}
				} else {
					// Next pos is valid
					nextPosValid = true
				}
			} else {
				inLastPos = true
				nextPosValid = true
			}
		}
		// Look for possible obstacles
		// Idea: Assume we are going to place an obstacle in the next position
		// Then, if going down that route, do we come back where we are (and do we hit the obstacle?)
		lab_O := make([]string, len(labBak))
		_ = copy(lab_O, labBak) // Fresh copy of initial maze
		nextDirInd := (currDirInd + 1) % len(dirs)
		if !inLastPos && encounterObstacle(labBak, currPos, nextDirInd) {
			// Add obstacle in maze copy
			lab_O[nextPos[0]] = functions.ReplaceInStr(lab_O[nextPos[0]], "O", nextPos[1])
			// fmt.Println("")
			// functions.PrintMatrix(lab_O)
			if !functions.ArrInMat(nextPos, indObstacles) {
				if findLoop(lab_O, start, 0) {
					indObstacles = append(indObstacles, nextPos)
                    count2 = len(indObstacles)
				}
			}
		}
		currPos = nextPos[:]
	}
	return coundIndPos, count2
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	checkErr(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var labyrinth []string
	var startPos []int
	for scanner.Scan() {
		line := scanner.Text()
		pos, err := functions.CharInStr("^", line)
		checkErr(err)
		if pos > -1 {
			startPos = []int{len(labyrinth), pos}
		}
		labyrinth = append(labyrinth, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if startPos == nil {
		log.Fatal("Starting position not found")
	}

	ans1, ans2 := solveMaze(labyrinth, startPos)
	functions.PrintMatrix(labyrinth)
	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
