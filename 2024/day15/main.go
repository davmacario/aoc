package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	_ "strings"
)

type Point struct {
	x, y int
}

func (p Point) getStateCh(state []string) string {
	return string(state[p.y][p.x])
}

func (p Point) setStateCh(state []string, newCh string) {
	if len(newCh) > 1 {
		log.Fatal("Replacement string must be of length 1")
	}
	state[p.y] = state[p.y][:p.x] + newCh + state[p.y][p.x+1:]
}

func (p Point) moveInDir(d Dir) Point {
	return Point{x: p.x + d.x, y: p.y + d.y}
}

func (p Point) getGPSCoord() int {
	return 100*p.y + p.x
}

type Dir struct {
	x, y int
}

func (d Dir) opposite() Dir {
	return Dir{x: -d.x, y: -d.y}
}

var dirMap = make(map[rune]Dir)

func init() {
	dirMap['^'] = Dir{x: 0, y: -1}
	dirMap['>'] = Dir{x: 1, y: 0}
	dirMap['v'] = Dir{x: 0, y: 1}
	dirMap['<'] = Dir{x: -1, y: 0}
}

// -----

func swapChars(state []string, p1 Point, p2 Point) {
	c1 := p1.getStateCh(state)
	c2 := p2.getStateCh(state)
	p2.setStateCh(state, c1)
	p1.setStateCh(state, c2)
}

func recursiveExplore(state []string, pos Point, movement Dir) bool {
	var shouldISwap bool
	curr_char := pos.getStateCh(state)
	if curr_char == "#" {
		shouldISwap = false
	} else if curr_char == "." {
		// Swap prev
		prev := pos.moveInDir(movement.opposite())
		if prev.getStateCh(state) != "#" {
			swapChars(state, pos, prev)
		}
		shouldISwap = true
	} else {
		// Recur
		shouldISwap = recursiveExplore(state, pos.moveInDir(movement), movement)
		if shouldISwap {
			prev := pos.moveInDir(movement.opposite())
			if curr_char != "@" {
				swapChars(state, pos, prev)
			}
		}
	}
	return shouldISwap
}

func execMovement(state []string, startPos Point, movement Dir) Point {
	// Look along movement direction, until either '.' or '#'
	// If '.': move all encountered items by 1 (swapChars)
	// If '#': do nothing
	moved := recursiveExplore(state, startPos, movement)
	if moved {
		return startPos.moveInDir(movement)
	}
	return startPos
}

func getTotGPS(state []string) (out int) {
	h := len(state)
	w := len(state[0])
	for j := 1; j < h-1; j++ {
		for i := 1; i < w-1; i++ {
			if string(state[j][i]) == "O" {
				out += Point{x: i, y: j}.getGPSCoord()
			}
		}
	}
	return out
}

func findRobot(state []string) Point {
	h := len(state)
	w := len(state[0])
	for j := 1; j < h-1; j++ {
		for i := 1; i < w-1; i++ {
			if string(state[j][i]) == "@" {
				return Point{x: i, y: j}
			}
		}
	}
	return Point{}
}

func printState(state []string) {
	for _, l := range state {
		fmt.Println(l)
	}
}

// Steps:
// - Find initial position
// - Iterate over `movements`
// - Execute movement (given curr. pos., state, curr. mov.)
// - Evaluate result
func solve(state []string, movements string) (ans1 int) {
	// Initial position
	robotPos := findRobot(state)
	for _, d := range movements {
		// fmt.Println("\nMovement", i, "-", string(d), dirMap[d])
		robotPos = execMovement(state, robotPos, dirMap[d])
	}
	// printState(state)

	// Calculate gps coordinates
	return getTotGPS(state)
}

type Box struct {
	a, b Point
}

// Run this after each movement on every single box (debugging only)
func (b Box) check() {
	if b.a.y != b.b.y {
		log.Fatal("Box has inconsistent coordinates (y coordinates of points must match)")
	}
}

func upscaleString(line string) string {
	out := ""
	for _, c := range line {
		switch c {
		case '.':
			out += ".."
		case '#':
			out += "##"
		case 'O':
			out += "[]"
		case '@':
			out += "@."
		}
	}
	return out
}

// Returns:
// - Upscaled (on x) state
// - StateMap (given point, return box ID); NOTE: boxID is 1-indexed as 0 is the default value for the map
// - Boxes list
func upscaleState(state []string) ([]string, map[Point]int, []Box) {
	h := len(state)
	w := len(state[0])
	upscaledState := make([]string, 0)
	stateMap := make(map[Point]int)
	boxesList := make([]Box, 0)
	for j := 0; j < h; j++ {
		upscaledState = append(upscaledState, upscaleString(state[j]))
		for i := 0; i < w; i++ {
			if state[j][i] == 'O' {
				// New Box
				new_a := Point{x: 2 * i, y: j}
				new_b := Point{x: 2*i + 1, y: j}
				new_box := Box{a: new_a, b: new_b}
				new_box.check()
				boxesList = append(boxesList, new_box)
				stateMap[new_a] = len(boxesList)
				stateMap[new_b] = len(boxesList)
			}
		}
	}
	return upscaledState, stateMap, boxesList
}

