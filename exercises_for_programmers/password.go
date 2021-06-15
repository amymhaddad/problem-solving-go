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

func normalize_password(password string) string {

	return strings.ToLower(password)
}

func password_length(password string) string {
	length := len(password) - standard_password_length

	if length < standard_password_length {
		return "short password length"
	} else {

		return "standard password length"
	}
}

func very_weak_password(password string) bool {
	short_length := password_length(password) == "short password length"
	contains_all_numbers, _ := regexp.MatchString("(^[0-9]+$)", password)

	return short_length && contains_all_numbers
}

func weak_password(password string) bool {
	short_length := password_length(password) == "short password length"
	contains_all_letters, _ := regexp.MatchString("([a-z]+)", normalize_password(password))
	
	return short_length && contains_all_letters
}

func strong_password(password string) bool {
	standard_length := password_length(password) == "standard password length"
	contains_letters, _ := regexp.MatchString("[a-z]", normalize_password(password))
	contains_number, _ := regexp.MatchString("[0-9]+", password)
	contains_special_characters, _ := regexp.MatchString("[^a-z0-9]", normalize_password(password))

	return standard_length && contains_letters && contains_number && !contains_special_characters
}

func very_strong_password(password string) bool {
	standard_length := password_length(password) == "standard password length"

	contains_letters, _ := regexp.MatchString("[a-z]", normalize_password(password))
	contains_number, _ := regexp.MatchString("[0-9]+", password)
	contains_special_characters, _ := regexp.MatchString("[^a-z0-9]", normalize_password(password))

	return standard_length && contains_letters && contains_number && contains_special_characters
}

func password_strength(password string) string {
	switch {
		case very_weak_password(password):
			return "very weak"
		case weak_password(password):
			return "weak"
		case strong_password(password):
			return "strong"
		case very_strong_password(password):
			return "very strong"
		default:
			return "Invalid password"
		}
}

func main() {
	password := user_password()
	user_password_strength := password_strength(password)
	fmt.Printf("The password '%s' is a '%s' password.\n", password, user_password_strength)
}
