package main

import (
	"math/rand"
	"fmt"
)

/* 
Random numbers are generated by a Source. Top-level functions, such as Float64 and Int, use a default shared Source that produces a deterministic sequence of values each time a program is run.
*/

func main() {
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
}