// Recursion is a bit more challenging
// Recur down a binary tree (dfs)
// NOTE: unlike other explore function, here we swap with **next** element
// If on a box, only swap if from both branches of the recursion tree we get 'true' (and)
func recursiveExplore2(state []string, pos Point, movement Dir, stateMap map[Point]int, boxesList []Box) bool {
	var canIMove bool
	boxID := stateMap[pos]
	if boxID == 0 {
		// Not a box, what is it then
		curr_char := pos.getStateCh(state)
		switch curr_char {
		case ".":
			// Can move here
			return true
		case "#":
			return false
		default:
			log.Fatal("Shouldn't be here")
		}
	} else {
		// We are on a box!
		current_box := boxesList[boxID-1]
		can_move_a := false
		can_move_b := false
		if movement.y == 0 {
			if movement.x > 0 {
				// Recur on b
				can_move_a = true
				can_move_b = recursiveExplore2(state, current_box.b.moveInDir(movement), movement, stateMap, boxesList)
			} else {
				// Recur on a
				can_move_b = true
				can_move_a = recursiveExplore2(state, current_box.a.moveInDir(movement), movement, stateMap, boxesList)
			}
		} else {
			// Recur on box.a and box.b
			can_move_a = recursiveExplore2(state, current_box.a.moveInDir(movement), movement, stateMap, boxesList)
			can_move_b = recursiveExplore2(state, current_box.b.moveInDir(movement), movement, stateMap, boxesList)
		}

		if can_move_a && can_move_b {
			canIMove = true
			// - move the chars on the map
			new_pt_a := current_box.a.moveInDir(movement)
			new_pt_b := current_box.b.moveInDir(movement)
			// Move first the box 'side' which is "forward" (this is only useful for movements along x)
			if movement.x >= 0 {
				swapChars(state, current_box.b, new_pt_b)
			}
			swapChars(state, current_box.a, new_pt_a)
			if movement.x < 0 {
				swapChars(state, current_box.b, new_pt_b)
			}
			// - update boxesList content (need check, but should be in-place)
			boxesList[boxID-1] = Box{a: new_pt_a, b: new_pt_b}
			// - update stateMap
			stateMap[current_box.a] = 0
			stateMap[current_box.b] = 0
			stateMap[new_pt_a] = boxID
			stateMap[new_pt_b] = boxID
		} else {
			canIMove = false
		}
	}
	return canIMove
}

func copyMap(newMap, originalMap map[Point]int) {
	for key, value := range originalMap {
		newMap[key] = value
	}
}

func execMovement2(state []string, startPos Point, movement Dir, stateMap map[Point]int, boxesList []Box) (Point, []string, map[Point]int, []Box) {
	state_bak := make([]string, len(state))
	copy(state_bak, state)
	sm_bak := make(map[Point]int)
	copyMap(sm_bak, stateMap)
	boxList_bak := make([]Box, len(boxesList))
	copy(boxList_bak, boxesList)
	moved := recursiveExplore2(state, startPos.moveInDir(movement), movement, stateMap, boxesList)
	if moved {
		// Update position of '@' in state
		swapChars(state, startPos, startPos.moveInDir(movement))
		return startPos.moveInDir(movement), state, stateMap, boxesList
	} else {
		// Undo all movements
		return startPos, state_bak, sm_bak, boxList_bak
	}
}

func getTotGPS2(boxesList []Box) (out int) {
	for _, b := range boxesList {
		out += b.a.getGPSCoord()
	}
	return out
}

// Steps:
// - Upscale on x
// - Get stateMap (point -> boxID)
func solve2(state []string, movements string) (ans2 int) {
	newState, stateMap, boxesList := upscaleState(state)
	robotPos := findRobot(newState)

	for _, d := range movements {
		// fmt.Println("\nMovement", i, "-", string(d), dirMap[d])
		robotPos, newState, stateMap, boxesList = execMovement2(newState, robotPos, dirMap[d], stateMap, boxesList)
	}
	// printState(newState)
	ans2 = getTotGPS2(boxesList)
	return ans2
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	lookingAtState := true
	state := make([]string, 0)
	movements := ""
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			lookingAtState = false
		} else if lookingAtState {
			state = append(state, line)
		} else {
			movements += line
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	state_bak := make([]string, len(state))
	copy(state_bak, state)
	// printState(state)
	// fmt.Println("\nMovements: ", movements)

	sol1 := solve(state, movements)
	fmt.Println("Part 1:", sol1)

	sol2 := solve2(state_bak, movements)
	fmt.Println("Part 2:", sol2)
}
