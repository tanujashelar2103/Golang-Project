package main

import (
	"fmt"
)

const conferenceTickets = 50
const conferenceName = "Go Conference"

var remainingTickets uint = conferenceTickets
var bookings []UserData

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have a total of %d tickets and %d are still available.\n",
		conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	for {
		firstName, lastName, email, userTickets := getUserInput()

		if userTickets > remainingTickets {
			fmt.Printf("Sorry, we only have %d tickets remaining.\n", remainingTickets)
			continue
		}

		bookTickets(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is fully booked. Come back next year!")
			break
		}
	}
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName, lastName, email string) {
	remainingTickets -= userTickets

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings: %v\n", bookings)
	fmt.Printf("Thank you %s %s for booking %d tickets. "+
		"You will receive a confirmation email at %s\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%d tickets remaining for %s\n", remainingTickets, conferenceName)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
