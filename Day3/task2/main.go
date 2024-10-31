package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var test_sets = []string{
	"0123456789",
	"abcdefghiklmnopqrstvxyz",
	"ABCDEFGHIKLMNOPQRSTVXYZ",
	"_!@#$%^&",
}
var err_types = []string{
	"digits",
	"lowercase",
	"uppercase",
	"special",
}

func checkPassword(input string, sets []map[rune]bool) (bool, error) {
	var msg string
	var flags = []bool{false, false, false, false}

	l := len([]rune(input))
	if l < 8 || l > 15 {
		msg += "password bust be from 8 to 15 symbols\n"
	}

	for _, c := range input {
		for i, tset := range sets {
			_, ok := tset[c]
			if ok {
				flags[i] = true
				break
			}
		}
	}

	for i, flag := range flags {
		if !flag {
			msg += "no " + err_types[i] + " symbols\n"
		}
	}
	if len(msg) == 0 {
		return true, nil
	} else {
		return false, errors.New(msg)
	}
}

func main() {

	var tests []map[rune]bool
	for _, test_set := range test_sets {
		chars := []rune(test_set)
		next_set := make(map[rune]bool, 0)
		for _, c := range chars {
			next_set[c] = true
		}
		tests = append(tests, next_set)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter eligible password")
	for i := 1; i <= 5; i++ {
		password, _ := reader.ReadString('\n')
		_, err := checkPassword(password, tests)
		if err != nil {
			fmt.Println("Error for attempt №" + strconv.Itoa(i) + " :")
			fmt.Println(err.Error())
			fmt.Println("try again")
		} else {
			fmt.Println("Password accepted")
			break
		}
	}
	fmt.Println("out of tries")

}
