package main

import (
	"bufio"
	ds "day21/data_structures"
	f "day21/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	_ "strings"
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

// Move from point `a` to point `b` minimizing the changes of direction
// FIXME: no need to modify the changes in direction. Need to find a way to
// minimize the overall distance between points that need to be traversed to
// enter the sequence -> <<^ takes "longer" than ^<< because in the former we
// have to "go back" to ^ after <.
// This means that we will always "refer" to the dir_keypad
func MoveInBetweenPoints(a, b, forbidden ds.Point) string {
	dx := b.GetX() - a.GetX()
	dy := b.GetY() - a.GetY()
	dir_lr := ds.MakeDir(f.Sign(dx), 0)
	dir_ud := ds.MakeDir(0, f.Sign(dy))
	char_lr := ds.DirToStr[dir_lr]
	char_ud := ds.DirToStr[dir_ud]

	// New implementation
	// We know we only have to move along 2 directions (l OR r, u OR d)
	// starting from A (2, 1), see what of the 2 dirs' key is closer (Manhattan)
	// Go to that one
	out := ""
	n_lr_keys := f.IntAbs(dx)
	n_ud_keys := f.IntAbs(dy)
	n_keys_to_be_pressed := n_lr_keys + n_ud_keys
	pointLR := dir_keypad[string(char_lr)]
	pointUD := dir_keypad[string(char_ud)]
	currPt := dir_keypad["A"]
	for n_keys_to_be_pressed > 0 {
		// Find closer to currPt between pointLR and pointUD
		dLR := ds.ManhattanDist(currPt, pointLR)
		dUD := ds.ManhattanDist(currPt, pointUD)
		if (dLR < dUD && n_lr_keys > 0) || n_ud_keys <= 0 {
			// Move towards key for lr position
			out += string(char_lr)
			currPt = pointLR
			n_lr_keys--
		} else {
			// Assert still need to move towards UD
			out += string(char_ud)
			currPt = pointUD
			n_ud_keys--
		}
		n_keys_to_be_pressed = n_lr_keys + n_ud_keys
	}

	// Try to invert the sequence if you would encounter the forbidden point
	// I'm sure there is a 100x more efficient way
	cpt := a
	ind := 0
	for cpt != b {
		next := cpt.MoveInDir(ds.ChToDir[rune(out[ind])])
		if next == forbidden {
			fmt.Print("BANG ", out, " ")
			out = f.ReverseString(out)
			fmt.Println(out)
			break
		}
		ind++
		cpt = next
	}

	return out + "A"
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
func writeSequenceNum(num_seq string, startFrom string, memo map[string]string) string {
	out := ""
	forbidden := ds.NewPoint(0, 0)
	// Add starting point
	complete_seq := startFrom + num_seq
	for i, c := range complete_seq[:len(complete_seq)-1] {
		curr_pair := string(c) + string(complete_seq[i+1])
		fmt.Println("Current move:", curr_pair)
		if memo[curr_pair] == "" {
			curr_start := num_keypad[string(c)]
			curr_end := num_keypad[string(complete_seq[i+1])]
			// Calculate movement (avoiding forbidden position and minimizing
			// dir changes to minimize input movements on the dir pad)
			chosenPath := MoveInBetweenPoints(curr_start, curr_end, forbidden)
			// Path in memo moves from start to end, then presses enter
			// Hence, the last key pressed should be 'A'
			memo[curr_pair] = chosenPath
		}
		out += memo[curr_pair]
		// fmt.Println("->", memo[curr_pair], "\n")
	}
	return out
}

func writeSequenceDir(dir_seq string, startFrom string, memo map[string]string) string {
	out := ""
	forbidden := ds.NewPoint(0, 1)
	complete_seq := startFrom + dir_seq
	for i, c := range complete_seq[:len(complete_seq)-1] {
		curr_pair := string(c) + string(complete_seq[i+1])
		if memo[curr_pair] == "" {
			curr_start := dir_keypad[string(c)]
			curr_end := dir_keypad[string(complete_seq[i+1])]
			// Calculate movement (avoiding forbidden position and minimizing
			// dir changes to minimize input movements on the dir pad)
			chosenPath := MoveInBetweenPoints(curr_start, curr_end, forbidden)
			// Path in memo moves from start to end, then presses enter
			// Hence, the last key pressed should be 'A' - done in function
			memo[curr_pair] = chosenPath
		}
		out += memo[curr_pair]
	}
	return out
}

func printMemo(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, "->", v)
	}
}

// Return solution for part 1: sum of complexities for entering each code
func solve1(instructions []string) (out int) {
	// Memoize, for each consecutive couple of positions (keys) in the sequence
	// (both on num and dir keypads), the minimum path (as string)
	memo_num := make(map[string]string)
	memo_dir := make(map[string]string) // e.g., ">v" -> "<A"

	for _, inst := range instructions {
		fmt.Println(inst)
		dir_sequence_1 := writeSequenceNum(inst, "A", memo_num)
		fmt.Println(dir_sequence_1)
		dir_sequence_2 := writeSequenceDir(dir_sequence_1, "A", memo_dir)
		fmt.Println(dir_sequence_2)
		dir_sequence_3 := writeSequenceDir(dir_sequence_2, "A", memo_dir)
		fmt.Println(dir_sequence_3)

		// Evaluate complexity for current sequence
		numericalPart, err := strconv.Atoi(f.GetNumCharsOnly(inst))
		if err != nil {
			panic("No num found in current sequence")
		}
		curr_complex := len(dir_sequence_3) * numericalPart
		fmt.Println("Complexity of", inst, "->", len(dir_sequence_3), "*", numericalPart, "=", curr_complex)
		out += curr_complex

		// if i == 0 {
		// 	log.Fatal("---")
		// }
		fmt.Println()
	}

	printMemo(memo_num)
	fmt.Println("")
	printMemo(memo_dir)

	return out
}

func main() {
	input_file := "./in_small.txt"
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

	res1 := solve1(instructions)
	fmt.Println("Part 1:", res1)
}
