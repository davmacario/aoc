package main

import (
	"bufio"
	"day05/functions"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Handle error
func errHdl(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Returns true if a follows b
func follows(a int, b int, m map[int][]int) bool {
	if functions.InSlice(a, m[b]) {
		return true
	}
	return false
}

// Returns true if a precedes b
func precedes(a int, b int, m map[int][]int) bool {
	if functions.InSlice(b, m[a]) {
		return true
	}
	return false
}

// Checks order given map
func checkOrder(line []int, m map[int][]int) bool {
	// Recur on subarrays
	if len(line) > 1 {
		return follows(line[1], line[0], m) && checkOrder(line[1:], m)
	} else {
		return true
	}
}

// Returns the middle element
func getMidElem(arr []int) int {
	return arr[len(arr)/2]
}

// Fix order of the elements
// Returns the array with ordered items, given map m
// NOTE: this occurred to me in a vision
func fixOrder(line []int, m map[int][]int) []int {
	if len(line) <= 1 {
		return line
	}
	i := 1
	for i < len(line) {
		for j, np := range line[:i] { // j < i
			if follows(np, line[i], m) {
				// Swap them
				swapF := reflect.Swapper(line)
				swapF(i, j)
				i = j
				break
			}
		}
		i++
	}
	return line
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Scanner to read the file
	scanner := bufio.NewScanner(f)

	// Iterate over lines of file
	is_still_ordering := true
	nums_map := make(map[int][]int) // Maps number to the ones that follow it
	var ans1 int
	var ans2 int
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			is_still_ordering = false
		}

		if is_still_ordering {
			nums := strings.Split(line, "|")
			a, e_a := strconv.Atoi(nums[0])
			b, e_b := strconv.Atoi(nums[1])

			errHdl(e_a)
			errHdl(e_b)

			lst, ok := nums_map[a]
			if ok { // a is valid key
				nums_map[a] = append(lst, b)
			} else {
				nums_map[a] = []int{b}
			}
		} else if line != "" {
			// Update
			pages := make([]int, 0)
			for _, val := range strings.Split(line, ",") {
				n, e := strconv.Atoi(val)
				errHdl(e)
				pages = append(pages, n)
			}
			if checkOrder(pages, nums_map) {
				ans1 += getMidElem(pages)
			} else {
				fixed_ord := fixOrder(pages, nums_map)
				ans2 += getMidElem(fixed_ord)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result 1:", ans1)

	// Part 2

	// Ideas:
	// - Use what done so far (map) to correct order:
	// - For each item in the update pages (index i)
	// - check each previous item in list (index j):
	//  - if should be after, swap and set i to the new position (swapped with) and continue
	//  - Else, j++
	// - i++
	fmt.Println("Result 2:", ans2)
}
