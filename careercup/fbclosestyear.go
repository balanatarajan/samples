// Enter a two digit number and find the year closest to that.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	str, err := scanner.ReadString('\n')
	if err != nil {
		panic("Something wrong")
	}

	ins, _ := strconv.Atoi(strings.TrimSpace(str))

	yr := time.Now().Year()
	d2 := yr % 100

	if ins > d2 {
		fmt.Println(1900 + ins)
	} else {
		fmt.Println(2000 + ins)
	}

}
