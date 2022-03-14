package main

import "math/rand"

func Greeting() string {
	return "Hello World!"
}

func Times(s rand.Source) int {
	return rand.New(s).Intn(10)
}
