package main

import (
	"bufio"
	"day18/functions"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	s []Point
}

func (s *Queue) Pop() (p Point, ok bool) {
	if s.Length() < 1 {
		return p, false
	}
	p = s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return p, true
}

func (s *Queue) Push(p Point) bool {
	s.s = append([]Point{p}, s.s...)
	return true
}

func (s Queue) Length() int {
	return len(s.s)
}

type Point struct {
	x, y int
}

func (p Point) moveInDir(d Dir) Point {
	return Point{x: p.x + d.x, y: p.y + d.y}
}

func (p Point) inBounds(w, h int) bool {
	if p.x < 0 || p.y < 0 || p.x >= w || p.y >= h {
		return false
	}
	return true
}

type Dir struct {
	x, y int
}

var up = Dir{x: 0, y: -1}
var right = Dir{x: 1, y: 0}
var down = Dir{x: 0, y: 1}
var left = Dir{x: -1, y: 0}
var dirs = []Dir{left, up, right, down}

func initMinDistMap(w, h int) map[Point]int {
	out := make(map[Point]int)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			out[Point{x: j, y: i}] = w * h // Upper bound: n. cells
		}
	}
	return out
}

func initMap(w, h int) []string {
	out := make([]string, 0)
	for i := 0; i < h; i++ {
		ln := ""
		for j := 0; j < w; j++ {
			ln += "."
		}
		out = append(out, ln)
	}
	return out
}

func addByteToMap(mem []string, bp Point) {
	mem[bp.y] = functions.ReplaceCharInString(mem[bp.y], bp.x, '#')
}

func getMemCont(mem []string, p Point) string {
	return string(mem[p.y][p.x])
}

// Returns path length
func shortestPath(mem []string, startPt, endPt Point) int {
	h := len(mem)
	w := len(mem[0])
	minDistPoint := initMinDistMap(w, h)
	minDistPoint[startPt] = 0
	visitQueue := Queue{s: []Point{startPt}}
	for visitQueue.Length() > 0 {
		currPt, _ := visitQueue.Pop()
		currScore := minDistPoint[currPt]
		for _, d := range dirs {
			nextPt := currPt.moveInDir(d)
			if nextPt.inBounds(w, h) && getMemCont(mem, nextPt) != "#" && currScore+1 < minDistPoint[nextPt] {
				minDistPoint[nextPt] = currScore + 1
				visitQueue.Push(nextPt)
			}
		}
	}
	return minDistPoint[endPt]
}


func solve(bytesList []Point, w, h, cons int) int {
	mem := initMap(w, h)

	for i, b := range bytesList {
		if i >= cons {
			break
		}
		addByteToMap(mem, b)
	}

	// Find shortest path
	minLenPt := shortestPath(mem, Point{x: 0, y: 0}, Point{x: w - 1, y: h - 1})
	if minLenPt == w*h {
		// Cannot reach point
		fmt.Println("Unable to reach destination")
	}
	return minLenPt
}

func main() {
	input_file := "./in.txt"
	var w, h, cons int
	if input_file == "./in_small.txt" {
		w = 7
        cons = 12
	} else {
		w = 71
        cons = 1024
	}
	h = w
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bytearr := make([]Point, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// Get array of fields (str)
		xy := strings.Split(line, ",")
		x_int, _ := strconv.Atoi(xy[0])
		y_int, _ := strconv.Atoi(xy[1])
		bytearr = append(bytearr, Point{x: x_int, y: y_int})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res1 := solve(bytearr, w, h, cons)
	fmt.Println("Part 1:", res1)

    res2 := 0
    ind := cons
    for res2 < w*h {
        ind++
        res2 = solve(bytearr, w, h, ind)
    }
    fmt.Println("Part 2:", bytearr[ind - 1])
}
