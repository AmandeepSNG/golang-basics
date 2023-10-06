package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Array section")
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println("value of array is =>", arr)

	// shorthand to declare and assign value of array
	shorthandArr := [2]int{3, 4}
	fmt.Println("shorthandArray =>", shorthandArr)

	// Array of strings
	arrOfStrings := [4]string{"firstElement", "secondElement", "ThirdElement"}
	fmt.Println("Array of strings ", arrOfStrings)

	// looping on array
	for i := 0; i < len(arrOfStrings); i++ {
		fmt.Println(arrOfStrings[i], "---- element in loop -----")
	}

	fmt.Println("Slices in golang")

}

func reverseArray(arr []int) []int {
	var reverseArr []int
	for i := arr[len(arr)-1]; i >= 0; i-- {
		reverseArr = append(reverseArr, arr[i])
	}
	return reverseArr
}
