package structs

import (
	"errors"
	"fmt"
)

type Person struct {
	firstName string // private fields
	lastName  string
	age       int
}

// constructor
func NewPerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

// receiver function
func (person Person) ToString() string {
	return fmt.Sprintf("%v %v %v", person.firstName, person.lastName, person.age)
}

// receiver function
func (person Person) Walk() string {
	return person.firstName + " " + person.lastName + " likes walking"
}

// Getter
func (person Person) GetFirstName() string {
	return person.firstName
}

// Getter
func (person Person) GetLastName() string {
	return person.lastName
}

// Getter
func (person Person) GetAge() int {
	return person.age
}

// Setter
// Pointer based on receiver functions
func (person *Person) SetFirstName(firstName string) error {
	if len(firstName) < 3 {
		return errors.New("invalid first name")
	}
	person.firstName = firstName

	return nil
}

// Setter
func (person *Person) SetLastName(lastName string) {
	person.lastName = lastName
}

// Setter
func (person *Person) SetAge(age int) {
	person.age = age
}
