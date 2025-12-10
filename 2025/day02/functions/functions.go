package functions

import (
	"log"
	"strconv"
	"strings"
	"day2/utils"
)

func CommonPrefix(a, b string) string {
	if len(a) != len(b) {
		return ""
	}

	prefix := ""
	i := 0
	for i < len(a) {
		if a[i] != b[i] {
			return prefix
		}
		prefix += string(a[i])
	}
	return a
}

// SplitFunc to pass to bufio.Scanner item (using Scanner.Split()) that allows
// to split the scanned string over commas (",").
func SplitCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the index of the input of a comma
	if i := strings.Index(string(data), ","); i >= 0 {
		return i + 1, data[0:i], nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), data, nil
	}

	return
}

// Given a string representing a range (e.g., "100-128"), returns the 2 numbers
func SplitRange(s string) (int, int) {
	numSlices := strings.Split(s, "-")
	if len(numSlices) != 2 {
		log.Fatalf("Expected 2 numbers, found %d!", len(numSlices))
	}
	a, err := strconv.Atoi(numSlices[0])
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(numSlices[1])
	if err != nil {
		log.Fatal(err)
	}
	return a, b
}

// Returns true if the number is invalid, i.e., it is made up of a number
// repeated twice, e.g., 123123, 67366736, 1111, 9091
func IsInvalid(n int) bool {
	numString := strconv.Itoa(n)
	numLen := len(numString)
	if numLen%2 != 0 {
		return false
	}
	halfway := numLen / 2

	return numString[:halfway] == numString[halfway:]
}

// Returns true if the number is invalid (for point 2), i.e., if it is made up
// of a repetition of a number any n. of times.
func IsInvalid2(n int) bool {
	if IsInvalid(n) {
		return true
	}

	numString := strconv.Itoa(n)
	numLen := len(numString)
	// i: number of times we want to divide the string. Upper bound is
	// the string length (1 chars)
	for i := 3; i <= numLen; i++ {
		if numLen%i != 0 {
			continue
		}
		if utils.AllEqualElements(utils.SplitStringEqual(numString, i)) {
			return true
		}
	}
	return false
}
