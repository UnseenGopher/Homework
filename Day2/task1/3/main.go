package main

import "fmt"

func main() {
	var inp uint
	fmt.Println("enter integer, 1 to 9")
	_, err := fmt.Scanf("%d", &inp)

	if err != nil {
		fmt.Println("Error, not a valid integer")
		return
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", inp, i, int(inp)*i)
	}

}
