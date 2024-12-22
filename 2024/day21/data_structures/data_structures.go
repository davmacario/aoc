package datastructures

import (
	"log"
	"math"
)

// Utils
func sign(a int) int {
	if a > 0 {
		return +1
	} else if a < 0 {
		return -1
	} else {
		return 0
	}
}

func intAbs(a int) int {
	return a * sign(a)
}

// -----------------------------------------------------------------------------
// Stack implementation in Golang
// -----------------------------------------------------------------------------
type Stack[T any] struct {
	s []T
}

// Returns the number of elements in the stack
func (s Stack[T]) Length() int {
	return len(s.s)
}

// Push an element on top of the stack
func (s *Stack[T]) Push(el T) {
	s.s = append(s.s, el)
}

// Pops the top element of the stack
func (s *Stack[T]) Pop() (T, bool) {
	if s.Length() <= 0 {
		var out T
		return out, false
	}
	out := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return out, true
}

// Returns the top element of the stack, without popping it
func (s *Stack[T]) Top() (T, bool) {
	if s.Length() <= 0 {
		var out T
		return out, false
	}
	return s.s[len(s.s)-1], true
}

// -----------------------------------------------------------------------------
// Queue (FIFO) implementation in Golang
// -----------------------------------------------------------------------------
type Queue[T any] struct {
	s []T
}

// Returns the number of elements in the queue
func (s Queue[T]) Length() int {
	return len(s.s)
}

// Push an element on top of the queue
func (s *Queue[T]) Push(el T) {
	s.s = append(s.s, el)
}

// Pops the top element of the queue
func (s *Queue[T]) Pop() (T, bool) {
	if s.Length() <= 0 {
		var out T
		return out, false
	}
	out := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return out, true
}

// Returns the top element of the queue, without popping it
func (s *Queue[T]) Top() (T, bool) {
	if s.Length() <= 0 {
		var out T
		return out, false
	}
	return s.s[len(s.s)-1], true
}

// -----------------------------------------------------------------------------
// Heap implementation in Golang
// -----------------------------------------------------------------------------
type HeapItem struct {
	p  int
	el interface{}
}

type Heap struct {
	s          []HeapItem
	comparator func(int, int) bool // Returns true if elem1 is >|< elem2
}

func NewHeap(comparator func(int, int) bool) *Heap {
	return &Heap{
		s:          []HeapItem{},
		comparator: comparator,
	}
}

func (h Heap) Length() int {
	return len(h.s)
}

func (h Heap) IsEmpty() bool {
	return h.Length() == 0
}

func swapSliceItems[T any](sl []T, a, b int) {
	ln := len(sl)
	if a >= ln || b >= ln {
		panic("Index outside of bounds")
	}
	if a == b {
		return
	}
	sl[b], sl[a] = sl[a], sl[b]
}

func parentInd(i int) int {
	return (i - 1) / 2
}

func childrenInd(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}

func (h *Heap) Insert(priority int, item interface{}) {
	h.s = append(h.s, HeapItem{p: priority, el: item})
	ind := len(h.s) - 1
	parent_ind := parentInd(ind)
	for parent_ind >= 0 && ind > 0 && h.comparator(h.s[ind].p, h.s[parent_ind].p) {
		swapSliceItems(h.s, ind, parent_ind)
		ind = parent_ind
		parent_ind = parentInd(ind)
	}
}

// Restore heap property by rearranging the slice elements accordingly
// Call this function after popping
func (h *Heap) Heapify() {
	ind := 0
	for ind < len(h.s) {
		// Look at both children, get max
		c1, c2 := childrenInd(ind)
		next_ind := ind
		if c1 < len(h.s) && h.comparator(h.s[c1].p, h.s[ind].p) {
			next_ind = c1
		}
		if c2 < len(h.s) && h.comparator(h.s[c2].p, h.s[next_ind].p) {
			next_ind = c2
		}

		if next_ind != ind {
			swapSliceItems(h.s, ind, next_ind)
			ind = next_ind
		} else {
			return
		}
	}
}

// Remove high/low-est priority item
func (h *Heap) Pop() (HeapItem, bool) {
	if len(h.s) <= 0 {
		return HeapItem{}, false
	}
	out := h.s[0]
	h.s = h.s[1:]
	h.Heapify()
	return out, true
}

// Read heap root
func (h *Heap) Top() (HeapItem, bool) {
	if len(h.s) <= 0 {
		return HeapItem{}, false
	}
	return h.s[0], true
}

// -----------------------------------------------------------------------------
// Points - 2D integer arrays
// -----------------------------------------------------------------------------
type Point struct {
	x, y int
}

func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

func (p *Point) Set(x, y int) {
	p.x = x
	p.y = y
}

func (p *Point) SetX(x int) {
	p.x = x
}

func (p *Point) SetY(y int) {
	p.y = y
}

func (p *Point) GetX() int {
	return p.x
}

func (p *Point) GetY() int {
    return p.y
}

func (p Point) MoveInDir(d Dir) Point {
	return Point{x: p.x + d.x, y: p.y + d.y}
}

func (p Point) InsideBounds(w, h int) bool {
	if p.x < 0 || p.y < 0 || p.x >= w || p.y >= h {
		return false
	}
	return true
}

// Get slice of points at "Manhattan" distance `d` from `p`
func (p Point) GetPointsAtDist(d int) []Point {
	totPts := 4 * d
	pts := make([]Point, 0)
	for i := 0; i <= d; i++ {
		j := d - i
		pts = append(pts, Point{x: p.x + i, y: p.y + j})
		if i > 0 && j > 0 {
			pts = append(pts, Point{x: p.x - i, y: p.y - j})
		}
		if i > 0 {
			pts = append(pts, Point{x: p.x - i, y: p.y + j})
		}
		if j > 0 {
			pts = append(pts, Point{x: p.x + i, y: p.y - j})
		}
	}
	if len(pts) != totPts {
		log.Fatal("Expected ", totPts, " elements, found ", len(pts))
	}
	return pts
}

func (p Point) GetPointsAtDistBounds(d, w, h int) []Point {
	i := 0
	pts := p.GetPointsAtDist(d)
	for i < len(pts) {
		if !pts[i].InsideBounds(w, h) {
			pts = append(pts[:i], pts[i+1:]...)
		} else {
			i++
		}
	}
	return pts
}

// Return the direction to follow to move from p to q.
// Equivalent to q-p
func (p Point) GetDir(q Point) Dir {
	return Dir{x: q.x - p.x, y: q.y - p.y}
}

func GetCharInPoint(mat []string, p Point) string {
	return string(mat[p.y][p.x])
}

// L1 norm of the distance
func ManhattanDist(p1, p2 Point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}


// -----------------------------------------------------------------------------
// Directions - 2D integer arrays
// -----------------------------------------------------------------------------

type Dir struct {
	x, y int
}

func NewDir(x, y int) Dir {
	return Dir{x: x, y: y}
}

func (d Dir) Opposite() Dir {
	return Dir{x: -d.x, y: -d.y}
}

var Up = Dir{x: 0, y: 1}
var Right = Dir{x: 1, y: 0}
var Down = Dir{x: 0, y: -1}
var Left = Dir{x: -1, y: 0}
var Dirs = []Dir{Left, Up, Right, Down}

var DirToStr = map[Dir]rune{
	Up:    '^',
	Down:  'v',
	Left:  '<',
	Right: '>',
}

var ChToDir = map[rune]Dir{
	'^': Up,
	'v': Down,
	'<': Left,
	'>': Right,
}
