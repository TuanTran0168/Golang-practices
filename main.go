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

	fmt.Printf("Welcome to %v booking application!", conferenceName)
	fmt.Print("Get your tickets here to attend!\n")

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	fmt.Print("\n")
}
