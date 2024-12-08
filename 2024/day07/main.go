package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Pad the given string with '0' on the left until the desired length is reached
func PadZerosLhs(str string, length int) string {
    for len(str) < length {
        str = "0" + str
    }
    return str
}

// Determine if there exists a valid operation (+ or * only)
func ValidOperation1(nums []int, res int) bool {
	nOps := len(nums) - 1
	nAttempts := int(math.Pow(2, float64(nOps)))
	for i := 0; i < nAttempts; i++ {
		resTest := nums[0]
		binOps := PadZerosLhs(strconv.FormatInt(int64(i), 2), nOps)

		for j, o := range binOps {
			if o == '0' {
				resTest += nums[j+1]
			} else {
				resTest *= nums[j+1]
			}
			if resTest > res {
				break
			}
		}

		if resTest == res {
			// Stop at first match
			return true
		}
	}

	return false
}

func ValidOperation2(nums []int, res int) bool {
	nOps := len(nums) - 1
	nCombinations := int(math.Pow(3, float64(nOps)))
	for i := 0; i < nCombinations; i++ {
		resTest := nums[0]
		binOps := PadZerosLhs(strconv.FormatInt(int64(i), 3), nOps)

		for j, o := range binOps {
			if o == '0' {
				resTest += nums[j+1]
			} else if o == '1' {
				resTest *= nums[j+1]
			} else if o == '2' {
				lhs := strconv.Itoa(resTest)
				rhs := strconv.Itoa(nums[j+1])
				resTestStr, err := strconv.Atoi(lhs + rhs)
				checkErr(err)
				resTest = resTestStr
			} else {
				log.Fatal("Wrong type of operation: ", o)
			}
			if resTest > res {
				break
			}
		}

		if resTest == res {
			// Stop at first match
			return true
		}
	}

	return false
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	checkErr(err)

	defer f.Close()

	// Scanner to read the file
	scanner := bufio.NewScanner(f)

	// Iterate over lines of file
	var count1 int
	var count2 int
	for scanner.Scan() {
		line := scanner.Text()
		// Get array of fields (str)
		lineSlice := strings.Split(line, ":")
		res, err := strconv.Atoi(lineSlice[0])
		checkErr(err)
		strNums := strings.Fields(lineSlice[1])
		nums := make([]int, len(strNums))
		for i, ln := range strNums {
			nums[i], _ = strconv.Atoi(ln)
		}
		if ValidOperation1(nums, res) {
			count1 += res
		}
		if ValidOperation2(nums, res) {
			count2 += res
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", count1)
	fmt.Println("Part 2:", count2)
}
