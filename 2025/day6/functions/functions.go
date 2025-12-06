package functions

import (
	"day6/utils"
	"log"
	"strconv"
)

func PerformOperation(nums []int, op string) int {
	var res int
	switch op {
	case "+":
		for _, n := range nums {
			res += n
		}
	case "*":
		res = 1
		for _, n := range nums {
			res *= n
		}
	}
	return res
}

// Returns true if for all line, element 'ind' is whitespace
func allIndexPointToSpace(lines []string, ind int) bool {
	for i := range len(lines) {
		if string(lines[i][ind]) != " " {
			return false
		}
	}
	return true
}

// Perform format conversion for part 2.
//
// Approach: travel all lines (containing numbers) simultaneously, using single
// index.
// The switch to a new operation corresponds to a column of all whitespaces.
func ConvertOperationFormat(lines []string) [][]int {
	nLines := len(lines)

	index := 0

	lineLengths := make([]int, nLines)
	for i := range lines {
		lineLengths[i] = len(lines[i])
	}
	maxLineLen, err := utils.MaxInSlice(lineLengths)

	if err != nil {
		log.Fatal(err)
	}

	// Pad lines to have all same length
	for i, l := range lines {
		lines[i] = utils.PadRight(l, maxLineLen, " ")
	}

	out := make([][]int, 0)
	for index < maxLineLen {
		numbersCurrOp := []int{}
		for index < maxLineLen && !allIndexPointToSpace(lines, index) {
			newNumberStr := ""
			for i := range nLines {
				currChar := string(lines[i][index])
				if currChar != " " {
					newNumberStr += currChar
				}
			}
			newNumberInt, err := strconv.Atoi(newNumberStr)
			if err != nil {
				log.Fatal(err)
			}
			numbersCurrOp = append(numbersCurrOp, newNumberInt)
			index += 1
		}
		index += 1

		out = append(out, numbersCurrOp)
	}

	return out
}
