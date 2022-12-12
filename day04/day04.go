package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	totalOverlaps := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		splitFunc := func(r rune) bool { return r == '-' || r == ',' }
		var pairs []int
		for _, s := range strings.FieldsFunc(scanner.Text(), splitFunc) {
			i, _ := strconv.Atoi(s)
			pairs = append(pairs, i)
		}

		if pairs[0] > pairs[3] || pairs[2] > pairs[1] { // no overlap
			continue
		}

		// Remove the check entirely for part 2
		//
		// NB: overlapLength is off by 1 compared to the counting in the
		// question, but it doesn't matter. E.g. we count 2-3 and 3-4 as
		// overlapping with "length" 0, but because we say that 2-3 only has
		// "length" 1, we're consistent and the answer is correct.
		overlapLength := min(pairs[1], pairs[3]) - max(pairs[0], pairs[2])
		if overlapLength == pairs[1]-pairs[0] || overlapLength == pairs[3]-pairs[2] {
			totalOverlaps += 1
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(totalOverlaps)
}
