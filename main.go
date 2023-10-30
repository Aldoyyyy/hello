package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// function for checking is parameter more than 5
	result := isMoreThanFive(9)
	fmt.Println(result)

	// Function for checking is parameter isLetter A
	resultS := isLetterA("A")
	fmt.Println(resultS)

	// Function for logic sum for 2 parameters
	resultInt := sum(50, 190)
	fmt.Println(resultInt)

	// function for logic multiply for 2 parameters
	resultInt = multiply(245, 100)
	println(resultInt)

	// fucntion for loop
	loop(5)

	// function for choose club Home
	chooseClubHome("MC")
}

func isMoreThanFive(input int64) bool {
	if input > 5 {
		return true
	} else {
		return false
	}
}

func isLetterA(input string) string {
	if input == "A" {
		return "true"
	} else {
		return "false"
	}
}

func sum(input1 int64, input2 int64) int64 {
	total := input1 + input2
	return total
}

func multiply(input1, input2 int64) int64 {
	total := input1 * input2
	return total
}

func loop(input int64) {
	// find the difference between loop 1-10 and 10-1
	for i := int64(1); i <= input; i++ {
		println(i)
	}
}

func chooseClubHome(club string)  {
	switch club {
	case "MU":
		println("Old Trafford")
	case "MC":
		println("Etihad")
	case "Madrid":
		println("Santiago Bernabeu")
	case "Barca":
		println("Camp Nou")
	default:
		println("club not found")
	}
}