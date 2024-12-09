package main

import "fmt"

var data int = 3
var ndata = 232

const pi float32 = 3.14

var arr1 = [3]int{1, 2, 3}

func main() {
	name := "Go Developers by njena"
	age := 12

	newsum := 12 + 13

	var slicearr = [5]int{1, 2, 3, 4, 5}
	arrayslice := slicearr[1:4]

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(data)
	fmt.Println(ndata)
	fmt.Println(pi)
	fmt.Println(arr1)
	fmt.Println(arr1[1])
	fmt.Println(len(arr1))
	fmt.Println(myslice2[2])
	fmt.Println(arrayslice)
	fmt.Println(newsum)

	x := 120
	y := 13

	if x < y {
		fmt.Println("x-", +x, "is smaller than y-", +y)
	} else {
		fmt.Println("y-", +y, "is smaller than x-", +x)
	}

	cat()

	day := 2
	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	}

	//Multi swich case
	day1 := 5
	switch day1 {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	//For loop

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//Range function
	fruits := [3]string{"apple", "banana", "orange"}

	for idx, val := range fruits {
		fmt.Println(idx, val)
	}

	//function with parameter
	database("njena")

	fmt.Println(sum(12, 13))

}

func cat() {
	name := "welcome to the real world"
	fmt.Println(name)
}

func database(fname1 string) {
	fmt.Println("Hey", fname1, "welcome to the real world and escape the matrix")
}

func sum(x int, y int) int {
	return x + y
}
