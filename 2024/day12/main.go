package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	_ "strings"
)

type Pos struct {
	x, y int
}

func (xy Pos) moveInDir(d Dir) Pos {
	return Pos{x: xy.x + d.x, y: xy.y + d.y}
}

func (xy Pos) inBounds(w, h int) bool {
	if xy.x < 0 || xy.y < 0 || xy.x >= w || xy.y >= h {
		return false
	}
	return true
}

// a -> b determines direction
type Side struct {
	a, b Pos
}

func (s Side) getSideDir() Dir {
	return Dir{x: s.b.x - s.a.x, y: s.b.y - s.a.y}
}

func totLen(mat [][]Side) (out int) {
	for _, r := range mat {
		out += len(r)
	}
	return out
}

// Return an "ordered" list of sides, i.e., for 2 consecutive items x and y,
// x.b = y.a
func sortSideList(in []Side) [][]Side {
	out := make([][]Side, 0)
	ext_ind := 0
	sideFlag := make(map[Side]bool)
	for i, s := range in {
		if i == 0 {
			sideFlag[s] = true
		} else {
			sideFlag[s] = false // set to true if already seen
		}
	}
	out = append(out, []Side{in[0]})
	for totLen(out) < len(in) {
		broken := false
		lookingFor := out[ext_ind][len(out[ext_ind])-1].b
		for _, s := range in {
			if s.a == lookingFor && !sideFlag[s] {
				sideFlag[s] = true
				out[ext_ind] = append(out[ext_ind], s)
				broken = true
				break
			}
		}
		if !broken {
			// Close current loop
			// Look for non-flagged element, add it to next out element
			// -> It is an additional boundary (e.g., internal)
			var first_not_taken Side
			for _, s := range in {
				if !sideFlag[s] {
					first_not_taken = s
					break
				}
			}
			out = append(out, []Side{first_not_taken})
			sideFlag[first_not_taken] = true
			ext_ind++
		}
	}
	return out
}

func getNumSides(in []Side) (count int) {
	if len(in) < 1 {
		log.Fatal("Not enough sides")
	}
	in_s := sortSideList(in)
	for ext_ind := 0; ext_ind < len(in_s); ext_ind++ {
		for i, s1 := range in_s[ext_ind] {
			// If direction changes, then add 1 side
			s2 := in_s[ext_ind][(i+1)%len(in_s[ext_ind])]
			d1 := s1.getSideDir()
			d2 := s2.getSideDir()

			if d1.opposite() == d2 {
				log.Fatal("The list of sides is not ordered. d1: ", d1, "; d2: ", d2, "; current item: ", i, "/", len(in_s[ext_ind])-1)
			} else if d1 != d2 {
				// Change in dir
				count++
			}
		}
	}
	return count
}

type Point struct {
	xy    Pos
	d_arr int
	plant string
}

// NOTE: returns p itself (updates in-place)
func (p *Point) setPlant(pl string) {
	p.plant = pl
}

func (p *Point) setArrDir(d int) {
	p.d_arr = d
}

func (p Point) moveInDir(d Dir) Point {
	newPos := p.xy.moveInDir(d)
	return Point{xy: newPos, d_arr: p.d_arr, plant: p.plant}
}

func (p Point) inBounds(w, h int) bool {
	return p.xy.inBounds(w, h)
}

type Region struct {
	letter string
	index  int
	points []*Point
}

/*
 * Idea: traverse region using a stack (dfs)
 */

type PointStack []Point

func (s *PointStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *PointStack) Push(p Point) {
	*s = append(*s, p)
}

func (s *PointStack) Pop() (Point, bool) {
	l := len(*s)
	if s.IsEmpty() {
		return Point{}, false
	}
	out := (*s)[l-1]
	*s = (*s)[:l-1]
	return out, true
}

type Dir struct {
	x, y int
}

func (d Dir) opposite() Dir {
	return Dir{x: -d.x, y: -d.y}
}

