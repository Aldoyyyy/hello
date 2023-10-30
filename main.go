package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	result := isMoreThanFive(3)
	fmt.Println(result)
}

func isMoreThanFive(input int64) bool {
	if input > 5 {
		return true
	} else {
		return false
	}
}
