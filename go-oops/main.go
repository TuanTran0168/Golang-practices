package main

import (
	"fmt"
	"go-oops/structs"
)

func main() {
	person := structs.Person{}

	err := person.SetFirstName("Tuấn")

	if err != nil {
		fmt.Println(err.Error())
	}

	person.SetLastName("Trần")
	person.SetAge(22)

	person_1 := structs.NewPerson("Tuấn", "Trần", 22)

	fmt.Println(person, person_1)
}
