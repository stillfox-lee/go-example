package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 100
	fmt.Println(rand.Intn(max-min+1)+min)
}
