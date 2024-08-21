/*
- Which package this .go file belongs to.
- And this file to also belong to package main
- run command: go run main.go helper.go
- run command: go run . (run all .go files in the project)
*/
// package main // old version

package helper

import (
	"strings"
)

/*
func validateUserInput (not export)

- 	========== Exporting a variable ==========
-> Make it available for all packages in the app
-> Capitalize first letter
-> Go can export: variables, functions, constants, types,...
==> func ValidateUserInput (Capitalize first letter to export function)
*/
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
