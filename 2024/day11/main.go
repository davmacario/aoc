package main

import (
	"bufio"
	"day11/functions"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	_ "sync"
)

var NIter2 = 75

func blink(stones []string) []string {
	out := make([]string, 0)
	for _, s := range stones {
		len_s := len(s)
		if s == "0" {
			out = append(out, "1")
		} else if len_s%2 == 0 {
			p1 := s[:len_s/2]
			p2_int, _ := strconv.Atoi(s[len_s/2:])
			p2 := strconv.Itoa(p2_int)
			out = append(out, p1)
			out = append(out, p2)
		} else {
			p_int, _ := strconv.Atoi(s)
			p := strconv.Itoa(p_int * 2024)
			out = append(out, p)
		}
	}
	return out
}

func part1(stones []string) int {
	// var count int
	upd_stones := stones
	for i := 0; i < 25; i++ {
		upd_stones = blink(upd_stones)
		// fmt.Println(upd_stones)
	}
	return len(upd_stones)
}

// At every recursion, returns a slice of length missing_iters = N - i
// Returned slice contains, for index i, the number of stones after i blinks
// NOTE: i = 0 -> 1 (always)
// The result is basically just sum([memo[s][-1] for s in stones]), pythonically
func blinkRecur(s string, n_iter int, max_iter int, memo map[int][]int) []int {
	s_int, _ := strconv.Atoi(s)
	n, ok := memo[s_int]
	missing_iters := max_iter - n_iter
	if missing_iters <= 0 {
		return []int{1}
	}
	if ok && missing_iters < len(n) {
		// Answer is there
		return n[:missing_iters+1]
	} else {
		next_stones := blink([]string{s})
		ret_follow := make([]int, missing_iters)
		for _, s_i := range next_stones {
			follow := blinkRecur(s_i, n_iter+1, max_iter, memo)
			ret_follow = functions.SumSlices(ret_follow, follow)
		}
		// Either create or replace with longer, otherwise would've matched prev cond
		memo[s_int] = append([]int{1}, ret_follow...)
		return memo[s_int]
	}
}

func part2(stones []string) int {
	// var count int
	n_stones_out := 0
	memo := make(map[int][]int)
	for _, s := range stones {
		followers := blinkRecur(s, 0, NIter2, memo)
		n_stones_out += followers[len(followers)-1]
		// fmt.Println(memo)
	}
	return n_stones_out
}

// ############################################################################
// Other (better) solution
// Much better in terms of memory usage
func betterSolution(stones []string) int {
	s_map := make(map[int]int)
	for _, s := range stones {
		s_map[functions.StrToInt(s)] += 1
	}
	for range NIter2 {
		s_map = newBlink(s_map)
	}
	return nOfStones(s_map)
}

// The 'stones' map contains, after each blink, the number of times each stone
// (key) appears in the sequence.
// The order does not matter, and equal stones produce equal outputs after each
// blink, so if stone 'n' produces stones 'a' and 'b', then 100 stones 'n' will
// yield 100 stones 'a' and 100 stones 'b'
// Therefore, we just need to iterate on the current 'stones' map and, for each
// stone 'k' appearing 'count' times, add 'count' to the value associated with
// the stone(s) generated by 'k' after a blink in the new map.

func newBlink(stones map[int]int) map[int]int {
	out := make(map[int]int)
	for k, count := range stones {
		k_str := strconv.Itoa(k)
		if k == 0 {
			// All zeros will become ones
			out[1] += count
		} else if len(k_str)%2 == 0 {
			out[functions.StrToInt(k_str[:len(k_str)/2])] += count
			out[functions.StrToInt(k_str[len(k_str)/2:])] += count
		} else {
			out[k*2024] += count
		}
	}
	return out
}

func nOfStones(stones map[int]int) int {
	var out int
	for _, v := range stones {
		out += v
	}
	return out
}

// ############################################################################

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	stones := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		stones = strings.Fields(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(stones)
	res1 := part1(stones)
	fmt.Println("Part 1:", res1)
	res2 := part2(stones)
	fmt.Println("Part 2:", res2)
	resBetter := betterSolution(stones)
	fmt.Println("Part 2, better:", resBetter)
}
