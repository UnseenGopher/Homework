package main

import "fmt"

func main() {
	var sum uint
	for i := 1; i <= 4; i++ {
		var mark uint
		fmt.Printf("Enter mark for exam №%d\n", i)
		_, err := fmt.Scan(&mark)
		if err != nil {
			fmt.Println("invalid mark for exam", i)
			return
		}
		sum += mark
	}
	fmt.Println("sum of your marks: ", sum)
}
