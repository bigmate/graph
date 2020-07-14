package graph

import (
	"testing"
)

func TestUWGraph_Add(t *testing.T) {
	var g = NewUWGraph()
	type args struct {
		v UVertex
	}
	tests := []struct {
		name string
		g    UWGraph
		args args
	}{
		{
			name: "A",
			g:    g,
			args: args{v: newUV("A")},
		},
		{
			name: "B",
			g:    g,
			args: args{v: newUV("B")},
		},
		{
			name: "C",
			g:    g,
			args: args{v: newUV("C")},
		},
		{
			name: "D",
			g:    g,
			args: args{v: newUV("D")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Add(tt.args.v)
		})
	}
	for _, tt := range tests {
		if !tt.g.Has(tt.args.v) {
			t.Errorf("Expected %s", tt.args.v.Repr())
		}
	}
}

func TestUWGraph_Connect(t *testing.T) {
	var g = NewUWGraph()
	var vertices = []UVertex{
		newUV("A"),
		newUV("B"),
		newUV("C"),
		newUV("D"),
		newUV("E"),
	}
	for _, v := range vertices {
		g.Add(v)
	}
	type args struct {
		from   UVertex
		to     UVertex
		weight float64
	}
	tests := []struct {
		name    string
		g       UWGraph
		args    args
		wantErr bool
	}{
		{
			name: "A->W",
			g:    g,
			args: args{
				from:   newUV("A"),
				to:     newUV("W"),
				weight: 1,
			},
			wantErr: true,
		},
		{
			name: "A->B",
			g:    g,
			args: args{
				from:   newUV("A"),
				to:     newUV("B"),
				weight: 2,
			},
			wantErr: false,
		},
		{
			name: "B->E",
			g:    g,
			args: args{
				from:   newUV("B"),
				to:     newUV("E"),
				weight: 1,
			},
			wantErr: false,
		},
		{
			name: "Q->A",
			g:    g,
			args: args{
				from:   newUV("Q"),
				to:     newUV("A"),
				weight: 1,
			},
			wantErr: true,
		},
		{
			name: "E->A",
			g:    g,
			args: args{
				from:   newUV("E"),
				to:     newUV("A"),
				weight: 5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.Connect(tt.args.from, tt.args.to, tt.args.weight); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if !g.Connected(vertices[0], vertices[1]) {
		t.Errorf("Connection between %s and %s has not been established",
			vertices[0].Repr(), vertices[1].Repr())
	}
	if t.Failed() {
		t.Log(g.repr())
	}
}

func TestUWGraph_Disconnect(t *testing.T) {
	var g = NewUWGraph()
	var vertices = []UVertex{
		newUV("A"),
		newUV("B"),
		newUV("C"),
	}
	for _, v := range vertices {
		g.Add(v)
	}
	g.Connect(newUV("A"), newUV("B"), 2)
	g.Connect(newUV("B"), newUV("C"), 2)

	type args struct {
		from UVertex
		to   UVertex
	}
	tests := []struct {
		name string
		g    UWGraph
		args args
	}{
		{
			name: "A<->B",
			g:    g,
			args: args{
				from: newUV("A"),
				to:   newUV("B"),
			},
		},
		{
			name: "W<->B",
			g:    g,
			args: args{
				from: newUV("W"),
				to:   newUV("B"),
			},
		},
		{
			name: "B<->C",
			g:    g,
			args: args{
				from: newUV("B"),
				to:   newUV("C"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Disconnect(tt.args.from, tt.args.to)
			if g.Connected(tt.args.from, tt.args.to) {
				t.Errorf("%s is still connected to %s", tt.args.from.Repr(), tt.args.to.Repr())
			}
		})
	}
	if t.Failed() {
		t.Log(g.repr())
	}
}

func TestUWGraph_Has(t *testing.T) {
	var g = NewUWGraph()
	var vertices = []UVertex{
		newUV("A"),
		newUV("B"),
		newUV("C"),
	}
	for _, v := range vertices {
		g.Add(v)
	}
	type args struct {
		v UVertex
	}
	tests := []struct {
		name string
		g    UWGraph
		args args
		want bool
	}{
		{
			name: "A",
			g:    g,
			args: args{v: newUV("A")},
			want: true,
		},
		{
			name: "B",
			g:    g,
			args: args{v: newUV("C")},
			want: true,
		},
		{
			name: "C",
			g:    g,
			args: args{v: newUV("C")},
			want: true,
		},
		{
			name: "D",
			g:    g,
			args: args{v: newUV("D")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Has(tt.args.v); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUWGraph_Path(t *testing.T) {
	var g = NewUWGraph()
	var vertices = []UVertex{
		newUV("A"),
		newUV("B"),
		newUV("C"),
		newUV("D"),
		newUV("E"),
	}
	for _, v := range vertices {
		g.Add(v)
	}
	g.Connect(newUV("A"), newUV("B"), 5)
	g.Connect(newUV("A"), newUV("D"), 2)
	g.Connect(newUV("D"), newUV("E"), 1)
	g.Connect(newUV("B"), newUV("E"), 3)
	g.Connect(newUV("E"), newUV("C"), 4)
	type args struct {
		from UVertex
		to   UVertex
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "A-Q",
			args: args{
				from: newUV("A"),
				to:   newUV("Q"),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "A-C",
			args: args{
				from: newUV("A"),
				to:   newUV("C"),
			},
			want:    7,
			wantErr: false,
		},
		{
			name: "C-C",
			args: args{
				from: newUV("C"),
				to:   newUV("C"),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "D-B",
			args: args{
				from: newUV("D"),
				to:   newUV("B"),
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := g.Path(tt.args.from, tt.args.to)
		t.Run(tt.name, func(t *testing.T) {
			if (err != nil) != tt.wantErr {
				t.Errorf("Path() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.Weight() != tt.want {
				t.Errorf("Path() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUWGraph_Cyclic(t *testing.T) {
	var g = NewUWGraph()
	var vertices = []UVertex{
		newUV("A"),
		newUV("B"),
		newUV("C"),
		newUV("D"),
		newUV("E"),
	}
	for _, v := range vertices {
		g.Add(v)
	}
	g.Connect(newUV("A"), newUV("B"), 5)
	g.Connect(newUV("A"), newUV("D"), 2)
	g.Connect(newUV("D"), newUV("E"), 1)
	g.Connect(newUV("B"), newUV("E"), 3)
	g.Connect(newUV("E"), newUV("C"), 4)
	if !g.Cyclic() {
		t.Error("Expected cyclic graph")
	}
	g.Disconnect(newUV("D"), newUV("A"))
	if g.Cyclic() {
		t.Error("Expected non-cyclic graph")
	}
	if t.Failed() {
		t.Log(g.repr())
	}
}

func TestUWGraph_MinTree(t *testing.T) {
	var g = NewUWGraph()
	var vertices = []UVertex{
		newUV("A"),
		newUV("B"),
		newUV("C"),
		newUV("D"),
	}
	for _, v := range vertices {
		g.Add(v)
	}
	g.Connect(newUV("A"), newUV("B"), 2)
	g.Connect(newUV("A"), newUV("D"), 7)
	g.Connect(newUV("D"), newUV("B"), 3)
	g.Connect(newUV("B"), newUV("C"), 4)
	g.Connect(newUV("D"), newUV("C"), 6)
	var tree = g.MinTree()
	var tests = []struct {
		from UVertex
		to   UVertex
		want bool
	}{
		{
			from: newUV("A"),
			to:   newUV("B"),
			want: true,
		},
		{
			from: newUV("C"),
			to:   newUV("D"),
			want: false,
		},
		{
			from: newUV("A"),
			to:   newUV("D"),
			want: false,
		},
		{
			from: newUV("C"),
			to:   newUV("B"),
			want: true,
		},
	}
	for _, tt := range tests {
		if tree.Connected(tt.from, tt.to) != tt.want {
			t.Errorf("Connected(%s, %s) != %v", tt.from.Id(), tt.to.Id(), tt.want)
		}
	}
	if t.Failed() {
		t.Logf("\n%s", tree.repr())
	}
}
