package main

import "fmt"

var chora = 50

func main() {
	x := 42
	y := "James Bond"
	z := 100
	z , chora := 200, 300

	fmt.Println(x, y, z)
	fmt.Println(chora)
}
