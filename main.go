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

	// Function println Array with for range
	logicArray()

	// Function Print Slice
	printSlice()

	// function print map
	printMaps()

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

func chooseClubHome(club string) {
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

func setListClub(clubs []string) {
	for index, value := range clubs {
		println(index, value)
	}
}

// Kalau depan kapital itu global func, kalau kecil private func
func PrintListInt(listInt []int64) {
	for index := range listInt {
		value := listInt[index]
		println(value, index)
	}
}

func logicArray() {
	clubs := []string{"MU", "MC"}
	setListClub(clubs)

	var clubList []string
	clubList = []string{"Persib", "Persija"}
	setListClub(clubList)

	clubs = append(clubs, "Persebaya", "Arema")
	setListClub(clubs)

	println("test", clubs[2])

	cars := []string{"BMW", "Ford", "Honda"}
	setListClub(cars)

	shoesSize := []int64{40, 44}
	PrintListInt(shoesSize)
}

func printSlice() {
	array := [6]int{1, 2, 3, 4, 5, 6}

	var s []int = array[2:6]
	fmt.Println(s)
}

func printMaps() {
	// Key itu yang kiri (contoh : brand, model, year)
	// Value itu yang kanan ( contoh : ford, mustang, 1964)
	var a = map[string]string{
		"brand": "Ford",
		"model": "Mustang",
		"year":  "1964",
	}
	b := map[string]int{
		"Oslo":      1,
		"Bergen":    2,
		"Trondheim": 3,
		"Stavanger": 4,
	}
	c := map[int64]string{
		3: "MU",
		2: "MC",
	}
	// %v bisa untuk semua tipe data
	// %s untuk string
	// %d untuk int
	// %f untuk float atau desimal
	// \t adalah Tab
	// \n adalah new line

	fmt.Println(a["brand"])
	fmt.Println(b["Trondheim"])
	fmt.Println(c[2])

	c[5] = "Barca"
	a["city"] = "New York"

	for key, value := range c {
		println(key, value)

	}
	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	// Jika mau menghapus dalah satu key maka menggunakan delete(nama map, key sesuai data title)
	delete(c, 3)
	fmt.Printf("c\t%v\n", c)
}
