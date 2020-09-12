package ctci

import (
	"github.com/anthonykrivonos/al-go/graph"
	"github.com/anthonykrivonos/al-go/list"
	"github.com/anthonykrivonos/al-go/queue"
	"github.com/anthonykrivonos/al-go/set"
	"github.com/anthonykrivonos/al-go/tree"
	"github.com/anthonykrivonos/al-go/utils"
	"math"
)

// 4.1 Route Between Nodes
func RouteBetweenNodes(g *graph.Graph, start, end *graph.Node) bool {
	var nodeQueue queue.Queue = list.NewArrayList()
	visitedSet := set.NewSet()

	nodeQueue.Unshift(start)

	for nodeQueue.Length() > 0 {
		top := nodeQueue.Poll().(*graph.Node)
		visitedSet.Insert(top)
		for _, neighbor := range g.Edges()[top] {
			if !visitedSet.Has(neighbor) {
				nodeQueue.Unshift(neighbor)
			}
		}
	}

	return visitedSet.Has(end)
}

// 4.2 Minimal Tree
func MinimalTree(arr []int) tree.BinarySearchTree {
	root := minimalTree(arr, 0, len(arr) - 1)
	return tree.NewBinarySearchTree(root, utils.IntAscComp)
}

func minimalTree(arr []int, left, right int) tree.BinaryTreeNode {
	if left < right {
		mid := (left + right) / 2
		midNode := tree.NewNode(arr[mid], nil, nil)
		midNode.SetLeft(minimalTree(arr, left, mid - 1))
		midNode.SetRight(minimalTree(arr, mid + 1, right))
		return midNode
	}
	return nil
}

// 4.3 List of Depths
func ListOfDepths(binaryTree tree.BinaryTree) []list.SinglyLinkedList {
	depthList := make([]list.SinglyLinkedList, 0)

	// Inserts a node at a given height
	insert := func (height int, node tree.BinaryTreeNode) {
		if len(depthList) >= height {
			depthList[height].Push(node)
		} else {
			for len(depthList) < height {
				depthList = append(depthList, list.NewSinglyLinkedList(nil))
			}
		}
	}
	listOfDepths(binaryTree.Root(), 0, insert)

	return depthList
}
func listOfDepths(node tree.BinaryTreeNode, height int, insert func (height int, node tree.BinaryTreeNode)) {
	if node != nil {
		listOfDepths(node.Left(), height + 1, insert)
		insert(height, node)
		listOfDepths(node.Right(), height + 1, insert)
	}
}

// 4.4 Check Balanced
func CheckBalanced(binaryTree tree.BinaryTree) bool {
	height := getHeight(binaryTree.Root())
	return checkBalanced(binaryTree.Root(), height)
}
func checkBalanced(node tree.BinaryTreeNode, height int) bool {
	if node == nil {
		return true
	}

	leftHeight := getHeight(node.Left())
	rightHeight := getHeight(node.Right())

	if height - leftHeight > 1 || height - rightHeight > 1 {
		return false
	}

	return checkBalanced(node.Left(), leftHeight) && checkBalanced(node.Right(), rightHeight)
}

func getHeight(node tree.BinaryTreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + utils.Max(getHeight(node.Left()), getHeight(node.Right()))
}

// 4.5 Validate BST
func ValidateBST(binaryTree tree.BinaryTree) bool {
	return validateBST(binaryTree.Root(), math.MaxInt8, math.MinInt8)
}
func validateBST(node tree.BinaryTreeNode, max, min int) bool {
	if node == nil {
		return true
	}

	if node.Value().(int) < min || node.Value().(int) > max {
		return false
	}

	return validateBST(node.Left(), node.Value().(int) - 1, min) && validateBST(node.Right(), max, node.Value().(int) + 1)
}

// 4.7 Build Order
func BuildOrder(projects []rune, dependencies [][]rune) []rune {

	// Create a node for each project
	projectGraph := graph.Graph{}
	runeToNode := make(map[rune]*graph.Node)
	for _, project := range projects {
		runeToNode[project] = graph.NewNode(project)
		projectGraph.AddNode(runeToNode[project])
	}

	// Add edges between projects
	for _, projects := range dependencies {
		dep := runeToNode[projects[0]]
		main := runeToNode[projects[1]]
		// Add edge pointing from main package to its dependency
		projectGraph.AddEdge(main, dep)
	}

	// By now, the most important dependency is the last node in each path
	// Do non-tail recursion DFS
	depsInOrder := make([]rune, 0)
	visitedSet := set.NewSet()

	// DFS from every node as a starting node
	for _, n := range projectGraph.Nodes() {
		buildOrderDFS(n, projectGraph, visitedSet, depsInOrder)
	}

	// Get the resulting dependencies in order
	return depsInOrder
}
func buildOrderDFS(node *graph.Node, projectGraph graph.Graph, visited set.Set, res []rune) {
	if !visited.Has(node) {
		visited.Insert(node)
		for _, e := range projectGraph.Edges()[node] {
			if visited.Has(e.To()) {
				buildOrderDFS(e.To(), projectGraph, visited, res)
			}
		}
		res = append(res, node.Value().(rune))
	}
}

// 4.8 First Common Ancestor
func FirstCommonAncestor(tree tree.BinaryTree, nodeA, nodeB tree.BinaryTreeNode) tree.BinaryTreeNode {
	return firstCommonAncestor(tree.Root(), nodeA, nodeB)
}
func firstCommonAncestor(current, nodeA, nodeB tree.BinaryTreeNode) tree.BinaryTreeNode {
	if current == nil {
		return nil
	} else if current == nodeA || current == nodeB {
		return current
	}

	leftAncestor := firstCommonAncestor(current.Left(), nodeA, nodeB)
	rightAncestor := firstCommonAncestor(current.Right(), nodeA, nodeB)

	if leftAncestor != nil && rightAncestor != nil {
		return current
	} else if leftAncestor != nil {
		return leftAncestor
	} else if rightAncestor != nil {
		return rightAncestor
	}
	return nil
}

// 4.9 BST Sequences
func BSTSequences(binaryTree tree.BinaryTree) [][]int {
	current := make([]int, 0)
	res := make([][]int, 0)
	bstSequences(binaryTree.Root(), current, res)
	return res
}
func bstSequences(root tree.BinaryTreeNode, current []int, res [][]int) {
	if root == nil {
		return
	} else if root.Left() == nil && root.Right() == nil {
		res = append(res, current)
	}

	current = append(current, root.Value().(int))

	leftCurrent := make([]int, len(current))
	copy(leftCurrent, current)
	rightCurrent := make([]int, len(current))
	copy(rightCurrent, current)

	bstSequences(root.Left(), leftCurrent, res)
	bstSequences(root.Right(), rightCurrent, res)
}
