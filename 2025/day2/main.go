package main

import (
	"bufio"
	"day2/functions"
	"fmt"
	"log"
	"os"
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
	scanner.Split(functions.SplitCommas)

	result1 := 0
	result2 := 0
	// Iterate over lines of file
	for scanner.Scan() {
		// Get range "a-b", cleaning up from \n (only applies to last item)
		numRange := strings.TrimSuffix(scanner.Text(), "\n")
		a, b := functions.SplitRange(numRange)

		for i := a; i <= b; i++ {
			if functions.IsInvalid(i) {
				fmt.Printf("Found invalid number: %d\n", i)
				result1 += i
				result2 += i
			} else if functions.IsInvalid2(i) {
				fmt.Printf("Found invalid(2) number: %d\n", i)
				result2 += i
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result 1: %d\n", result1)
	fmt.Printf("Result 2: %d\n", result2)
}
