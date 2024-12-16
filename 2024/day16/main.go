package main

import (
	"bufio"
	"day16/functions"
	"fmt"
	"log"
	"os"
	_ "slices"
	_ "strings"
)

type Point struct {
	x, y int
}

func (p Point) moveInDir(d Dir) Point {
	return Point{x: p.x + d.x, y: p.y + d.y}
}

type Dir struct {
	x, y int
}

func (d Dir) Opposite() Dir {
	return Dir{x: -d.x, y: -d.y}
}

var up = Dir{x: 0, y: -1}
var right = Dir{x: 1, y: 0}
var down = Dir{x: 0, y: 1}
var left = Dir{x: -1, y: 0}

var dirs = []Dir{left, up, right, down}

func findStart(maze []string) (out Point) {
	i, j := functions.FindInMat(maze, 'S')
	return Point{x: j, y: i}
}

func findEnd(maze []string) (out Point) {
	i, j := functions.FindInMat(maze, 'E')
	return Point{x: j, y: i}
}

type Queue struct {
	s []Point
	curr_score []int // matches with q, score up until that point
	arr_dir []int // Arrival direction
}

func (s *Queue) Pop() (p Point, sc int, d int, ok bool) {
	if s.Length() < 1 {
		return p, sc, d, false
	}
	p = s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	sc = s.curr_score[len(s.curr_score)-1]
	s.curr_score = s.curr_score[:len(s.curr_score)-1]
	d = s.arr_dir[len(s.arr_dir)-1]
	s.arr_dir = s.arr_dir[:len(s.arr_dir)-1]
	if len(s.s) != len(s.arr_dir) {
		log.Fatal("Lengths of the slices don't match")
	}
	return p, sc, d, true
}

func (s *Queue) Push(p Point, sc int, d int) bool {
	s.s = append([]Point{p}, s.s...)
	s.curr_score = append([]int{sc}, s.curr_score ...)
	s.arr_dir = append([]int{d}, s.arr_dir...)
	return true
}

func (s Queue) Length() int {
	if len(s.s) != len(s.arr_dir) {
		log.Fatal("Lengths of the slices don't match")
	}
	return len(s.s)
}

func pointInSlice(p Point, l []Point) int {
	for i, p1 := range l {
		if p1 == p {
			return i
		}
	}
	return -1
}

func replaceCharInMaze(maze []string, p Point, c rune) {
	maze[p.y] = functions.ReplaceCharInString(maze[p.y], p.x, c)
}

// Backtrack - from end to start, only move towards points that are greatest smaller nums
func getPath(maze []string, pointsMap map[Point]int) (sortedPts []Point) {
	startPos := findStart(maze)
	endPos := findEnd(maze)

	currPt := endPos
	currCost := pointsMap[endPos]
	out := make([]Point, 0)
	out = append([]Point{endPos}, out...)
	for currPt != startPos {
		// fmt.Println(">", currPt)
		min_cost_around := -1
		nextPtCandidate := currPt
		for i := range len(dirs) {
			nextPt := currPt.moveInDir(dirs[i])
			nextCost := pointsMap[nextPt]
			// fmt.Println("  >", i, nextPt, currCost)
			if nextCost != -1 && nextCost < currCost {
				if min_cost_around == -1 {
					min_cost_around = nextCost
					nextPtCandidate = nextPt
				} else if nextCost < currCost && nextCost > min_cost_around {
					min_cost_around = nextCost
					nextPtCandidate = nextPt
				}
			}
		}
		if currPt == nextPtCandidate {
			log.Fatal("Didn't update point")
		}
		out = append([]Point{currPt}, out...)
		currPt = nextPtCandidate
		currCost = pointsMap[currPt]
		// log.Fatal("")
	}
	return out
}

// func getAllPaths(maze []string, pointsMap map[Point]bool, minCostNeighb map[Point][]Point) int {
// 	startPos := findStart(maze)
// 	endPos := findEnd(maze)
//
// 	currPt := endPos
// 	currCost := pointsMap[endPos]
// 	out := make([]Point, 0)
// 	out = append([]Point{endPos}, out...)
// }

func printVisited(maze []string, seenPts []Point) {
	mazeBak := make([]string, len(maze))
	copy(mazeBak, maze)
	for _, p := range seenPts {
		replaceCharInMaze(mazeBak, p, 'X')
	}

	for _, l := range mazeBak {
		fmt.Println(l)
	}
}

func initMapMinScore(maze []string) map[Point]int {
	out := make(map[Point]int)
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[0]); j++ {
			out[Point{x: j, y: i}] = -1
		}
	}
	return out
}

// Idea: use queue to flood fill the maze (bfs) as we want to find the shortest
// path.
// Only keep lowest price per cell, add cell to the stack iff lower price has
// been found
func part1(maze []string) int {
	startPos := findStart(maze)
	endPos := findEnd(maze)
	travelQueue := Queue{}
	travelQueue.Push(startPos, 0, 0)
	visitedMinScore := initMapMinScore(maze)
	visitedMinScore[startPos] = 0
    // For each point, contains the neighbors from which it can be reached with minimum score
    allMinScore := make(map[Point][]Point)
	for travelQueue.Length() > 0 {
		currPos, givenScore, arrDir, ok := travelQueue.Pop()
		currScore := visitedMinScore[currPos]
		// fmt.Println(currPos, arrDir, currScore)

		// printVisited(maze, []Point{currPos})
		// fmt.Println("")
		if ok && currPos != endPos && currScore == givenScore {
            // Ignore if popped item has higher score than what contained in the map
			for i := 0; i < len(dirs); i++ {
				nextPos := currPos.moveInDir(dirs[i])
				newScore := currScore
                incr := 0
				if i == functions.Mod(arrDir+2, len(dirs)) {
					// Going straight
					incr = 1
				} else if i == arrDir {
					// Turning 180 deg - only allowed at 1st iter
					// In non-1st iter, cost for the following position would increase
					incr = 2001
				} else {
					// Turning 90 deg
					incr = 1001
				}
                newScore += incr
				if maze[nextPos.y][nextPos.x] != '#' && (visitedMinScore[nextPos] == -1 || newScore <= visitedMinScore[nextPos]) {
					// fmt.Println("Moving to", nextPos, "with new lowest of", newScore, "(vs. old", visitedMinScore[nextPos], ")")
					// if nextPos == endPos {
					// 	fmt.Println("Found better sol:", newScore)
					// }
					// tgtpt := Point{x: 6, y: 7}
					// if nextPos == tgtpt {
					// 	fmt.Println("From", currPos, "to", nextPos, "moving in direction", i, "coming from direction", arrDir, "New cost", newScore, "increased by", incr)
					// }
                    if newScore < visitedMinScore[nextPos] {
					    visitedMinScore[nextPos] = newScore
                        allMinScore[nextPos] = []Point{currPos}
                    } else {
                        allMinScore[nextPos] = append(allMinScore[nextPos], currPos)
                    }
					travelQueue.Push(nextPos, newScore, functions.Mod(i+2, len(dirs)))
				}
			}
		} else if currPos != endPos && currScore == givenScore {
			log.Fatal("Shouldn't be here - popped from empty stack")
		}
	}

    count_on_path := getAllPaths(maze, visitedMinScore, allMinScore)
	traversed := getPath(maze, visitedMinScore)
	// fmt.Println("Path:")
	// for _, p := range traversed {
	// 	fmt.Println(">", p, "- cost:", visitedMinScore[p])
	// }
	printVisited(maze, traversed)

	return visitedMinScore[endPos]
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	maze := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ans1 := part1(maze)
	fmt.Println("Part 1:", ans1)
}
