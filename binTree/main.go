package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func buildTree(size int) *Node {
	if size < 1 {
		return nil
	}
	var TreeNode *Node

	for i := 0; i < size; i++ {
		TreeNode = insert(TreeNode, rand.Intn(100))
	}

	return TreeNode
}

func insert(treeNode *Node, value int) *Node {
	if treeNode == nil {
		return &Node{value, nil, nil}
	}

	if treeNode.Value == value {
		return treeNode
	}

	if treeNode.Value > value {
		treeNode.Left = insert(treeNode.Left, value)
		return treeNode
	}
	treeNode.Right = insert(treeNode.Right, value)
	return treeNode
}

func traverse(TreeNode *Node) {
	if TreeNode == nil {
		return
	}
	traverse(TreeNode.Left)
	fmt.Printf("%d ", TreeNode.Value)
	traverse(TreeNode.Right)
}

func main() {
	var myTree = buildTree(20)

	traverse(myTree)
}
