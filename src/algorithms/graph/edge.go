package graph

import "fmt"
import "extensions/strext"

type IEdge interface {
	fmt.Stringer
	IsDirected() bool
	GetWeight() (weight interface{})
	SetWeight(weight interface{})
	GetVertices() (origin, destination IVertex)
	GetOpposite(vertex IVertex) (opposite IVertex)
	HasVertex(vertex IVertex) bool
	HasOrigin(origin IVertex) bool
	HasDestination(destination IVertex) bool
}

type Edge struct {
	origin      IVertex
	destination IVertex
	weight      interface{}
	isDirected  bool
}

func NewEdge(origin, destination IVertex, weight interface{}, isDirected bool) *Edge {
	if origin == nil {
		panic("Argument null: origin")
	}
	if destination == nil {
		panic("Argument null: destination")
	}
	return &Edge{origin, destination, weight, isDirected}
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

func (edge *Edge) GetOpposite(vertex IVertex) (opposite IVertex) {
	if edge.HasOrigin(vertex) {
		return edge.origin
	} else if edge.HasDestination(vertex) {
		return edge.destination
	} else {
		return nil
	}
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

func (edge *Edge) HasVertex(vertex IVertex) bool {
	return edge.HasOrigin(vertex) || edge.HasDestination(vertex)
}

func (edge *Edge) HasOrigin(origin IVertex) bool {
	return strext.ToString(origin) == strext.ToString(edge.origin)
}

func (edge *Edge) HasDestination(destination IVertex) bool {
	return strext.ToString(destination) == strext.ToString(edge.destination)
}
