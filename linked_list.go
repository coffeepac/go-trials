package main

import "fmt"
import "errors"

type Node struct {
	value int
	next  *Node
}

func delete(n *Node) error {
	if n.next == nil {
		return errors.New("No next node")
	}

	n.value = n.next.value
	n.next = n.next.next
	return nil
}

func printList(n *Node) {
	var cur *Node
	cur = n
	fmt.Print(cur.value)
	for cur.next != nil {
		cur = cur.next
		fmt.Print("->", cur.value)
	}
	fmt.Println()
}

func reverse(n *Node) *Node {
	cur := n
	next := n.next
	var prev *Node
	for next != nil {
		cur.next = prev
		prev = cur
		cur = next
		next = next.next
	}

	cur.next = prev

	return cur

}

func main() {
	root := &Node{value: 10}
	a := &Node{value: 1}
	root.next = a

	b := &Node{value: 2}
	a.next = b

	c := &Node{value: 3}
	b.next = c

	printList(root)

	rev := reverse(root)

	printList(rev)

	one := &Node{value: 666}
	printList(one)
	rone := reverse(one)
	printList(rone)
}
