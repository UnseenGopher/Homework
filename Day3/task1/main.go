package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func codeToString(inp string) (string, error) {
	len := len(inp)
	if len <= 2 || len%2 != 0 {
		return "", errors.New("string len must be more than 2 and even")
	}

	out := make([]byte, 0)

	for i := 0; i < len; i += 2 {
		ind, err := strconv.Atoi(string(inp[i : i+2]))
		if err != nil {
			return "", err
		}
		//(ab)using ASCII encoding, classic
		if ind >= 0 && ind <= 25 {
			out = append(out, byte(ind+int('a')))
		} else if ind == 26 {
			out = append(out, byte(' '))
		} else {
			return "", errors.New("invalid number: " + strconv.Itoa(ind))
		}
	}
	return string(out), nil
}

func stringtoCode(inp string) (string, error) {
	var out string
	chars := []byte(inp)
	for _, c := range chars {
		var ind int
		if c >= 'a' && c <= 'z' {
			ind = int(c - byte('a'))
		} else if c == ' ' {
			ind = 26
		} else {
			return "", errors.New("invalid, symbol in input: " + string(c))
		}
		var next string
		if ind < 10 {
			next += "0" + strconv.Itoa(ind)
		} else {
			next += strconv.Itoa(ind)
		}
		out += next

	}
	return out, nil

}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter input to encode")
	if reader.Scan() {
		input := reader.Text()
		encoded, err1 := codeToString(input)
		if err1 != nil {
			fmt.Println(err1)
			return
		}

		fmt.Println("Encoded: ", encoded)

		decoded, err2 := stringtoCode(encoded)

		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println("Decoded back: ", decoded)

		fmt.Println("Enter input to decode")

		if reader.Scan() {

			input2 := reader.Text()
			decoded, err3 := stringtoCode(input2)
			if err3 != nil {
				fmt.Println(err3)
				return
			}
			encoded, err4 := codeToString(decoded)
			if err4 != nil {
				fmt.Println(err3)
				return
			}

			fmt.Println("Decoded: ", decoded)
			fmt.Println("Encoded back: ", encoded)

		}
	}

}
