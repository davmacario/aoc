package main

import (
	"bufio"
	"day5/functions"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()


	var passedHalf bool
	rangesStr := make([]string, 0)
	numbers := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")

		if line == "" {
			passedHalf = true
			continue
		}

		if !passedHalf {
			// Get ranges
			rangesStr = append(rangesStr, line)
		} else {
			numInt, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, numInt)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rangesSlice := functions.RangesFromStrings(rangesStr)
	simplifiedRanges := functions.SimplifyRanges(rangesSlice)

	var result1 int
	for _, n := range numbers {
		for _, r := range simplifiedRanges {
			if r.Contains(n) {
				result1 += 1
				break
			}
		}
	}

	var result2 int
	for _, r := range simplifiedRanges {
		result2 += r.Span()
	}

	fmt.Printf("Solution 1: %d\n", result1)
	fmt.Printf("Solution 2: %d\n", result2)
}
