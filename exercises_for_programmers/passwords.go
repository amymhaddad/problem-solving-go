/*
Create a simple program that validates user login credentials. The program must prompt the user for a username and password. The program should compare the password given by the user to a known password. If the password matches, the program should display “Welcome!” If it doesn’t match, the program should display “I don’t know you.”
*/

package main

import "fmt"

func main() {
	fmt.Println(passwords())
}


//Define constants in or outside of the function? 
//Need to add username question once i learn how to ask for multiple inputs
func passwords() string {
	const known_password = "123abc"

	fmt.Println("Enter your password: ")
	var entered_password string
	fmt.Scanln(&entered_password)

	if known_password == entered_password {
		return "Welcome"
	} else {
		return "I don't know you"
	}

}
