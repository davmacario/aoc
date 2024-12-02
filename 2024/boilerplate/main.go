package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
	for scanner.Scan() {
		line := scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
