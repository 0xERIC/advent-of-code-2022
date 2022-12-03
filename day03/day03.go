package main

import (
	"bufio"
	"fmt"
	"os"
)

func getPriority(char byte) int {
	if char >= 97 && char <= 122 { // lower
		return int(char) - 96
	} else { // upper
		return int(char) - 38
	}
}

const elfGroupSize = 3 // 1 for part 1, 3 for part 2

func main() {
	elfIndex := 0

	m := make(map[byte]int)
	totalPriority := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if elfIndex == 0 {
			m = make(map[byte]int)
		}

		rucksack := scanner.Text()

		// Part 1

		// for i := 0; i < len(rucksack)/2; i += 1 {
		// 	m[rucksack[i]] = 0
		// }
		// for j := len(rucksack) / 2; j < len(rucksack); j += 1 {
		// 	c := rucksack[j]

		// 	_, found := m[c]
		// 	if found {
		// 		totalPriority += getPriority(c)
		// 		break
		// 	}
		// }

		// Part 2

		for i := 0; i < len(rucksack); i += 1 {
			c := rucksack[i]

			lastElfIndex, found := m[c]
			if found && lastElfIndex == elfIndex-1 {
				m[c] = elfIndex
			} else if elfIndex == 0 {
				m[c] = 0
			}
		}

		if elfIndex == elfGroupSize-1 {
			for c := range m {
				if m[c] == elfGroupSize-1 {
					totalPriority += getPriority(c)
				}
			}
		}

		elfIndex = (elfIndex + 1) % elfGroupSize
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(totalPriority)
}
