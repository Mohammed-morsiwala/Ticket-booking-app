package helper

import "strings"

// here the capital letter Vlidation is to export packages in golang

func Validation(fName string, email string, lName string, bookedTicket uint, remainigTicket uint) (bool, bool, bool) {
	isValidname := len(fName) >= 2 && len(lName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := bookedTicket > 0 && bookedTicket <= remainigTicket

	return isValidname, isValidEmail, isValidTicket
}
