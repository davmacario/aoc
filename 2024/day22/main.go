package main

import (
	"bufio"
	ds "day22/data_structures"
	"day22/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	_ "strings"
)

func prune(a int) int {
	return utils.Mod(a, 16777216)
}

func mix(a, b int) int {
	return a ^ b
}

func getNextSecret(start int) int {
	// Mult by 64, mix with input and prune
	s := prune(mix(start*64, start))
	// Divide by 32, mix and prune
	s = prune(mix(s/32, s))
	// Mult by 2048, mix and prune
	s = prune(mix(s*2048, s))

	return s
}

func solve1(initial []int) (out int) {
	memoNext := make(map[int]int)
	var next int
	for _, prev := range initial {
		for i := 0; i < 2000; i++ {
			next_mem, ok := memoNext[prev]
			if ok {
				next = next_mem
			} else {
				next = getNextSecret(prev)
				memoNext[prev] = next
			}
			prev = next
		}
		out += prev
	}
	return out
}

// Appends the new element of the changes slice, and drops the first one if
// len > 4
func pushAndSlide(el int, sl []int) []int {
	sl = append(sl, el)
	for len(sl) > 4 {
		sl = sl[1:]
	}
	return sl
}

func solve2(initial []int) (out int) {
	memoNext := make(map[int]int)
	changesToPrice := make([]map[string]int, 0) // Map the encoding of changes slice to the associated price (for each element)
	stringToChanges := make(map[string][]int)   // Map a string to a slice of changes
	var next int
	for _, prev := range initial {
		changesSlice := make([]int, 0)
		currChToP := make(map[string]int)
		for i := 0; i < 2000; i++ {
			// fmt.Println(changesSlice)
			next_mem, ok := memoNext[prev]
			if ok {
				next = next_mem
			} else {
				next = getNextSecret(prev)
				memoNext[prev] = next
			}

			if i > 0 {
				// Can calculate change
				change := utils.Mod(next, 10) - utils.Mod(prev, 10)
				changesSlice = pushAndSlide(change, changesSlice)
			}

			if len(changesSlice) == 4 {
				changesSliceString := utils.SliceToStr(changesSlice)
				_, ok := currChToP[changesSliceString]
				// If current changesSliceString is already a valid key, don't update it (only care about 1st occurrence)
				if !ok {
					cpy := make([]int, 4)
					copy(cpy, changesSlice) // need to change reference
					stringToChanges[changesSliceString] = cpy
					currChToP[changesSliceString] = utils.Mod(next, 10)
				}
			} else if len(changesSlice) > 4 {
				panic("Too many elements in changes slice")
			}

			prev = next
		}
		changesToPrice = append(changesToPrice, currChToP)
		out += prev
	}

	// Get the unique keys of all maps in changesToPrice (use custom Set implementation)
	uniqueKeys := ds.NewSet([]string{})
	for i := 0; i < len(changesToPrice); i++ {
		for k := range changesToPrice[i] {
			uniqueKeys.Add(k)
		}
	}

	max_price := 0
	for _, k := range uniqueKeys.Elements() {
		curr_price := 0
		for i := 0; i < len(changesToPrice); i++ {
			p, ok := changesToPrice[i][k]
			if ok {
				curr_price += p
			}
		}
		if curr_price > max_price {
			max_price = curr_price
		}
	}

	return max_price
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	initialNums := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		n_curr, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Can't convert ", line, " to int")
		}
		initialNums = append(initialNums, n_curr)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ans1 := solve1(initialNums)
	fmt.Println("Part 1:", ans1)

	ans2 := solve2(initialNums)
	fmt.Println("Part 2:", ans2)
}
