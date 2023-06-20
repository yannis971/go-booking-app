package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func conditional_loop() {
	var i uint = 0
	for i < 10 {
		fmt.Printf("i = %d\n", i)
		i++
	}
}

func infinite_loop() {
	var exitFunction string
	for {
		_ = readString("type exit to end loop : ", &exitFunction)
		if strings.Compare(exitFunction, "exit") == 0 {
			break
		}
	}
}

func for_each_loop() {
	var fullNames [3]string
	fullNames[0] = "Lewis Hamilton"
	fullNames[1] = "Lebron James"
	fullNames[2] = "Michael Jordan"
	var firtNames []string
	for _, fullName := range fullNames {
		var names = strings.Fields(fullName)
		firtNames = append(firtNames, names[0])
	}
	fmt.Printf("The firstnames are : %v\n", firtNames)

}

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//var arrayOfNames = [50]string{"Jack", "Nicole", "Jenny"}
	var arrayOfNames [50]string
	var sliceOfNames []string

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

	arrayOfNames[0] = firstName + " " + lastName
	sliceOfNames = append(sliceOfNames, firstName+" The GOAT !")

	fmt.Printf("The whole array : %v\n", arrayOfNames)
	fmt.Printf("The first value : %v\n", arrayOfNames[0])
	fmt.Printf("arrayOfNames is type : %T\n", arrayOfNames)
	fmt.Printf("len(arrayOfNames) is : %v\n", len(arrayOfNames))

	fmt.Printf("The whole slice : %v\n", sliceOfNames)
	fmt.Printf("The first value : %v\n", sliceOfNames[0])
	fmt.Printf("sliceOfNames is type : %T\n", sliceOfNames)
	fmt.Printf("len(sliceOfNames) is : %v\n", len(sliceOfNames))

	var chaine string

	_ = readString("Enter a string of characters : ", &chaine)
	fmt.Println("You typed : ", chaine)
	entier, err := readInteger("Enter an integer : ")
	if err == nil {
		fmt.Println("The integer typed is :", entier)
	} else {
		fmt.Println(err)
	}

	infinite_loop()

	for_each_loop()

	conditional_loop()

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
