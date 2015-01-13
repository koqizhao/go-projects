package graph

import "extensions/strext"

type intVertex struct {
	IVertex
	id int
}

func (vertex intVertex) String() string {
	return strext.ToString(vertex.id)
}

func newIntVertex(id int) *intVertex {
	v := NewVertex()
	r := &intVertex{v, id}
	v.SetReal(r)
	return r
}

func lessThan(x, y interface{}) bool {
	v1 := x.(int)
	v2 := y.(int)
	return v1 < v2
}
