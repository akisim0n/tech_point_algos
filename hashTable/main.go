package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func buildTable(table *HashTable, size int) *HashTable {
	if size < 1 {
		return nil
	}
	table.Size = size
	table.Table = make(map[int]*Node, size)
	for i := 0; i < size; i++ {
		insert(table, rand.Intn(20))
	}

	return table
}

func insert(table *HashTable, value int) {
	index := value % table.Size

	if table.Table[index] == nil {
		table.Table[index] = &Node{value, nil}
		return
	}

	if table.Table[index].Value == value {
		return
	}

	head := table.Table[index]
	for head.Next != nil {
		if head.Value == value {
			return
		}
		head = head.Next
	}
	head.Next = &Node{value, nil}
}

func lookup(table *HashTable, value int) bool {
	index := value % table.Size

	if table.Table[index] == nil {
		return false
	}
	head := table.Table[index]

	for head != nil {
		if head.Value == value {
			return true
		}
		head = head.Next
	}
	return false
}

func printTable(table *HashTable) {
	for key, node := range table.Table {
		fmt.Printf("For key %d: ", key)
		for node != nil {
			fmt.Printf("%d,", node.Value)
			node = node.Next
		}
		fmt.Println()
	}
}

func main() {

	var myTable HashTable

	buildTable(&myTable, 10)

	printTable(&myTable)

	fmt.Println(lookup(&myTable, 1))
	fmt.Println(lookup(&myTable, 3))
	fmt.Println(lookup(&myTable, 5))
	fmt.Println(lookup(&myTable, 7))

}
