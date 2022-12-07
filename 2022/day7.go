package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	Alit "github.com/Alit54/General/gofunc"
)

type Tree struct {
	root *TreeNode
}

type TreeNode struct {
	// DIR will have sons.
	// FILES will have son = nil
	// Both DIR and FILES will have a size equals to the sum of his sons
	name   string
	father *TreeNode
	son    []*TreeNode
	isDir  bool
	size   int
}

func main() {
	fmt.Println(Day7(os.Args[1]))
}

func Day7(input string) (int, int) {
	return Part7A(input), Part7B(input)
}

func Part7A(input string) int {
	sum := 0
	tree := createTree(input)
	updateFolders(tree)
	sum = sumFilter(tree.root)
	return sum
}

func Part7B(input string) int {
	tree := createTree(input)
	updateFolders(tree)
	totalSpace := tree.root.size
	spaceRequired := 30000000 - (70000000 - totalSpace)
	dirDeleted := getDirectory(tree.root, tree.root, spaceRequired)
	return dirDeleted.size
}

func updateFolders(tree *Tree) {
	for _, v := range tree.root.son {
		tree.root.size += calculateSize(v)
	}
}

func getDirectory(leaf, min *TreeNode, limit int) *TreeNode {
	if leaf.isDir && leaf.size > limit && leaf.size < min.size {
		min = leaf
	}
	for _, v := range leaf.son {
		min = getDirectory(v, min, limit)
	}
	return min
}

func calculateSize(leaf *TreeNode) int {
	if leaf.isDir {
		for _, v := range leaf.son {
			leaf.size += calculateSize(v)
		}
	}
	return leaf.size
}

func sumFilter(file *TreeNode) int {
	sum := 0
	if file.isDir && file.size < 100000 {
		sum += file.size
	}
	for _, v := range file.son {
		sum += sumFilter(v)
	}
	return sum
}

func createTree(input string) *Tree {
	rows := Alit.ReadRows(input)
	var tree Tree
	slice := []string{"dir", "/"}
	tree.root = createLeaf(nil, slice)
	dir := tree.root
	// Create the Tree
	for i := 1; i < len(rows); i++ {
		row := strings.Split(rows[i], " ")
		if row[1] == "cd" {
			// There is a command: & cd ".." OR "name of folder"
			if row[2] == ".." {
				dir = dir.father
			} else {
				for _, v := range dir.son {
					if v.name == row[2] {
						dir = v
						break
					}
				}
			}
		}
		if row[1] == "ls" {
			// There is a command: & ls
			for i = i + 1; i < len(rows); i++ {
				row = strings.Split(rows[i], " ")
				if row[0] == "$" {
					i--
					break
				}
				dir.son = append(dir.son, createLeaf(dir, row))
			}
		}
	}
	return &tree
}

// Returns a leaf for the current directory
func createLeaf(dir *TreeNode, data []string) *TreeNode {
	var leaf TreeNode
	var err error
	if dir != nil {
		leaf.father = dir
	}
	leaf.name = data[1]
	leaf.size, err = strconv.Atoi(data[0]) // if this gives error, than leaf is a directory and therefor size is zero for now.
	if err != nil {
		leaf.isDir = true
	} else {
		leaf.isDir = false
	}
	return &leaf
}
