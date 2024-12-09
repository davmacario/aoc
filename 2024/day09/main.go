package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	_ "strings"
)

// Returns the string 0..11..[...]
func getFileSystemString(in string) string {
	out := ""
	var diskID int
	for i, val := range in {
		n_blocks, _ := strconv.Atoi(string(val))
		if i%2 == 0 {
			// File
			for j := 0; j < n_blocks; j++ {
				out += string(strconv.Itoa(diskID))
			}
			diskID++
		} else {
			// Free space
			for j := 0; j < n_blocks; j++ {
				out += "."
			}
		}
	}
	return out
}

func removeDiskGaps(fileSystem string) string {

}

func main() {
	input_file := "./in_small.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Scanner to read the file
	scanner := bufio.NewScanner(f)

	// Iterate over lines of file
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(line)

	fileSystemStart := getFileSystemString(line)
	fmt.Println(fileSystemStart)
}
