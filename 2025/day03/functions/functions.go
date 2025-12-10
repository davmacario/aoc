package functions

import (
	"day3/utils"
	"log"
	"strconv"
)

// Returns the highest achievable "joltage" by selecting 2 digits in the given
// slice. Order is respected.
// Idea: first digit is max in 0:n-1, second digit is max in `<ind_first>:n`
func PickTwoBatteries(bank []int) int {
	digit1, ind1, err := utils.MaxWithFirstIndex(bank[:len(bank)-1])
	if err != nil {
		log.Fatal(err)
	}
	digit2, _, err := utils.MaxWithFirstIndex(bank[ind1+1:])
	if err != nil {
		log.Fatal(err)
	}
	numStr := strconv.Itoa(digit1) + strconv.Itoa(digit2)
	out, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func PickTwelveBatteries(bank []int) int {
	if len(bank) < 12 {
		log.Fatal("Provided bank contains less than 12 elements")
	}
	digits := make([]int, 12)
	ind := 0
	bankLen := len(bank)
	for i := range 12 {
		d, new_ind, err := utils.MaxWithFirstIndex(bank[ind : bankLen-(11-i)])
		if err != nil {
			log.Fatal(err)
		}
		digits[i] = d
		ind = ind + new_ind + 1
	}
	return utils.MergeIntSlice(digits)
}
