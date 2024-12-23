package main

import (
	"bufio"
	ds "day21/data_structures"
	f "day21/utils"
	"fmt"
	"log"
	"os"
	"sort"
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

var num_keypad_min_dist = map[string]string{
	"AA": "A",
	"A0": "<A",
	"A1": "^<<A",
	"A2": "^<A",
	"A3": "^A",
	"A4": "^^<<A",
	"A5": "",
	"A6": "",
	"A7": "",
	"A8": "",
	"A9": "",
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
func MoveInBetweenPoints(a, b ds.Point) (string, string) {
	dx := b.GetX() - a.GetX()
	dy := b.GetY() - a.GetY()
	dir_lr := ds.NewDir(f.Sign(dx), 0)
	dir_ud := ds.NewDir(0, f.Sign(dy))
	char_lr := ds.DirToStr[dir_lr]
	char_ud := ds.DirToStr[dir_ud]

	substr_lr := strings.Repeat(string(char_lr), f.IntAbs(dx))
	substr_ud := strings.Repeat(string(char_ud), f.IntAbs(dy))
	return (substr_lr + substr_ud + "A"), (substr_ud + substr_lr + "A")
}

// Approach: given the chars to be pressed to move, generate all permutations.
// Then, select best one based on the cost to be typed (CalcPathCost)
func MoveInBetweenPoints2(a, b, forbidden ds.Point, num bool) string {
	dx := b.GetX() - a.GetX()
	dy := b.GetY() - a.GetY()
	dir_lr := ds.NewDir(f.Sign(dx), 0)
	dir_ud := ds.NewDir(0, f.Sign(dy))
	char_lr := ds.DirToStr[dir_lr]
	char_ud := ds.DirToStr[dir_ud]

	// Need to place |dx|*char_lr and |dy|*char_ud
	tmp := strings.Repeat(string(char_lr), f.IntAbs(dx)) + strings.Repeat(string(char_ud), f.IntAbs(dy))
	all_perm := f.PermutationsOfString(tmp)
	// Only keep ones that don't pass through forbidden point
	var i int
	for i < len(all_perm) {
		if GoesThroughPoint(all_perm[i], a, forbidden) {
			all_perm = f.RemoveFromSlice(all_perm, i)
		} else {
			i++
		}
	}

	sort.Slice(all_perm, func(i, j int) bool {
		return f.CalcCharChanges(all_perm[i]) < f.CalcCharChanges(all_perm[j])
	})
	out := all_perm[0]
	min_score := CalcPathCost(out)
	for _, s := range all_perm {
		c_score := CalcPathCost(s)
		if c_score < min_score {
			min_score = c_score
			out = s
		}
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
	complete_seq := startFrom + num_seq
	for i, c := range complete_seq[:len(complete_seq)-1] {
		curr_pair := string(c) + string(complete_seq[i+1])
		fmt.Println("Current move:", curr_pair)
		if memo[curr_pair] == "" && curr_pair[0] == curr_pair[1] {
			memo[curr_pair] = "A"
		} else if memo[curr_pair] == "" {
			curr_start := num_keypad[string(c)]
			curr_end := num_keypad[string(complete_seq[i+1])]
			// Calculate movement (avoiding forbidden position and minimizing
			// dir changes to minimize input movements on the dir pad)
			chosenPath := MoveInBetweenPoints2(curr_start, curr_end, forbidden, true)
			// Path in memo moves from start to end, then presses enter
			memo[curr_pair] = chosenPath
		}
		out += memo[curr_pair]
	}
	return out
}

func writeSequenceDir(dir_seq string, startFrom string, memo map[string]string) string {
	out := ""
	forbidden := ds.NewPoint(0, 1)
	complete_seq := startFrom + dir_seq
	for i, c := range complete_seq[:len(complete_seq)-1] {
		curr_pair := string(c) + string(complete_seq[i+1])
		// fmt.Println("Current move:", curr_pair)
		if memo[curr_pair] == "" && curr_pair[0] == curr_pair[1] {
			memo[curr_pair] = "A"
		} else if memo[curr_pair] == "" {
			curr_start := dir_keypad[string(c)]
			curr_end := dir_keypad[string(complete_seq[i+1])]
			// Calculate movement (avoiding forbidden position and minimizing
			// dir changes to minimize input movements on the dir pad)
			chosenPath := MoveInBetweenPoints2(curr_start, curr_end, forbidden, false)
			// Path in memo moves from start to end, then presses enter
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
	memo_num["37"] = "<<^^A"
	memo_dir := make(map[string]string) // e.g., ">v" -> "<A"
	// memo_dir["A<"] = "<v<A"

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

	// printMemo(memo_num)
	// fmt.Println("")
	// printMemo(memo_dir)

	return out
}

func main() {
	input_file := "./in_test.txt"
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
