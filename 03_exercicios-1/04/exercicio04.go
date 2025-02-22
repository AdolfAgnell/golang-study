package main

import (
	"fmt"
)
type sauron int
var x sauron

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Println(x)
}
