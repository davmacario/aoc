package main

import (
	"bufio"
	"day02/functions"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkAbsInRange(n int) bool {
	abs_n := functions.Abs(n)
	if abs_n >= 1 && abs_n <= 3 {
		return true
	}
	return false
}

func checkSafe(v []int) bool {
	diffs := make([]int, len(v)-1)
	for i, n := range v[1:] {
		diffs[i] = n - v[i]
	}
	check := func(num int) bool { return (checkAbsInRange(num) && functions.Sign(num) == functions.Sign(diffs[0])) }
	return functions.All(diffs, check)
}

func checkSafe2(v []int) bool {
	if checkSafe(v) {
		return true
	}
	slice := v[:]
	// fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		arr_but_one := make([]int, 0, len(slice)-1)
		arr_but_one = append(arr_but_one, slice[:i]...)
		arr_but_one = append(arr_but_one, slice[i+1:]...)
		// fmt.Println("> ", arr_but_one)
		if checkSafe(arr_but_one) {
            // fmt.Println("->>Safe!")
			return true
		}
	}
	return false
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
	var countSafe int
	for scanner.Scan() {
		line := scanner.Text()
		// Get array of fields (str)
		fields := strings.Fields(line)
		numbers := make([]int, len(fields))
		for i, s := range fields {
			numbers[i], _ = strconv.Atoi(s)
		}

		// Assuming n. fields > 0
		if checkSafe(numbers) {
			countSafe++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of safe lines (part 1): " + strconv.Itoa(countSafe))

	// Part II
	f2, err2 := os.Open(input_file)
	if err2 != nil {
		log.Fatal(err2)
	}

	defer f2.Close()
	scanner2 := bufio.NewScanner(f2)

	var countSafe2 int
	for scanner2.Scan() {
		line := scanner2.Text()
		// Get array of fields (str)
		fields := strings.Fields(line)
		numbers := make([]int, len(fields))
		for i, s := range fields {
			numbers[i], _ = strconv.Atoi(s)
		}

		if checkSafe2(numbers) {
			countSafe2++
		}
	}
	if err := scanner2.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of safe lines (part 2): " + strconv.Itoa(countSafe2))
}

