package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewReader(os.Stdin)

	str, err := scanner.ReadString('\n')

	if err != nil {
		panic("Something wrong")
	}

	str = strings.TrimSpace(str)
	mctr, _ := strconv.Atoi(str)
	fmt.Println(mctr)

	str, err = scanner.ReadString('\n')

	if err != nil {
		panic("Something wrong")
	}

	str = strings.TrimSpace(str)
	b := []byte(str)

	for i, j := 0, len(b)-1; i < len(b); i, j = i+1, j-1 {

		if b[i] != b[j] {
			if mctr == 0 {
				mctr = -1
				break
			} else {
				b[j] = b[i]
				mctr--
			}
		}
	}

	if mctr == -1 {
		fmt.Println(mctr)
	} else {
		fmt.Println(string(b))
	}

}
