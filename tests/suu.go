package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewReader(os.Stdin)
	str, err := scanner.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		panic("Some Error while reading the string")
	}

	b := []byte(str)
	b = bytes.TrimSpace(b)

	var t []byte

	for i := 0; i != len(b); {

		if b[i] != b[i+1] {
			t = append(t, b[i])
		} else {
			k := b[i+2:]
			b = k
			i = i + 2
		}

		i = i + 1
	}

	fmt.Println(t)

}
