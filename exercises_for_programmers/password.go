package main

import (
	"fmt"
	"regexp"
)

const standardPasswordLength = 8

var containsNumbers = regexp.MustCompile("([0-9]+)")
var containsLetters = regexp.MustCompile("([a-zA-Z]+)")
var containsSpecialChars = regexp.MustCompile("[^a-zA-Z0-9]")

func userPassword() string {
	var password string

	fmt.Printf("Enter your password: ")
	fmt.Scanf("%s", &password)

	return password
}

func veryWeakPassword(password string) bool {
	return containsNumbers.MatchString(password)
}

func weakPassword(password string) bool {
	return containsLetters.MatchString(password)
}

func strongPassword(password string) bool {
	containsLettersNums := containsLetters.MatchString(password) && containsNumbers.MatchString(password)
	specialChars := containsSpecialChars.MatchString(password)

	return containsLettersNums && !specialChars
}

func veryStrongPassword(password string) bool {
	containsLettersNumsSpecChars := containsLetters.MatchString(password) && containsNumbers.MatchString(password) && containsSpecialChars.MatchString(password)

	return containsLettersNumsSpecChars
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
