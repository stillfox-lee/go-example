package main

import "fmt"

type Role int

const (
	Admin Role = iota
	User
	Guest
)

func main() {
	fmt.Println(Admin, User, Guest)
}
