package main

import (
	"bufio"
	"day17/functions"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solve(input []int, reg map[string]int) []int {
	var instPointer int
	out_list := make([]int, 0)
	for instPointer < len(input)-1 {
		// fmt.Print(instPointer)
		opcode := input[instPointer]
		lit_op := input[instPointer+1]

		var combo int
		switch lit_op {
		case 4:
			combo = reg["A"]
		case 5:
			combo = reg["B"]
		case 6:
			combo = reg["C"]
		default:
			combo = lit_op
		}

		// var res int
		switch opcode {
		case 0:
			reg["A"] = reg["A"] / int(math.Pow(2, float64(combo)))
			// res = reg["A"]
			instPointer += 2
		case 1:
			reg["B"] = reg["B"] ^ lit_op
			// res = reg["B"]
			instPointer += 2
		case 2:
			reg["B"] = functions.Mod(combo, 8)
			// res = reg["B"]
			instPointer += 2
		case 3:
			if reg["A"] > 0 {
				instPointer = lit_op
				// res = input
			} else {
				instPointer += 2
			}
		case 4:
			reg["B"] = reg["B"] ^ reg["C"]
			// res = reg["B"]
			instPointer += 2
		case 5:
			out_list = append(out_list, functions.Mod(combo, 8))
			// res = functions.Mod(input, 8)
			instPointer += 2
		case 6:
			reg["B"] = reg["A"] / int(math.Pow(2, float64(combo)))
			// res = reg["B"]
			instPointer += 2
		case 7:
			reg["C"] = reg["A"] / int(math.Pow(2, float64(combo)))
			// res = reg["C"]
			instPointer += 2
		default:
			panic("Should not be here!")
		}
		// fmt.Println(" > operation", opcode, "- Operand:", operand, ":", input, ";", reg, "--- Result", res)
	}
	return out_list
}

/*
* 2,4
* 1,2
* 7,5
* 4,7
* 1,3
* 5,5 --> print
* 0,3
* 3,0
*
*2,4,1,2,7,5,4,7,1,3,5,5,0,3,3,0
* 0330553174572142 base8
*
* Comp. by hand: 119137285677840 b10
* = 3305531745721420 b8
 */

func formatOutput(out_sl []int) string {
	out := ""
	for i, n := range out_sl {
		out += strconv.Itoa(n)
		if i < len(out_sl)-1 {
			out += ","
		}
	}
	return out
}

func main() {
	input_file := "./in.txt"
	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var linenr int
	// var ind int
	register := make(map[string]int)
	input_list := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(line) > 0 {
			if linenr < 3 {
				reg_id := string(fields[1][0])
				val, err := strconv.Atoi(fields[len(fields)-1])
				if err != nil {
					log.Fatal(err)
				}
				register[reg_id] = val
				// if reg_id == "A" {
				// 	ind = val
				// }
			} else {
				for _, n := range strings.Split(fields[1], ",") {
					int_val, _ := strconv.Atoi(n)
					input_list = append(input_list, int_val)
				}
			}
			linenr++
		}
	}
	for k, v := range register {
		fmt.Println(string(k), ": ", v)
	}
	fmt.Println(input_list)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// log.Fatal()

	res1 := solve(input_list, register)
	fmt.Println("Part 1:", formatOutput(res1))
	fmt.Println(res1)
	fmt.Println(strconv.FormatInt(int64(35200350), 8))
	fmt.Println("")

	oct_string := "10355"
	fmt.Println(input_list)
	i_dec, err := strconv.ParseInt(oct_string, 8, 0)
	fmt.Println("A =", i_dec)
	if err != nil {
		log.Fatal(err)
	}

	register["A"] = int(i_dec)
	register["B"] = 0
	register["C"] = 0
	sol_new := solve(input_list, register)
	fmt.Println(sol_new)

	// log.Fatal()

	// The number of octal digits of ans2 is the same as the number of digits in the target output (== n. in input)
	// Appending digits to the octal input, we prepend digits at the output (the others are NOT changed)
	// IDEA: start building the number from the last digits of the output
	n_digits_octal2 := len(input_list)
	out_str_octal := ""
	// curr_digit_number := 1
    start_digit := 0
	for len(out_str_octal) < n_digits_octal2 {
		// Iterate on all octal numbers with 1 digits (0 to 7)
		len_curr := len(out_str_octal)
		// curr_n_iter, err := strconv.ParseInt(functions.MultString("7", curr_digit_number), 8, 0)
		if err != nil {
			log.Fatal(err)
		}
		for i := start_digit; i < 8; i++ {
		    fmt.Println(out_str_octal, i)
            i_oct := strconv.FormatInt(int64(i), 8)
			i_dec, err := strconv.ParseInt(out_str_octal+i_oct, 8, 0)
			if err != nil {
				log.Fatal(err)
			}

			register["A"] = int(i_dec)
			register["B"] = 0
			register["C"] = 0
			sol_new := solve(input_list, register)
			// fmt.Println(sol_new)
			// fmt.Println(input_list[len(input_list)-len(out_str_octal)-1:])
			if slices.Equal(sol_new, input_list[len(input_list)-len(out_str_octal)-1:]) {
				out_str_octal = out_str_octal + strconv.Itoa(i)
				fmt.Println(out_str_octal)
				fmt.Println("matching")
                start_digit = 0
                break
			}
		}
		if len(out_str_octal) == len_curr {
            // Remove last digit
            out_str_octal = out_str_octal[:len(out_str_octal)-1]
            // Grab second to last digit, it will be starting point for iter
            start_digit, _ = strconv.Atoi(string(out_str_octal[len(out_str_octal)-1]))
            fmt.Println("HERE")
		}
	}

	res2, err := strconv.ParseInt(out_str_octal, 8, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 2:", res2)
	fmt.Println(input_list)
	register["A"] = int(res2)
	register["B"] = 0
	register["C"] = 0
	fmt.Println(solve(input_list, register))

	// start := 119137285669000
	// start := 100000
	// for i:=start; i<start+9000; i++ {
	//     register := make(map[string]int)
	//     register["A"] = i
	//     register["B"] = 0
	//     register["C"] = 0
	//     sol_new := solve(input_list, register)
	//     fmt.Println(i, sol_new)
	//     if slices.Equal(sol_new, input_list) {
	//         fmt.Println("Found sol")
	//         break
	//     }
	// }
	// res2 := res1
	// fmt.Println(input_list)
	// fmt.Println(res2)
	// for !slices.Equal(input_list, res2) {
	// 	register := make(map[string]int)
	// 	if functions.Mod(ind, 1000000) == 0 {
	// 		fmt.Println(ind)
	// 	}
	// 	register["A"] = ind
	// 	register["B"] = 0
	// 	register["C"] = 0
	// 	// fmt.Println(ind, solve(input_list, register))
	// 	res2 = solve(input_list, register)
	// 	ind++
	// }
}
