package lib

import (
	"fmt"
	"strconv"
)

// A Graph is the interface implemented by graphs that
// this algorithm can run on.
type Graph interface {
	Vertices() []Vertex
	Neighbors(v Vertex) []Vertex
	Weight(u, v Vertex) int
}

// Nonnegative integer ID of vertex
type Vertex int

// MutableGraph is a graph of integers that satisfies the Graph interface.
type MutableGraph struct {
	vert  []Vertex
	edges map[Vertex]map[Vertex]int
}

func NewGraph() MutableGraph {
	return MutableGraph{vert: nil, edges: map[Vertex]map[Vertex]int{}}
}

func (g *MutableGraph) Add(a, b int, weight int) {
	if !g.has(a) {
		g.vert = append(g.vert, Vertex(a))
	}
	if !g.has(b) {
		g.vert = append(g.vert, Vertex(b))
	}
	g.edge(Vertex(a), Vertex(b), weight)
}

func (g *MutableGraph) has(a int) bool {
	for _, v := range g.vert {
		if v == Vertex(a) {
			return true
		}
	}
	return false
}

func (g MutableGraph) edge(u, v Vertex, w int) {
	if _, ok := g.edges[u]; !ok {
		g.edges[u] = make(map[Vertex]int)
	}
	g.edges[u][v] = w
}
func (g MutableGraph) Vertices() []Vertex { return g.vert }
func (g MutableGraph) Neighbors(v Vertex) (vs []Vertex) {
	for k := range g.edges[v] {
		vs = append(vs, k)
	}
	return vs
}
func (g MutableGraph) Weight(u, v Vertex) int { return g.edges[u][v] }
func (g MutableGraph) path(vv []Vertex) (s string) {
	if len(vv) == 0 {
		return ""
	}
	s = strconv.Itoa(int(vv[0]))
	for _, v := range vv[1:] {
		s += " -> " + strconv.Itoa(int(v))
	}
	return s
}

const Infinity = int(^uint(0) >> 1)

func FloydWarshall(g Graph) (dist map[Vertex]map[Vertex]int, next map[Vertex]map[Vertex]*Vertex) {
	vert := g.Vertices()
	dist = make(map[Vertex]map[Vertex]int)
	next = make(map[Vertex]map[Vertex]*Vertex)
	for _, u := range vert {
		dist[u] = make(map[Vertex]int)
		next[u] = make(map[Vertex]*Vertex)
		for _, v := range vert {
			dist[u][v] = Infinity
		}
		dist[u][u] = 0
		for _, v := range g.Neighbors(u) {
			v := v
			dist[u][v] = g.Weight(u, v)
			next[u][v] = &v
		}
	}
	for _, k := range vert {
		for _, i := range vert {
			for _, j := range vert {
				if dist[i][k] < Infinity && dist[k][j] < Infinity {
					if dist[i][j] > dist[i][k]+dist[k][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
						next[i][j] = next[i][k]
					}
				}
			}
		}
	}
	return dist, next
}

func Path(u, v Vertex, next map[Vertex]map[Vertex]*Vertex) (path []Vertex) {
	if next[u][v] == nil {
		return
	}
	path = []Vertex{u}
	for u != v {
		u = *next[u][v]
		path = append(path, u)
	}
	return path
}

func ExampleFloyd() {
	g := MutableGraph{[]Vertex{1, 2, 3, 4}, make(map[Vertex]map[Vertex]int)}
	g.edge(1, 3, -2)
	g.edge(3, 4, 2)
	g.edge(4, 2, -1)
	g.edge(2, 1, 4)
	g.edge(2, 3, 3)

	dist, next := FloydWarshall(g)
	fmt.Println("pair\tdist\tpath")
	for u, m := range dist {
		for v, d := range m {
			if u != v {
				fmt.Printf("%d -> %d\t%3d\t%s\n", u, v, d, g.path(Path(u, v, next)))
			}
		}
	}
}
