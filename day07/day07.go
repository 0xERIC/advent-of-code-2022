package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sizes := make(map[string]int)
	currentPath := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				if tokens[2] == "/" {
					currentPath = []string{}
				} else if tokens[2] == ".." {
					currentPath = currentPath[:len(currentPath)-1]
				} else {
					currentPath = append(currentPath, tokens[2])
				}
			}
		} else if tokens[0] == "dir" {
			// Do nothing; if it has contents we'll pick it up later
		} else {
			addedSize, _ := strconv.Atoi(tokens[0])

			for i := 0; i <= len(currentPath); i++ {
				path := strings.Join(currentPath[:i], "/")

				currentSize, ok := sizes[path]
				if !ok {
					sizes[path] = 0
				}

				sizes[path] = currentSize + addedSize
			}
		}
	}

	// Part 1
	//
	// totalSize := 0
	// for _, size := range sizes {
	// 	if size <= 100000 {
	// 		totalSize += size
	// 	}
	// }
	// fmt.Println(totalSize)

	// Part 2

	targetSize := sizes[""] - 40000000
	closestTarget := sizes[""]
	for _, size := range sizes {
		if size > targetSize && size < closestTarget {
			closestTarget = size
		}
	}
	fmt.Println(closestTarget)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
