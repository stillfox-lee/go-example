package main

import (
	"fmt"
	"time"
)

func zombie() {
	for {
		fmt.Println("zzzz")
		time.Sleep(time.Second / 2)
	}
}

func callZombie() {
	go zombie()
	fmt.Println("call zombie return")
}

func main() {
	callZombie()
	time.Sleep(time.Second * 5)
	fmt.Println("main return")
}
