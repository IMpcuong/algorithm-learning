package main

type Node struct {
	Name      string
	Adjacents map[string][]Node
}

type Graph struct {
	Root Node
}

func NewGraph() *Graph {
	adjacents := make(map[string][]Node, 0)
	g := &Graph{Node{"root-node", adjacents}}
	return g
}
