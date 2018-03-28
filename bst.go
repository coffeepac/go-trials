package main

import (
	"fmt"
	"github.com/coffeepac/go-trials/stack"
	"math/rand"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Bst interface {
	add(int) bool
	delete(int) bool
	find(int) int
	print_breath()
	print_depth()
}

func genInputs(size int) []int {
	intputs := make([]int, size)
	for i := 0; i < size; i++ {
		intputs[i] = rand.Intn(size)
	}
	return intputs
}

func (n *Node) add(value int) (added bool) {
	if value == n.Data {
		added = false
	} else {
		if value < n.Data {
			if n.Left != nil {
				return n.Left.add(value)
			} else {
				n.Left = &Node{Data: value}
				added = true
			}
		} else {
			if n.Right != nil {
				return n.Right.add(value)
			} else {
				n.Right = &Node{Data: value}
				added = true
			}
		}
	}
	return
}

func (n *Node) print_breath() {
	previous := -1
	var stack stack.Stack
	stack.Push(n)
	for stack.Head != nil {
		elem := stack.Pop()
		eNode := elem.Data.(*Node)
		if eNode.Left != nil {
			stack.Push(eNode.Left)
		}

		if eNode.Right != nil {
			stack.Push(eNode.Right)
		}

		if eNode.Data < previous {
			fmt.Println()
		}

		fmt.Print(eNode.Data, " ")
		previous = eNode.Data
	}
	fmt.Println()
}

func (n *Node) print_depth() {
	if n.Left != nil {
		n.Left.print_depth()
	}

	fmt.Print(n.Data, " ")

	if n.Right != nil {
		n.Right.print_depth()
	}
}

func main() {
	inputs := genInputs(200)
	fmt.Println("inputs: ", inputs)

	root := &Node{Data: inputs[0]}

	for _, val := range inputs {
		root.add(val)
	}

	fmt.Println("breath")
	root.print_breath()
	fmt.Println("depth")
	root.print_depth()
	fmt.Println()
}
