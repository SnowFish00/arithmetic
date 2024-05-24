package main

import (
	"fmt"
)

// Edge 定义了图中的边
type Edge struct {
	To     *Vertex // 边指向的节点
	Weight int     // 边的权重，如果没有权重可以为0或忽略
}

// Vertex 定义了图中的节点
type Vertex struct {
	Key   string  // 节点的唯一标识
	Edges []*Edge // 与该节点相连的边列表
}

// Graph 表示一个图
type Graph struct {
	Vertices map[string]*Vertex // 所有节点的映射
}

// NewGraph 创建一个新的图
func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[string]*Vertex),
	}
}

// AddVertex 添加一个新的节点到图中
func (g *Graph) AddVertex(key string) {
	if _, found := g.Vertices[key]; !found {
		g.Vertices[key] = &Vertex{
			Key:   key,
			Edges: make([]*Edge, 0),
		}
	}
}

// AddEdge 添加一条边到图中，连接两个节点
func (g *Graph) AddEdge(from, to string, weight int) {
	fromVertex := g.Vertices[from]
	toVertex := g.Vertices[to]

	// 确保两个节点都存在
	if fromVertex == nil || toVertex == nil {
		fmt.Println("One of the vertices does not exist.")
		return
	}

	// 添加边
	fromVertex.Edges = append(fromVertex.Edges, &Edge{
		To:     toVertex,
		Weight: weight,
	})
}

// RemoveVertex 从图中移除一个节点及其所有关联的边
func (g *Graph) RemoveVertex(key string) {
	vertex, found := g.Vertices[key]
	if !found {
		fmt.Println("Vertex not found.")
		return
	}

	// 移除所有指向该节点的边
	for _, v := range g.Vertices {
		for i, e := range v.Edges {
			if e.To == vertex {
				v.Edges = append(v.Edges[:i], v.Edges[i+1:]...)
				break
			}
		}
	}

	// 移除节点
	delete(g.Vertices, key)
}

// RemoveEdge 移除两个节点之间的边
func (g *Graph) RemoveEdge(from, to string) {
	vertex, found := g.Vertices[from]
	if !found {
		fmt.Println("From vertex not found.")
		return
	}

	for i, e := range vertex.Edges {
		if e.To.Key == to {
			vertex.Edges = append(vertex.Edges[:i], vertex.Edges[i+1:]...)
			return
		}
	}

	fmt.Println("Edge not found.")
}

// FindVertex 查找一个节点
func (g *Graph) FindVertex(key string) *Vertex {
	vertex, found := g.Vertices[key]
	if !found {
		fmt.Println("Vertex not found.")
		return nil
	}
	return vertex
}

// FindEdge 查找两个节点之间的边
func (g *Graph) FindEdge(from, to string) *Edge {
	vertex := g.Vertices[from]
	if vertex == nil {
		fmt.Println("From vertex not found.")
		return nil
	}

	for _, e := range vertex.Edges {
		if e.To.Key == to {
			return e
		}
	}

	fmt.Println("Edge not found.")
	return nil
}

func main() {
	g := NewGraph()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")

	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "C", 2)

	vertex := g.FindVertex("A")
	if vertex != nil {
		fmt.Printf("Found vertex: %s\n", vertex.Key)
	}

	edge := g.FindEdge("A", "B")
	if edge != nil {
		fmt.Printf("Found edge from %s to %s with weight %d\n", "A", "B", edge.Weight)
	}

	g.RemoveEdge("A", "B")
	g.RemoveVertex("C")

	// 打印图
	for _, v := range g.Vertices {
		fmt.Printf("%s -> ", v.Key)
		for _, e := range v.Edges {
			fmt.Printf("%s ", e.To.Key)
		}
		fmt.Println()
	}
}
