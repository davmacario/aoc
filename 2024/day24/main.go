package main

import (
	"bufio"
	ds "day24/data_structures"
	"day24/utils"
	"fmt"
	"log"
	"math"
	"os"
	_ "slices"
	"strconv"
	"strings"
)

type Operation struct {
	a, b, out, op string
}

func (o Operation) CanExec(valuesMap map[string]int) bool {
	_, ok_a := valuesMap[o.a]
	_, ok_b := valuesMap[o.b]
	return ok_a && ok_b
}

// Make sure the operands are sorted alphabetically
func (o *Operation) SortOperands() *Operation {
	if o.a > o.b {
		o.a, o.b = o.b, o.a
	}
	return o
}

func (o Operation) ToString() string {
	return o.a + " " + o.op + " " + o.b + " -> " + o.out
}

func (o Operation) CheckEqual(op2 Operation) bool {
	return o.out == op2.out && o.op == op2.op && (o.a == op2.a || o.a == op2.b) && (o.b == op2.b || o.b == op2.a)
}

// Needed to consider commuted operands
func (o Operation) InSlice(s []Operation) bool {
	o_cpy := Operation{a: o.b, b: o.a, op: o.op, out: o.out}
	return utils.GetIndex(o, s) != -1 || utils.GetIndex(o_cpy, s) != -1
}

// Returns the matching operation - needs to have the provided string as one of the factors
func FindOpByFactor(opList []Operation, factor string) (Operation, bool) {
	for _, op := range opList {
		if op.a == factor || op.b == factor {
			return op, true
		}
	}
	return Operation{}, false
}

// Given a list of operations, return the operation with that output
func FindOutput(sl []Operation, out string) (Operation, bool) {
	for _, op := range sl {
		if op.out == out {
			return op, true
		}
	}
	return Operation{}, false
}

func executeOP(a, b int, op string) int {
	switch op {
	case "AND":
		return a & b
	case "OR":
		return a | b
	case "XOR":
		return a ^ b
	default:
		panic("Invalid operation")
	}
}

func swapOutputs(op1, op2 *Operation) {
	op1.out, op2.out = op2.out, op1.out
}

