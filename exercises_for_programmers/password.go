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

func weak_password(password string) string {
	if contains_all_numbers(password) {
		return "very weak"
	} else if contains_all_letters(password) {
		return "weak"
	}
}

func strong_password(password string) string {
	if contains_number_and_letters(password) {
		return "strong"
	} else if contains_numbers_letters_special_chars(password) {
		return "very strong"
	}
}

func password_strength(password string) string {
	return "here"
	// if weak_password_length(password) {
	// 	weak_result := weak_password(password)
	// 	return weak_result
	// 	} else if strong_password_length(password) {
	// 		strong_result := strong_password(password)
	// 		return strong_result
	//  } 

}

func main() {

	password := user_password()
	
	user_password_strength := password_strength(password)
	fmt.Println(user_password_strength)
}



// func contains_special_chars(password string) bool {
// 	found, _ := regexp.MatchString("[^a-z0-9]", normalize_password(password))
// 	return found
// }


// func contains_number_and_letters(password string) bool {
// 	contains_letters, _ := regexp.MatchString("[a-z]", normalize_password(password))
// 	contains_number, _ := regexp.MatchString("[0-9]+", password)
// 	contains_special_characters := contains_special_chars(password)

// 	if contains_special_characters {
// 		return false
// 	} else if contains_letters && contains_number {
// 		return true
// 	}
// 	//Unclear why I need to have a return at the end for this function to run,
// 	//given that my conditional above handles the return?
// 	return false 
// }


