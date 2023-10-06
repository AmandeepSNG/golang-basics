package main

import (
	"fmt"
)

type User struct {
	firstName string
	lastName  string
	age       int
	email     string
}

func main() {
	fmt.Println("Working with Maps in golang")
	var userArray []User
	userArray = append(userArray, User{firstName: "Testing", lastName: "singh", age: 24, email: "testing@yopmail.com"})
	// var userObject = User{"Testing", "singh", 24, "test@yopmail.com"}
	fmt.Println("object is => ", userArray)
}
