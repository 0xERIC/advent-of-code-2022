package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stacks := make([][]byte, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		} else if strings.Count(line, "[") == 0 {
			continue
		}

		// Stacks are off by 1 (lower) relative to the instructions
		for stack := 0; stack <= len(line)/4; stack++ {
			if stack >= len(stacks) { // Only hit on the first iteration
				stacks = append(stacks, make([]byte, 0))
			}

			crate := line[stack*4+1]
			if crate != ' ' {
				stacks[stack] = append([]byte{crate}, stacks[stack]...)
			}
		}
	}

	for scanner.Scan() {
		parsed := strings.Split(scanner.Text(), " ")
		moveCount, _ := strconv.Atoi(parsed[1])
		moveSourceLabel, _ := strconv.Atoi(parsed[3]) // Off by 1; see above
		moveTargetLabel, _ := strconv.Atoi(parsed[5]) // OFf by 1; see above

		moveSource := &stacks[moveSourceLabel-1]
		moveTarget := &stacks[moveTargetLabel-1]

		for i := 0; i < moveCount; i++ {
			moved := (*moveSource)[len(*moveSource)-1]
			*moveSource = (*moveSource)[:len(*moveSource)-1]
			*moveTarget = append(*moveTarget, moved)
		}

		// Uncomment for part 1
		//
		// This just reverses the last moveCount crates in the target, which
		// imitates them having been moved as a unit (maintains order) rather
		// than 1 by 1 (flips order).
		for i, j := len(*moveTarget)-moveCount, len(*moveTarget)-1; i < j; i, j = i+1, j-1 {
			(*moveTarget)[i], (*moveTarget)[j] = (*moveTarget)[j], (*moveTarget)[i]
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, s := range stacks {
		fmt.Print(string(s[len(s)-1]))
	}
	fmt.Println()
}
