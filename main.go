package main

import (
	"fmt"

	b "github.com/root-root1/conTut/data/binarysearchtree"
)

func main() {
	tree := &b.BinarySearchTree{}
	tree.InsertElement(8, 8)
	tree.InsertElement(4, 3)
	tree.InsertElement(9, 1)
	tree.InsertElement(7, 75)
	tree.InsertElement(5, 3)
	tree.String()
	tree.InOrderTraversal(
		func(i int) {
			// fmt.Print("In order: ")
			fmt.Printf("%d ", i)
		},
	)
	fmt.Println()
	tree.PreOrderTraversal(
		func(i int) {
			// fmt.Print("Pre order: ")
			fmt.Printf("%d ", i)
		},
	)
	fmt.Println()
	tree.PostOrderTraversal(
		func(i int) {
			// fmt.Print("Post order: ")
			fmt.Printf("%d ", i)
		},
	)
	fmt.Println()
}
