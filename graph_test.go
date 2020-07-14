package graph

import (
	"testing"
)

func TestGraph_Add_Remove(t *testing.T) {
	var g = NewDirectedGraph()
	var v = &vertex{id: "A"}
	g.Add(v)
	if !g.Has(v) {
		t.Fatal("Vertex has not been added")
	}
	g.Remove(v)
	if len(g) > 0 {
		t.Errorf("Vertex has not been removed\n%s", g.repr())
	}

	// Test Clean Removal
	g.Add(v)
	g.Add(&vertex{id: "B"})
	g.Connect(v, &vertex{id: "B"})
	g.Remove(&vertex{id: "B"})
	var l, err = g.Edges(v)
	if err != nil {
		t.Fatalf("Vertex %s absent", v.Repr())
	}
	if l.head != nil {
		t.Errorf("Connection has not been removed")
	}
	if t.Failed() {
		t.Log(g.repr())
	}
}

func TestGraph_Connect_Disconnect(t *testing.T) {
	var g = NewDirectedGraph()
	var vs = []Vertex{
		&vertex{"A"},
		&vertex{"B"},
		&vertex{"C"},
	}
	g.Add(vs[0])
	g.Add(vs[1])
	g.Add(vs[2])
	var err = g.Connect(vs[0], &vertex{id: "B"})
	if err == nil {
		t.Fatalf("Expected Error")
	}
	err = g.Disconnect(vs[0], vs[1])
	if err != nil {
		t.Fatalf("Disconnect Error: %s", err)
	}
	g.Connect(vs[0], vs[2])
	var l, _ = g.Edges(vs[0])
	if !l.head.val.Equal(vs[2]) {
		t.Errorf("Expected: C, got: %s", l.head.val.Id())
	}
	if t.Failed() {
		t.Log(g.repr())
	}
}

func TestGraph_DFS(t *testing.T) {
	var g = NewDirectedGraph()
	var vs = []Vertex{
		&vertex{"A"},
		&vertex{"B"},
		&vertex{"C"},
		&vertex{"D"},
		&vertex{"E"},
	}
	for _, v := range vs {
		g.Add(v)
	}
	g.Connect(vs[0], vs[1])
	g.Connect(vs[0], vs[4])
	g.Connect(vs[1], vs[2])
	g.Connect(vs[2], vs[3])
	g.Connect(vs[3], vs[4])
	g.Connect(vs[4], vs[3])
	var expected = []string{"A", "E", "D", "B", "C"}
	var j int
	for v := range g.DFS(vs[0]) {
		if j > len(expected)-1 {
			t.Fatal("Unexpected path")
		}
		if expected[j] != v.Id() {
			t.Errorf("Unexpected vertex, expected: %s, got: %s", expected[j], v.Id())
		}
		j++
	}
}

func TestGraph_BFS(t *testing.T) {
	var g = NewDirectedGraph()
	var vs = []Vertex{
		&vertex{"A"},
		&vertex{"B"},
		&vertex{"C"},
		&vertex{"D"},
		&vertex{"E"},
		&vertex{"K"},
		&vertex{"W"},
		&vertex{"Y"},
		&vertex{"X"},
	}
	for _, v := range vs {
		g.Add(v)
	}
	g.Connect(vs[5], vs[0])
	g.Connect(vs[5], vs[1])
	g.Connect(vs[5], vs[2])
	g.Connect(vs[5], vs[3])
	g.Connect(vs[5], vs[4])
	g.Connect(vs[3], vs[6])
	g.Connect(vs[3], vs[7])
	g.Connect(vs[6], vs[8])
	g.Connect(vs[8], vs[4])
	var expected = []string{"K", "A", "B", "C", "D", "E", "W", "Y", "X"}
	var j int
	for v := range g.BFS(vs[5]) {
		if j > len(expected)-1 {
			t.Fatal("Unexpected path")
		}
		if expected[j] != v.Id() {
			t.Errorf("Unexpected vertex, expected: %s, got: %s", expected[j], v.Id())
		}
		j++
	}
	for v := range g.BFS(&vertex{"MK"}) {
		t.Errorf("Unexpected vertex: %s", v.Id())
	}
}

func TestGraph_Sorted(t *testing.T) {
	var g = NewDirectedGraph()
	var vs = []Vertex{
		&vertex{"P"},
		&vertex{"A"},
		&vertex{"B"},
		&vertex{"X"},
		&vertex{"Y"},
		&vertex{"W"},
	}
	for _, v := range vs {
		g.Add(v)
	}
	g.Connect(vs[3], vs[1])
	g.Connect(vs[3], vs[2])
	g.Connect(vs[1], vs[0])
	g.Connect(vs[2], vs[0])
	g.Connect(vs[3], vs[4])
	g.Connect(vs[4], vs[0])
	g.Connect(vs[5], vs[1])

	var ls, err = g.Sorted()
	if err != nil {
		t.Fatal(err)
	}
	if ls[0].Id() != "X" && ls[0].Id() != "W" {
		t.Errorf("Unexpected vertex: %s", ls[0].Id())
	}
	if ls[len(ls)-1].Id() != "P" {
		t.Errorf("Unexpected vertex: %s", ls[len(ls)-1].Id())
	}
	if t.Failed() {
		t.Log(g.repr())
	}
}

func TestGraph_Cyclic(t *testing.T) {
	var g = NewDirectedGraph()
	var vs = []Vertex{
		&vertex{"A"},
		&vertex{"B"},
		&vertex{"C"},
	}
	for _, v := range vs {
		g.Add(v)
	}
	g.Connect(vs[0], vs[1])
	g.Connect(vs[1], vs[2])
	g.Connect(vs[2], vs[0])
	if !g.Cyclic() {
		t.Errorf("Expected DirectedGraph to be cyclic")
	}
	g = NewDirectedGraph()
	g.Add(vs[0])
	g.Connect(vs[0], vs[0])
	if !g.Cyclic() {
		t.Errorf("Expected DirectedGraph to be cyclic")
	}

	if t.Failed() {
		t.Log(g.repr())
	}

	g = NewDirectedGraph()
	g.Add(vs[0])
	g.Connect(vs[0], vs[1])
	g.Connect(vs[0], vs[2])
	g.Connect(vs[1], vs[2])
	if g.Cyclic() {
		t.Errorf("Expected DirectedGraph not to be cyclic")
	}
	if t.Failed() {
		t.Log(g.repr())
	}
}
