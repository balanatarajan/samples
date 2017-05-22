package main

import (
	"math"
)

type TreeNode struct {
	l, r *TreeNode
	val  int
}

func height(n *TreeNode) int {

	if n == nil {
		return 0
	}

	lh := height(n.l)
	rh := height(n.r)

	var max int

	if lh >= rh {
		max = lh
	} else if rh > lh {

		max = rh
	}

	return max + 1
}

func (n *TreeNode) isBalanced() bool {
	lh := height(n.l)
	rh := height(n.r)

	if math.Abs(float64(rh-lh)) <= 1 {
		return true
	}

	return false
}

func main() {

	var n *TreeNode
	n.isBalanced()
}
