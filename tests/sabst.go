package main

// Sorted array as a BST
// Also has code for finding the depth of the tree
// Has code for checking whether the tree is balanced or not

import (
	"bufio"
	"fmt"
		"strings"
	"os"
	"strconv"

)

type node struct {
	l, r *node

	val int
}

func addNode(n *node, v int) *node {
	if n == nil {
		n = new(node)
		n.val = v
		n.l, n.r = nil, nil
		return n
	}

	if v <= n.val {
		n.l = addNode(n.l, v)
	} else if v > n.val {
		n.r = addNode(n.r, v)
	}

	return n
}

func createBST(sa []int) *node {

	rtindex := len(sa) / 2

	root := addNode(nil, sa[rtindex])
	for i := 0; i < rtindex; i++ {
		addNode(root, sa[i])
	}

	for i := rtindex; i < len(sa); i++ {
		addNode(root, sa[i])
	}

	return root
}

func ar2BST(sa []int) *node {

	if len(sa) == 0 {
		return nil
	}
	rtindex := len(sa) / 2

	root := addNode(nil, sa[rtindex])

	if rtindex == 0 {
		return root
	}

	root.l = ar2BST(sa[0:rtindex])
	root.r = ar2BST(sa[rtindex+1:])
	return root
}

func printBST(n *node) {

	if n == nil {
		return
	}

	printBST(n.l)
	fmt.Println(n.val)
	printBST(n.r)

}

func heightBST(n *node) int {

	if n == nil {
		return 0
	}

	lh := heightBST(n.l)
	rh := heightBST(n.r)

	var max int
	if lh > rh {
		max = lh
	} else {
		max = rh
	}

	return max + 1
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	str, err := scanner.ReadString('\n')

	if err != nil {
		panic("something went wrong")
	}

	strs := strings.SplitN(str, " ", -1)

	ints := make([]int, len(strs))

	for i, v := range strs {
		ints[i], _ = strconv.Atoi(strings.TrimSpace(v))
	}

	// rt := createBST(ints)
	rt := ar2BST(ints)

	fmt.Println("Printing back")

	// printBST(rt)

	fmt.Println(heightBST(rt))

	diff := heightBST(rt.l) - heightBST(rt.r)

	if diff < 0 {
		diff = -diff
	}

	if diff > 1 {
		fmt.Println("Not balanced")
	} else {
		fmt.Println("Balanced")
	}

}
