package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makeBST(ints []int) *node {
}

func main() {

	scanner := bufio.NewReader()
	str, err := scanner.ReadString('\n')

	if err != nil {
		panic("something on")
	}

	strs, _ := strings.SplitN(str, " ", -1)

	ints := make([]ints, len(strs))

	for k, v := range strs {

		ints[k] = strconv.Atoi(strings.TrimSpace(v))
	}

}
