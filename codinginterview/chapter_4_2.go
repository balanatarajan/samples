package main

import (
	"fmt"
	"graph"
)

var (
	visited  = make(map[int]bool)
	distance int
	found    bool
	// Do 4 as a start
	order = make([]int, 4)
)

func Distance(first, second *graph.GraphNode) {

	if first == nil {
		return
	}

	if first.Val == second.Val {
		found = true
		return
	}

	_, ok := visited[first.Val]

	if ok == true {
		return
	}

	visited[first.Val] = true
	order = append(order, first.Val)

	for _, v := range first.Adjacent() {
		distance++
		Distance(v, second)
	}
}

func DepthFirstOrder(rt *graph.GraphNode) {

	if rt == nil {
		return
	}
	v, _ := visited[rt.Val]

	if v == true {
		return
	}

	visited[rt.Val] = true

	fmt.Println(rt.Val)

	for _, val := range rt.Adjacent() {
		DepthFirstOrder(val)
	}
}

func main() {

	rt := graph.New(0)

	rt.Add(2)
	rt.Add(1)

	n := rt.Find(2)
	n.Add(3)
	n.AddNode(rt)

	s := n.Find(3)
	s.AddNode(s)

	// DepthFirstOrder(n)
	// Distance()

}
