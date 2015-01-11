package graph

import (
    "errors"
    )

type IEdge interface {
    GetOrigin() (origin IVertex)
    GetDestination() (destination IVertex)
    GetVertices() (origin, destination IVertex)
    GetWeight() (weight interface{})
    SetWeight(weight interface{})
}

type Edge struct {
    origin IVertex
    destination IVertex
    weight interface{}
}

func NewEdge(origin, destination IVertex, weight interface{}) (*Edge, error) {
    if origin == nil {
        return nil, errors.New("Argument null: origin")
    }
    if destination == nil {
        return nil, errors.New("Argument null: destination")
    }
    return &Edge{ origin, destination, weight }, nil
}

func (edge *Edge) GetOrigin() IVertex {
    return edge.origin
}

func (edge *Edge) GetDestination() IVertex {
    return edge.destination
}

func (edge *Edge) GetVertices() (origin, destination IVertex) {
    return edge.origin, edge.destination
}

func (edge *Edge) GetWeight() interface{} {
    return edge.weight
}

func (edge *Edge) SetWeight(weight interface{}) {
    edge.weight = weight
}
