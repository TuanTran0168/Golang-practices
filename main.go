package main

import "fmt"

func main() {
	fmt.Print("\n")
	var conferenceName = "Tuan Tran conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application!", conferenceName)
	fmt.Print("Get your tickets here to attend!\n")

	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Print("Get your tickets here to attend!\n")

	fmt.Print("\n")
}
