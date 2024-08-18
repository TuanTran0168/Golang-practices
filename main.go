package main

import "fmt"

func main() {
	fmt.Print("\n")

	// var conferenceName string = "Tuan Tran Conference"
	conferenceName := "Tuan Tran Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 // not negative
	// Data type
	fmt.Printf("ConferenceTickets is (%T) and remainingTickets is (%T) and conferenceName is (%T)\n\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Print("Get your tickets here to attend!\n")

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	var firstName string
	var lastName string
	var userTickets uint
	var email string
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("Variable types: %T, %T, %T, %T\n", firstName, lastName, userTickets, email)
	fmt.Printf("Pointer types: %T, %T, %T, %T\n", &firstName, &lastName, &userTickets, &email)
	fmt.Println("Pointer: ", &firstName, &lastName, &userTickets, &email)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you (%v %v) for booking (%v) tickets. You will receive a confirmation email at (%v)\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Print("\n")
}
