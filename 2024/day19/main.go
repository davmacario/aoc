package main

import (
	"bufio"
	"day19/functions"
	"fmt"
	"log"
	"os"
	"strings"
)

// Returns true if the current design is valid, given the available towels
// Recursion strategy:
func isValidDesign(design string, patterns []string) (out bool) {
	for _, p := range patterns {
		if functions.StartsWith(design, p) {
			// fmt.Println(design, "starts with", p)
			if len(p) < len(design) {
				if isValidDesign(design[len(p):], patterns) {
					return true
				}
			} else {
				return true
			}
		}
	}
	return false
}

// Similar approach to isValidDesign. Instead, count the total number of
// valid patterns for the current design
func countTotWays(design string, patterns []string, memo map[string]int) (out int) {
	if memo[design] != 0 {
		return memo[design]
	}
	for _, p := range patterns {
		if functions.StartsWith(design, p) {
			// fmt.Println(design, "starts with", p)
			if len(p) < len(design) {
				out += countTotWays(design[len(p):], patterns, memo)
			} else {
				out++
			}
		}
	}
	// fmt.Println(">", design, "can be made in", out, "ways")
	memo[design] = out
	return out
}

// Returns the number of patterns that are valid out of the design
func solve1(availPatterns, designs []string) (out int) {
	for _, d := range designs {
		if isValidDesign(d, availPatterns) {
			fmt.Println(d, "is valid")
			out++
		} else {
			fmt.Println(d, "is not valid")
		}
	}
	return out
}

func solve2(availPatterns, designs []string) (out int) {
	// Store n. for patterns
	memo := make(map[string]int)
	for _, d := range designs {
		n_valid := countTotWays(d, availPatterns, memo)
		out += n_valid
		// fmt.Println("")
	}
	return out
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	got_patterns := false
	availPatterns := make([]string, 0)
	designs := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			got_patterns = true
		} else {
			if !got_patterns {
				availPatterns = strings.Split(line, ", ")
			} else {
				designs = append(designs, line)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Available towels:", availPatterns)

	ans1 := solve1(availPatterns, designs)
	fmt.Println("Part 1:", ans1)
	ans2 := solve2(availPatterns, designs)
	fmt.Println("Part 2:", ans2)
}
