package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	str, err := scanner.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		panic("Some Error while reading the string")
	}

	ctr := 0
	b := []byte(str)
	b = bytes.TrimSpace(b)

	for _, v := range b {

		if unicode.IsUpper(rune(v)) {
			ctr++
		}
	}

	fmt.Println(ctr + 1)
}
