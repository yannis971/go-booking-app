package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Println("Number of conference tickets is", conferenceTickets)
	fmt.Println("Number of tickets available is", remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println(conferenceTickets)
}
