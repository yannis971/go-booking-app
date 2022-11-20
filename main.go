package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Println("Number of conference tickets is", conferenceTickets)
	fmt.Println("Number of tickets available is", remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println(conferenceTickets)

	fmt.Printf("Enter your firstname : ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your lastname : ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email : ")
	fmt.Scan(&email)
	fmt.Printf("How many tickets you want to buy ? ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	fmt.Println("Hello", firstName, lastName, "!")
	fmt.Println("Your email is", email)
	fmt.Println("You want", userTickets, "tickets")
	fmt.Println("Memory address of variable userTickets is", &userTickets)
	fmt.Printf("userTickets is type %T, conferenceName is %T\n", userTickets, conferenceName)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
