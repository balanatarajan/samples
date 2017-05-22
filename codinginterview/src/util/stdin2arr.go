package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Stdin2arr() []int {

	scanner := bufio.NewReader(os.Stdin)

	str, err := scanner.ReadString('\n')

	if err != nil {
		fmt.Println("Something went wrong", err)
		return nil
	}

	strs := strings.SplitN(str, ",", -1)

	ints := make([]int, len(strs))

	for i, v := range strs {

		ints[i], _ = strconv.Atoi(strings.TrimSpace(v))
	}

	return ints[:]
}
