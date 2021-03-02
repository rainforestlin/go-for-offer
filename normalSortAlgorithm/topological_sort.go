package main

import (
	"errors"
	"fmt"
)

type graph struct {
	vertex     int
	vertexList map[int][]int
}

func NewGraph(v int) *graph {
	g := new(graph)
	g.vertex = v
	g.vertexList = map[int][]int{}
	for i := 1; i <= v; i++ {
		g.vertexList[i] = make([]int, 0)
	}
	return g

}

func (g *graph) addVertex(k, v int) {
	g.vertexList[k] = push(g.vertexList[k], v)
}

func push(list []int, value int) []int {
	result := append(list, value)
	return result
}

func pop(list []int) (int, []int, error) {
	if len(list) > 0 {
		value := list[0]
		remain := list[1:]
		return value, remain, nil
	} else {
		return -1, list, errors.New("empty list cannot pop out anything")
	}
}

func (g *graph) KhanSort() {
	var inDegree = make(map[int]int)
	var queue []int
	// 计算入度
	for i := 1; i <= g.vertex; i++ {
		for _, v := range g.vertexList[i] {
			inDegree[v]++
		}
	}

	for i := 1; i <= g.vertex; i++ {
		if inDegree[i] == 0 {
			queue = push(queue, i)
		}
	}
	for len(queue) > 0 {
		var now int
		var err error
		now, queue, err = pop(queue)
		if err != nil {
			panic(err)
		}
		fmt.Println("->", now)
		for _, k := range g.vertexList[now] {
			if inDegree[k] >= 0 {
				inDegree[k]--
			}
			if inDegree[k] == 0 {
				queue = push(queue, k)
			}
		}

	}
}
