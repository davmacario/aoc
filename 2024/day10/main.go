package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	_ "strings"
)

type Point struct {
	x, y int
}

func (a Point) relativePos(b Point) Point {
	return Point{x: a.x - b.x, y: a.y - b.y}
}

func (a Point) add(b Point) Point {
	return Point{x: a.x + b.x, y: a.y + b.y}
}

func (a Point) insideBounds(w, h int) bool {
	if a.x < 0 || a.y < 0 || a.x >= w || a.y >= h {
		return false
	}
	return true
}

func (a Point) opposite() Point {
	return Point{x: -a.x, y: -a.y}
}

func (a Point) in(sl []Point) bool {
	for _, p := range sl {
		if p == a {
			return true
		}
	}
	return false
}

var Up = Point{x: 0, y: -1}
var Right = Point{x: 1, y: 0}
var Down = Point{x: 0, y: 1}
var Left = Point{x: -1, y: 0}
var NullDir = Point{x: 0, y: 0}

var Dirs = [4]Point{Up, Right, Down, Left}

// Recursively find all the numbers of trailheads
func countTrailheadScore1(topo []string, start Point, ind int, arr_dir Point, reached *[]Point) int {
	h := len(topo)
	w := len(topo[0])
	if !start.insideBounds(w, h) || string(topo[start.y][start.x]) != strconv.Itoa(ind) {
		return 0
	} else if string(topo[start.y][start.x]) == "9" {
		if !start.in(*reached) {
			*reached = append(*reached, start)
			return 1
		} else {
			return 0
		}
	}

	var out int
	for _, d := range Dirs {
		if d != arr_dir.opposite() || arr_dir == NullDir {
			out += countTrailheadScore1(topo, start.add(d), ind+1, d, reached)
		}
	}
	return out
}

func countTrailheadScore2(topo []string, start Point, ind int, arr_dir Point) int {
	h := len(topo)
	w := len(topo[0])
	if !start.insideBounds(h, w) || string(topo[start.y][start.x]) != strconv.Itoa(ind) {
		return 0
	} else if string(topo[start.y][start.x]) == "9" && strconv.Itoa(ind) == "9" {
		return 1
	}

	var out int
	for _, d := range Dirs {
        if d != arr_dir.opposite() || arr_dir == NullDir {
		    out += countTrailheadScore2(topo, start.add(d), ind+1, d)
        }
	}
	return out
}

/*
 * Steps:
 * - Iterate until a "0" is found
 * - Perform recursive DFS-like mechanism to find the number of trails
 */
func solve(topo []string) (int, int) {
	h := len(topo)
	w := len(topo[0])
	var out1 int
	var out2 int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if string(topo[i][j]) == "0" {
				reached := make([]Point, 0)
				out1 += countTrailheadScore1(topo, Point{x: j, y: i}, 0, Point{x: 0, y: 0}, &reached)
			}
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if string(topo[i][j]) == "0" {
                out2 += countTrailheadScore2(topo, Point{x: j, y: i}, 0, Point{x: 0, y: 0})
			}
		}
	}
	return out1, out2
}

func main() {
    input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Topographic map
	topo := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		topo = append(topo, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	res1, res2 := solve(topo)
	fmt.Println("Part 1:", res1)
	fmt.Println("Part 2:", res2)
}
