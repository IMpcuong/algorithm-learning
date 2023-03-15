package main

type Node struct {
	Label     string
	Adjacents map[string][]Node
}

type Graph struct {
	Root Node
}

func NewGraph() (g *Graph) {
	adjacents := make(map[string][]Node, 0)
	g = &Graph{Node{"root", adjacents}}
	return
}

func (g *Graph) AddVertexToRoot(v Node) *Graph {
	rootLabel := g.Root.Label
	if rootLabel == v.Label {
		return g
	}
	for _, rootChild := range g.Root.Adjacents[rootLabel] {
		if rootChild.Label == v.Label {
			return g
		}
	}
	g.Root.Adjacents[v.Label] = v.Adjacents[v.Label]
	return g
}
