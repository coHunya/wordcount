package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := readInput()
	result := returnCount(str)
	fmt.Println(result)
}

func readInput() string {
	var str string
	if len(os.Args) > 1 {
		str = os.Args[1]
	}
	return str
}

func returnCount(str string) int {
	counter := 0
	if len(str) == 0 {
		return counter
	}
	words := strings.Split(str, " ")
	for _, w := range words {
		if len(w) != 0 {
			counter += 1
		}
	}
	return counter
}
