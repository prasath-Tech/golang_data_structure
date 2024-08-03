package operations

import (
	"fmt"
)

type ListNode struct {
	data int
	next *ListNode
}

type TreeNode struct {
	data int
	left *TreeNode
	right *TreeNode
}

type LinkedList struct {
	head *ListNode
}

type BST struct {
	root *TreeNode
}

//Linked list

func (list *LinkedList) Push(data int) {
	newNode := &ListNode{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *LinkedList) Print() {
	var current *ListNode = list.head
	for current != nil {
		fmt.Print(current.data, "\t")
		current = current.next
	}
	fmt.Println()
}

func (list *LinkedList) Length() int32 {
	var count int32
	current := list.head
	for(current != nil){
		current = current.next
		count++
	}

	return count
}

func (list *LinkedList) Find(value int) (int,bool) {
	count := 0
	presence := false
	current := list.head
	for(current != nil){
		if(current.data == value){
			presence = true
			return count,presence
		}
		count++
		current = current.next
	}
	return count,presence
}


//Binary tree
func (bst *BST) Push(val int){
	bst.InsertRec(bst.root, val)
}

func (bst *BST) InsertRec(node *TreeNode, val int) *TreeNode{
	if bst.root == nil {
        bst.root = &TreeNode{val, nil, nil}
        return bst.root
    }
	if node == nil {
		return &TreeNode{val, nil,nil}
	}
	if val <= node.data {
        node.left = bst.InsertRec(node.left, val)
    }
    if val > node.data {
        node.right = bst.InsertRec(node.right, val)
    }
    return node
}

func (bst *BST)Inorder(node *TreeNode){
	if(node == nil){
		return 
	}else{
		bst.Inorder(node.left)
		fmt.Print(node.data," ")
		bst.Inorder(node.right)
	}
}

func (bst *BST) Provideroot() *TreeNode{
	return bst.root
}

func (bst *BST) LevelOrder() {
	fmt.Print("\n")
	if(bst.root == nil){
		return
	}
	nodeList := make([]*TreeNode,0)
	nodeList = append(nodeList, bst.root)
	for !(len(nodeList)==0){
		current := nodeList[0]
		fmt.Print(current.data,"\t")
		if(current.left != nil){
			nodeList = append(nodeList, current.left)
		}
		if(current.right != nil){
			nodeList = append(nodeList, current.right)
		}
		nodeList = nodeList[1:]
	}
}

func (bst *BST) Find(node *TreeNode, val int) bool {
    if node.data == val {
        return true
    }
    if node == nil {
        return false
    }
    if val < node.data {
        return bst.Find(node.left, val)
    }
    if val > node.data {
        return bst.Find(node.right, val)
    }
    return false
}

