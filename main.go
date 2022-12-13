package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Bardiesel/weird-cli-go.git/helper"
)

const conferenceTickets = 50

var conferenceName = "GO Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicket {
		bookTicket(firstName, lastName, email, userTickets)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The following people have booked tickets: %s\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Sorry we are sold out")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("Please enter a valid name")
		}
		if !isValidEmail {
			fmt.Println("Please enter a valid email")
		}
		if !isValidTicket {
			fmt.Println("Please enter a valid number of tickets")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %s\n", conferenceName)
	fmt.Printf("We have total of %d tickets and remaining Tickets are %d\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Please enter your name")
	fmt.Scan(&firstName)
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email")
	fmt.Scan(&email)
	fmt.Println("Please enter number of tickets you want to buy")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v\n", bookings)

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have total of %s tickets and remaining Tickets are %d\n", conferenceName, remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %s %s\n", userTickets, firstName, lastName)
	fmt.Println("------------------------------------")
	fmt.Printf("Sending Ticket:\n %v \nto email address %s\n", ticket, email)
	fmt.Println("------------------------------------")
	wg.Done()
}
