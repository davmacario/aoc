package main

import (
	"bufio"
	"day4/functions"
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

	matrix := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result1, upd_matrix := functions.CountAccessibleRollsAndRemove(matrix)
	count_rem := result1
	result2 := result1
	for count_rem > 0 {
		count_rem, upd_matrix = functions.CountAccessibleRollsAndRemove(upd_matrix)
		result2 += count_rem
	}

	fmt.Printf("Result 1: %d\n", result1)
	fmt.Printf("Result 2: %d\n", result2)
}