var up = Dir{x: 0, y: -1}
var right = Dir{x: 1, y: 0}
var down = Dir{x: 0, y: 1}
var left = Dir{x: -1, y: 0}

// This defines the visit order of the cells nearby
// If arriving from `up`, start visiting from dir 0
// In general, if arriving from direction d (int), explore like:
//
//	for k := 0; k < len(visitDirs); k++ {
//	  next_dir_ind := (d+k) % len(visitDirs)
//	  next_dir := visitDirs[next_dir_ind]
//	  arr_dir := (next_dir_ind + 2) % len(visitDirs) // Next `d`
//	}
var visitDirs = []Dir{up, right, down, left}

func getSide(pt Pos, d_i int) Side {
	a := pt
	b := pt.moveInDir(right)
	c := pt.moveInDir(right).moveInDir(down)
	d := pt.moveInDir(down)
	switch d_i {
	case 0:
		return Side{a: a, b: b}
	case 1:
		return Side{a: b, b: c}
	case 2:
		return Side{a: c, b: d}
	case 3:
		return Side{a: d, b: a}
	}
	log.Fatal("Direction index should be in range [0, 3] - ", d_i)
	return Side{}
}

func printVisited(garden []string, visited map[Pos]bool) {
	for j := 0; j < len(garden); j++ {
		for i := 0; i < len(garden[0]); i++ {
			if visited[Pos{x: i, y: j}] {
				fmt.Print(string(garden[j][i]))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func floodFill(garden []string, start_pt Point, visited map[Pos]bool) (area, perimeter, n_sides int) {
	h := len(garden)
	w := len(garden[0])
	inCurrRegion := make(map[Pos]bool)
	sides := make([]Side, 0)
	visitStack := make(PointStack, 0)
	visitStack.Push(start_pt)
	for len(visitStack) > 0 {
		curr_pt, ok := visitStack.Pop()
		if !ok {
			log.Fatal("Tried to pop from empty stack!")
		}
		seen, ok := visited[curr_pt.xy]
		if !ok || !seen {
			// fmt.Println("New point ", curr_pt)
			visited[curr_pt.xy] = true
			inCurrRegion[curr_pt.xy] = true
			d := curr_pt.d_arr
			// is_curr_bound := false
			area++
			for k := 0; k < len(visitDirs); k++ {
				next_dir_ind := (d + k) % len(visitDirs)
				next_dir := visitDirs[next_dir_ind]
				next_pt := curr_pt.moveInDir(next_dir)
				arr_dir := (next_dir_ind + 2) % len(visitDirs) // Next `d`
				if next_pt.inBounds(w, h) && string(garden[next_pt.xy.y][next_pt.xy.x]) == curr_pt.plant {
					next_pt.setPlant(curr_pt.plant)
					next_pt.setArrDir(arr_dir)
					if !visited[next_pt.xy] {
						visitStack.Push(next_pt)
					}
				} else { // Boundary in `next_dir` direction
					perimeter++
					newSide := getSide(curr_pt.xy, next_dir_ind)
					sides = append(sides, newSide)
				}
			}

		}
	}
	n_sides = getNumSides(sides)
	return area, perimeter, n_sides
}

func explore(garden []string) (res1, res2 int) {
	h := len(garden)
	w := len(garden[0])
	visited := make(map[Pos]bool)

	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			curr_pos := Pos{x: i, y: j}
			curr_pt := Point{xy: curr_pos, plant: string(garden[j][i])}
			seen, ok := visited[curr_pos]
			if !ok || !seen {
				area, perimeter, n_sides := floodFill(garden, curr_pt, visited)
				res1 += area * perimeter
				res2 += area * n_sides
			}
		}
	}
	return res1, res2
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	garden := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		garden = append(garden, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	res1, res2 := explore(garden)
	fmt.Println("Part1:", res1)
	fmt.Println("Part2:", res2)
}
