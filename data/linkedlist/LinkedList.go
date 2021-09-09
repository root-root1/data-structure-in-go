package linkedlist

import "fmt"

type Node struct { // node class
	data     int
	nextNode *Node
}

type LinkedList struct { // linked list class
	head *Node
}

func (linkedlist *LinkedList) AddTOHead(data int) {
	var node = &Node{}
	node.data = data
	if linkedlist.head != nil {
		node.nextNode = linkedlist.head
	}
	linkedlist.head = node
}

func (linkedlist *LinkedList) Show() {
	// var node *Node
	for node := linkedlist.head; node != nil; node = node.nextNode {
		fmt.Println(node.data)
	}
}

func (linkedlist *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedlist.head; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (linkedlist *LinkedList) AddTOEnd(data int) {
	var node = &Node{}
	node.data = data
	node.nextNode = nil
	// var lastNode *Node
	var lastNode = linkedlist.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

func (linkedlist *LinkedList) NodeWithValue(data int) *Node {
	var node *Node
	var nodeWith *Node

	for node = linkedlist.head; node != nil; node = node.nextNode {
		if node.data == data {
			nodeWith = node
			break
		}
	}
	return nodeWith

}

func (linkedlist *LinkedList) AddAfter(nodeData, data int) {
	var node = &Node{}
	node.data = data
	node.nextNode = nil
	var nodewith = linkedlist.NodeWithValue(nodeData)
	if nodewith != nil {
		node.nextNode = nodewith.nextNode
		nodewith.nextNode = node
	}
}
