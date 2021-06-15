package main

import (
	"fmt"
	"regexp"
	"strings"
)

const standard_password_length = 8

func user_password() string {
	var password string

	fmt.Printf("Enter your password: ")
	fmt.Scanf("%s", &password)

	return password
}

func very_weak_password(password string, pw_length bool) bool {
	//short_length := password_length(password) == "short password length"
	contains_all_numbers, _ := regexp.MatchString("(^[0-9]+$)", password)

	return pw_length && contains_all_numbers
}

func weak_password(password string, pw_length bool) bool {
	//short_length := password_length(password) == "short password length"
	contains_all_letters, _ := regexp.MatchString("([a-z]+)", password)
	contains_special_characters, _ := regexp.MatchString("[^a-z0-9]", password)
	contains_number, _ := regexp.MatchString("[0-9]+", password)
	
	return pw_length && contains_all_letters && !contains_special_characters && !contains_number
}

func strong_password(password string, pw_length bool) bool {
//standard_length := password_length(password) == "standard password length"
	contains_letters, _ := regexp.MatchString("[a-z]", password)
	contains_number, _ := regexp.MatchString("[0-9]+", password)
	contains_special_characters, _ := regexp.MatchString("[^a-z0-9]", password)

	return pw_length && contains_letters && contains_number && !contains_special_characters
}

func very_strong_password(password string, pw_length bool) bool {
//	standard_length := password_length(password) == "standard password length"

	contains_letters, _ := regexp.MatchString("[a-z]", password)
	contains_number, _ := regexp.MatchString("[0-9]+", password)
	contains_special_characters, _ := regexp.MatchString("[^a-z0-9]", password)

	return pw_length && contains_letters && contains_number && contains_special_characters
}

func password_strength(password string, pw_length bool) string {
	fmt.Println("PW: ", password)
	switch {
		case very_weak_password(password, pw_length):
			return "very weak"
		case weak_password(password, pw_length):
			return "weak"
		case strong_password(password, pw_length):
			return "strong"
		case very_strong_password(password, pw_length):
			return "very strong"
		default:
			return "Invalid password"
			}
}

func main() {
	password := user_password()
	normalized_password := strings.ToLower(password)
	pw_length := len(normalized_password) < standard_password_length

	user_password_strength := password_strength(normalized_password, pw_length)
	fmt.Printf("The password '%s' is a '%s' password.\n", password, user_password_strength)
}
