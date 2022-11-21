package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	var chaine string

	_ = readString("Enter a string of characters : ", &chaine)
	fmt.Println("You typed : ", chaine)
	entier, err := readInteger("Enter an integer : ")
	if err == nil {
		fmt.Println("The integer typed is :", entier)
	} else {
		fmt.Println(err)
	}

}

func readString(prompt string, result *string) (err error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)

	*result, err = reader.ReadString('\n')

	if err == nil {
		*result = strings.Replace(*result, "\n", "", 1)
	}

	return err
}

func readInteger(prompt string) (result int64, err error) {

	var text string
	err = readString(prompt, &text)

	if err == nil {
		result, err = strconv.ParseInt(text, 10, 64)
	}

	return result, err
}
