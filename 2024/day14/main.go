package main

import (
	"bufio"
	"day14/functions"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

// Used to move each second
func (p Pos) moveInDir(d Pos, dims ...int) Pos {
	if len(dims) == 2 {
		return Pos{x: functions.Mod((p.x + d.x), dims[0]), y: functions.Mod((p.y + d.y), dims[1])}
	}
	return Pos{x: p.x + d.x, y: p.y + d.y}
}

func (p Pos) sum(q Pos) Pos {
	return Pos{x: p.x + q.x, y: p.y + q.y}
}

// Euclidean distance between 2 points
func (p Pos) distance(q Pos) float64 {
	return float64(math.Sqrt(math.Pow(float64(p.x-q.x), 2.) + math.Pow(float64(p.y-q.y), 2.)))
}

func (p Pos) inBounds(w, h int) bool {
	if p.x < 0 || p.y < 0 || p.x >= w || p.y >= h {
		return false
	}
	return true
}

// Given a string of the type "X,Y", returns Pos{x: X, y: Y}
func getPosFromStr(in string) Pos {
	vals := strings.Split(in, ",")
	v1, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Fatal("Unable to convert '", vals[0], "' to integer")
	}
	v2, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Fatal("Unable to convert '", vals[1], "' to integer")
	}
	return Pos{x: v1, y: v2}
}

// Given the position and the area dimensions, get the quadrant.
// This function assumes w and h are odd numbers.
// Returns -1 if on the boundary lines
func (p Pos) getQuadrant(w, h int) int {
	if p.x < w/2 {
		if p.y < h/2 {
			return 0
		} else if p.y > h/2 {
			return 1
		}
	} else if p.x > w/2 {
		if p.y < h/2 {
			return 3
		} else if p.y > h/2 {
			return 2
		}
	}
	return -1
}

type Robot struct {
	start_pos, curr_pos, vel Pos
}

func printCurrLayout(robots []Robot, w, h int) {
	// initialize with '.'
	layout := make([]string, h)
	for i := 0; i < h; i++ {
		layout[i] = strings.Repeat(".", w)
	}
	for _, v := range robots {
		x := v.curr_pos.x
		y := v.curr_pos.y
		if layout[y][x] == '.' {
			// It becomes '1'
			layout[y] = functions.ReplaceAtIndex(layout[y], '1', x)
		} else {
			// Add 1
			old_count, _ := strconv.Atoi(string(layout[y][x]))
			new_count := old_count + 1
			layout[y] = functions.ReplaceAtIndex(layout[y], []rune(strconv.Itoa(new_count))[0], x)
		}
	}
	for i := 0; i < h; i++ {
		fmt.Println(layout[i])
	}
	fmt.Println("")
}

func nextSecond(robots []Robot, w, h int) {
	for i, r := range robots {
		// Move robot along direction of velocity
		robots[i].curr_pos = r.curr_pos.moveInDir(r.vel, w, h)
	}
}

func getSafetyFactor(robots []Robot, w, h int) (out int) {
	out = 1
	n_in_bounds := make(map[int]int)
	for _, r := range robots {
		n_in_bounds[r.curr_pos.getQuadrant(w, h)]++
	}
	for i := 0; i < 4; i++ {
		out *= n_in_bounds[i]
	}
	return out
}

// I.e., get mean distance from centroid
func getAggregationCoeff(robots []Robot) (out float64) {
	// Get centroid position (int coord. is fine)
	var centroid Pos // Init to 0, 0
	for _, r := range robots {
		centroid = centroid.sum(r.curr_pos)
	}
	centroid = Pos{x: centroid.x / len(robots), y: centroid.y / len(robots)}

	// Calculate mean centroid distance
	for _, r := range robots {
		out += r.curr_pos.distance(centroid)
	}
	return out / float64(len(robots))
}

func solve(robots []Robot, w, h int) (int, int) {
	fmt.Println("Start:")
	robots2 := make([]Robot, len(robots))
	copy(robots2, robots)
	printCurrLayout(robots, w, h)

	n_seconds := 100
	for i := 0; i < n_seconds; i++ {
		nextSecond(robots, w, h)
		// fmt.Println(robots[0])
	}
	fmt.Println("After", n_seconds, "seconds:")
	printCurrLayout(robots, w, h)

	// Part 2
	// Idea: use measure of "how close together" the points are
	var mcd, mcd_prev float64
	stop_cond := false
	var iter_ind int
	for !stop_cond {
		iter_ind++
		nextSecond(robots2, w, h)
		if iter_ind == 0 {
			mcd = getAggregationCoeff(robots2)
			mcd_prev = mcd
		} else {
			mcd_prev = mcd
			mcd = 0.5*mcd + 0.5*getAggregationCoeff(robots2)
		}
		fmt.Println("> 2:", iter_ind, "MCD:", mcd)
		// Expect mean centroid distance to decrease "notably"
		if mcd*1.2 < mcd_prev { // Tune hyperparameter
			fmt.Println("Possibly found tree! Iter", iter_ind)
			printCurrLayout(robots2, w, h)
			stop_cond = true
		}
	}

	return getSafetyFactor(robots, w, h), iter_ind
}

func main() {
	input_file := "./in.txt"
	var h, w int
	if input_file == "./in_small.txt" {
		h = 7
		w = 11
	} else {
		h = 103
		w = 101
	}
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	robots := make([]Robot, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			fields := strings.Fields(line)
			// Get start position
			start_pos_str := strings.Split(fields[0], "=")
			start_pos := getPosFromStr(start_pos_str[1])
			// Get velocity
			vel_str := strings.Split(fields[1], "=")
			vel := getPosFromStr(vel_str[1])
			robots = append(robots, Robot{start_pos: start_pos, curr_pos: start_pos, vel: vel})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res1, res2 := solve(robots, w, h)
	fmt.Println("Part 1:", res1)
	fmt.Println("Part 2:", res2)
}
