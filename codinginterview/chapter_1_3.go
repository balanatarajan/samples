package main

import "fmt"

func main() {

	s := "aaabbb"

	b := []byte(s)

	for k, v := range b {

		for i := 0; i < k; i++ {
			if b[i] == v {

				b[i] = ' '
			}
		}
	}

	fmt.Println(string(b))
	for k, v := range b {
		if v == ' ' && k < (len(b)-1) {
			b[k] = b[k+1]
		}
	}

	i := 0
	i--
	fmt.Println(i)

}
