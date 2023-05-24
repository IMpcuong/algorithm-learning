package main

import (
	"algorithm-learning/utils"
	"fmt"
	"math"
)

type Vertex struct {
	label     int
	adjacents []*Vertex
}

type Graph struct {
	vertices []*Vertex
}

func (g *Graph) existsVertexWith(label int) bool {
	for _, v := range g.vertices {
		if v.label == label {
			return true
		}
	}
	return false
}

func (g *Graph) getVertexBy(label int) (*Vertex, error) {
	if !g.existsVertexWith(label) {
		return &Vertex{}, fmt.Errorf("ERROR: Vertex %d not found", label)
	}
	for i, v := range g.vertices {
		if v.label == label {
			return g.vertices[i], nil
		}
	}
	return &Vertex{}, fmt.Errorf("ERROR: Vertex's label %d was out of bound", label)
}

func (g *Graph) AddVertex(label int) error {
	if g.existsVertexWith(label) {
		return fmt.Errorf("ERROR: Vertex %d already exists", label)
	}
	nextVertex := &Vertex{
		label: label,
	}
	g.vertices = append(g.vertices, nextVertex)
	return nil
}

func (g *Graph) AddEdge(srcLabel, destLabel int) error {
	src, _ := g.getVertexBy(srcLabel)
	dest, _ := g.getVertexBy(destLabel)
	if dest == nil || src == nil || srcLabel == destLabel {
		return fmt.Errorf("ERROR: Not a valid edge from [%d] --> [%d]", srcLabel, destLabel)
	} else if utils.Contains(src.adjacents, dest) {
		return fmt.Errorf("WARN: Edge from vertex [%d] --> [%d] already exists", srcLabel, destLabel)
	}
	src.adjacents = append(src.adjacents, dest)
	return nil
}

func (g *Graph) Print() {
	for _, vertex := range g.vertices {
		fmt.Printf("Vertex[%d]:\n", vertex.label)
		for _, v := range vertex.adjacents {
			fmt.Printf("+ Adjacent's label: %d\n", v.label)
		}
		fmt.Println()
	}
}

func exampleDirectedGraph() {
	g := &Graph{}
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)
	g.AddVertex(8)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)

	g.AddEdge(2, 5)
	g.AddEdge(2, 6)

	g.AddEdge(3, 6)

	g.AddEdge(4, 7)
	g.AddEdge(4, 8)

	g.Print()
}

func bfsWithFixedData() {
	weightedGraph := make(map[string]map[string]int)
	weightedGraph["start"] = map[string]int{}
	// weightedGraph["start"]["a"] = 6
	// weightedGraph["start"]["b"] = 2

	// Exp2:
	weightedGraph["start"]["a"] = 5
	weightedGraph["start"]["b"] = 2

	weightedGraph["a"] = map[string]int{}
	// weightedGraph["a"]["fin"] = 1

	// Exp2:
	weightedGraph["a"]["c"] = 4
	weightedGraph["a"]["d"] = 2

	weightedGraph["b"] = map[string]int{}
	// weightedGraph["b"]["a"] = 3
	// weightedGraph["b"]["fin"] = 5

	// Exp2:
	weightedGraph["b"]["a"] = 8
	weightedGraph["b"]["d"] = 7

	weightedGraph["c"] = map[string]int{}
	weightedGraph["c"]["d"] = 6
	weightedGraph["c"]["fin"] = 3

	weightedGraph["d"] = map[string]int{}
	weightedGraph["d"]["fin"] = 1

	weightedGraph["fin"] = map[string]int{}

	costGraph := make(map[string]int)
	// costGraph["a"] = 6
	// costGraph["b"] = 2
	// costGraph["fin"] = math.MaxInt

	// Exp2:
	costGraph["a"] = 5
	costGraph["b"] = 2
	costGraph["c"] = math.MaxInt
	costGraph["d"] = math.MaxInt
	costGraph["fin"] = math.MaxInt

	parentGraph := make(map[string]string)
	// parentGraph["a"] = "start"
	// parentGraph["b"] = "start"
	// parentGraph["fin"] = ""

	// Exp2:
	parentGraph["a"] = "start"
	parentGraph["b"] = "start"
	parentGraph["c"] = ""
	parentGraph["d"] = ""
	parentGraph["fin"] = ""

	findLowestCostNode := func(costGraph map[string]int, processed []string) string {
		lowestCost := math.MaxInt
		lowestCostNode := ""
		for nodeName, cost := range costGraph {
			isNotExist := func(list []string, v string) bool {
				for _, n := range list {
					if n == v {
						return false
					}
				}
				return true
			}(processed, nodeName)
			if lowestCost > cost && isNotExist {
				lowestCost = cost
				lowestCostNode = nodeName
			}
		}
		fmt.Println(lowestCostNode)
		return lowestCostNode
	}

	processedNode := make([]string, 0, len(weightedGraph)-1)
	curNodeName := findLowestCostNode(costGraph, processedNode)
	for curNodeName != "" {
		cost := costGraph[curNodeName]
		neighbors := weightedGraph[curNodeName]
		for neighborName := range neighbors {
			newCost := cost + neighbors[neighborName]
			if costGraph[neighborName] > newCost {
				costGraph[neighborName] = newCost
				parentGraph[neighborName] = curNodeName
			}
		}
		processedNode = append(processedNode, curNodeName)
		curNodeName = findLowestCostNode(costGraph, processedNode)
	}

	fmt.Println(parentGraph)
	fmt.Println(costGraph)
}
