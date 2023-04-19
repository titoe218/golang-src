package main

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	visited  bool
	neightbors []*Vertex
}

func (g *Graph) AddVertex(key int) {
	vertex := &Vertex{key: key}
	g.vertices = append(g.vertices, vertex)
}

func (g *Graph) AddEdge(from, to int) {
	v1 := g.getVertex(from)
	v2 := g.getVertex(to)

	v1.neightbors = append(v1.neightbors, v2)
	v2.neightbors = append(v2.neightbors, v1)
}

func (g *Graph) getVertex(key int) *Vertex {
	for i := 0; i < len(g.vertices); i++ {
		if g.vertices[i].key == key {
			return g.vertices[i]
		}
	}

	return nil
}

func (g *Graph) bfs(start *Vertex) {
	queue := []*Vertex{}
	queue = append(queue, start)
	start.visited = true

	for len(queue) > 0 {
		// Pop vertex from queue
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", node.key)

		// Add neighbors to queue
		for i := 0; i < len(node.neightbors); i++ {
			if !node.neightbors[i].visited {
				node.neightbors[i].visited = true
				queue = append(queue, node.neightbors[i])
			}
		}
	}
}

func main() {
	graph := &Graph{}
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(2, 5)

	fmt.Println("Breath First Traversal (starting from vertex 1):")
	graph.bfs(graph.getVertex(1))
}
