package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	rd := bufio.NewReader(os.Stdin)

	str, err := rd.ReadString('\n')

	if err != nil {
		panic("Couldn't read")
	}

	arr := strings.SplitN(str, ",", -1)
	iarr := make([]int, len(arr))

	for i := range arr {

		iarr[i], _ = strconv.Atoi(strings.TrimSpace(arr[i]))

	}

	sum4(iarr, 8)

}

func sum4(iaar []int, wntd int) {

	for i, j := 0, len(iaar)-1; i < len(iaar); {

		sum := iaar[i] + iaar[j]
		if sum < wntd {
			i = i + 1
			continue
		}
		if sum > wntd {
			j = j - 1
			if j < 0 {
				break
			}
			continue
		}
		if sum == wntd {
			fmt.Println(iaar[i], iaar[j])
			i = i + 1
			continue
		}

	}

}
