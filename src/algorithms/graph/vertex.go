package graph

import (
	"extensions/strext"
	"fmt"
)

type IVertex interface {
	fmt.Stringer
	AddEdge(edge IEdge)
	GetEdge(other IVertex) (edge IEdge)
	RemoveEdge(edge IEdge)
	GetEdges() (edges []IEdge)
	GetOutDegree() (outDegree int)
	GetInDegree() (inDegree int)
}

type Vertex struct {
	undirectedEdges map[string]IEdge
	outgoingEdges   map[string]IEdge
	incomingEdges   map[string]IEdge
}

func NewVertex() *Vertex {
	return &Vertex{make(map[string]IEdge), make(map[string]IEdge), make(map[string]IEdge)}
}

func (vertex *Vertex) String() string {
	return fmt.Sprintf("%p", vertex)
}

func (vertex *Vertex) AddEdge(edge IEdge) {
	origin, dest := edge.GetVertices()
	thisId, originId, destId := strext.ToString(vertex), strext.ToString(origin), strext.ToString(dest)
	if thisId != originId && thisId != destId {
		panic("The edge is not incident to the vertex.")
	}
	if edge.IsDirected() {
		if thisId == originId {
			vertex.outgoingEdges[destId] = edge
		}

		if thisId == destId {
			vertex.incomingEdges[originId] = edge
		}
	} else {
		otherId := originId
		if thisId == originId {
			otherId = destId
		}
		vertex.undirectedEdges[otherId] = edge
	}
}

func (vertex *Vertex) GetEdge(other IVertex) IEdge {
	otherId := strext.ToString(other)
	if edge, found := vertex.undirectedEdges[otherId]; found {
		return edge
	} else if edge, found := vertex.outgoingEdges[otherId]; found {
		return edge
	} else if edge, found := vertex.incomingEdges[otherId]; found {
		return edge
	} else {
		return nil
	}
}

func (vertex *Vertex) RemoveEdge(edge IEdge) {
	origin, dest := edge.GetVertices()
	thisId, originId, destId := strext.ToString(vertex), strext.ToString(origin), strext.ToString(dest)
	if thisId != originId && thisId != destId {
		panic("The edge is not incident to the vertex.")
	}
	if edge.IsDirected() {
		if thisId == originId {
			if _, found := vertex.outgoingEdges[destId]; found {
				delete(vertex.outgoingEdges, destId)
			}
		}

		if thisId == destId {
			if _, found := vertex.incomingEdges[originId]; found {
				delete(vertex.incomingEdges, originId)
			}
		}
	} else {
		otherId := originId
		if thisId == originId {
			otherId = destId
		}
		if _, found := vertex.undirectedEdges[otherId]; found {
			delete(vertex.undirectedEdges, otherId)
		}
	}
}

func (vertex *Vertex) GetEdges() []IEdge {
	var edges []IEdge
	for _, edge := range vertex.undirectedEdges {
		edges = append(edges, edge)
	}
	for _, edge := range vertex.outgoingEdges {
		edges = append(edges, edge)
	}
	for _, edge := range vertex.incomingEdges {
		edges = append(edges, edge)
	}
	return edges
}

func (vertex *Vertex) GetOutDegree() int {
	return len(vertex.undirectedEdges) + len(vertex.outgoingEdges)
}

func (vertex *Vertex) GetInDegree() int {
	return len(vertex.undirectedEdges) + len(vertex.incomingEdges)
}
