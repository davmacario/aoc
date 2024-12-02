package main

import (
	"bufio"
	"day01/functions"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input_file string

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Scanner to read the file
	scanner := bufio.NewScanner(f)

	// Initialize slices
	v1 := []int{}
	v2 := []int{}
	var lineNr int
	var line string
	var strNumbers []string
	var numbers []int
	var a int
	var b int

	// Loop over all lines
	for scanner.Scan() {
		// Line (string)
		line = scanner.Text()
		// Slice of strings containing the numbers
		strNumbers = strings.Fields(line)
		// Convert slice of str to slice of ints
		numbers = make([]int, len(strNumbers))
		for i, s := range strNumbers {
			numbers[i], _ = strconv.Atoi(s)
		}

		a = numbers[0]
		b = numbers[1]

		if lineNr == 0 {
			// First iteration, add item to array
			v1 = append(v1, a)
			v2 = append(v2, b)
		} else {
			pos_a := functions.BinSearch(v1, a)
			v1 = append(v1[:pos_a+1], append([]int{a}, v1[pos_a+1:]...)...)

			pos_b := functions.BinSearch(v2, b)
			v2 = append(v2[:pos_b+1], append([]int{b}, v2[pos_b+1:]...)...)
		}
		lineNr++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int
	for i := 0; i < len(v1); i++ {
		// Difference between values in v2 and v1
		sum += functions.Abs(v2[i] - v1[i])
	}

	fmt.Println("Result (part 1): " + strconv.Itoa(sum))

	count_val_v2 := functions.UniqueValuesCount(v2)
	var sum_2 int
	for i, val := range v1 {
		count, ok := count_val_v2[val]
		if ok {
			sum_2 += count * v1[i]
		}
	}

	fmt.Println("Result (part2): " + strconv.Itoa(sum_2))
}
