package graph

import (
	"fmt"
	"github.com/anthonykrivonos/al-go/list"
	"github.com/anthonykrivonos/al-go/queue"
	"github.com/anthonykrivonos/al-go/utils"
	"math"
	"sync"
)

type Graph struct {
	nodes []*Node
	edges map[*Node][]*Edge
	lock sync.RWMutex
}

func (g *Graph) Nodes() []*Node {
	return g.nodes
}

func (g *Graph) Edges() map[*Node][]*Edge {
	return g.edges
}

func (g *Graph) Add(v interface{}) *Node {
	g.lock.Lock()
	defer g.lock.Unlock()

	n := Node{v}
	g.nodes = append(g.nodes, &n)

	return &n
}

func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddWeightedEdge(n1 *Node, n2 *Node, weight float64) {
	g.lock.Lock()
	defer g.lock.Unlock()

	if g.edges == nil {
		g.edges = make(map[*Node][] *Edge)
	}

	e := NewEdge(n1, n2, weight)

	// Overwrite if n1 is already connected to n2
	includes := false
	for i, edge := range g.edges[n1] {
		if edge.to.value == n2.value {
			includes = true
			g.edges[n1][i] = e
			break
		}
	}

	// Otherwise, add to the list of edges
	if !includes {
		g.edges[n1] = append(g.edges[n1], e)
	}
}

func (g *Graph) AddEdge(n1 *Node, n2 *Node) {
	g.AddWeightedEdge(n1, n2, 0.0)
}

func (g *Graph) AddWeightedBidirectionalEdge(n1 *Node, n2 *Node, weight float64) {
	g.AddWeightedEdge(n1, n2, weight)
	g.AddWeightedEdge(n2, n1, weight)
}

func (g *Graph) AddBidirectionalEdge(n1 *Node, n2 *Node) {
	g.AddWeightedBidirectionalEdge(n1, n2, 0.0)
}

func (g *Graph) Neighbors(n *Node) []*Edge {
	g.lock.RLock()
	defer g.lock.RUnlock()

	if _, ok := g.edges[n]; !ok {
		return make([]*Edge, 0)
	}
	return g.edges[n]
}

func (g *Graph) String() (s string) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + ":\n"
		near := g.edges[g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += "  " + near[j].String() + "\n"
		}
	}
	return
}

func (g *Graph) DepthFirstPrint(start *Node) {
	visited := new(map[Node]bool)
	for _, n := range g.nodes {
		(*visited)[*n] = false
	}
	g.depthFirstPrint(start, visited)
}

func (g *Graph) depthFirstPrint(node *Node, visited *map[Node]bool) {
	if node == nil {
		return
	}
	fmt.Printf("%t ", node.value)
	(*visited)[*node] = true
	for _, e := range g.edges[node] {
		adjacent := e.to
		if !(*visited)[*adjacent] {
			g.depthFirstPrint(node, visited)
		}
	}
}

func (g *Graph) BreadthFirstPrint(start *Node) {
	if start == nil {
		return
	}

	visited := new(map[Node]bool)
	var visitQueue queue.Queue = list.NewArrayList()

	visitQueue.Unshift(start)

	for visitQueue.Length() > 0 {
		top := visitQueue.Poll().(*Node)
		(*visited)[*top] = true
		fmt.Printf("%t ", top.value)
		for _, e := range g.edges[top] {
			adjacent := e.to
			if !(*visited)[*adjacent] {
				visitQueue.Unshift(adjacent)
			}
		}
	}
}

func (g *Graph) DijkstraShortestPath(start *Node, end *Node) []*Node {
	// Initialize a mapping from nodes to their previous nodes
	previousMap := make(map[*Node] *Node)

	// Initialize distances
	distanceMap := make(map[*Node] float64)
	for _, node := range g.nodes {
		if node != start {
			distanceMap[node] = math.MaxFloat64
			previousMap[node] = nil
		}
	}
	distanceMap[start] = 0

	// Create a priority queue that sorts by distance
	comparator := func(a, b interface{}) int {
		return utils.Float64AscComp(distanceMap[a.(*Node)], distanceMap[b.(*Node)])
	}
	distanceQueue := queue.NewPriorityQueue(comparator)

	// Add the starting node to the queue
	distanceQueue.Unshift(start)

	// Traverse distances
	for distanceQueue.Length() > 0 {
		curr := distanceQueue.Poll().(*Node)
		for _, e := range g.edges[curr] {
			next := e.to
			if distanceMap[curr] + e.weight < distanceMap[next] {
				distanceMap[next] = distanceMap[curr] + e.weight
				previousMap[next] = curr
				distanceQueue.Unshift(next)
			}
		}
	}

	// Create path
	path := make([]*Node, 1)
	path[0] = end
	fin := previousMap[end]
	for fin != nil {
		path = append([]*Node{fin}, path...)
		fin = previousMap[fin]
	}

	return path
}

func (g *Graph) AStarShortestPath(start *Node, end *Node, heuristic func(a, b *Node) float64) []*Node {
	// Initialize a mapping from nodes to their previous nodes
	previousMap := make(map[*Node] *Node)

	// Initialize distances
	distanceMap := make(map[*Node] float64)
	for _, node := range g.nodes {
		if node != start {
			distanceMap[node] = math.MaxFloat64
			previousMap[node] = nil
		}
	}
	distanceMap[start] = 0

	// Create a priority queue that sorts by distance
	comparator := func(a, b interface{}) int {
		return utils.Float64AscComp(distanceMap[a.(*Node)], distanceMap[b.(*Node)])
	}
	distanceQueue := queue.NewPriorityQueue(comparator)

	// Add the starting node to the queue
	distanceQueue.Unshift(start)

	// Traverse distances
	for distanceQueue.Length() > 0 {
		curr := distanceQueue.Poll().(*Node)
		for _, e := range g.edges[curr] {
			next := e.to
			// Only difference from Dijkstra is the extra heuristic distance
			if distanceMap[curr] + e.weight + heuristic(curr, next) < distanceMap[next] {
				distanceMap[next] = distanceMap[curr] + e.weight + heuristic(curr, next)
				previousMap[next] = curr
				distanceQueue.Unshift(next)
			}
		}
	}

	// Create path
	path := make([]*Node, 1)
	path[0] = end
	fin := previousMap[end]
	for fin != nil {
		path = append([]*Node{fin}, path...)
		fin = previousMap[fin]
	}

	return path
}
