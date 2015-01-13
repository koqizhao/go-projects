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
	HasEdge(edge IEdge) bool
	GetEdges() (edges []IEdge)
	GetEdgeCount() (edgeCount int)
	DepthFirstSearch(visitVertex func(IVertex))
	BreadthFirstSearch(visitVertex func(IVertex))
	GenerateMininumSpanningTreeInPrimWay(lessThan func(interface{}, interface{}) bool) (mst IGraph)
	GenerateMininumSpanningTreeInKruskalWay(lessThan func(interface{}, interface{}) bool) (mst IGraph)
}

type Graph struct {
	vertices map[string]IVertex
}

func NewGraph() *Graph {
	return &Graph{make(map[string]IVertex)}
}

func (graph *Graph) String() string {
	vertices, edges := "", ""
	for _, v := range graph.vertices {
		if vertices != "" {
			vertices += ", "
		}
		vertices += strext.ToString(v)
	}
	for _, e := range graph.GetEdges() {
		if edges != "" {
			edges += ", "
		}
		edges += strext.ToString(e)
	}
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
	if origin != nil && dest != nil {
		origin.RemoveEdge(edge)
		dest.RemoveEdge(edge)
	}
}

func (graph *Graph) HasEdge(edge IEdge) bool {
	origin, dest := edge.GetVertices()
	originId := strext.ToString(origin)
	origin = graph.vertices[originId]
	if origin == nil {
		return false
	}
	return origin.GetEdge(dest) != nil
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

func (graph *Graph) GenerateMininumSpanningTreeInPrimWay(lessThan func(interface{}, interface{}) bool) IGraph {
	mst := NewGraph()
	vertexCount := len(graph.vertices)
	if vertexCount == 0 {
		return mst
	}

	for k, v := range graph.vertices {
		mst.vertices[k] = v.New()
		break
	}

	for len(mst.vertices) < vertexCount {
		var minEdge IEdge
		var existVertex, otherVertex IVertex
		for k, v := range mst.vertices {
			for _, e := range graph.vertices[k].GetEdges() {
				if e.IsDirected() {
					if e.HasOrigin(v) {
						dest := e.GetOpposite(v)
						if mst.HasVertex(dest) {
							continue
						}
						if minEdge == nil || lessThan(e.GetWeight(), minEdge.GetWeight()) {
							minEdge, existVertex, otherVertex = e, v, dest
						}
					}
				} else if dest := e.GetOpposite(v); dest != nil {
					if mst.HasVertex(dest) {
						continue
					}
					if minEdge == nil || lessThan(e.GetWeight(), minEdge.GetWeight()) {
						minEdge, existVertex, otherVertex = e, v, dest
					}
				}
			}
		}
		if minEdge == nil {
			panic("The graph is not a simple connected graph.")
		}
		newOtherVertex := otherVertex.New()
		var newEdge IEdge
		if minEdge.HasOrigin(existVertex) {
			newEdge = minEdge.New(existVertex, newOtherVertex)
		} else {
			newEdge = minEdge.New(newOtherVertex, existVertex)
		}
		mst.AddEdge(newEdge)
	}

	return mst
}

func (graph *Graph) GenerateMininumSpanningTreeInKruskalWay(lessThan func(interface{}, interface{}) bool) IGraph {
	vertexCount := len(graph.vertices)
	if vertexCount <= 1 {
		mst := NewGraph()
		for k, v := range graph.vertices {
			mst.vertices[k] = v.New()
		}
		return mst
	}

	edges := graph.GetEdges()
	trees := []*Graph{}
	for i, l := 0, vertexCount-1; i < l; i++ {
		j := 0
		var minEdge IEdge
		for k, e := range edges {
			if minEdge == nil || lessThan(e.GetWeight(), minEdge.GetWeight()) {
				minEdge = e
				j = k
			}
		}
		if minEdge == nil {
			panic("The graph is not a simple connected graph.")
		}
		newEdges := append([]IEdge{}, edges[0:j]...)
		edges = append(newEdges, edges[j+1:]...)

		origin, dest := minEdge.GetVertices()
		originId, destId := strext.ToString(origin), strext.ToString(dest)
		var originTree, destTree *Graph
		originTreeIndex, destTreeIndex := -1, -1
		for index, tree := range trees {
			if tree.HasVertex(origin) {
				origin, originTree, originTreeIndex = tree.vertices[originId], tree, index
			}
			if tree.HasVertex(dest) {
				dest, destTree, destTreeIndex = tree.vertices[destId], tree, index
			}
		}
		if originTreeIndex != -1 && originTreeIndex == destTreeIndex {
			continue
		}
		if originTreeIndex != -1 && destTreeIndex != -1 {
			for _, e := range originTree.GetEdges() {
				destTree.AddEdge(e)
			}
			destTree.AddEdge(minEdge.New(origin, dest))

			newTrees := append([]*Graph{}, trees[0:originTreeIndex]...)
			trees = append(newTrees, trees[originTreeIndex+1:]...)
		} else if originTreeIndex == -1 && destTreeIndex == -1 {
			tree := NewGraph()
			tree.AddEdge(minEdge.New(nil, nil))
			trees = append(trees, tree)
		} else if originTreeIndex == -1 {
			destTree.AddEdge(minEdge.New(origin.New(), dest))
		} else {
			originTree.AddEdge(minEdge.New(origin, dest.New()))
		}
	}

	return trees[0]
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
