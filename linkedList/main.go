package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func buildList(size int) *LinkedList {
	if size <= 0 {
		return nil
	}

	var list LinkedList

	list.Head = &Node{
		Value: 0,
		Next:  nil,
	}

	currentNode := list.Head

	for i := 1; i < size; i++ {
		currentNode.Next = &Node{Value: i}
		currentNode = currentNode.Next
	}

	return &list
}

func lookup(list *LinkedList, value int) bool {
	if list == nil || list.Head == nil {
		return false
	}

	head := list.Head
	for head != nil {
		if head.Value == value {
			return true
		}
		head = head.Next
	}
	return false
}

func insertBack(list *LinkedList, value int) {
	if list == nil || list.Head == nil {
		return
	}

	head := list.Head
	for head.Next != nil {
		if head.Value == value {
			return
		}
		head = head.Next
	}
	head.Next = &Node{Value: value}
}

func insertFront(list *LinkedList, value int) {
	if list == nil || list.Head == nil {
		return
	}

	head := list.Head
	for head.Next != nil {
		if head.Next.Value == value {
			return
		}
		head = head.Next
	}

	list.Head = &Node{Value: value, Next: list.Head}
}

func printList(list *LinkedList) {
	if list == nil {
		return
	}
	head := list.Head
	for head != nil {
		fmt.Printf("%d ", head.Value)
		head = head.Next
	}
	fmt.Printf("\n")
}

func main() {

	myList := buildList(10)

	printList(myList)

	insertFront(myList, 10)
	insertFront(myList, 100)

	printList(myList)

	insertBack(myList, 20)
	insertBack(myList, 200)

	printList(myList)

	fmt.Println(lookup(myList, 20))
	fmt.Println(lookup(myList, 10))

}
