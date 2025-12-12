package main

import "fmt"

type BinaryTreeNode struct {
	Val                   int
	leftChild, rightChild *BinaryTreeNode
}

func NewBinaryTreeNode(val int) *BinaryTreeNode {
	return &BinaryTreeNode{Val: val, leftChild: nil, rightChild: nil}
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(val int) BinaryTree {
	return BinaryTree{Root: NewBinaryTreeNode(val)}
}

/*We will use level order Insertion*/
func (b *BinaryTree) InsertBinaryTree(val int) {
	newNode := NewBinaryTreeNode(val)
	if b.Root == nil {
		b.Root = newNode
		return
	}

	// S1. Create a queue and insert root address
	queue := []*BinaryTreeNode{b.Root}

	for len(queue) > 0 {

		// S2. Take out the address of root give to temp
		temp := queue[0]

		// pop out 
		queue = queue[1:]

		// visit left child 
		if temp.leftChild == nil {
			temp.leftChild = newNode
			return
		} else {
			queue = append(queue, temp.leftChild)
		}

		if temp.rightChild == nil {
			temp.rightChild = newNode
			return
		} else {
			queue = append(queue, temp.rightChild)

		}

	}

}

// PreOrder Traversal
func PrintNode(node *BinaryTreeNode) {
	if node != nil {
		fmt.Printf("%d\n", node.Val)
		PrintNode(node.leftChild)
		PrintNode(node.rightChild)
	}
}

// PostOrder Traversal 

func PrintPostOrder(node *BinaryTreeNode) {
	if node != nil {
		fmt.Printf("%d\n", node.Val)
		PrintNode(node.leftChild)
		PrintNode(node.rightChild)
	}
}

/*
PreOrder : visit(node), Preorder (left subtree), Preorder(right subtree)

Inorder : Inorder(left), visit(node), Inorder(right)

PostOrder : Postorder(left), Postorder(right), visit(node)

Level order : L

*/

func main() {
	myTree := NewBinaryTree(10)

	myTree.InsertBinaryTree(5)
	myTree.InsertBinaryTree(15)
	myTree.InsertBinaryTree(2)
	myTree.InsertBinaryTree(7)

	fmt.Print("Tree Nodes (Pre-order): \n")

	PrintNode(myTree.Root)

	// The structure is:
    //      10
    //     /  \
    //    5    15
    //   / \
    //  2   7
}
