package main

import (
	"fmt"
)

type sm struct {
	value int
	ptr   *sm
}

func al(s *sm) *sm {

	if s == nil {
		t := sm{1, nil}
		return &t
	}

	return s
}

func main() {
	s := &sm{1, nil}

	s.ptr = al(s.ptr)
	fmt.Println(s.ptr)
}
