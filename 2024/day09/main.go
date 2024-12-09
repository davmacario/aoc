package main

import (
	"bufio"
	"day09/functions"
	"fmt"
	"log"
	"os"
	"strconv"
	_ "strings"
)

// Returns the string 0..11..[...]
func getFileSystemString(in string) []string {
	out := make([]string, 0)
	var diskID int
	for i, val := range in {
		n_blocks, _ := strconv.Atoi(string(val))
		if i%2 == 0 {
			// File
			for j := 0; j < n_blocks; j++ {
				out = append(out, string(strconv.Itoa(diskID)))
			}
			diskID++
		} else {
			// Free space
			for j := 0; j < n_blocks; j++ {
				out = append(out, ".")
			}
		}
	}
	return out
}

func removeDiskGaps1(fileSystem []string) []string {
	i := 0
	j := len(fileSystem) - 1

	for i < j {
		if string(fileSystem[i]) == "." && string(fileSystem[j]) != "." {
			functions.SwapChars(fileSystem, i, j)
			// fmt.Println(fileSystem)
			i++
			j--
		} else {
			if string(fileSystem[i]) != "." {
				i++
			}
			if string(fileSystem[j]) == "." {
				j--
			}
		}
	}
	return fileSystem
}

// Return index of 1st gap of length l in fileSystem starting from startPos
// Returns -1 if not found
func findGap(fileSystem []string, startPos int, l int) int {
	ind := startPos
	for ind < len(fileSystem)-l {
		foundGap := true
		var j int
		for j = 0; j < l; j++ {
			if fileSystem[ind+j] != "." {
				// fmt.Println("nogap", ind, j, fileSystem[ind+j])
				foundGap = false
				break
			}
		}
		if foundGap {
			return ind
		}
		ind += j + 1
	}
	return -1
}

func removeDiskGaps2(fileSystem []string) []string {
	i := 0
	j := len(fileSystem) - 1

	// Find last file (last file: highest id)
	var file_id string
	for fileSystem[j] == "." {
		j--
	}
	file_id = fileSystem[j]
	file_id_int, _ := strconv.Atoi(file_id)

	for file_id_int >= 0 {
		// Go back until target file is found
		for j >= 0 && fileSystem[j] != file_id {
			j--
		}
		// Get curr file length
		j_e := j
		for j_e >= 0 && fileSystem[j_e] == file_id {
			j_e--
		}
		len_des := j - j_e // Desired block length
		j_e++

		// i - find first gap in filesystem that can fit the desired length
		i = findGap(fileSystem, 0, len_des)
		if i < 0 || i > j {
			// No gap of that size "earlier"
			i = 0
			// decrease j (to first empty block before)
			j = j_e - 1
		} else {
			// Found space,can swap
			for k := 0; k < len_des; k++ {
				functions.SwapChars(fileSystem, i+k, j-k)
			}
            // Jump backwards
			j = j_e - 1
		}
        // Update next target file
		file_id = strconv.Itoa(file_id_int - 1)
		file_id_int--
	}
	return fileSystem
}

func calcChecksum1(fileSystem []string) int64 {
	var out int64
	var i int
	for string(fileSystem[i]) != "." {
		curr_int, _ := strconv.Atoi(string(fileSystem[i]))
		out += (int64(curr_int) * int64(i))
		i++
	}
	return out
}

func calcChecksum2(fileSystem []string) int64 {
	var out int64
	for i := 0; i < len(fileSystem); i++ {
		if fileSystem[i] != "." {
			curr_int, _ := strconv.Atoi(string(fileSystem[i]))
			out += (int64(curr_int) * int64(i))
		}
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
	// fmt.Println(line)

	fileSystemStart := getFileSystemString(line)
	// fmt.Println(fileSystemStart)

	fileSystemSorted1 := removeDiskGaps1(fileSystemStart)
	ans1 := calcChecksum1(fileSystemSorted1)
	fmt.Println("Part 1:", ans1)

	fileSystemStart = getFileSystemString(line)
	fileSystemSorted2 := removeDiskGaps2(fileSystemStart)
	ans2 := calcChecksum2(fileSystemSorted2)
	fmt.Println("Part 2:", ans2)
}
