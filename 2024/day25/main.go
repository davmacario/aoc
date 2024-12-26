package main

import (
	"bufio"
	"day25/utils"
	"fmt"
	"log"
	"os"
	_ "strings"
)

func IsLock(in []string) bool {
	return in[0][0] == '#'
}

func KeyFitsInLock(key, lock []int) bool {
	return utils.All(utils.SumSlice(key, lock), func(x int) bool { return x <= 5 })
}

func MatToArr(k []string) []int {
	out := make([]int, 0)
	h := len(k)
	w := len(k[0])
	for i := 0; i < w; i++ {
		curr_count := 0
		for j := 0; j < h; j++ {
			if k[j][i] == '#' {
				curr_count++
			}
		}
		out = append(out, curr_count-1)
	}
	return out
}

func part1(keys, locks [][]string) int {
	// Translate all keys and locks to slices of ints (or arrays...)
	keySlices := make([][]int, len(keys))
	// fmt.Println("Keys:")
	for i, mat := range keys {
		keySlices[i] = MatToArr(mat)
		// fmt.Println(keySlices[i])
	}
	lockSlices := make([][]int, len(locks))
	// fmt.Println("Locks:")
	for i, mat := range locks {
		lockSlices[i] = MatToArr(mat)
		// fmt.Println(lockSlices[i])
	}

	// Nested loop - keys and locks --- may be improved...
	var out int
	for _, k := range keySlices {
		for _, l := range lockSlices {
			if KeyFitsInLock(k, l) {

				out++
			}
		}
	}
	return out
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	inputLocks := make([][]string, 0)
	inputKeys := make([][]string, 0)
	currElem := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			// "close" current lock
			if IsLock(currElem) {
				inputLocks = append(inputLocks, currElem)
			} else {
				inputKeys = append(inputKeys, currElem)
			}
			currElem = make([]string, 0)
		} else {
			currElem = append(currElem, line)
		}
	}
	if IsLock(currElem) {
		inputLocks = append(inputLocks, currElem)
	} else {
		inputKeys = append(inputKeys, currElem)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ans1 := part1(inputKeys, inputLocks)
	fmt.Println("Part 1:", ans1)
}
