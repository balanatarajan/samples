package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func longest(bs []byte) int {

	l := 1
	b := bs[0]
	le := len(bs)
	for k, v := range bs {
		tl := 1

		for i := 1; k+i < le; i++ {

			if v != bs[k+i] {
				break
			} else {
				tl = tl + 1
			}
		}

		if tl > l {
			l = tl
		}
	}

	return l
}

func main() {

	scanner := bufio.NewReader(os.Stdin)
	str, err := scanner.ReadString('\n')

	if err != nil {

		panic("Something went wrong")
	}

	str = strings.TrimSpace(str)
	bs := make([]byte, len(str))

	for k, v := range str {
		bs[k] = byte(v)
	}

	fmt.Println(longest(bs))

}
