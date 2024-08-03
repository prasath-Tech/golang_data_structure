package main

import (
	"fmt"
	"data_struct/operations"
)

func main() {
	var demoarr = [10]int {23,87,34,56,57,98,53,56,73,10}
	list := operations.LinkedList{}
	bst := operations.BST{}
	for _,interger := range(demoarr){
		list.Push(interger)
		bst.Push(interger)
	}
	fmt.Printf("Length of the linked list %v \n",list.Length())
	list.Print()
	index,presence := list.Find(56)
	fmt.Printf("Index %d, Presence %t \n", index,presence)
	bst.Inorder(bst.Provideroot())
	bst.LevelOrder()
	fmt.Println("\n",bst.Find(bst.Provideroot(),56))
}
