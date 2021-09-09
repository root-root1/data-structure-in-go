package doublelinked

import "fmt"

type Node struct {
	data     int
	nextNode *Node
	prevNode *Node
}

type DLinkedList struct {
	head *Node
}

func (dlinkedlist *DLinkedList) NodeBetweenValue(firstData int, secondData int) *Node {
	var node *Node
	var withValue *Node

	for node = dlinkedlist.head; node != nil; node = node.nextNode {
		if node.prevNode != nil && node.nextNode != nil {
			if node.prevNode.data == firstData && node.nextNode.data == secondData {
				withValue = node
			}
		}
	}
	return withValue
}

func (linkedlist *DLinkedList) AddTOHead(data int) {
	var node = &Node{}
	node.data = data
	node.nextNode = nil
	if linkedlist.head != nil {
		node.nextNode = linkedlist.head
		linkedlist.head.prevNode = node
	}
	linkedlist.head = node
}

func (linkedlist *DLinkedList) nodeWithValue(data int) *Node {
	var node *Node
	var withValue *Node
	for node = linkedlist.head; node != nil; node = node.nextNode {
		if node.data == data {
			withValue = node
		}
	}
	return withValue
}

func (linkedlist *DLinkedList) AddAfter(nodeData, data int) {
	var node = &Node{}
	node.data = data
	node.nextNode = nil
	var withValue = linkedlist.nodeWithValue(nodeData)
	if withValue != nil {
		node.nextNode = withValue.nextNode
		node.prevNode = withValue
		withValue.nextNode = node
	}
}

func (linkedlist *DLinkedList) lastNode() *Node {
	var node *Node
	var lastValue *Node

	for node = linkedlist.head; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastValue = node
		}
	}
	return lastValue
}

func (linkedlist *DLinkedList) AddTOEnd(data int) {
	var node = &Node{}
	node.data = data
	node.nextNode = nil
	var lastNode = linkedlist.lastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.prevNode = lastNode
	}
}

func (linkedlist *DLinkedList) Show() {
	var node *Node
	for node = linkedlist.head; node != nil; node = node.nextNode {
		fmt.Printf("%v -> ", node.data)
	}
	fmt.Println("Nil")
}
