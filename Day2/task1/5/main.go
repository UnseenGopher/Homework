package main

import "fmt"

func main() {

	var input, sum float64
	fmt.Println("enter numbers, when -1 is endtered, programs stops")
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error, not a number", input)
			return
		}
		if input == -1 {
			break
		}
		sum += input
	}
	fmt.Println("sum of inputs", sum)

}
