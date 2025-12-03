package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Head *Node
	Size int
}

func buildQueue(size int) *Queue {

	if size <= 0 {
		return nil
	}

	var queue Queue

	queue.Head = &Node{Value: 0, Next: nil}
	queue.Size++

	currentNode := queue.Head
	for i := 1; i < size; i++ {
		currentNode.Next = &Node{Value: i, Next: nil}
		queue.Size++
		currentNode = currentNode.Next
	}
	return &queue
}

func pop(queue *Queue) *Node {
	if queue.Head == nil {
		return nil
	}

	retNode := queue.Head
	queue.Head = queue.Head.Next
	queue.Size--

	return retNode
}

func push(queue *Queue, value int) {

	if queue == nil {
		return
	}

	currentNode := queue.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = &Node{Value: value, Next: nil}
	queue.Size++
}

func printQueue(queue *Queue) {
	if queue.Head == nil {
		return
	}

	currentNode := queue.Head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println()
}

func main() {

	myQueue := buildQueue(10)

	printQueue(myQueue)

	push(myQueue, 100)
	push(myQueue, 200)

	printQueue(myQueue)

	removedEl := pop(myQueue)

	fmt.Println("removed el:", removedEl.Value)

	printQueue(myQueue)

}
