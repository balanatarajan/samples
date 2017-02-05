package main

import (
	"bufio"
	"os"
	"strings"
)

type node struct {
	l, r *node
	val,index  int
}


func formHeap(ar []int) *node{
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


}
