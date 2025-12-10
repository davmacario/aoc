package main

import (
	"bufio"
	"day7/functions"
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

	scanner := bufio.NewScanner(f)

	matrix := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result2 := functions.FollowQuantumBeams(matrix)
	result1 := functions.FollowBeams(matrix)

	// for _, r := range matrix {
	// 	fmt.Println(r)
	// }

	fmt.Printf("Result 1: %d\n", result1)
	fmt.Printf("Result 2: %d\n", result2)
}
