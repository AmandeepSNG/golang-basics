package main

import (
	"fmt"
	"strconv"
)

func main() {
	printLine()
}

func printLine() {
	fmt.Println("------ method => 1 to declare variables in golang.")
	// first method to declare variables
	var integerVariable = 1           // automatically type assign to a => int
	var floatVariable = 1.5           // automatically type assign to b => float 64
	var strinVariable = "hello world" // automatically type asign to c => string
	fmt.Println("value of integerVariable is :", integerVariable)
	fmt.Println("value of floatVariable is :", floatVariable)
	fmt.Println("value of strinVariable is :", strinVariable)
	// print type of variable
	fmt.Printf("type of integerVariable is %T\n", integerVariable)
	fmt.Printf("type of floatVariable is %T\n", floatVariable)
	fmt.Printf("type of strinVariable is %T\n", strinVariable)

	// second method to decalare variables
	fmt.Println("------ method => 2 to declare variables ------")
	var integerVariable_2 int = 27
	var floatVariable_2 float64 = 27.4
	var stringVariable_2 string = "hello from string variable"
	fmt.Println("intergerVariable_2", integerVariable_2)
	fmt.Println("floatVariable_2", floatVariable_2)
	fmt.Println("stringVariable_2", stringVariable_2)

	// third shorthand method to declare varibales in golang
	fmt.Println("----- method => 3 to declare variables in golang.")
	integerVariable_3 := 4
	stringVariable_3 := "hello from string variable 3"
	floatVariable_3 := 33.4
	fmt.Println("integer variable 3 => ", integerVariable_3)
	fmt.Println("string variable 3 =>", stringVariable_3)
	fmt.Println("float variable 3 => ", floatVariable_3)

	// concatenation in golang
	/*
		we cannot concatenate strings and number directly like we used to do in js. In go we have to use Itoa function of strconv package which will convert a number to string
	*/
	fmt.Println("My age is " + strconv.Itoa(25))

	fmt.Println("------ constants in golang ---")
	const normalString = "I am constant string"
	fmt.Println("value of constant is => ", normalString)
}