// Find forbidden operations in current list of unique ops
func checkCurrBitOps(ops []Operation, ind int, revMap map[string]Operation) *ds.Set[Operation] {
	// out := make([]Operation, 0)
	outSet := ds.NewSet([]Operation{})
	ind_str := utils.ZeroPadLeft(strconv.Itoa(ind), 2)
	prev_ind := ind - 1
	prev_ind_str := utils.ZeroPadLeft(strconv.Itoa(prev_ind), 2)
	for _, op := range ops {
		// opsNoCurr := append(ops[:i], ops[i+1:]...)
		switch op.op {
		case "XOR": // Need 2 XORs
			if op.out == "z"+ind_str {
				if op.a[0] == 'x' || op.b[0] == 'y' {
					if ind > 0 {
						outSet.Add(op)
					}
				}
			} else {
                if op.a[0] != 'x' {
					outSet.Add(op)
                } else if op.a != "x"+ind_str || op.b != "y"+ind_str {
					// no need to add it - maybe the op itself is correct, just
					// misplaced because of wrong upstream connection
					// outSet.Add(op)
				}
			}
		case "AND":
			if op.out[0] == 'z' {
				outSet.Add(op)
			} else {
				// Another case: if upstream result is used in 'OR', which is wrong, don't add it
				upstreamOp, ok := FindOpByFactor(ops, op.out)
				if op.a[0] == 'x' && op.a[1:] != prev_ind_str {
					if !ok || !outSet.Contains(upstreamOp) {
						outSet.Add(op)
					}
				} else if op.b[0] == 'y' && op.b[1:] != prev_ind_str {
					if !ok || !outSet.Contains(upstreamOp) {
						outSet.Add(op)
					}
				}
			}
		case "OR":
			if op.out[0] == 'z' && ind != 45 {
				// Not @ output
				outSet.Add(op)
				break
			} else {
				next, ok := FindOpByFactor(ops, op.out)

				if !ok {
					if ind != 45 {
                        fmt.Println("--", ops)
                        fmt.Println(">", op)
						outSet.Add(op)
						break
					}
				} else {
					if ind != 45 && (next.op != "XOR" || next.out != "z"+ind_str) {
						outSet.Add(op)
						break
					}
				}

				// Its output must be inside the one operation with current z bit as output
				currOut := revMap["z"+ind_str]
				// In general this ^ will always be 1st element of slice.
				if outSet.Contains(currOut) {
					if op.out != currOut.a && op.out != currOut.b {
						// This OR is a wrong op
						outSet.Add(op)
						break
					}
				}
				// If here, conn to output (z) is good
				// Need to check that one of the factors is x(ind-1) AND y(ind-1)
				// opsMinusThis := append(ops[:i], ops[i+1:]...)
				child1, ok1 := revMap[op.a]
				child2, ok2 := revMap[op.b]
				if !(ok1 && ok2) { // Also catches case in which a or b contain x/y
					outSet.Add(op)
				} else {
					// Check that, if the child contains x and y bits, it is AND, and ind = ind-1
					if child1.op != "AND" {
						outSet.Add(child1)
					} else {
						if child1.a[0] == 'x' && child1.a[1:] != prev_ind_str {
							outSet.Add(child1)
						} else if child1.b[0] == 'y' && child1.b[1:] != prev_ind_str {
							outSet.Add(child1)
						}
					}

					if child2.op != "AND" {
						outSet.Add(child2)
					} else {
						if child2.a[0] == 'x' && child2.a[1:] != prev_ind_str {
							outSet.Add(child2)
						} else if child2.b[0] == 'y' && child2.b[1:] != prev_ind_str {
							outSet.Add(child2)
						}
					}
				}
			}

		}
	}
	return outSet
}

func checkOperations(ops []Operation, nBitsInput int) bool {
	// nBitsOutput := nBitsInput + 1
	allZeros := strings.Repeat("0", nBitsInput)
	yZeros := strings.Repeat("0", nBitsInput)
	yOnes := strings.Repeat("1", nBitsInput)
	for i := range nBitsInput {
		tmpMap := make(map[string]int)
		// Activate one bit at a time
		currXBin := utils.ReplaceCharInString(allZeros, nBitsInput-i-1, '1')
		currXInt, _ := strconv.ParseInt(currXBin, 2, 0)
		for i, n := range []string{yZeros, yOnes, ""} {
			op_cpy := make([]Operation, len(ops))
			copy(op_cpy, ops)
			ind := 0
			for ind < nBitsInput {
				c_k := "x" + utils.ZeroPadLeft(strconv.Itoa(ind), 2)
				_, ok := tmpMap[c_k]
				if ok {
					conv, _ := strconv.Atoi(string(currXBin[nBitsInput-1-ind]))
					tmpMap[c_k] = conv
					ind++
				} else {
					panic("Shouldn't be here")
				}
			}
			var yInt int64
			if i == 2 {
				ind = 0
				for ind < nBitsInput {
					c_k := "y" + utils.ZeroPadLeft(strconv.Itoa(ind), 2)
					_, ok := tmpMap[c_k]
					if ok {
						conv, _ := strconv.Atoi(string(currXBin[nBitsInput-1-ind]))
						tmpMap[c_k] = conv
						ind++
					} else {
						panic("Shouldn't be here")
					}
				}
				yInt = currXInt
			} else {
				ind = 0
				for ind < nBitsInput {
					c_k := "y" + utils.ZeroPadLeft(strconv.Itoa(ind), 2)
					_, ok := tmpMap[c_k]
					if ok {
						conv, _ := strconv.Atoi(string(n[nBitsInput-1-ind]))
						tmpMap[c_k] = conv
						ind++
					} else {
						panic("Shouldn't be here")
					}
				}
				yInt, _ = strconv.ParseInt(n, 2, 0)
			}
			targetResult := int(yInt + currXInt)
			targetResBin := utils.ZeroPadLeft(utils.IntToBin(targetResult), nBitsInput+1)

			evalOutput(tmpMap, op_cpy)

			// Get output string
			_, out_z_str := getOutput(tmpMap)

			// Figure out what are the wrong bits in the output
			// fmt.Println("", currXBin, "\n+\n", n, "\n=")
			// fmt.Println("A:", targetResBin)
			// fmt.Println("O:", out_z_str)
			rev_target := utils.ReverseString(utils.ZeroPadLeft(targetResBin, nBitsInput+1))
			rev_result := utils.ReverseString(utils.ZeroPadLeft(out_z_str, len(rev_target)))
			diffBits := utils.FindDifferences(rev_target, rev_result)
			if len(diffBits) > 0 {
				fmt.Println("> Flipped Bits:", diffBits)
				return false
			}
		}
		fmt.Println()
	}
	return true
}

