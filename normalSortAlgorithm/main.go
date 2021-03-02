package main

func main() {
	g := NewGraph(8)
	g.addVertex(2, 1)
	g.addVertex(3, 1)
	g.addVertex(7, 1)
	g.addVertex(4, 2)
	g.addVertex(5, 2)
	g.addVertex(8, 7)
	g.KhanSort()
}
