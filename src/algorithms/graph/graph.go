package graph

import "extensions/strext"

type IGraph interface {
    AddVertex(vertex IVertex)
    RemoveVertex(vertex IVertex)
    HasVertex(vertex IVertex) bool
    GetVertices() (vertices []IVertex)
    GetVertexCount() (vertexCount int)
    AddEdge(edge IEdge)
    RemoveEdge(edge IVertex)
    GetEdges() (edges []IEdge)
    GetEdgeCount() (edgeCount int)
}

type Graph struct {
    vertices map[string]IVertex
}

func NewGraph() *Graph {
    return &Graph { make(map[string]IVertex) }
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
