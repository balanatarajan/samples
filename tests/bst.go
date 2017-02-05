package main

// COnstructing a BST
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	value int
	l     *node
	r     *node
}

func newNode(v int) *node {
	n := new(node)
	n.value = v
	n.l, n.r = nil, nil

	return n
}

func addNode(n *node, v int) *node {

	if n == nil {
		return newNode(v)
	}

	if v > n.value {
		n.r = addNode(n.r, v)
	} else if v <= n.value {
		n.l = addNode(n.l, v)
	}

	return n
}

func createBST(a []int) *node {

	var root *node
	for k, v := range a {
		if k == 0 {
			root = newNode(v)
			continue
		}
		addNode(root, v)
	}

	return root
}

func findSubtree(v int, r *node) *node {
	if r == nil {
		return nil
	}

	if v == r.value {
		return r
	}

	if v > r.value {
		return findSubtree(v, r.r)
	} else if v <= r.value {
		return findSubtree(v, r.l)
	}
	return nil
}

func insertNode(v int, r *node) *node {

	if r == nil {
		return newNode(v)
	}

	if v > r.value {
		r.r = insertNode(v, r.r)
	} else if v <= r.value {
		r.l = insertNode(v, r.l)
	}

	return r
}

func printBST(n *node) {
	if n == nil {
		return
	}

	printBST(n.l)
	printBST(n.r)
	fmt.Println(n.value)
}

func main() {

	scanner := bufio.NewReader(os.Stdin)

	str, err := scanner.ReadString('\n')

	if err != nil {
		panic("something wrong")
	}

	// Break the strings
	strs := strings.SplitN(str, " ", -1)

	ins := make([]int, len(strs))

	for k, v := range strs {
		ins[k], _ = strconv.Atoi(strings.TrimSpace(v))
	}

	rt := createBST(ins)
	printBST(rt)

	// n := findSubtree(6, rt)

	rt = insertNode(60, rt)

	fmt.Println("New Tree")
	printBST(rt)
}
