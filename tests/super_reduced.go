package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func reduce(b []byte) []byte {

	var t []byte

	for i := 0; i < len(b); {

		incr := false
		if i+2 <= len(b) {

			if b[i] != b[i+1] {
				t = append(t, b[i])
			} else {

				incr = true
			}
		} else {
			t = append(t, b[i])
		}

		if incr == true {
			i = i + 2
		} else {
			i = i + 1
		}
	}

	return t
}

func main() {

	scanner := bufio.NewReader(os.Stdin)
	str, err := scanner.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		panic("Some Error while reading the string")
	}

	b := []byte(str)
	b = bytes.TrimSpace(b)

	t := b

	for {
		tmp := reduce(t)

		if len(t) == len(tmp) || len(t) == 0 {
			break
		}

		t = tmp
	}

	if len(t) == 0 {
		fmt.Println("Empty String")
	} else {
		fmt.Println(string(t))
	}
}
