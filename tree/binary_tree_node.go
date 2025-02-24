package tree

type BinaryTreeNode interface {
	Node
	Left() BinaryTreeNode
	Right() BinaryTreeNode
	SetLeft(n BinaryTreeNode)
	SetRight(n BinaryTreeNode)
}

type binaryTreeNode struct {
	value interface{}
	left *binaryTreeNode
	right *binaryTreeNode
}

func NewNode(value interface{}, left BinaryTreeNode, right BinaryTreeNode) BinaryTreeNode {
	n := &binaryTreeNode{}
	n.value = value
	if left != nil {
		n.left = left.(*binaryTreeNode)
	}
	if right != nil {
		n.right = right.(*binaryTreeNode)
	}
	return n
}

func (n binaryTreeNode) Value() interface{} {
	return n.value
}

func (n binaryTreeNode) Left() BinaryTreeNode {
	return n.left
}

func (n binaryTreeNode) Right() BinaryTreeNode {
	return n.right
}

func (n binaryTreeNode) SetLeft(left BinaryTreeNode) {
	n.left = left.(*binaryTreeNode)
}

func (n binaryTreeNode) SetRight(right BinaryTreeNode) {
	n.right = right.(*binaryTreeNode)
}
