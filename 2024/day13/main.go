package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ClawMachine struct {
	count, x, y int
	M           map[byte][]int
}

func (m ClawMachine) solve() (n_a, n_b int, ok bool) {
	mat := make([][]int, 2)
	for _, c := range "AB" {
		curr_arr := m.M[byte(c)]
		mat[0] = append(mat[0], curr_arr[0])
		mat[1] = append(mat[1], curr_arr[1])
	}
	// fmt.Println(mat, "= [", m.x, m.y, "]")
	det_mat := mat[0][0]*mat[1][1] - mat[1][0]*mat[0][1]
	// fmt.Println("Determinant:", det_mat)
	if det_mat == 0 {
		ok = false
	} else {
		ok = true
		n_a = (mat[1][1]*m.x - mat[0][1]*m.y) / det_mat
		n_b = (-mat[1][0]*m.x + mat[0][0]*m.y) / det_mat
		if (mat[0][0]*n_a+mat[0][1]*n_b) != m.x || (mat[1][0]*n_a+mat[1][1]*n_b) != m.y {
			ok = false
		}
		if n_a < 0 || n_b < 0 {
			ok = false
		}
	}
	return n_a, n_b, ok
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var sol1, sol2 int
	current_machine := ClawMachine{count: 0, M: make(map[byte][]int)}
	current_machine2 := ClawMachine{count: 0, M: make(map[byte][]int)}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			// fmt.Println(line)
			sliceSc := strings.Split(line, ": ")
			if sliceSc[0][:5] == "Prize" {
				if sliceSc[0] != "Prize" {
					log.Fatal("Expected 'Prize', got ", sliceSc[0])
				}
				xy := strings.Split(sliceSc[1], ", ")
				str_x := strings.Split(xy[0], "=")[1]
				current_machine.x, _ = strconv.Atoi(str_x)
				current_machine2.x = current_machine.x + 10000000000000
				str_y := strings.Split(xy[1], "=")[1]
				current_machine.y, _ = strconv.Atoi(str_y)
				current_machine2.y = current_machine.y + 10000000000000
			} else {
				buttonID := sliceSc[0][len(sliceSc[0])-1]
				xy := strings.Split(sliceSc[1], ", ")
				str_x := strings.Split(xy[0], "+")[1]
				val_x, _ := strconv.Atoi(str_x)
				str_y := strings.Split(xy[1], "+")[1]
				val_y, _ := strconv.Atoi(str_y)
				current_machine.M[buttonID] = []int{val_x, val_y}
				current_machine2.M[buttonID] = []int{val_x, val_y}
			}
		} else {
			// Solve current machine
			n_a, n_b, ok := current_machine.solve()
			if ok {
				fmt.Println("Machine", current_machine.count, "is winnable:", n_a, "x A,", n_b, "x B")
				sol1 += 3*n_a + n_b
			} else {
				fmt.Println("Machine", current_machine.count, "is NOT winnable:", n_a, "x A,", n_b, "x B")
			}
			n_a2, n_b2, ok2 := current_machine2.solve()
			if ok2 {
				fmt.Println("Machine", current_machine2.count, "is winnable:", n_a2, "x A,", n_b2, "x B")
				sol2 += 3*n_a2 + n_b2
			} else {
				fmt.Println("Machine", current_machine2.count, "is NOT winnable:", n_a2, "x A,", n_b2, "x B")
			}
			fmt.Println("")
			// Create new machine
			current_machine = ClawMachine{count: current_machine.count + 1, M: make(map[byte][]int)}
			current_machine2 = ClawMachine{count: current_machine2.count + 1, M: make(map[byte][]int)}
		}
	}

	n_a, n_b, ok := current_machine.solve()
	if ok {
		fmt.Println("Machine", current_machine.count, "is winnable:", n_a, "x A,", n_b, "x B")
		sol1 += 3*n_a + n_b
	} else {
		fmt.Println("Machine", current_machine.count, "is NOT winnable:", n_a, "x A,", n_b, "x B")
	}
	n_a2, n_b2, ok2 := current_machine2.solve()
	if ok2 {
		fmt.Println("Machine", current_machine2.count, "is winnable:", n_a2, "x A,", n_b2, "x B")
		sol2 += 3*n_a2 + n_b2
	} else {
		fmt.Println("Machine", current_machine2.count, "is NOT winnable:", n_a2, "x A,", n_b2, "x B")
	}
	fmt.Println("")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", sol1)
	fmt.Println("Part 2:", sol2)
}
