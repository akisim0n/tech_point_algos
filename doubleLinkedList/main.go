package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type DoubleLinkedList struct {
	Head *Node
}

func buildDoubleList(size int) *DoubleLinkedList {
	if size <= 0 {
		return nil
	}

	var list DoubleLinkedList

	list.Head = &Node{
		Value: 0,
		Next:  nil,
		Prev:  nil,
	}

	currentNode := list.Head

	for i := 1; i < size; i++ {
		currentNode.Next = &Node{Value: i, Prev: currentNode, Next: nil}
		currentNode = currentNode.Next
	}

	return &list
}

func lookup(list *DoubleLinkedList, value int) bool {
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

func insert(list *DoubleLinkedList, value int) {
	if list == nil || list.Head == nil {
		return
	}

	if list.Head.Value > value {
		list.Head = &Node{Value: value, Prev: nil, Next: list.Head}
		return
	} else if list.Head.Value == value {
		fmt.Printf("Value: %d exists in list\n", value)
		return
	}

	head := list.Head
	for head.Next != nil && head.Next.Value <= value {
		if head.Next.Value == value {
			fmt.Printf("Value: %d exists in list\n", value)
			return
		}
		head = head.Next
	}

	temp := &Node{Value: value, Prev: head, Next: head.Next}
	if head.Next != nil {
		head.Next.Prev = temp
	}
	head.Next = temp
}

func deleteNode(list *DoubleLinkedList, value int) {
	if list == nil || list.Head == nil {
		return
	}

	head := list.Head
	for head != nil {
		if head.Value == value {
			if head.Next != nil {
				head.Next.Prev = head.Prev
			}
			if head.Prev != nil {
				head.Prev.Next = head.Next
			} else {
				list.Head = head.Next
			}
			return
		}
		head = head.Next
	}

	fmt.Printf("Value: %d does not exist in list\n", value)
}

func insertBack(list *DoubleLinkedList, value int) {
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
	head.Next = &Node{Value: value, Prev: head, Next: nil}
}

func insertFront(list *DoubleLinkedList, value int) {
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

	list.Head = &Node{Value: value, Next: list.Head, Prev: nil}
}

func printList(list *DoubleLinkedList) {
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

	myList := buildDoubleList(10)

	printList(myList)

	insert(myList, 10)
	insert(myList, 100)

	printList(myList)

	insert(myList, 20)
	insert(myList, 200)
	insert(myList, 0)
	insert(myList, -5)
	insert(myList, 11)

	printList(myList)

	deleteNode(myList, 10)
	deleteNode(myList, -5)
	deleteNode(myList, 200)

	printList(myList)

	fmt.Println(lookup(myList, 20))
	fmt.Println(lookup(myList, 10))

}
