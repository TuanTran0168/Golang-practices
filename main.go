package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("\n")
	// var conferenceName string = "Tuan Tran Conference"
	conferenceName := "Tuan Tran Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 // not negative
	// var bookings = [50]string{} // Array
	// var bookings []string // Slice
	// var bookings = []string{} // Slice
	bookings := []string{} // Slice

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

	// for condition {}
	for remainingTickets > 0 && len(bookings) < 50 {
		fmt.Println("=========================================")
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you (%v %v) for booking (%v) tickets. You will receive a confirmation email at (%v)\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			// for index, booking := range bookings {}
			for _, booking := range bookings { // (Blank Identifier) Use underscore '_' to ignore variable I don't want to use
				names := strings.Fields(booking)
				firstName := names[0]

				firstNames = append(firstNames, firstName)
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come back next year!")
				break
			}

		} else {
			fmt.Printf("we have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		fmt.Print("\n")
	}

	// fmt.Printf("Variable types: %T, %T, %T, %T\n", firstName, lastName, userTickets, email)
	// fmt.Printf("Pointer types: %T, %T, %T, %T\n", &firstName, &lastName, &userTickets, &email)
	// fmt.Println("Pointer: ", &firstName, &lastName, &userTickets, &email)

	// fmt.Printf("The whole Slice: %v\n", bookings)
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("First element: %v\n", bookings[0])
	// fmt.Printf("Slice lenght: %v\n", len(bookings))
	// fmt.Printf("Slice: %v\n", bookings[0:])
}
