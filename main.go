package main

import (
	"fmt"

	set "github.com/root_root1/conTut/data/test/set"
)

func main() {
	// var list linkedlist.LinkedList
	// list = linkedlist.LinkedList{}
	var setData *set.Set
	setData = &set.Set{}
	setData.New()
	setData.AddElement(1)
	setData.AddElement(2)
	fmt.Println("Initial Set ", setData)
	fmt.Println(setData.ContainsElement(1))
	var anotherSet *set.Set
	anotherSet = &set.Set{}
	anotherSet.New()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)
	fmt.Println(setData.Intersect(anotherSet))
	fmt.Println(setData.Union(anotherSet))
}
