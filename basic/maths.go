package main

import (
	"fmt"
)

func add(x float64, y float64) float64{
	return x + y
}
func main() {
	var num1 float64 = 2.1
	var num2 float64 = 3.2
	var sum float64 = add(num1, num2)
	fmt.Println(sum)
}

// OR

// func add(x , y float64) float64{
// 	return x + y
// }
// func main() {
// 	var num1, num2 float64 = 2.1, 3.2
// 	var sum float64 = add(num1, num2)
// 	fmt.Println(sum)

// 	// To convert type

// 	var a int = 10;
// 	var b float64 = float64(a)
// 	fmt.Println(b)
// }
