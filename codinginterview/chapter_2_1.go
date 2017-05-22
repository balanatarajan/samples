package main

import (
	"container/list"
	"fmt"
)

func findElement(e *list.Element, l *list.List) *list.Element {

	for f := l.Front(); f != nil; f = f.Next() {

		if e != f && f.Value.(int) == e.Value.(int) {
			return f
		}
	}

	return nil
}

func main() {

	l := list.New()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(10)
	l.PushBack(2)
	l.PushBack(12)
	l.PushBack(2)

	for e := l.Front(); e != nil; e = e.Next() {
		if e != l.Front() {

			r := findElement(e, l)

			if r != nil {
				l.Remove(r)
			}
		}
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
