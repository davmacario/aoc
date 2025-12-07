package datastructures

import (
	"errors"
	"fmt"
	"log"
	"math"
)

// -----------------------------------------------------------------------------
// Utils
// -----------------------------------------------------------------------------

// Replace `i`-th character of string `s` with `c`
func replaceCharInString(s string, i int, c rune) string {
	rowRunes := []rune(s)
	rowRunes[i] = c
	return string(rowRunes)
}

// Useful interface

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64
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
	el any
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

func (h *Heap) Insert(priority int, item any) {
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
	X, Y int
}

func MakePoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p *Point) Set(x, y int) {
	p.X = x
	p.Y = y
}

func (p *Point) SetX(x int) {
	p.X = x
}

func (p *Point) SetY(y int) {
	p.Y = y
}

func (p Point) MoveInDir(d Dir) Point {
	return Point{X: p.X + d.X, Y: p.Y + d.Y}
}

// Note:
//	- w: X dimennsion
//	- h: Y dimenstion
func (p Point) InsideBounds(w, h int) bool {
	if p.X < 0 || p.Y < 0 || p.X >= w || p.Y >= h {
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
		pts = append(pts, Point{X: p.X + i, Y: p.Y + j})
		if i > 0 && j > 0 {
			pts = append(pts, Point{X: p.X - i, Y: p.Y - j})
		}
		if i > 0 {
			pts = append(pts, Point{X: p.X - i, Y: p.Y + j})
		}
		if j > 0 {
			pts = append(pts, Point{X: p.X + i, Y: p.Y - j})
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
	return Dir{X: q.X - p.X, Y: q.Y - p.Y}
}

func GetCharInPoint(mat any, p Point) (string, error) {
	switch m := mat.(type) {
	case []string:
		if !p.InsideBounds(len(m[0]), len(m)) {
			err := errors.New("Point is outside bounds of provided matrix")
			return "", err
		}
		return string(m[p.Y][p.X]), nil
	case [][]string:
		if !p.InsideBounds(len(m[0]), len(m)) {
			err := errors.New("Point is outside bounds of provided matrix")
			return "", err
		}
		return m[p.Y][p.X], nil
	default:
		return "", fmt.Errorf("Unsupported type for mat (%T)", m)
	}
}

// Given a matrix `mat` and a Point `p`, set the value of `mat` in `p` to `c`.
// This function operates in-place.
func SetCharInPoint(mat any, p Point, c string) error {
	if len([]rune(c)) != 1 {
		return errors.New("c must be a single character")
	}
	r := []rune(c)[0]

	switch m := mat.(type) {
	case []string:
		h := len(m)
		if h == 0 {
			return errors.New("matrix has no rows")
		}
		rowRunes := []rune(m[p.Y])
		w := len(rowRunes)
		if !p.InsideBounds(w, h) {
			return fmt.Errorf("Point %v is outside bounds (w=%d, h=%d)", p, w, h)
		}
		m[p.Y] = replaceCharInString(m[p.Y], p.X, r)
		return nil
	case [][]string:
		h := len(m)
		if h == 0 {
			return errors.New("matrix has no rows")
		}
		w := len(m[0])
		if !p.InsideBounds(w, h) {
			return fmt.Errorf("Point %v is outside bounds (w=%d, h=%d)", p, w, h)
		}
		m[p.Y][p.X] = string(r)
		return nil
	default:
		return fmt.Errorf("Unsupported type for mat (%T)", m)
	}
}

// L1 norm of the distance
func ManhattanDist(p1, p2 Point) int {
	return int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
}

// -----------------------------------------------------------------------------
// Directions - 2D integer arrays
// -----------------------------------------------------------------------------

type Dir struct {
	X, Y int
}

func MakeDir(x, y int) Dir {
	return Dir{X: x, Y: y}
}

func (d Dir) Opposite() Dir {
	return Dir{X: -d.X, Y: -d.Y}
}

var Up = Dir{X: 0, Y: -1}
var Right = Dir{X: 1, Y: 0}
var Down = Dir{X: 0, Y: 1}
var Left = Dir{X: -1, Y: 0}
var Dirs = []Dir{Left, Up, Right, Down}
var DirsAll = []Dir{Left, Up, Right, Down, {X: 1, Y: 1}, {X: -1, Y: 1}, {X: 1, Y: -1}, {X: -1, Y: -1}}

// -----------------------------------------------------------------------------
// Set
// -----------------------------------------------------------------------------
// Set implementation in Go - basically Maps from T to bool; true if in set,
// else false

type Set[T comparable] struct {
	S map[T]bool
}

func (s Set[T]) Length() int {
	return len(s.S)
}

func NewSet[T comparable](sl []T) *Set[T] {
	init_s := make(map[T]bool)
	out := &Set[T]{S: init_s}
	for _, k := range sl {
		out.S[k] = true
	}
	return out
}

func (s *Set[T]) GetFirstKey() (T, bool) {
	for k, v := range s.S {
		if v {
			return k, true
		}
	}
	var t T // Empty
	return t, false
}

func (s *Set[T]) Add(el T) {
	s.S[el] = true
}

func (s *Set[T]) Remove(el T) {
	delete(s.S, el)
}

func (s *Set[T]) Contains(el T) bool {
	_, exists := s.S[el]
	return exists
}

func (s *Set[T]) Elements() []T {
	elements := make([]T, 0, len(s.S))
	for elem := range s.S {
		elements = append(elements, elem)
	}
	return elements
}

func SetUnion[T comparable](s, t *Set[T]) *Set[T] {
	union := NewSet([]T{})
	for k := range s.S {
		union.Add(k)
	}
	for k := range t.S {
		union.Add(k)
	}
	return union
}

func SetIntersect[T comparable](s, t *Set[T]) *Set[T] {
	intersection := NewSet([]T{})
	if len(t.S) > len(s.S) {
		s, t = t, s
	}
	for k := range s.S {
		if t.S[k] {
			intersection.Add(k)
		}
	}
	return intersection
}

// -----------------------------------------------------------------------------
// Integer Ranges
// -----------------------------------------------------------------------------

type Range[T Number] struct {
	Start, End T
}

func (r Range[T]) Span() T {
	return r.End - r.Start + 1
}

// Returns True if ranges overlap
func RangesOverlap[T Number](r1, r2 Range[T]) bool {
	if r1.Start >= r2.Start && r1.Start <= r2.End {
		return true
	}
	if r2.Start >= r1.Start && r2.Start <= r1.End {
		return true
	}
	if r1.End >= r2.Start && r1.End <= r2.End {
		return true
	}
	if r2.End >= r1.Start && r2.End <= r1.End {
		return true
	}
	return false
}

// Merge ranges `r1` and `r2`.
//
// Note: does not check for overlapping!
func MergeRanges[T Number](r1, r2 Range[T]) Range[T] {
	return Range[T]{Start: min(r1.Start, r2.Start), End: max(r1.End, r2.End)}
}

func (r Range[T]) Contains(n T) bool {
	return n >= r.Start && n <= r.End
}
