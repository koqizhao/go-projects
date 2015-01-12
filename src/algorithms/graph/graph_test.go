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
	var edge = intEdge{NewEdge(vertex, vertex, 1, true)}
	actual, expected := strext.ToString(edge), "<1, 1>"
	if actual != expected {
		t.Error("Expected", expected, "Actual", actual)
	}

	var vertex2 = intVertex{NewVertex(), 2}
	var edge2 = intEdge{NewEdge(vertex, vertex2, 1, false)}
	actual, expected = strext.ToString(edge2), "(1, 2)"
	if actual != expected {
		t.Error("Expected", expected, "Actual", actual)
	}
}
