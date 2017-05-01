package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func prod(ar []int) int {

	retval := 1

	for _, v := range ar {
		retval = retval * v
	}

	return retval
}

func bigProd(ar []int) (int, int) {

	s, e, max := 0, 0, 0

	for i := 0; i < len(ar)/2; i++ {

		for j := len(ar) - 1; j >= len(ar)/2; j-- {
			o := prod(ar[i:j])

			if o > max {
				max = o
				s = i
				e = j
			}
		}
	}

	return s, e
}

func main() {

	rd := bufio.NewReader(os.Stdin)

	str, err := rd.ReadString('\n')

	if err != nil {
		panic("Gone")
	}

	strs := strings.SplitN(str, ",", -1)

	is := make([]int, len(strs))

	for k, v := range strs {

		is[k], _ = strconv.Atoi(strings.TrimSpace(v))
	}

	fmt.Println(bigProd(is))
}
