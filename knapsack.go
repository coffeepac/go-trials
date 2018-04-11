package main

import "fmt"

type cake struct {
	weight int
	value  int
}

func max_duffel_bag_value(cakes []cake, capacity int) int {
	maxValues := make(map[int]int, capacity+1)
	for i := 0; i <= capacity; i++ {
		for _, cake := range cakes {
			if cake.weight <= i {
				//for j := (i - cake.weight); j >= 0 && j >= (i-2*cake.weight); j-- {
				if (cake.value + maxValues[i-cake.weight]) > maxValues[i] {
					maxValues[i] = cake.value + maxValues[i-cake.weight]
				}
				//}
			}
		}
	}
	fmt.Println(maxValues)
	return maxValues[capacity]
}

func main() {
	cake_tuples := make([]cake, 3)
	cake_tuples[0] = cake{weight: 3, value: 20}
	cake_tuples[1] = cake{weight: 5, value: 70}

	fmt.Println("Max value: ", max_duffel_bag_value(cake_tuples, 9))
}
