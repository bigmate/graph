package graph

import (
	"reflect"
	"testing"
)

func TestFixedHeap_Pop(t *testing.T) {
	var a = newEdge(newUV("O"), newUV("A"), 4)
	var b = newEdge(newUV("O"), newUV("B"), 8)
	var c = newEdge(newUV("O"), newUV("C"), 9)
	var d = newEdge(newUV("O"), newUV("D"), 5)
	var e = newEdge(newUV("O"), newUV("E"), 2)
	var f = newEdge(newUV("O"), newUV("F"), 1)
	var tests = []struct {
		name string
		edge Edge
		pop  Edge
	}{
		{
			name: "a",
			edge: a,
			pop:  f,
		},
		{
			name: "b",
			edge: b,
			pop:  e,
		},
		{
			name: "c",
			edge: c,
			pop:  a,
		},
		{
			name: "d",
			edge: d,
			pop:  d,
		},
		{
			name: "f",
			edge: f,
			pop:  b,
		},
		{
			name: "e",
			edge: e,
			pop:  c,
		},
	}
	var h = NewFixedHeap(len(tests))
	for _, tt := range tests {
		h.Push(tt.edge)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := h.Pop(); !reflect.DeepEqual(got, tt.pop) {
				t.Errorf("Pop() = %v, want %v", got, tt.pop)
			}
		})
	}
	if t.Failed() {
		t.Log(h.repr())
	}
}

func TestFixedHeap_Push(t *testing.T) {
	var a = newEdge(newUV("O"), newUV("A"), 10)
	var b = newEdge(newUV("O"), newUV("B"), 8)
	var c = newEdge(newUV("O"), newUV("C"), 9)
	var d = newEdge(newUV("O"), newUV("D"), 4)
	var e = newEdge(newUV("O"), newUV("E"), 2)
	var f = newEdge(newUV("O"), newUV("F"), 1)
	var tests = []struct {
		name string
		edge Edge
		peek Edge
	}{
		{
			name: "a",
			edge: a,
			peek: a,
		},
		{
			name: "b",
			edge: b,
			peek: b,
		},
		{
			name: "c",
			edge: c,
			peek: b,
		},
		{
			name: "d",
			edge: d,
			peek: d,
		},
		{
			name: "f",
			edge: f,
			peek: f,
		},
		{
			name: "e",
			edge: e,
			peek: f,
		},
	}
	var p = NewFixedHeap(len(tests))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p.Push(tt.edge)
			if !reflect.DeepEqual(p.arr[1], tt.peek) {
				t.Errorf("Expected %s got %s", tt.peek, p.arr[1])
			}
		})
	}
	if t.Failed() {
		t.Log(p.repr())
	}
}
