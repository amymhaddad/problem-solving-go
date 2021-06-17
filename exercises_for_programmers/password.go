package main

import (
	"fmt"
	"regexp"
)

const standardPasswordLength = 8

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
	weakPasswordLength := len(password) < standardPasswordLength
	strongPasswordLength := len(password) >= standardPasswordLength

	switch {
		case strongPasswordLength && strongPassword(password):
			return "strong"
		case strongPasswordLength && veryStrongPassword(password):
			return "very strong"
		case weakPasswordLength && weakPassword(password):
			return "weak"
		case weakPasswordLength && veryWeakPassword(password):
			return "very weak"
		default:
			return "Invalid password"
	}
}

func main() {
	password := userPassword()
	userPasswordStrength := passwordStrength(password)

	fmt.Printf("The password '%s' is a '%s' password.\n", password, userPasswordStrength)
}
