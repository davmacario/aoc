package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func blink(stones []string) []string {
	out := make([]string, 0)
	for _, s := range stones {
		len_s := len(s)
		if s == "0" {
			out = append(out, "1")
		} else if len_s%2 == 0 {
			p1 := s[:len_s/2]
			p2_int, _ := strconv.Atoi(s[len_s/2:])
			p2 := strconv.Itoa(p2_int)
			out = append(out, p1)
			out = append(out, p2)
		} else {
			p_int, _ := strconv.Atoi(s)
			p := strconv.Itoa(p_int * 2024)
			out = append(out, p)
		}
	}
	return out
}

func part1(stones []string) []string {
	// var count int
	upd_stones := stones
	for i := 0; i < 25; i++ {
		upd_stones = blink(upd_stones)
		// fmt.Println(upd_stones)
	}
	return upd_stones
}

func part2(stones []string) int {
	// var count int
	upd_stones := stones
	for i := 0; i < 75; i++ {
		fmt.Println(i + 1)
		upd_stones = blink(upd_stones)
		// fmt.Println(upd_stones)
	}
	return len(upd_stones)
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
	stones := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		stones = strings.Fields(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(stones)
	res1 := part1(stones)
	fmt.Println("Part 1:", len(res1))
	// Fuck it, multithread (don't wanna think too much rn)
	results := make(chan int, len(res1))
	var res2 int
	var wg sync.WaitGroup
	for i, n := range stones {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			result := part2([]string{n})
			results <- result
		}(i)
	}
    wg.Wait()
    close(results)
    for res:= range results {
        res2 += res
    }
	fmt.Println("Part 2:", res2)
}
