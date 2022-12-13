package main

import (
	"bufio"
	"fmt"
	"os"
)

// Set to 4 for part 1; 14 for part 2
const windowSize = 4

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		window := make(map[rune]int)
		for i, c := range line {
			_, ok := window[c]
			if !ok {
				window[c] = 0
			}

			window[c] = window[c] + 1

			if i >= windowSize {
				removed := line[i-windowSize]
				removedCount := window[rune(removed)] - 1
				window[rune(removed)] = removedCount

				if removedCount == 0 {
					delete(window, rune(removed))
				}
			}

			if len(window) == windowSize {
				fmt.Println(i + 1)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
