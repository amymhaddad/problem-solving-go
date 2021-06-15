package main

import (
	"fmt"
	"regexp"
)

const standard_password_length = 8

func user_password() string {
	var password string

	fmt.Printf("Enter your password: ")
	fmt.Scanf("%s", &password)

	return password
}

func very_weak_password(password string) bool {
	contains_all_numbers, _ := regexp.MatchString("(^[0-9]+$)", password)

	return contains_all_numbers
}

func weak_password(password string) bool {
	contains_all_letters, _ := regexp.MatchString("([a-zA-Z]+)", password)
	contains_special_characters, _ := regexp.MatchString("[^a-zA-Z0-9]", password)
	contains_number, _ := regexp.MatchString("[0-9]+", password)
	
	return contains_all_letters && !contains_special_characters && !contains_number
}

func strong_password(password string) bool {
	contains_letters, _ := regexp.MatchString("[a-zA-Z]", password)
	contains_number, _ := regexp.MatchString("[0-9]+", password)
	contains_special_characters, _ := regexp.MatchString("[^a-zA-Z0-9]", password)

	return contains_letters && contains_number && !contains_special_characters
}

func very_strong_password(password string) bool {
	contains_letters, _ := regexp.MatchString("[a-zA-Z]", password)
	contains_number, _ := regexp.MatchString("[0-9]+", password)
	contains_special_characters, _ := regexp.MatchString("[^a-zA-Z0-9]", password)

	return contains_letters && contains_number && contains_special_characters
}

func password_strength(password string) string {
//can delete pw_legnth 

	if len(password) >= standard_password_length {
		if strong_password(password){
			return "strong"
		} else if very_strong_password(password) {
			return "very strong"
		} 
		return "Invalid password"
	}
	
	if len(password) < standard_password_length {
		if very_weak_password(password) {
			return "very weak"
		} else if weak_password(password) {
			return "weak"
		}
		return "Invalid password"
	}
	
	return "Invalid Password"
	
}



func main() {
	password := user_password()

	user_password_strength := password_strength(password)
	fmt.Printf("The password '%s' is a '%s' password.\n", password, user_password_strength)
}
