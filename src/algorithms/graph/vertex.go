package graph

import (
    "fmt"
    "extensions/strext"
    )

type IVertex interface {
    fmt.Stringer
    AddOutgoingEdge(outgoingEdge IEdge)
    GetOutgoingEdge(destination IVertex) (edge IEdge)
    RemoveOutgoingEdge(destination IVertex) (edge IEdge)
    GetOutgoingEdges() (outgoingEdges []IEdge)
    GetOutDegree() (outDegrees int)
    AddIncomingEdge(incomingEdge IEdge)
    GetIncomingEdge(origin IVertex) (edge IEdge)
    RemoveIncomingEdge(origin IVertex) (edge IEdge)
    GetIncomingEdges() (incomingEdges []IEdge)
    GetInDegree() (inDegrees int)
}

type Vertex struct {
    outgoingEdges map[string]IEdge
    incomingEdges map[string]IEdge
}

func NewVertex() *Vertex {
    return &Vertex{ make(map[string]IEdge), make(map[string]IEdge) }
}

func (vertex *Vertex) String() string {
    return ""
}

func (vertex *Vertex) AddOutgoingEdge(outgoingEdge IEdge) {
    destId := strext.ToString(outgoingEdge.GetDestination())
    vertex.outgoingEdges[destId] = outgoingEdge
}

func (vertex *Vertex) GetOutgoingEdge(destination IVertex) IEdge {
    return vertex.outgoingEdges[strext.ToString(destination)]
}

func (vertex *Vertex) RemoveOutgoingEdge(destination IVertex) IEdge {
    destId := strext.ToString(destination)
    edge, found := vertex.outgoingEdges[destId]
    if found {
        delete(vertex.outgoingEdges, destId)
    }
    return edge
}

func (vertex *Vertex) GetOutgoingEdges() []IEdge {
    var edges []IEdge
    for _, edge := range vertex.outgoingEdges {
        edges = append(edges, edge)
    }
    return edges
}

func (vertex *Vertex) GetOutDegree() int {
    return len(vertex.outgoingEdges)
}

func (vertex *Vertex) AddIncomingEdge(incomingEdge IEdge) {
    originId := strext.ToString(incomingEdge.GetDestination())
    vertex.incomingEdges[originId] = incomingEdge
}

func (vertex *Vertex) GetIncomingEdge(origin IVertex) IEdge {
    return vertex.incomingEdges[strext.ToString(origin)]
}

func (vertex *Vertex) RemoveIncomingEdge(origin IVertex) IEdge {
    originId := strext.ToString(origin)
    edge, found := vertex.incomingEdges[originId]
    if found {
        delete(vertex.incomingEdges, originId)
    }
    return edge
}

func (vertex *Vertex) GetIncomingEdges() []IEdge {
    var edges []IEdge
    for _, edge := range vertex.incomingEdges {
        edges = append(edges, edge)
    }
    return edges
}

func (vertex *Vertex) GetInDegree() int {
    return len(vertex.incomingEdges)
}
