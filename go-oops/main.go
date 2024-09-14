package main

import (
	"fmt"
	"go-oops/structs"
)

func main() {
	person := structs.Person{}

	person.SetFirstName("Tuấn")
	person.SetLastName("Trần")
	person.SetAge(22)

	fmt.Println(person)
}
