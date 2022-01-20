package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

//creating wait group bocoz at a time main func got executed it will end the program not wait for other threads to complete
var wg = sync.WaitGroup{}

const confrenceTicket = 50

// sugar syntax only for int
var confrenceName = "Go conference"
var remainigTicket uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()
	for remainigTicket > 0 {

		fName, lName, email, bookedTicket := getUsersinput()

		// fmt.Println(isValidname)
		isValidname, isValidEmail, isValidTicket := helper.Validation(lName, email, fName, bookedTicket, remainigTicket)

		if isValidname && isValidEmail && isValidTicket {
			bookedTickets(bookedTicket, fName, lName, email)

			//adding scnchronization
			wg.Add(1)
			go sendTickets(bookedTicket, fName, lName, email)
			//print tickets holder first name
			firstNames := printFnames()

			fmt.Printf("The list of the ticket holders so far is: %v \n", firstNames)

			if remainigTicket == 0 {
				fmt.Println("Our connference is fully booked come next year")
				break
			}
		} else {
			if !isValidname {
				fmt.Println("first name and last name should be atleast of 2 character long !!")
			}
			if !isValidEmail {
				fmt.Println("email should be consist of '@' !!")
			}
			if !isValidTicket {
				fmt.Printf("we just have %v remaining out of %v\n", remainigTicket, confrenceTicket)
			}
		}

	}

	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v boooking application\n", confrenceName)
	fmt.Printf("We are having %v remaining out of total %v tickets\n", confrenceTicket, remainigTicket)
	fmt.Println("Get your tickets from here")
}

func printFnames() []string {
	var firstNames = []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUsersinput() (string, string, string, uint) {
	// var booking [confrenceTicket]string
	// we are using slice instead of array due to its dynamic nature
	var fName string
	var lName string
	var email string
	var bookedTicket uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&fName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter your how much ticket you willing to buy: ")
	fmt.Scan(&bookedTicket)

	return fName, lName, email, bookedTicket
}

func bookedTickets(bookedTicket uint, fName string, lName string, email string) {
	remainigTicket = remainigTicket - bookedTicket
	//using struct
	var userData = userData{
		firstName:       fName,
		lastName:        lName,
		email:           email,
		numberOfTickets: bookedTicket,
	}

	//using maps
	// userData["firstName"] = fName
	// userData["lastName"] = lName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(bookedTicket), 10)
	bookings = append(bookings, userData)
	fmt.Printf("here is the userData : %v\n", bookings)
	fmt.Printf("hey %v %v you have book %v tickets\n", fName, lName, bookedTicket)
	fmt.Printf("remaining tickets are %v\n", remainigTicket)

}

func sendTickets(bookedTicket uint, fName string, lName string, email string) {
	time.Sleep(10 * time.Second)

	fmt.Printf("#########################\n To : %v", email)
	var ticket = fmt.Sprintf("Hey %v %v here is your %v tickets for the conference Enjoy the Ride\n", fName, lName, bookedTicket)
	fmt.Printf("Go Reference Booking Tickets :\n %v", ticket)
	fmt.Printf("#########################\n")
	wg.Done()
}
