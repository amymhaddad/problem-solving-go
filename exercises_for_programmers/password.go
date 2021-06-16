package main

import (
	"fmt"
	"regexp"
)

const standard_password_length = 8

func userPassword() string {
	var password string

	fmt.Printf("Enter your password: ")
	fmt.Scanf("%s", &password)

	return password
}

func veryWeakPassword(password string) bool {
	containsAllNumbers, _ := regexp.MatchString("(^[0-9]+$)", password)

	return containsAllNumbers
}

func weakPassword(password string) bool {
	containsAllLetters, _ := regexp.MatchString("([a-zA-Z]+)", password)
	containsSpecialChars, _ := regexp.MatchString("[^a-zA-Z0-9]", password)
	containsNumber, _ := regexp.MatchString("[0-9]+", password)
	
	return containsAllLetters && !containsSpecialChars && !containsNumber
}

func strongPassword(password string) bool {
	containsLetters, _ := regexp.MatchString("[a-zA-Z]", password)
	containsNumber, _ := regexp.MatchString("[0-9]+", password)
	containsSpecialChars, _ := regexp.MatchString("[^a-zA-Z0-9]", password)

	return containsLetters && containsNumber && !containsSpecialChars
}

func veryStrongPassword(password string) bool {
	containsLetters, _ := regexp.MatchString("[a-zA-Z]", password)
	containsNumber, _ := regexp.MatchString("[0-9]+", password)
	containsSpecialChars, _ := regexp.MatchString("[^a-zA-Z0-9]", password)

	return containsLetters && containsNumber && containsSpecialChars
}

func passwordStrength(password string) string {
	if len(password) >= standard_password_length {
		if strongPassword(password){

			return "strong"

		} else if veryStrongPassword(password) {

			return "very strong"
		} 
	}
	
	if len(password) < standard_password_length {
		if veryWeakPassword(password) {

			return "very weak"

		} else if weakPassword(password) {

			return "weak"
		}
	}
	
	return "Invalid Password"
	
}

func main() {
	password := userPassword()
	userPasswordStrength := passwordStrength(password)

	fmt.Printf("The password '%s' is a '%s' password.\n", password, userPasswordStrength)
}
