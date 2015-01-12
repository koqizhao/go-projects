package graph

import "extensions/strext"

type intVertex struct {
	IVertex
	id int
}

type intEdge struct {
	IEdge
}

type intGraph struct {
	IGraph
}

func (vertex intVertex) String() string {
	return strext.ToString(vertex.id)
}
