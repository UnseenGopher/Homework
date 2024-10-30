package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("input a phrase")
	input, _ := reader.ReadString('\n')
	runes := []rune(input)
	runes[len(runes)-1] = ' '
	// runes = append(runes, ' ')

	var start, end int
	var max_len, word_start int

	_ = word_start
	for i, c := range runes {

		if c != ' ' {
			end = i

		} else {
			if end-start > max_len {
				max_len = end - start
				word_start = start
			}
			start = i + 1

		}
	}
	fmt.Println("longest word: \"", string(runes[word_start:word_start+max_len+1]), "\" with len", max_len+1)

}
