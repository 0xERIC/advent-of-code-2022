package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

// Basically stolen from:
// https://pkg.go.dev/container/heap#example-package-IntHeap

type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// type IntMinHeap

const numTopElves = 3 // 1 for part 1, 3 for part 2

func main() {
	currentCalories := 0
	topCalories := &IntMinHeap{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch nextLine := scanner.Text(); nextLine {
		case "":
			heap.Push(topCalories, currentCalories)
			if topCalories.Len() > numTopElves {
				heap.Pop(topCalories)
			}
			currentCalories = 0

		default:
			nextCalories, err := strconv.Atoi(nextLine)
			if err != nil {
				panic(err)
			}
			currentCalories += nextCalories
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	totalCalories := 0
	for topCalories.Len() > 0 {
		totalCalories += topCalories.Pop().(int)
	}
	fmt.Print(totalCalories)
}
