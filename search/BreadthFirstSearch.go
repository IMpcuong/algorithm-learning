package main

import (
	"algorithm-learning/utils"
	"fmt"
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
