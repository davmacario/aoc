package main

import (
	"bufio"
	ds "day21/data_structures"
	f "day21/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
// |   | 0 | A |
// +---+---+---+
// ^ (0, 0)
var num_keypad = map[string]ds.Point{
	"A": ds.NewPoint(2, 0),
	"0": ds.NewPoint(1, 0),
	"1": ds.NewPoint(0, 1),
	"2": ds.NewPoint(1, 1),
	"3": ds.NewPoint(2, 1),
	"4": ds.NewPoint(0, 2),
	"5": ds.NewPoint(1, 2),
	"6": ds.NewPoint(2, 2),
	"7": ds.NewPoint(0, 3),
	"8": ds.NewPoint(1, 3),
	"9": ds.NewPoint(2, 3),
}

// +---+---+---+
// |   | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+
// ^ (0,0)
var dir_keypad = map[string]ds.Point{
	"<": ds.NewPoint(0, 0),
	"v": ds.NewPoint(1, 0),
	">": ds.NewPoint(2, 0),
	"^": ds.NewPoint(1, 1),
	"A": ds.NewPoint(2, 1),
}

// Higher priority means chosen first
var dir_priority = map[string]int{
	"<": 3,
	"v": 1,
	">": 0,
	"^": 2,
}

// Find best path (in term of minimum output string length)
// Idea: get the 2 possible paths minimizing the changes in dir.
// If one of the 2 is impossible, return the other
// Else, recur on writeSequenceDir on both, until one of the 2 has a lower score
func MoveInBetweenPoints(a, b, forbidden ds.Point) string {
	dx := b.GetX() - a.GetX()
	dy := b.GetY() - a.GetY()
	dir_lr := ds.NewDir(f.Sign(dx), 0)
	dir_ud := ds.NewDir(0, f.Sign(dy))
	char_lr := ds.DirToStr[dir_lr]
	char_ud := ds.DirToStr[dir_ud]

	path1 := strings.Repeat(string(char_lr), f.IntAbs(dx)) + strings.Repeat(string(char_ud), f.IntAbs(dy)) + "A"
	path2 := strings.Repeat(string(char_ud), f.IntAbs(dy)) + strings.Repeat(string(char_lr), f.IntAbs(dx)) + "A"
	if GoesThroughPoint(path1, a, forbidden) {
		return path2
	}
	if GoesThroughPoint(path2, a, forbidden) {
		return path1
	}

	if dir_priority[string(char_lr)] >= dir_priority[string(char_ud)] {
		return path1
	} else {
		return path2
	}
}

func GoesThroughPoint(movements string, start, p ds.Point) bool {
	currPoint := start
	for _, c := range movements {
		if currPoint == p {
			return true
		}
		currPoint = currPoint.MoveInDir(ds.ChToDir[c])
	}
	return false
}

// Returns the number of operations needed to type a specific path (and insert A after)
// A + >>^^ + A
func CalcPathCost(movements string) int {
	out := len(movements)
	movements_complete := movements
	for i, c := range movements_complete[:len(movements_complete)-1] {
		out += ds.ManhattanDist(dir_keypad[string(c)], dir_keypad[string(movements_complete[i+1])])
	}
	return out
}

// mov: string "<src><dst>" where source and destination are keypad keys
// This function also updates the map
func MapKeypadMovements(mov string, memoKeypad map[string]string) string {
	forbidden := ds.NewPoint(0, 1)
	_, ok := memoKeypad[mov]
	if !ok {
		if mov[0] == mov[1] {
			memoKeypad[mov] = "A"
		} else if memoKeypad[mov] == "" {
			curr_start := dir_keypad[string(mov[0])]
			curr_end := dir_keypad[string(mov[1])]
			// Calculate movement (avoiding forbidden position and minimizing
			// dir changes to minimize input movements on the dir pad)
			chosenPath := MoveInBetweenPoints(curr_start, curr_end, forbidden)
			// Path in memo moves from start to end, then presses enter
			memoKeypad[mov] = chosenPath
		}
	}
	return memoKeypad[mov]
}

// IDEA: create functions to translate a string (to be "entered" on a specific
// keypad type) into movements to be executed on the directional keypad
// controlling the arm that enters it.
// To translate, consider each char and its successor. Given any couple of
// chars, the minimum (valid) path is always the same, so we may be able to
// generate a mapping (or something O(1) to get the movements to go from one
// char to the other)
// For N keys, the number of possible movements is N*(N-1) (direction counts)

// Given a string to be entered in the numpad, get series of keys to be pressed
// on the dir_pad (in string form)
// Inputs:
// - num_seq: sequence of numbers to be entered; at the end, "A" will be appended
// - startFrom: Point where the arm will start to generate the seq (should
// always be "A")
// - memo: memoization string - given a string "ab", will save the combination
// of presses on the dir_pad to go from key a to key b
func writeSequenceNum(num_seq string, startFrom string, memoNum map[string]string) string {
	out := ""
	forbidden := ds.NewPoint(0, 0)
	complete_seq := startFrom + num_seq
	for i, c := range complete_seq[:len(complete_seq)-1] {
		curr_pair := string(c) + string(complete_seq[i+1])
		// fmt.Println("Current move:", curr_pair)
		if memoNum[curr_pair] == "" && curr_pair[0] == curr_pair[1] {
			memoNum[curr_pair] = "A"
		} else if memoNum[curr_pair] == "" {
			curr_start := num_keypad[string(c)]
			curr_end := num_keypad[string(complete_seq[i+1])]
			// Calculate movement (avoiding forbidden position and minimizing
			// dir changes to minimize input movements on the dir pad)
			chosenPath := MoveInBetweenPoints(curr_start, curr_end, forbidden)
			// Path in memo moves from start to end, then presses enter
			memoNum[curr_pair] = chosenPath
		}
		out += memoNum[curr_pair]
	}
	return out
}

func writeSequenceDir(dir_seq string, startFrom string, memo map[string]string) string {
	out := ""
	complete_seq := startFrom + dir_seq
	for i, c := range complete_seq[:len(complete_seq)-1] {
		curr_pair := string(c) + string(complete_seq[i+1])
		out += MapKeypadMovements(curr_pair, memo)
	}
	return out
}

func printMemo(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, "->", v)
	}
}

