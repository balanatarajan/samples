package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"
import "io"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewReader(os.Stdin)
	_, err := scanner.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		panic("Some Error while reading the string")
	}

	// fmt.Println(st)

	l, err := scanner.ReadString('\n')

	if err != nil && err != io.EOF {
		panic("Some Error while reading the string")
	}

	sts := strings.SplitN(l, " ", -1)

	fmt.Println(sts[len(sts)-1])

	ns := make([]int, len(sts))
	for i := range sts {
		ns[i], err = strconv.Atoi(strings.TrimSpace(sts[i]))
	}

	for i := len(ns); i != 0; i-- {

		fmt.Printf("%d", ns[i-1])
	}
	fmt.Printf("\n")
}
