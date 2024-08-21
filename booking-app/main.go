package main

import (
	"fmt"
	"strings"
)

/*
Package Level Variables (like Global Variables in another programing languages)
In Package Level Variables can't use (:=) like (conferenceName := "Tuan Tran Conference")
*/
var conferenceName = "Tuan Tran Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {
	fmt.Print("\n")

	greetUsers(conferenceName, conferenceTickets, int(remainingTickets))

	// for condition {}
	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			firstNames := printFirstName(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// var noRemainingTickets bool = remainingTickets == 0
			// noRemainingTickets = remainingTickets == 0
			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come back next year!")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short!")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign!")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid!")
			}
		}

		fmt.Print("\n")
	}
}

func greetUsers(confName string, confTickets int, remainingTickets int) {
	fmt.Println("========================== START ==========================")
	fmt.Printf("Welcome to %v booking application!\n", confName)
	fmt.Print("Get your tickets here to attend!\n")
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func printFirstName(bookings []string) []string {
	firstNames := []string{}
	// for index, booking := range bookings {}
	// (Blank Identifier) Use underscore '_' to ignore variable I don't want to use
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstName := names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("=============== User Information ===============")
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you (%v %v) for booking (%v) tickets. You will receive a confirmation email at (%v)\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// switch case in Golang
/*
	city := "Ho Chi Minh"
	switch city {
	case "London", "Paris":
		fmt.Println("London & Paris")
	case "Viet Nam":
		fmt.Println("Ho Chi Minh")
	case "Hong Kong":
		fmt.Println("Hong Kong")
	case "New York":
		fmt.Println("New York")
	default:
		fmt.Println("No valid city selected!")
	}
*/
