package main

import (
	"booking-app/helper" // booking-app is the name of module in go.mod
	"fmt"
	"sync"
	"time"
)

/*
Package Level Variables (like Global Variables in another programing languages)
In Package Level Variables can't use (:=) like (conferenceName := "Tuan Tran Conference")
*/
var conferenceName = "Tuan Tran Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

/*
- 			Creating a slice with "make"
- Alternative way to create a slice
- We need to define the initial size of the slice (example: make([]map[string]string, 0) )
*/
// var bookings = make([]map[string]string, 0)

// Struct in Go
type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var bookings = make([]userData, 0)
var isContinue string = "yes"

/*
-                   Goroutines - Concurrency
- WaitGroup waits for a collection of goroutines to finish.
- Add(): Sets the number of goroutines to wait for (increase the counter by the provide number).
- Wait(): Blocks until the WaitGroup counter is zero.
- Done(): Decrements the WaitGroup counter by 1. This is called by Goroutines to indicate that it's finished.
*/
var waitGroup = sync.WaitGroup{}

func main() {
	fmt.Print("\n")

	greetUsers(conferenceName, conferenceTickets, int(remainingTickets))

	// for condition {}
	for remainingTickets > 0 && len(bookings) < 50 && isContinue == "yes" {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			// Add(): Sets the number of goroutines to wait for (increase the counter by the provide number).
			waitGroup.Add(1)
			/*
				- "go ..." start a new Goroutine
				- A Goroutine is a lightweight thread managed by the Go runtime.
			*/
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstName(bookings)
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

		fmt.Printf("Do you want to continue? (yes/no): ")
		fmt.Scan(&isContinue)

		if isContinue == "no" {
			fmt.Println("Bye...")
			break
		}

		fmt.Print("\n")
	}
	// Wait(): Blocks until the WaitGroup counter is zero.
	waitGroup.Wait()
}

func greetUsers(confName string, confTickets int, remainingTickets int) {
	fmt.Println("========================== START ==========================")
	fmt.Printf("Welcome to %v booking application!\n", confName)
	fmt.Print("Get your tickets here to attend!\n")
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstName(bookings []userData) []string {
	firstNames := []string{}
	// for index, booking := range bookings {}
	// (Blank Identifier) Use underscore '_' to ignore variable I don't want to use
	for _, booking := range bookings {
		/*
			// Old version with slice of strings
			names := strings.Fields(booking)
			firstName := names[0]
			firstNames = append(firstNames, firstName)
		*/

		/*
			//Old version with slice of maps
			firstName := booking["firstName"]
		*/

		// New version with struct
		firstName := booking.firstName
		firstNames = append(firstNames, firstName)
	}
	return firstNames
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

	/*
		old version: create a map for user's data
		var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		** Keys of map are sorted by alphabet (ASC) **
	*/

	// new version: create a struct for user's data
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you (%v %v) for booking (%v) tickets. You will receive a confirmation email at (%v)\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	/*
		-                     Save formatted string into a variable (Sprintf)
		-	2 way:
			-	tickets_string := strconv.FormatUint(uint64(userTickets), 10) + " tickets for " + firstName + " " + lastName
			-	tickets := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	*/

	time.Sleep(10 * time.Second)
	tickets := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("=============== Sending email... ===============")
	fmt.Printf("Sending tickets:\n%v \nto email address %v\n", tickets, email)
	fmt.Println("=============== Email sent! ===============")

	// Done(): Decrements the WaitGroup counter by 1. This is called by Goroutines to indicate that it's finished.
	waitGroup.Done()
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
