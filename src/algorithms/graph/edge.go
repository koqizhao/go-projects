package graph

import "fmt"

type IEdge interface {
    fmt.Stringer
    IsDirected() bool
    GetVertices() (origin, destination IVertex)
    GetWeight() (weight interface{})
    SetWeight(weight interface{})
}

type Edge struct {
    origin IVertex
    destination IVertex
    weight interface{}
    isDirected bool
}

func NewEdge(origin, destination IVertex, weight interface{}, isDirected bool) *Edge {
    if origin == nil {
        panic("Argument null: origin")
    }
    if destination == nil {
        panic("Argument null: destination")
    }
    return &Edge{ origin, destination, weight, isDirected }
}

func (edge *Edge) String() string {
    if edge.isDirected {
        return fmt.Sprintf("<%s, %s>", edge.origin, edge.destination)
    } else {
        return fmt.Sprintf("(%s, %s)", edge.origin, edge.destination)
    }
}

func (edge *Edge) GetVertices() (IVertex, IVertex) {
    return edge.origin, edge.destination
}

func (edge *Edge) GetWeight() interface{} {
    return edge.weight
}

func (edge *Edge) SetWeight(weight interface{}) {
    edge.weight = weight
}

func (edge *Edge) IsDirected() bool {
    return edge.isDirected
}
