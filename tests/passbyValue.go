package main

import "fmt"

type User struct {
	name string
}

func modify(u *User) {

	u.name = "Bala"
}

func main() {

	u := &User{name: "amba"}

	fmt.Println(u)

	modify(u)

	fmt.Println(u)
}
