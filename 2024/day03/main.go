package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func evalExpr(expr string) int {
	re1 := regexp.MustCompile(`^mul\((\d+),`)
	re2 := regexp.MustCompile(`,(\d+)\)$`)

	val1 := re1.FindStringSubmatch(expr)[1]
	val2 := re2.FindStringSubmatch(expr)[1]

	if val1 == "" || val2 == "" {
		log.Fatal("Expression did not match")
	}
	int1, _ := strconv.Atoi(val1)
	int2, _ := strconv.Atoi(val2)
	return int1 * int2
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
	var res1 int
	for scanner.Scan() {
		line := scanner.Text()
		subLine := line
		var startInd int
		// Use regex to match expression `mul(n,m)`
		re := regexp.MustCompile(`mul\(\d+,\d+\)`)
		for startInd < len(line) {
			subLine = subLine[startInd:]
			inds := re.FindStringIndex(subLine)
			if inds != nil {
				start, end := inds[0], inds[1]
				expr := subLine[start:end]
				// Eval matched expr
				res1 += evalExpr(expr)
				startInd = end
			} else {
				startInd = len(line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result (1): ", res1)

	// Part 2
	f2, err2 := os.Open(input_file)
	if err2 != nil {
		log.Fatal(err2)
	}

	defer f2.Close()

	// Scanner to read the file
	scanner2 := bufio.NewScanner(f2)

	// Iterate over lines of file
	var res2 int
    doEnabled := true // Flag to decide whether to add or not
	for scanner2.Scan() {
		line := scanner2.Text()
		subLine := line
		var startInd int
		// Use regex to match expression `mul(n,m)`
		re := regexp.MustCompile(`mul\(\d+,\d+\)`)
		re_do := regexp.MustCompile(`do\(\)`)
		re_dont := regexp.MustCompile(`don't\(\)`)

		for startInd < len(line) {
			subLine = subLine[startInd:]

			if doEnabled {
				// Look for both expr and dont, consider the one with earlier start
				inds := re.FindStringIndex(subLine)
				inds_dont := re_dont.FindStringIndex(subLine)

				if inds_dont != nil && inds != nil {
					if inds[0] < inds_dont[0] {
                        fmt.Println("Found mul")
						start, end := inds[0], inds[1]
						expr := subLine[start:end]
						fmt.Println(" -> ", expr)
						// Eval matched expr
						res2 += evalExpr(expr)
						startInd = end
					} else {
                        fmt.Println("Found don't")
						_, end := inds_dont[0], inds_dont[1]
						doEnabled = false
						startInd = end
					}
				} else if inds_dont != nil {
                    fmt.Println("Found don't")
					doEnabled = false
					// No mul expression anymore, just skip to eol
					startInd = len(line)
				} else if inds != nil {
                    fmt.Println("Found mul")
					start, end := inds[0], inds[1]
					expr := subLine[start:end]
					fmt.Println(expr)
					// Eval matched expr
					res2 += evalExpr(expr)
					startInd = end
				} else {
					startInd = len(line)
				}
			} else {
				// Look for do
				inds_do := re_do.FindStringIndex(subLine)
				if inds_do != nil {
                    fmt.Println("Found do")
					doEnabled = true
					_, end := inds_do[0], inds_do[1]
					startInd = end
				} else {
					startInd = len(line)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result (2): ", res2)
}