// Approach: remove operations from the list once executed.
// Iterate on the list of operations - O(n^2) - for each one run .CanExec()
func solve1(known map[string]int, operations []Operation) int {
	for len(operations) > 0 {
		for i, op := range operations {
			if op.CanExec(known) {
				// fmt.Println(op)
				a := known[op.a]
				b := known[op.b]
				known[op.out] = executeOP(a, b, op.op)
				operations, _ = utils.RemoveFromSlice(operations, i)
				break
			}
		}
	}
	// Evaluate the output
	ind_z := 0
	out := 0
	for true {
		key := "z" + utils.ZeroPadLeft(strconv.Itoa(ind_z), 2)
		val, ok := known[key]
		if ok {
			out += val * int(math.Pow(2, float64(ind_z)))
			ind_z++
		} else {
			break
		}
	}
	return out
}

func getInt(prefix string, known map[string]int) (out int) {
	ind_z := 0
	for true {
		key := prefix + utils.ZeroPadLeft(strconv.Itoa(ind_z), 2)
		val, ok := known[key]
		if ok {
			out += val * int(math.Pow(2, float64(ind_z)))
			ind_z++
		} else {
			break
		}
	}
	return out
}

func evalOutput(known map[string]int, operations []Operation) {
	for len(operations) > 0 {
		for i, op := range operations {
			if op.CanExec(known) {
				// fmt.Println(op.ToString())
				a := known[op.a]
				b := known[op.b]
				known[op.out] = executeOP(a, b, op.op)
				operations, _ = utils.RemoveFromSlice(operations, i)
				break
			}
		}
	}
}

// Read the output of the z register; returns the integer in base 10 and the
// base 2 representation as string
func getOutput(known map[string]int) (int, string) {
	out := 0
	out_str := ""
	ind := 0
	for true {
		key := "z" + utils.ZeroPadLeft(strconv.Itoa(ind), 2)
		val, ok := known[key]
		if ok {
			out_str = strconv.Itoa(val) + out_str
			out += val * int(math.Pow(2, float64(ind)))
			ind++
		} else {
			break
		}
	}
	return out, out_str
}

