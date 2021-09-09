package linkedlist

import "fmt"

type Node struct {
	data     int
	nextNode *Node
}

type LinkedList struct {
	head *Node
}

// add to head function take data and insert it to head of linkedlist
func (linkedlist *LinkedList) AddTOHead(data int) {
	var node = &Node{}
	node.data = data
	node.nextNode = nil
	if linkedlist.head != nil {
		node.nextNode = linkedlist.head
	}
	linkedlist.head = node
}

func (linkedlist *LinkedList) Show() {
	var node *Node
	for node = linkedlist.head; node != nil; node = node.nextNode {
		fmt.Printf("%v -> ", node.data)
	}
	fmt.Println("None")
}

func (linkedlist *LinkedList) lastNode() *Node {
	var node *Node
	var lastData *Node
	for node = linkedlist.head; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastData = node
		}
	}
	return lastData
}

func (linkedlist *LinkedList) AddTOEnd(data int) {
	var node = &Node{}
	node.data = data
	node.nextNode = nil
	var lastNode = linkedlist.lastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

func (linkedlist *LinkedList) nodeWithValue(data int) *Node {
	var node *Node
	var valueWith *Node
	for node = linkedlist.head; node != nil; node = node.nextNode {
		if node.data == data {
			valueWith = node
		}
	}
	return valueWith
}

func (linkedlist *LinkedList) AddAfter(nodeData, data int) {
	var node = &Node{}
	node.data = data
	node.nextNode = nil

	var withValue = linkedlist.nodeWithValue(nodeData)
	if withValue != nil {
		// withValue.nextNode = node
		node.nextNode = withValue.nextNode
		withValue.nextNode = node

	}
}
