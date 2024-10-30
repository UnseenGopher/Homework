package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Println("Enter two numbers,a and b, b must be more than a")
	fmt.Scanf("%f%f", &a, &b)
	if b < a {
		fmt.Println("Erorr, b must be more than a")
	}
	a_int := math.Ceil(a)
	b_int := math.Floor(b)
	for i := a_int; i < b_int; i++ {
		fmt.Printf("%.0f\t", i)
	}
}
