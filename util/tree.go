package util

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

type Tree struct {
	Root *TreeNode
}

// Generates a tree (complete or semi-complete) from an array.
func ArrayToTree(arr []int, i int) (root *TreeNode) {
	if i < 0 {
		panic("Index cannot be less than 0")
	}
	if i >= len(arr) {
		return nil
	}
	root = NewLeaf(arr[i])
	root.Left = ArrayToTree(arr, 2*i+1)
	root.Right = ArrayToTree(arr, 2*i+2)
	return root
}

// Creates a new leaf for a tree.
func NewLeaf(val int) *TreeNode {
	return &TreeNode{nil, nil, val}
}

// Prints a Tree in one row.
func PrintTree(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Value, "")
	if node.Left == nil && node.Right == nil {
		return
	}
	fmt.Print(" [")
	if node.Left != nil {
		PrintTree(node.Left)
	} else {
		fmt.Print("Null")
	}
	fmt.Print(", ")
	if node.Right != nil {
		PrintTree(node.Right)
	} else {
		fmt.Print("Null")
	}
	fmt.Print("]")
}
