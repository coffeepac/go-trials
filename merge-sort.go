package main

import (
	"fmt"
	"math/rand"
)

func buildSlice(size int) []int {
	bSlice := make([]int, size)
	for i := 0; i < size; i++ {
		bSlice[i] = rand.Intn(size)
	}
	return bSlice
}

func sortSlice(input []int) []int {
	if len(input) > 1 {
		splitter := len(input) / 2
		lsclice := sortSlice(input[:splitter])
		rsclice := sortSlice(input[splitter:])

		tempSlice := make([]int, len(input))
		i, l, r := 0, 0, 0
		for i < len(input) && l < len(lsclice) && r < len(rsclice) {
			if lsclice[l] <= rsclice[r] {
				tempSlice[i] = lsclice[l]
				l++
				i++
			} else {
				tempSlice[i] = rsclice[r]
				r++
				i++
			}
		}

		if l < len(lsclice) {
			for l < len(lsclice) {
				tempSlice[i] = lsclice[l]
				l++
				i++
			}
		}

		if r < len(rsclice) {
			for r < len(rsclice) {
				tempSlice[i] = rsclice[r]
				r++
				i++
			}
		}

		copy(input, tempSlice)
	}

	return input
}

func main() {
	slice := buildSlice(10)
	fmt.Println("input:", slice)
	slice = sortSlice(slice)

	fmt.Println("sorted:", slice)
}
