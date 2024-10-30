package main

import (
	"fmt"
	"strconv"
)

func main() {
	//since console input is text anyway
	var input string
	var sum uint
	fmt.Println("enter natural number")
	fmt.Scan(&input)
	for i, c := range input {
		var digit int
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Printf("error, not a digit at position %d in input", i)
		}
		sum += uint(digit)
	}
	fmt.Println("Sum of digits", sum)

}
