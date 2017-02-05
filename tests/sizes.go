package main

import (
	"fmt"
	"go/types"
	"io"
)

func main() {

	var s types.StdSizes
	fmt.Println("Size of ", s.Sizeof(types.Int32))
}
