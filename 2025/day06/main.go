package main

import (
	"bufio"
	"day6/functions"
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

	// Scanner to read the file
	scanner := bufio.NewScanner(f)

	// nums contains, on each line (1st index) slices of all numbers that will
	// be part of the same operation
	var nums [][]int
	operations := []string{}
	firstLine := true

	// Part 2: store all lines in memory first
	var lines []string

	// Iterate over lines of file
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
		// Get array of fields (str)
		fields := strings.Fields(line)

		_, err := strconv.Atoi(fields[0])
		if err != nil {
			// Last line, operations
			operations = fields
			continue
		}

		if firstLine {
			// Get length of nums (i.e., how many operations)
			nums = make([][]int, len(fields))
			firstLine = false
		}

		// Default, append to pre-existing lists
		for i, n := range fields {
			nInt, err := strconv.Atoi(n)

			if err != nil {
				log.Fatal(err)
			}

			nums[i] = append(nums[i], nInt)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result1 int
	for i := range nums {
		// fmt.Println(nums[i])
		// fmt.Println(operations[i])

		result1 += functions.PerformOperation(nums[i], operations[i])
	}

	fmt.Printf("Result 1: %d\n", result1)

	// NOTE: operations are the same
	formattedOps := functions.ConvertOperationFormat(lines[:len(lines)-1])
	var result2 int
	for i, opNums := range formattedOps {
		// fmt.Println(opNums)
		// fmt.Println(operations[i])

		result2 += functions.PerformOperation(opNums, operations[i])
	}
	fmt.Printf("Result 2: %d\n", result2)
}
