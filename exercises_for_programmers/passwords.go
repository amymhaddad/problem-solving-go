/*
Exercise 15
Create a simple program that validates user login credentials. The program must prompt the user for a username and password. The program should compare the password given by the user to a known password. If the password matches, the program should display “Welcome!” If it doesn’t match, the program should display “I don’t know you.”
*/

package main

import "fmt"

const known_password = "123abc"

func passwords() string {
	var username, user_password string

	fmt.Printf("Enter your username: ")
	fmt.Scanf("%s/n", &username)

	fmt.Printf("Enter your password: ")
	fmt.Scanf("%s/n", &user_password)

	if known_password == user_password {
		return "Welcome"
	} else {
		return "I don't know you"
	}
}

func main() {
	fmt.Println(passwords())
}

