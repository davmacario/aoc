package main

import (
	"bufio"
	"day08/functions"
	"fmt"
	"log"
	"os"
	// "strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Point struct {
	x, y int
}

func (p Point) InSlice(sl []Point) bool {
	for _, o := range sl {
		if o == p {
			return true
		}
	}
	return false
}

func (p Point) InsideOfBounds(h int, w int) bool {
	if p.x < 0 || p.y < 0 || p.x >= w || p.y >= h {
		return false
	}
	return true
}

// NOTE
//  --> x
// |
// v
// y

// Given 2 points (should have the same char, unchecked), return the 2
// antinode positions
// Work with dx dy
func calcAntinodesFromPair(p1 Point, p2 Point) []Point {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	po1 := Point{x: p1.x + dx, y: p1.y + dy}
	po2 := Point{x: p2.x - dx, y: p2.y - dy}
	return []Point{po1, po2}
}

func calcAntinodesFromPair2(p1 Point, p2 Point, bounds Point) []Point {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	out := make([]Point, 0)
	// Harmonics in '+' direction
	po1 := Point{x: p1.x + dx, y: p1.y + dy}
	for po1.InsideOfBounds(bounds.y, bounds.x) {
		out = append(out, po1)
		po1 = Point{x: po1.x + dx, y: po1.y + dy}
	}
	// Harmonics in '-' direction
	po2 := Point{x: p2.x - dx, y: p2.y - dy}
	for po2.InsideOfBounds(bounds.y, bounds.x) {
		out = append(out, po2)
		po2 = Point{x: po2.x - dx, y: po2.y - dy}
	}
	return out
}

func findAntinodes(mat []string, mapChars map[rune][]Point) ([]Point, []Point) {
	out := make([]Point, 0)
	out2 := make([]Point, 0)
	h := len(mat)
	w := len(mat[0])
	dims := Point{x: w, y: h}
	matCpy := make([]string, h)
	_ = copy(matCpy, mat)
	// Iter on the subarrays for each rune
	for _, v := range mapChars {
		// Get all pairs
		for i := 0; i < len(v); i++ {
            // Add antenna position unless only 1 of that type exist
            if !v[i].InSlice(out2) && len(v) > 1 {
                out2 = append(out2, v[i])
            }
			for j := i + 1; j < len(v); j++ {
				// fmt.Println()
				antLoc := calcAntinodesFromPair(v[i], v[j])
				for _, pp := range antLoc { // always 2 iters
					if pp.InsideOfBounds(h, w) && !pp.InSlice(out) {
						// fmt.Println("Found antinode", pp)
						// matCpy[pp.y] = functions.ReplaceInStr(matCpy[pp.y], "#", pp.x)
						out = append(out, pp)
					}
				}

				antLoc2 := calcAntinodesFromPair2(v[i], v[j], dims)
				for _, pp := range antLoc2 {
					if !pp.InSlice(out2) {
						// fmt.Println("Found antinode", pp)
						if matCpy[pp.y][pp.x] == '.' {
							matCpy[pp.y] = functions.ReplaceInStr(matCpy[pp.y], "#", pp.x)
						}
						out2 = append(out2, pp)
					}
				}
			}
		}
	}
	functions.PrintMatrix(matCpy)
	return out, out2
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	mat := make([]string, 0)
	// Map each char to a list of points which have it
	mapChars := make(map[rune][]Point)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		for j, c := range line {
			if c != '.' {
				// Add it to the map
				sl := mapChars[c]
				mapChars[c] = append(sl, Point{x: j, y: i})
			}
		}
		mat = append(mat, line)
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(mapChars)
	antinodes1, antinodes2 := findAntinodes(mat, mapChars)
	ans1 := len(antinodes1)
	ans2 := len(antinodes2)
	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
