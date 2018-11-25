package main

import (
	"fmt"
)

func main() {
	x := 15
	y := &x
	fmt.Println(y)  // print memory address
	fmt.Println(*y) // print value
	*y = 10
	fmt.Println(x)
	*y = *y + *y
	fmt.Println(x)
}
