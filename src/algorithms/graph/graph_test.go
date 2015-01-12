package graph

import (
	"extensions/strext"
	"fmt"
	"testing"
)

func TestVertexToString(t *testing.T) {
	var vertex IVertex = NewVertex()
	actual, expected := strext.ToString(vertex), fmt.Sprintf("%p", vertex)
	if actual != expected {
		t.Error("Expected", expected, "Actual", actual)
	}

	var vertex2 = intVertex{NewVertex(), 1}
	actual, expected = strext.ToString(vertex2), "1"
	if actual != expected {
		t.Error("Expected", expected, "Actual", actual)
	}
}

func TestEdgeToString(t *testing.T) {
	var vertex = intVertex{NewVertex(), 1}
	var edge = NewEdge(vertex, vertex, 1, true)
	actual, expected := strext.ToString(edge), "<1, 1>"
	if actual != expected {
		t.Error("Expected", expected, "Actual", actual)
	}

	var vertex2 = intVertex{NewVertex(), 2}
	var edge2 = NewEdge(vertex, vertex2, 1, false)
	actual, expected = strext.ToString(edge2), "(1, 2)"
	if actual != expected {
		t.Error("Expected", expected, "Actual", actual)
	}
}

func TestConstructAGraph(t *testing.T) {
	var vertices []IVertex
	for i := 0; i < 4; i++ {
		vertices = append(vertices, newIntVertex(i))
	}

	var edges = []IEdge{
		NewEdge(vertices[0], vertices[1], 1, false),
		NewEdge(vertices[1], vertices[2], 2, false),
		NewEdge(vertices[2], vertices[3], 3, false),
		NewEdge(vertices[3], vertices[0], 4, false),
		NewEdge(vertices[1], vertices[3], 5, false),
		NewEdge(vertices[2], vertices[0], 6, false),
	}

	var graph IGraph = NewGraph()
	for _, v := range vertices {
		graph.AddVertex(v)
	}
	for _, e := range edges {
		graph.AddEdge(e)
	}

	var s = strext.ToString(graph)
	fmt.Println(s)
}