// Returns the number of keys to be pressed to achieve dir_sequence via n_rec
// keypads
func nextDirRecur(dir_sequence string, n_rec, tot_rec int, memoDirRecur map[string][]int64, memo_dir map[string]string) []int64 {
	mem, ok := memoDirRecur[dir_sequence]
	missing_it := tot_rec - n_rec
	// fmt.Println(missing_it)
	if missing_it <= 0 {
		return []int64{int64(len(dir_sequence))}
	}
	if ok && missing_it < len(mem) {
		// fmt.Println("HERE")
		return mem[:missing_it+1]
	} else {
		next_iter := writeSequenceDir(dir_sequence, "A", memo_dir)
		following := make([]int64, missing_it)
		full_next := next_iter
		for i, c := range full_next[:len(full_next)-1] {
			curr_couple := string(c) + string(full_next[i+1])
			// fmt.Println(curr_couple)
			follow := nextDirRecur(curr_couple, n_rec+1, tot_rec, memoDirRecur, memo_dir)
			// fmt.Println(follow)
			following = f.SumSlices(following, follow)
		}
		memoDirRecur[dir_sequence] = append([]int64{int64(len(dir_sequence))}, following...)
		return memoDirRecur[dir_sequence]
	}
}

// Return solution for part 1: sum of complexities for entering each code
func solve(instructions []string) (out int, out2 int64) {
	// Memoize, for each consecutive couple of positions (keys) in the sequence
	// (both on num and dir keypads), the minimum path (as string)
	memo_num := make(map[string]string)
	memo_dir := make(map[string]string)

	for _, inst := range instructions {
		fmt.Println(inst)
		dir_sequence_1 := writeSequenceNum(inst, "A", memo_num)
		dir_sequence_2 := writeSequenceDir(dir_sequence_1, "A", memo_dir)
		dir_sequence_3 := writeSequenceDir(dir_sequence_2, "A", memo_dir)
		// Evaluate complexity for current sequence
		numericalPart, err := strconv.Atoi(f.GetNumCharsOnly(inst))
		if err != nil {
			panic("No num found in current sequence")
		}
		curr_complex := len(dir_sequence_3) * numericalPart
		fmt.Println("Complexity of", inst, "->", len(dir_sequence_3), "*", numericalPart, "=", curr_complex)
		out += curr_complex
		fmt.Println()
	}

	// Build map of direction string -> slice of count after i iters
	memoDirSlice := make(map[string][]int64)
	for k, v := range memo_dir {
		memoDirSlice[k] = []int64{int64(len(k)), int64(len(v))}
	}

	n_iter_2 := 24
	var curr_complex int64
	for _, inst := range instructions {
		fmt.Println(inst)
		for num := range n_iter_2 {
			dir_sequence := writeSequenceNum(inst, "A", memo_num)
			complete_seq := "A" + dir_sequence
			var countKeys int64
			for i, c := range complete_seq[:len(complete_seq)-1] {
				curr_pair := string(c) + string(complete_seq[i+1])
				followers := nextDirRecur(curr_pair, 0, num, memoDirSlice, memo_dir)
				countKeys += followers[len(followers)-1]
			}
			fmt.Println(countKeys)
			numericalPart, _ := strconv.Atoi(f.GetNumCharsOnly(inst))
			curr_complex = countKeys * int64(numericalPart)
		}
	}
	// fmt.Println("Complexity of", inst, "->", countKeys, "*", numericalPart, "=", curr_complex)
	out2 += curr_complex

	return out, out2
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	instructions := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res1, res2 := solve(instructions)
	fmt.Println("Part 1:", res1)
	fmt.Println("Part 2:", res2)
}
