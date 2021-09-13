package binarysearchtree

import (
	"fmt"
	"strings"
	"sync"
)

type Node struct {
	key       int
	value     int
	leftNode  *Node
	rightNode *Node
}

type BinarySearchTree struct {
	rootNode *Node
	lock     sync.RWMutex
}

func (tree *BinarySearchTree) InsertElement(key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	treeNode := &Node{key: key, value: value}
	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertElement(tree.rootNode, treeNode)
	}
}

func insertElement(rootNode *Node, treeNode *Node) {
	if treeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = treeNode
		} else {
			insertElement(rootNode.leftNode, treeNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = treeNode
		} else {
			insertElement(rootNode.rightNode, treeNode)
		}
	}
}

func (tree *BinarySearchTree) InOrderTraversal(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Print("In Order: ")
	inOrderTraversal(tree.rootNode, function)
}

func inOrderTraversal(rootNode *Node, function func(int)) {
	if rootNode != nil {
		inOrderTraversal(rootNode.leftNode, function)
		function(rootNode.key)
		inOrderTraversal(rootNode.rightNode, function)
	}
}

func (tree *BinarySearchTree) PreOrderTraversal(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Print("Pre Order: ")
	preOrderTraversal(tree.rootNode, function)
}

func preOrderTraversal(rootNode *Node, function func(int)) {
	if rootNode != nil {
		function(rootNode.key)
		preOrderTraversal(rootNode.leftNode, function)
		preOrderTraversal(rootNode.rightNode, function)
	}
}

func (tree *BinarySearchTree) PostOrderTraversal(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Print("Post Order: ")
	postOrderTraversal(tree.rootNode, function)
}

func postOrderTraversal(rootNode *Node, function func(int)) {
	if rootNode != nil {
		postOrderTraversal(rootNode.leftNode, function)
		postOrderTraversal(rootNode.rightNode, function)
		function(rootNode.key)
	}
}

func (tree *BinarySearchTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	root := tree.rootNode
	if root == nil {
		return (*int)(nil)
	}
	for {
		if root.leftNode == nil {
			return &root.value
		}

		root = root.leftNode
	}
}

func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.Lock()
	root := tree.rootNode
	if root == nil {
		return (*int)(nil)
	}
	for {
		if root.rightNode == nil {
			return &root.value
		}
		root = root.rightNode
	}
}

func (tree *BinarySearchTree) Search(data int) bool {
	tree.lock.RLock()
	defer tree.lock.RLock()
	return search(tree.rootNode, data)
}

func search(rootNode *Node, data int) bool {
	if rootNode == nil {
		return false
	}
	if data < rootNode.key {
		return search(rootNode.leftNode, data)
	}
	if data > rootNode.key {
		return search(rootNode.rightNode, data)
	}
	return true
}

func (tree *BinarySearchTree) RemoveNode(data int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	removeNode(tree.rootNode, data)
}

func removeNode(rootNode *Node, data int) *Node {
	if rootNode == nil {
		return nil
	}
	if data < rootNode.key {
		rootNode.leftNode = removeNode(rootNode.leftNode, data)
	}
	if data > rootNode.key {
		rootNode.rightNode = removeNode(rootNode.rightNode, data)
	}

	if rootNode.leftNode == nil && rootNode.rightNode == nil {
		rootNode = nil
		return nil
	}

	if rootNode.leftNode == nil {
		rootNode = rootNode.leftNode
		return rootNode
	}

	if rootNode.rightNode == nil {
		rootNode = rootNode.rightNode
		return rootNode
	}

	leftMostRightNode := rootNode.rightNode
	for {
		if leftMostRightNode != nil && leftMostRightNode.leftNode != nil {
			leftMostRightNode = leftMostRightNode.leftNode
		} else {
			break
		}
	}
	rootNode.key, rootNode.value = leftMostRightNode.key, leftMostRightNode.value
	rootNode.rightNode = removeNode(rootNode.rightNode, rootNode.key)
	return rootNode
}

func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println(strings.Repeat("-", 48))
	stringify(tree.rootNode, 0)
	fmt.Println(strings.Repeat("-", 48))
}

func stringify(rootNode *Node, level int) {
	if rootNode != nil {
		format := ""
		for i := 0; i < level; i += 1 {
			format += " "
		}
		format += "---[ "
		level += 1
		stringify(rootNode.leftNode, level)
		fmt.Printf(format+"%d\n", rootNode.key)
		stringify(rootNode.rightNode, level)
	}
}