// Ideas:
// - "sort" the operands in the operations alphabetically to help in visualization
// - Create list of values (10 random sums) - check which bits are wrong at the output
// - Isolate wrong bits and climb up the chain of operations
func solve2(known map[string]int, operations []Operation) (out string) {
	strOperations := make([]string, len(operations))
	for i, op := range operations {
		operations[i] = *op.SortOperands()
		strOperations[i] = op.ToString()
	}
	// slices.Sort(strOperations)
	nBitsInput := 0
	for k := range known {
		if k[0] == 'x' {
			nBitsInput++
		}
	}

	// ---------
	// Translate x__ and y__ to actual integers
	xInt := getInt("x", known)
	yInt := getInt("y", known)

	// ---------
	// Create map from result to operation - this way, if we know the wrong
	// bit, we can "climb up"
	reverseMap := make(map[string]Operation)
	for _, op := range operations {
		_, ok := reverseMap[op.out]
		if ok {
			panic("Duplicate operation???")
		}
		reverseMap[op.out] = op
	}

	// Climb up from each output bit and get the input bits on which it depends
	nBitsOutput := nBitsInput + 1
	// Use set for the operations that have been seen (identified by output)
	// This will only print the operations that have to be performed ONLY for
	// the current output bit (we are starting from the LSB)
	seenOperations := ds.NewSet([]string{})

	wrongCandidates := ds.NewSet([]Operation{})
	// prevCarry := ""
	for i := 0; i < nBitsOutput; i++ {
		uniqueOps := make([]Operation, 0)
		top_level := make([]Operation, 0)
		bit_index := utils.ZeroPadLeft(strconv.Itoa(i), 2)
		outBit := "z" + bit_index
		fmt.Println("Climbing up", outBit)
		// Go up BFS style
		qUp := ds.NewQueue([]string{})
		qUp.Push(outBit)
		ops := make(map[string]Operation) // Gather all operations related to this output bit
		for qUp.Length() > 0 {            // Assumes every value is output of only 1 op (reasonable)
			currEl, _ := qUp.Pop()
			if !seenOperations.Contains(currEl) {
				// Get operation having currEl as result
				opRes, ok := reverseMap[currEl]
				if !ok {
					panic("Operation not found...")
				}
				uniqueOps = append(uniqueOps, opRes)
				ops[currEl] = opRes
				if opRes.a[0] != 'x' && opRes.a[0] != 'y' {
					qUp.Push(opRes.a)
				}
				if opRes.b[0] != 'x' && opRes.b[0] != 'y' {
					qUp.Push(opRes.b)
				}
				if opRes.a[0] == 'x' || opRes.a[0] == 'y' || opRes.b[0] == 'x' || opRes.b[0] == 'y' {
					top_level = append(top_level, opRes)
				}
				fmt.Println(opRes)
				seenOperations.Add(currEl)
			}
		}
		// fmt.Println(top_level)
		outOfPlace := checkCurrBitOps(uniqueOps, i, reverseMap)
		fmt.Println("Out of place:", outOfPlace.Elements())
		wrongCandidates = ds.SetUnion(wrongCandidates, outOfPlace)
		fmt.Println()
	}

	fmt.Println("Candidates:\n", wrongCandidates.Elements())

	// Run down candidates, figure out best way to swap them

	// If nothing works: brute force

	// Analyze original operation
	targetOriginal := xInt + yInt
	fmt.Println("\nOriginal operation:", xInt, "+", yInt, "=", targetOriginal)
	targetOrigBits := utils.ZeroPadLeft(utils.IntToBin(targetOriginal), nBitsInput)
	fmt.Println("Target:", targetOrigBits)
	evalOutput(known, operations)
	outOriginal, outOrigBin := getOutput(known)
	flipped := utils.FindDifferences(outOrigBin, targetOrigBits)
	fmt.Println("Obtain:", outOrigBin)
	fmt.Println("Base 10:", outOriginal)
	fmt.Println("Flipped bits:", flipped)
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
	var readingOps bool
	knownValues := make(map[string]int)
	operations := make([]Operation, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			readingOps = true
		} else {
			fields := strings.Fields(line)
			if readingOps {
				a := fields[0]
				op := fields[1]
				b := fields[2]
				out := fields[4]
				operations = append(operations, Operation{a: a, b: b, op: op, out: out})
			} else {
				varName := fields[0][:len(fields[0])-1]
				varValue, err := strconv.Atoi(fields[1])
				if err != nil {
					panic("Unable to convert str to int")
				}
				knownValues[varName] = varValue
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	copy_known := utils.CopyMap(knownValues)
	copy_operations := make([]Operation, len(operations))
	copy(copy_operations, operations)
	ans1 := solve1(copy_known, copy_operations)
	fmt.Println("Part 1:", ans1)

	ans2 := solve2(knownValues, operations)
	fmt.Println("Part 2:", ans2)
}
