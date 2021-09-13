package stack

import (
	"fmt"
	"strconv"
)

type Element struct {
	ElementValue int
}

func (element *Element) String() string {
	return strconv.Itoa(element.ElementValue)
}

func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

// stack class
type Stack struct {
	elements     []*Element
	elementCount int
}

func (stack *Stack) Push(data *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], data)
	stack.elementCount += 1
}

func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}
	length := len(stack.elements)
	element := stack.elements[length-1]

	if length > 1 {
		stack.elements = stack.elements[:length-1]
	} else {
		stack.elements = stack.elements[0:]
	}
	stack.elementCount = len(stack.elements)
	return element
}

func (stack *Stack) Show() {
	for data := range stack.elements {
		fmt.Printf("%d ", data)
	}
}
