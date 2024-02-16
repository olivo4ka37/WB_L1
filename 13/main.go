package main

import "fmt"

func main() {
	x1, x2 := 1, 2
	fmt.Println("x1 is:", x1, "x2 is:", x2)
	x1, x2 = x2, x1
	fmt.Println("x1 is:", x1, "x2 is:", x2)

}
