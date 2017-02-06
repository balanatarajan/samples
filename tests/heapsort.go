package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	l, r     *node
	idx, val int
}

func maxifyHeap(nodes []*node) ([]*node, bool) {

	changed := false
	st := len(nodes) / 2

	for i := st - 1; i >= 0; i-- {

		if nodes[i].l != nil && nodes[i].l.val > nodes[i].val {
			nodes[i].val, nodes[i].l.val = nodes[i].l.val, nodes[i].val
			changed = true
		}

		if nodes[i].r != nil && nodes[i].r.val > nodes[i].val {
			nodes[i].val, nodes[i].r.val = nodes[i].r.val, nodes[i].val
			changed = true
		}
	}

	return nodes, changed
}

func main() {

	scanner := bufio.NewReader(os.Stdin)
	str, err := scanner.ReadString('\n')

	if err != nil {
		panic("something went wrong")
	}

	strs := strings.SplitN(str, " ", -1)

	nodes := make([]*node, len(strs))

	for i, v := range strs {
		g, _ := strconv.Atoi(strings.TrimSpace(v))
		tmp := node{val: g, idx: i + 1}

		nodes[i] = &tmp
	}

	for i, v := range nodes {

		fmt.Println(v)
		if len(nodes) >= (2*(i+1) + 1) {
			t := 2 * (i + 1)
			v.l = nodes[t-1]
			v.r = nodes[t]
		}

	}

	var changed bool
	for {
		nodes, changed = maxifyHeap(nodes[0:])
		if changed == false {
			break
		}

		fmt.Println("here I go")
		for _, v := range nodes {
			fmt.Println(v)
		}
	}

}
