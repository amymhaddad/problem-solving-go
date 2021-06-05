package main

import (
	"fmt"
	"regexp"
	"strings"
)

const standard_password_length int = 8

func user_password() string {
	var password string
	fmt.Printf("Enter your password: ")
	fmt.Scanf("%s", &password)
	return password
}

func normalize_password(password string) string {
	return strings.ToLower(password)
}



func weak_password_length(password string) bool {
	return len(password) < standard_password_length
}

func strong_password_length(password string) bool {
	return len(password) >= standard_password_length
}

func password_length(password string) int {
	return len(password)
}

func contains_letters(password string) bool {
	found, _ := regexp.MatchString("([a-z]+)", normalize_password(password))
	return found
}

func contains_numbers(password string) bool {
	found, _ := regexp.MatchString("[0-9]+", password)
	return found
}

func contains_special_chars(password string) bool {
	found, _ := regexp.MatchString("[^a-z0-9]", normalize_password(password))
	return found
}

func contains_all_letters(password string) bool {
	found, _ := regexp.MatchString("(^[a-z]+$)", normalize_password(password))
	return found
}

func contains_all_numbers(password string) bool {
	found, _ := regexp.MatchString("(^[0-9]+$)", password)
	return found
}

func contains_number_and_letters(password string) bool {
	contains_letters, _ := regexp.MatchString("[a-z]", normalize_password(password))
	contains_number, _ := regexp.MatchString("[0-9]+", password)
	contains_special_characters := contains_special_chars(password)

	return contains_letters && contains_number && !contains_special_characters
}

func contains_numbers_letters_special_chars(password string) bool {
	contains_letters, _ := regexp.MatchString("[a-z]", normalize_password(password))
	contains_number, _ := regexp.MatchString("[0-9]+", password)
	contains_special_characters := contains_special_chars(password)

	return contains_letters && contains_number && contains_special_characters
}

func password_strength(password string) string {
	switch {
		case contains_all_numbers(password):
			return "very weak"
		case contains_all_letters(password):
			return "weak"
		case contains_number_and_letters(password):
			return "strong"
		case contains_numbers_letters_special_chars(password):
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
