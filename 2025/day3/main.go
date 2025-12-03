package main

import (
	"bufio"
	"day3/functions"
	"day3/utils"
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

	result1 := 0
	result2 := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		bank := utils.StringToIntSlice(line)
		result1 += functions.PickTwoBatteries(bank)
		jolt2 := functions.PickTwelveBatteries(bank)
		result2 += jolt2
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result, part 1: %d\n", result1)
	fmt.Printf("Result, part 2: %d\n", result2)
}
