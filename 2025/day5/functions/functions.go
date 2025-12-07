package functions

import (
	ds "day5/data_structures"
	"log"
	"slices"
	"strconv"
	"strings"
)

// Function passed to `slices.SortFunc` to sort ranges.
// Sorting is done based on start
func compareRanges(r1, r2 ds.Range[int]) int {
	return r1.Start - r2.Start
}

// Simplify ranges by merging overlapping ones.
//
// Strategy:
// 	1. Sort ranges by Start
// 	2. Iterate over sorted ranges. When the
func SimplifyRanges(ranges []ds.Range[int]) []ds.Range[int] {
	if len(ranges) <= 1 {
		return ranges
	}

	// Sort slice
	slices.SortFunc(ranges, compareRanges)

	out := []ds.Range[int]{ranges[0]}
	for i:=1; i<len(ranges); i++ {
		outTail := out[len(out) - 1]
		if ds.RangesOverlap(outTail, ranges[i]) {
			out[len(out) - 1] = ds.MergeRanges(outTail, ranges[i])
		} else {
			out = append(out, ranges[i])
		}
	}

	return out
}

func rangeFromString(s string) ds.Range[int] {
	split := strings.Split(s, "-")
	if len(split) != 2 {
		log.Fatal("Wrong range format")
	}

	startInt, err := strconv.Atoi(split[0])
	if err != nil {
		log.Fatal(err)
	}
	endInt, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}
	return ds.Range[int]{Start: startInt, End: endInt}
}

func RangesFromStrings(strSlice []string) []ds.Range[int] {
	out := make([]ds.Range[int], len(strSlice))

	for i, s := range strSlice {
		out[i] = rangeFromString(s)
	}

	return out
}
