package main

import "fmt"

func main(){
	fmt.Println("Starting the execution of the addressbook")

	testuser:= &Person{
		Name: "TestUser",
		Age: 23,
		EmailID: "something@agasf.com",


	}
}


