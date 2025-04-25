package main

type Node struct {
	Val int
	Left *Node
	Right *Node
}

func main() {
	
}

func insert(root *Node, val int) *Node {
	newNode := &Node{Val: val}
	node := root

	for node != nil {
		if val <= node.Val {
			if node.Left == nil {
				node.Left = newNode
				break
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = newNode
				break
			}
			node = node.Right
		}
	}

	return newNode
}

func search(root *Node, val int) *Node {
	node := root

	for node != nil {
		if node.Val == val {
			return node
		} else if val <= node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return nil	
}


// BST Deletion
// If Node is a leaf node, delete it
// If Node has one child, the child takes its place
// If Node has two children, the successor takes its place

// Successor is the smallest node in the right subtree
// Predecessor is the largest node in the left subtree

func delete(root *Node, val int) *Node {
	if root == nil{
		return nil
	}

	// parent of the node that will be deleted
	var delParent *Node

	// node that will be deleted
	delNode := root

	// find the node that will be deleted
	for delNode != nil {
		if val == delNode.Val {
			break
		} else if val < delNode.Val {
			delParent = delNode
			delNode = delNode.Left
		} else {
			delParent = delNode
			delNode = delNode.Right
		}
	}

	// if the node to be deleted is not found, return the root
	if delNode == nil {
		return root
	}

	// if the node to be deleted is a leaf node
	if delNode.Left == nil && delNode.Right == nil {
		// If the node to be deleted is the root, return nil
		if delParent == nil {
			return nil
		// If the node to be deleted is the left child of its parent, set the left child to nil
		} else if delParent.Left == delNode {
			delParent.Left = nil
		// If the node to be deleted is the right child of its parent, set the right child to nil
		} else {
			delParent.Right = nil
		}

		return root
	}

	if delNode.Right == nil {
		if delParent == nil {
			return delNode.Left
		} else if delParent.Left == delNode {
			delParent.Left = delNode.Left
		} else {
			delParent.Right = delNode.Left
		}
		
		return root
	}

	successor := findSuccessor(delNode)
	successor.Left = delNode.Left
	successor.Right = delNode.Right
	delNode.Left = nil
	delNode.Right = nil

	if delParent == nil {
		return successor
	} else if delParent.Left == delNode {
		delParent.Left = successor
	} else {
		delParent.Right = successor
	}

	return root
}

func findPredecessor(delNode *Node) *Node {
	if delNode.Left == nil {
		return nil
	}

	parent := delNode
	node := delNode.Left

	for node.Right != nil {
		parent = node
		node = node.Right
	}

	if parent == delNode {
		parent.Left = node.Left
	} else if node.Left != nil {
		parent.Right = node.Left
	} else {
		parent.Right = nil
	}

	return node
}

func findSuccessor(delNode *Node) *Node {
	if delNode.Right == nil {
		return nil
	}

	parent := delNode
	node := delNode.Right

	for node.Left != nil {
		parent = node
		node = node.Left
	}

	if parent == delNode {
		parent.Right = node.Right
	} else if node.Right != nil {
		parent.Left = node.Right
	} else {
		parent.Left = nil
	}

	return node
}