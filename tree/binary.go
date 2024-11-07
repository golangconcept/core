package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(Val int) {
	bst.InsertRec(bst.root, Val)
}

func (bst *BST) InsertRec(node *Node, val int) *Node {
	if bst.root == nil {
		bst.root = &Node{val, nil, nil}
	}
	if node == nil {
		return &Node{val, nil, nil}
	}
	if val <= node.data {
		node.left = bst.InsertRec(node.left, val)
	} else {
		node.right = bst.InsertRec(node.right, val)
	}
	return node
}

func (bst *BST) Search(Val int) bool {
	found := bst.SearchRec(bst.root, Val)
	return found
}
func (bst *BST) SearchRec(node *Node, val int) bool {
	if node.data == val {
		return true
	}
	if node == nil {
		return false
	}

	if val < node.data {
		return bst.SearchRec(node.left, val)
	} else {
		return bst.SearchRec(node.right, val)
	}
}
func (bst *BST) Inorder(node *Node) {
	if node == nil {
		return
	} else {
		bst.Inorder(node.left)
		fmt.Print(node.data, " ")
		bst.Inorder(node.right)
	}
}

func (bst *BST) LevelOrder() {
	if bst.root == nil {
		return
	}

	nodeList := make([](*Node), 0)
	nodeList = append(nodeList, bst.root)

	for !(len(nodeList) == 0) {
		current := nodeList[0]
		fmt.Println(current.data, " ")

		if current.left != nil {
			nodeList = append(nodeList, current.left)
		}
		if current.right != nil {
			nodeList = append(nodeList, current.right)
		}

		nodeList = nodeList[1:]
	}
}
func main() {
	bst := BST{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(20)
	bst.Insert(17)
	bst.Insert(4)
	bst.Inorder(bst.root)
	fmt.Println(bst.Search(5))
}
