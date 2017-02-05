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
	fmt.Println("Length is ", len(b), len(str))

	r := false
	for i, j := 0, len(b)-1; i != len(b); i, j = i+1, j-1 {
		if b[i] != b[j] {
			fmt.Println("Not a palindrome", i, j)
			r = true
			break
		}
	}

	if r == false {
		fmt.Println("Its a Palindrome")
	}

}
