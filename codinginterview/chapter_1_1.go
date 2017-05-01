package main

import (
	"bufio"
	"fmt"
)

func simpleCheck(str string) bool {

	m := make(map[rune]bool, len(str))

	for _, v := range str {

		_, ok := m[v]

		if ok == true {
			return false
		}

		m[v] = true
	}

	return true
}

func recursiveCheck(str string) bool {

	var r rune

	for _, v := range str {

	}
}

func main() {

	scanner := bufio.NewReader(os.Stdin)

	str, err := scanner.ReadString('\n')

	if err != nil {
		panic("something wrong")
	}

	if simpleCheck(str) == false {

		fmt.Println("Non unique strings")
	}

	if bitCheck(str) == false {

		fmt.Println("Non unique strings")
	}

}
