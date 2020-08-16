package graph

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var g Graph

func TestGraph(t *testing.T) {
	nA := &Node{"A"}
	nB := &Node{"B"}
	nC := &Node{"C"}
	nD := &Node{"D"}
	nE := &Node{"E"}

	g.AddNode(nA)
	g.AddNode(nB)
	g.AddNode(nC)
	g.AddNode(nD)
	g.AddNode(nE)
	nF := g.Add("F")

	assert.Equal(t, 6, len(g.Nodes()))

	g.AddEdge(nA, nB)
	g.AddEdge(nA, nC)
	g.AddEdge(nB, nE)
	g.AddEdge(nC, nE)
	g.AddEdge(nE, nF)
	g.AddEdge(nD, nA)

	assert.Equal(t, 2, len(g.Edges()[nA]))
	assert.Equal(t, 1, len(g.Edges()[nB]))
	assert.Equal(t, 1, len(g.Edges()[nB]))
	assert.Equal(t, 1, len(g.Edges()[nB]))
	assert.Equal(t, 1, len(g.Edges()[nB]))
}

func TestGraph_DijkstraShortestPath(t *testing.T) {
	newYork := &Node{"New York"}
	losAngeles := &Node{"Los Angeles"}
	sanFrancisco := &Node{"San Francisco"}
	chicago := &Node{"Chicago"}
	dallas := &Node{"Dallas"}
	miami := &Node{"Miami"}
	denver := &Node{"Denver"}
	seattle := &Node{"Seattle"}

	g.AddNode(newYork)
	g.AddNode(losAngeles)
	g.AddNode(sanFrancisco)
	g.AddNode(chicago)
	g.AddNode(dallas)
	g.AddNode(miami)
	g.AddNode(denver)
	g.AddNode(seattle)

	g.AddWeightedBidirectionalEdge(newYork, chicago, 750)
	g.AddWeightedBidirectionalEdge(newYork, chicago, 1278)
	g.AddWeightedBidirectionalEdge(miami, dallas, 1315)
	g.AddWeightedBidirectionalEdge(chicago, dallas, 967)
	g.AddWeightedBidirectionalEdge(chicago, seattle, 2042)
	g.AddWeightedBidirectionalEdge(dallas, seattle, 2096)
	g.AddWeightedBidirectionalEdge(chicago, denver, 1001)
	g.AddWeightedBidirectionalEdge(dallas, denver, 795)
	g.AddWeightedBidirectionalEdge(seattle, sanFrancisco, 807)
	g.AddWeightedBidirectionalEdge(seattle, denver, 1303)
	g.AddWeightedBidirectionalEdge(denver, sanFrancisco, 1248)
	g.AddWeightedBidirectionalEdge(denver, losAngeles, 1098)
	g.AddWeightedBidirectionalEdge(sanFrancisco, losAngeles, 382)

	path := g.DijkstraShortestPath(newYork, sanFrancisco)
	for _, node := range path {
		fmt.Printf("%s  " , node.String())
	}
}

type testCity struct {
	name string
	lat float64
	lng float64
}

func TestGraph_AStarShortestPath(t *testing.T) {
	newYork := &Node{testCity{name: "New York", lat: 40.7128, lng: 74.006}}
	losAngeles := &Node{testCity{name: "Los Angeles", lat: 34.0522, lng: 118.2437}}
	sanFrancisco := &Node{testCity{name: "San Francisco", lat: 37.7749, lng: 122.4194}}
	chicago := &Node{testCity{name: "Chicago", lat: 41.8781, lng: 87.6298}}
	dallas := &Node{testCity{name: "Dallas", lat: 32.7767, lng: 96.7970}}
	miami := &Node{testCity{name: "Miami", lat: 25.7617, lng: 80.1918}}
	denver := &Node{testCity{name: "Denver", lat: 39.7392, lng: 104.9903}}
	seattle := &Node{testCity{name: "Seattle", lat: 47.6062, lng: 122.3321}}

	haversine := func(a, b *Node) float64 {
		const R = 6371e3
		const rad = math.Pi / 180

		pointA := a.value.(testCity)
		pointB := a.value.(testCity)
		φ1 := pointA.lat * rad
		φ2 := pointB.lat * rad
		Δφ := (pointB.lat - pointA.lat) * rad
		Δλ := (pointB.lng - pointA.lng) * rad

		alpha := math.Sin(Δφ/2) * math.Sin(Δφ/2) + math.Cos(φ1) * math.Cos(φ2) * math.Sin(Δλ/2) * math.Sin(Δλ/2)
		beta := 2 * math.Atan2(math.Sqrt(alpha), math.Sqrt(1 - alpha));

		return R * beta
	}

	g.AddNode(newYork)
	g.AddNode(losAngeles)
	g.AddNode(sanFrancisco)
	g.AddNode(chicago)
	g.AddNode(dallas)
	g.AddNode(miami)
	g.AddNode(denver)
	g.AddNode(seattle)

	g.AddWeightedBidirectionalEdge(newYork, chicago, 750)
	g.AddWeightedBidirectionalEdge(newYork, chicago, 1278)
	g.AddWeightedBidirectionalEdge(miami, dallas, 1315)
	g.AddWeightedBidirectionalEdge(chicago, dallas, 967)
	g.AddWeightedBidirectionalEdge(chicago, seattle, 2042)
	g.AddWeightedBidirectionalEdge(dallas, seattle, 2096)
	g.AddWeightedBidirectionalEdge(chicago, denver, 1001)
	g.AddWeightedBidirectionalEdge(dallas, denver, 795)
	g.AddWeightedBidirectionalEdge(seattle, sanFrancisco, 807)
	g.AddWeightedBidirectionalEdge(seattle, denver, 1303)
	g.AddWeightedBidirectionalEdge(denver, sanFrancisco, 1248)
	g.AddWeightedBidirectionalEdge(denver, losAngeles, 1098)
	g.AddWeightedBidirectionalEdge(sanFrancisco, losAngeles, 382)

	path := g.AStarShortestPath(newYork, sanFrancisco, haversine)
	for _, node := range path {
		fmt.Printf("%s  " , node.value.(testCity).name)
	}
}
