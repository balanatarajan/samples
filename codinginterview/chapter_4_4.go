package main

import (
	"container/list"
	"graph"
	"util"
)

var (
	mp    = make(map[int]bool)
	lists = make([]container.List)
)

func BFS(nd *graph.GraphNode) {

	if nd == nil {
		return
	}

	ls := container.List.New()
	for _, v := range nd.Adjacent() {

		_, ok := mp[v.Val]
		if ok == true {
			continue
		}

		mp[v.Val] = true
		ls.PushBack(v)
	}

	lists := append(lists, ls)

	for _, v := range nd.Adjacent() {
		BFS(v)
	}
}

func main() {

}
