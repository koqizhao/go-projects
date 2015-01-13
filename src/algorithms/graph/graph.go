package graph

import "fmt"
import "extensions/strext"

type IGraph interface {
	fmt.Stringer
	AddVertex(vertex IVertex)
	RemoveVertex(vertex IVertex)
	HasVertex(vertex IVertex) bool
	GetVertices() (vertices []IVertex)
	GetVertexCount() (vertexCount int)
	AddEdge(edge IEdge)
	RemoveEdge(edge IEdge)
	GetEdges() (edges []IEdge)
	GetEdgeCount() (edgeCount int)
	DepthFirstSearch(visitVertex func(IVertex))
	BreadthFirstSearch(visitVertex func(IVertex))
}

type Graph struct {
	vertices map[string]IVertex
}

func NewGraph() *Graph {
	return &Graph{make(map[string]IVertex)}
}

func (graph *Graph) String() string {
	vertices, edges := "", ""
	visitVertex := func(v IVertex) {
		if vertices != "" {
			vertices += ", "
		}
		vertices += strext.ToString(v)
	}
	visitEdge := func(e IEdge) {
		if edges != "" {
			edges += ", "
		}
		edges += strext.ToString(e)
	}
	graph.depthFirstSearchWithEdge(visitVertex, visitEdge)
	return fmt.Sprintf("Vertices:\n%s\nEdges:\n%s", vertices, edges)
}

func (graph *Graph) AddVertex(vertex IVertex) {
	id := strext.ToString(vertex)
	if _, found := graph.vertices[id]; !found {
		graph.vertices[id] = vertex
	}
}

func (graph *Graph) RemoveVertex(vertex IVertex) {
	id := strext.ToString(vertex)
	if _, found := graph.vertices[id]; found {
		delete(graph.vertices, id)
	}
}

func (graph *Graph) HasVertex(vertex IVertex) bool {
	id := strext.ToString(vertex)
	_, found := graph.vertices[id]
	return found
}

func (graph *Graph) GetVertices() []IVertex {
	var vertices []IVertex
	for _, vertex := range graph.vertices {
		vertices = append(vertices, vertex)
	}
	return vertices
}

func (graph *Graph) GetVertexCount() int {
	return len(graph.vertices)
}

func (graph *Graph) AddEdge(edge IEdge) {
	origin, dest := edge.GetVertices()
	originId, destId := strext.ToString(origin), strext.ToString(dest)
	existedOrigin, existedDest := graph.vertices[originId], graph.vertices[destId]
	if existedOrigin == nil && existedDest == nil {
		graph.vertices[originId], graph.vertices[destId] = origin, dest
	} else if existedOrigin == nil {
		graph.vertices[originId] = origin
		dest = existedDest
	} else if existedDest == nil {
		graph.vertices[destId] = dest
		origin = existedOrigin
	} else {
		origin, dest = existedOrigin, existedDest
	}

	origin.AddEdge(edge)
	dest.AddEdge(edge)
}

func (graph *Graph) RemoveEdge(edge IEdge) {
	origin, dest := edge.GetVertices()
	originId, destId := strext.ToString(origin), strext.ToString(dest)
	origin, dest = graph.vertices[originId], graph.vertices[destId]
	if origin != nil || dest != nil {
		origin.RemoveEdge(edge)
		dest.RemoveEdge(edge)
	}
}

func (graph *Graph) GetEdges() (edges []IEdge) {
	m := map[string]IEdge{}
	for _, vertex := range graph.vertices {
		for _, edge := range vertex.GetEdges() {
			m[strext.ToString(edge)] = edge
		}
	}
	for _, edge := range m {
		edges = append(edges, edge)
	}
	return edges
}

func (graph *Graph) GetEdgeCount() (edgeCount int) {
	return len(graph.GetEdges())
}

func (graph *Graph) BreadthFirstSearch(visitVertex func(IVertex)) {
	visited := []string{}
	for _, v := range graph.vertices {
		visited = bfs(v, visited, visitVertex, nil)
	}
}

func (graph *Graph) breadthFirstSearchWithEdge(visitVertex func(IVertex), visitEdge func(IEdge)) {
	visited := []string{}
	for _, v := range graph.vertices {
		visited = bfs(v, visited, visitVertex, visitEdge)
	}
}

func (graph *Graph) DepthFirstSearch(visitVertex func(IVertex)) {
	visited := []string{}
	for _, v := range graph.vertices {
		visited = dfs(v, visited, visitVertex, nil)
	}
}

func (graph *Graph) depthFirstSearchWithEdge(visitVertex func(IVertex), visitEdge func(IEdge)) {
	visited := []string{}
	for _, v := range graph.vertices {
		visited = dfs(v, visited, visitVertex, visitEdge)
	}
}

func dfs(v IVertex, visited []string, visitVertex func(IVertex), visitEdge func(IEdge)) []string {
	k := strext.ToString(v)
	if strext.Contains(visited, k) {
		return visited
	}
	visited = append(visited, k)
	if visitVertex != nil {
		visitVertex(v)
	}
	for _, e := range v.GetEdges() {
		origin, dest := e.GetVertices()
		originId, destId := strext.ToString(origin), strext.ToString(dest)
		var other IVertex
		if originId == k {
			if strext.Contains(visited, destId) {
				continue
			}
			other = dest
		} else if e.IsDirected() {
			continue
		} else {
			if strext.Contains(visited, originId) {
				continue
			}
			other = origin
		}
		if visitEdge != nil {
			visitEdge(e)
		}
		visited = dfs(other, visited, visitVertex, visitEdge)
	}
	return visited
}

func bfs(v IVertex, visited []string, visitVertex func(IVertex), visitEdge func(IEdge)) []string {
	k := strext.ToString(v)
	if !strext.Contains(visited, k) {
		visited = append(visited, k)
		if visitVertex != nil {
			visitVertex(v)
		}
	}
	others := []IVertex{}
	for _, e := range v.GetEdges() {
		origin, dest := e.GetVertices()
		originId, destId := strext.ToString(origin), strext.ToString(dest)
		var other IVertex
		var otherId string
		if originId == k {
			if strext.Contains(visited, destId) {
				continue
			}
			other, otherId = dest, destId
		} else if e.IsDirected() {
			continue
		} else {
			if strext.Contains(visited, originId) {
				continue
			}
			other, otherId = origin, originId
		}
		visited = append(visited, otherId)
		if visitVertex != nil {
			visitVertex(other)
		}
		if visitEdge != nil {
			visitEdge(e)
		}
		others = append(others, other)
	}
	for _, v := range others {
		visited = bfs(v, visited, visitVertex, visitEdge)
	}
	return visited
}
