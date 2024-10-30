package main

import "fmt"

func main() {
	var vals [5]int
	var have_asc bool
	var have_desc bool
	var have_const bool
	fmt.Println("Enter five numbers ")
	for i := range vals {
		fmt.Scan(&vals[i])
	}

	for i := 1; i < 4; i++ {
		if vals[i] > vals[i-1] {
			have_asc = true
		} else if vals[i] < vals[i-1] {
			have_desc = true
		} else {
			have_const = true
		}

	}
	switch {
	case have_asc && !have_desc && !have_const:
		fmt.Println("ASCENDING")
	case have_asc && have_const && !have_desc:
		fmt.Println("WEAKLY ASCENDING")
	case have_asc && have_desc:
		fmt.Println("RANDOM")
	case have_desc && !have_asc && !have_const:
		fmt.Println("DESCENDING")
	case have_desc && !have_asc && have_const:
		fmt.Println("WEAKLY DESCENDING")
	default:
		fmt.Println("CONSTANT")
	}

}
