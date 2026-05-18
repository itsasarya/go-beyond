package main

import (
	"fmt"
)

func main() {
	name := "testing" // change input as per function called below
	phone := "0987654321"
	result := AddContact(name, phone) //change function to call
	find:= FindContact(name)
	up:= UpdateContact(name, "9879")
	list := ListContact()
	fmt.Println(result, find, up, list)
}
