package tree

type BinarySearchTree interface {
	BinaryTree
	Insert(x interface{})
	Has(x interface{}) bool
}

type binarySearchTree struct {
	binaryTree
	comparator func(a, b interface{}) int
}

func NewBinarySearchTree(comparator func(a, b interface{}) int) BinarySearchTree {
	b := &binarySearchTree{}
	b.comparator = comparator
	return b
}

func (b *binarySearchTree) Insert(x interface{}) {
	b.insert(x, b.root)
}

func (b *binarySearchTree) insert(x interface{}, n *binaryTreeNode) *binaryTreeNode {
	if n == nil {
		return &binaryTreeNode{value: x}
	}
	if b.comparator(x, n.value) < 0 {
		n.left = b.insert(x, n.left)
	} else {
		n.right = b.insert(x, n.right)
	}
	return n
}

func (b *binarySearchTree) Has(x interface{}) bool {
	return b.has(x, b.root)
}

func (b *binarySearchTree) has(x interface{}, n *binaryTreeNode) bool {
	if n == nil || n.value == x {
		return true
	}
	if b.comparator(x, n.value) < 0 {
		return b.has(x, n.left)
	}
	return b.has(x, n.right)
}

var _ BinarySearchTree = &binarySearchTree{}
