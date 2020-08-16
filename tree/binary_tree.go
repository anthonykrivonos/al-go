package tree

import "fmt"

type BinaryTree interface {
	Tree
	Root() BinaryTreeNode
	Flip()
}

type binaryTree struct {
	root *binaryTreeNode
}

func NewBinaryTree(root BinaryTreeNode) BinaryTree {
	b := &binaryTree{}
	if root != nil {
		b.root = root.(*binaryTreeNode)
	}
	return b
}

func (t *binaryTree) Root() BinaryTreeNode {
	if t.root != nil {
		return t.root
	}
	return nil
}

func (t *binaryTree) PreOrderPrint() {
	preOrderPrint(t.root)
}

func preOrderPrint(n *binaryTreeNode) {
	if n != nil {
		fmt.Printf("%t ", n.Value())
		preOrderPrint(n.left)
		preOrderPrint(n.right)
	}
}

func (t *binaryTree) InOrderPrint() {
	inOrderPrint(t.root)
}

func inOrderPrint(n *binaryTreeNode) {
	if n != nil {
		inOrderPrint(n.left)
		fmt.Printf("%t ", n.Value())
		inOrderPrint(n.right)
	}
}

func (t *binaryTree) PostOrderPrint() {
	postOrderPrint(t.root)
}

func postOrderPrint(n *binaryTreeNode) {
	if n != nil {
		postOrderPrint(n.left)
		postOrderPrint(n.right)
		fmt.Printf("%t ", n.Value())
	}
}

func (t *binaryTree) Flip() {
	t.root = flip(t.root)
}

func flip(node *binaryTreeNode) *binaryTreeNode {
	if node == nil || (node.left == nil && node.right == nil) {
		return node
	}

	leftFlippedNode := flip(node.left)

	node.left.left = node.right
	node.left.right = node
	node.left = nil
	node.right = nil

	return leftFlippedNode
}
