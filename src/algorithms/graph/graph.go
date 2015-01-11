package graph

import "extensions/strext"

type IGraph interface {
    AddVertex(vertex IVertex)
    RemoveVertex(vertex IVertex)
    HasVertex(vertex IVertex) bool
    GetVertices() (vertices []IVertex)
    GetVertexCount() (vertexCount int)
    AddEdge(edge IEdge)
    RemoveEdge(origin, destination IVertex) (edge IEdge)
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
    graph.vertices[id] = vertex
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
    origin, dest := edge.GetOrigin(), edge.GetDestination()
    originId, destId := strext.ToString(origin), strext.ToString(dest)
    graph.vertices[originId], graph.vertices[destId] = origin, dest
    origin.AddOutgoingEdge(edge)
    dest.AddOutgoingEdge(edge)
}

func (graph *Graph) RemoveEdge(origin, destination IVertex) (edge IEdge) {
    originId, destId := strext.ToString(origin), strext.ToString(destination)
    origin, dest := graph.vertices[originId], graph.vertices[destId]
    edge = origin.RemoveOutgoingEdge(dest)
    dest.RemoveIncomingEdge(origin)
    return edge
}

func (graph *Graph) GetEdges() (edges []IEdge) {
    return nil
}

func (graph *Graph) GetEdgeCount() (edgeCount int) {
    return 0
}
