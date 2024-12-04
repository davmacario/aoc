package main

import (
	"bufio"
	"day04/functions"
	"fmt"
	"log"
	"os"
	// "strings"
)

// Directions where to move
var N = [2]int{-1, 0}
var NE = [2]int{-1, 1}
var E = [2]int{0, 1}
var SE = [2]int{1, 1}
var S = [2]int{1, 0}
var SW = [2]int{1, -1}
var W = [2]int{0, -1}
var NW = [2]int{-1, -1}
var Dirs = [8][2]int{N, NE, E, SE, S, SW, W, NW}

var Word = []string{"X", "M", "A", "S"}

// Given the matrix and the position of a letter `X`, explore the letters
// around to find the word XMAS.
// It returns the number of XMAS words found for this X
func exploreAround(mat []string, start [2]int) int {
	if string(mat[start[0]][start[1]]) != Word[0] {
		log.Fatal("First letter is not ", Word[0])
	}

	h := len(mat)
	w := len(mat[0])
	dim_x := [2]int{0, w}
	dim_y := [2]int{0, h}
	var count_words int
	for _, direction := range Dirs {
		points := make([][]int, len(Word))
		for i := range Word {
			points[i] = append(points[i], functions.SumArr(start[:], functions.MultArray(direction[:], i))...)
		}
		for i, pos := range points {
			if !functions.InRange(pos[0], dim_y) || !functions.InRange(pos[1], dim_x) || string(mat[pos[0]][pos[1]]) != Word[i] {
				break
			} else if i == len(points)-1 {
				count_words++
			}
		}
	}
	return count_words
}

// Returns 1 if current A is center of X-MAS
func exploreAroundMAS(mat []string, start [2]int) int {
	y0, x0 := start[0], start[1]
	diag_1 := [3]string{string(mat[y0+1][x0-1]), "A", string(mat[y0-1][x0+1])}
	diag_2 := [3]string{string(mat[y0-1][x0-1]), "A", string(mat[y0+1][x0+1])}
    //fmt.Println(diag_1)
    //fmt.Println(diag_2)

	if (diag_1 == [3]string{"M", "A", "S"} || diag_1 == [3]string{"S", "A", "M"}) &&
		(diag_2 == [3]string{"M", "A", "S"} || diag_2 == [3]string{"S", "A", "M"}) {
		return 1
	}
	return 0
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
	var matrix []string

	// Iterate over lines of file
	for scanner.Scan() {
		line := scanner.Text()
		// Get array of fields (str)
		var row string
		row = line
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	h := len(matrix)
	w := len(matrix[0])
	fmt.Println("Dims:", w, h)

	var count1 int
	// Iterate over values, look for "X"
	for i, row := range matrix {
		for j, val := range row {
			if string(val) == "X" {
				pos := [2]int{i, j}
				count1 += exploreAround(matrix, pos)
			}
		}
	}

	// functions.DispMatrix(matrix)
	fmt.Println("Part 1 result: ", count1)

	// Part 2
	// Similar approach, find "A", then determine if X-MAS
	var count2 int
	for i, row := range matrix[1 : len(matrix)-1] {
		for j, val := range row[1 : len(row)-1] {
			if string(val) == "A" {
				pos := [2]int{i + 1, j + 1}
				count2 += exploreAroundMAS(matrix, pos)
			}
		}
	}
	fmt.Println("Part 2 result: ", count2)
}
