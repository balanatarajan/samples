// Given a binary tree (not a BST), find the inorder successor of a
// node.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	val int
	parent,
	left,
	right *TreeNode
}

func createBT(ar []int, p *TreeNode) *TreeNode {
	if len(ar) == 0 {
		return nil
	}

	root := len(ar) / 2

	retval := &TreeNode{val: ar[root], parent: p}

	lt := ar[:root]
	retval.left = createBT(lt, retval)

	rt := ar[root+1 : len(ar)]
	retval.right = createBT(rt, retval)

	return retval
}

/*
func findNode(rt *TreeNode, n int) *TreeNode {
	if rt == nil {
		return nil
	}

	left := findNode(rt.left, n)

	if left != nil {
		return left
	}

	if rt.val == n {
		return rt
	}

	right := findNode(rt.right, n)
	if right != nil {
		return right
	}

	return nil
}
*/

// 4.5 in career cup
func inOrderSuccessor(n *TreeNode) *TreeNode {

	if n == nil {
		return nil
	}

	return nil
}

func printBT(bt *TreeNode) {

	// fmt.Println("Started")
	if bt == nil {
		return
	}
	printBT(bt.left)
	fmt.Println(bt.val)
	printBT(bt.right)
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

	rt := createBT(ins, nil)
	printBT(rt)

	fmt.Println(findNode(rt, 6))
}
