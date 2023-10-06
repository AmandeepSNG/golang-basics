package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Creating Custom Array/slices functions in go")
	arr := []int{4, 2, 1, 3, 5, 6}
	// Reversing a slice
	fmt.Println("reversed Array is => ", reverseArray(arr))

	// sorting a slice
	fmt.Println("sorted Array in AscedingOrder =>", sortArray(arr, "asc"))
	fmt.Println("sorted Array in DescendingOrder =>", sortArray(arr, "desc"))

	// finding smallest number in the array of numbers
	fmt.Println("Smallest Number in the Array is =>", findSmallest(arr))

	// finding largest number in the array of numbers
	fmt.Println("Largest Number in the Array is =>", findLargest(arr))

	// couting total odd numbers in array
	fmt.Println("There are Total => " + strconv.Itoa(countOddNumbers(arr)) + " Odd numbers in Array")

	// couting total Even numbers in array
	fmt.Println("There are Total => " + strconv.Itoa(countEvenNumbers(arr)) + " Even numbers in Array")
}

/*
Function to reverse an array
*/
func reverseArray(arr []int) []int {
	fmt.Println()
	var reverseArr []int
	for i := len(arr) - 1; i >= 0; i-- {
		reverseArr = append(reverseArr, arr[i])
	}
	return reverseArr
}

/*
Function to sort Array
*/
func sortArray(arr []int, sortOrder string) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if sortOrder == "asc" {
				if arr[i] > arr[j] {
					temp := arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			} else {
				if arr[i] < arr[j] {
					temp := arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			}
		}
	}
	return arr
}

/*
Function to find the smallest number in the given Array
*/
func findSmallest(arr []int) int {
	smallestNumber := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallestNumber {
			smallestNumber = arr[i]
		}
	}
	return smallestNumber
}

/*
Function to find the largest number in the given Array
*/
func findLargest(arr []int) int {
	largestNumber := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > largestNumber {
			largestNumber = arr[i]
		}
	}
	return largestNumber
}

/*
Function to count the number of odd integers in a given array
*/
func countOddNumbers(arr []int) int {
	oddNumberCounter := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			oddNumberCounter++
		}
	}
	return oddNumberCounter
}

/*
Function to countEvenNumbers in an given array
*/
func countEvenNumbers(arr []int) int {
	evenNumberCounter := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			evenNumberCounter++
		}
	}
	return evenNumberCounter
}
