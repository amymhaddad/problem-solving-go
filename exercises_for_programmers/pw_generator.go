/*
Create a program that generates a secure password. Prompt the user for the minimum length, the number of special characters, and the number of numbers. Then generate a password for the user using those inputs.
*/

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)


var letters = [9]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
var numbers = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var special_chars = [9]string{"!", "@", "#", "$", "%", "^", "&", "*", "("}



func main() {
	pw_length, total_special_chars, total_nums, total_letters := user_pw_specs()
//	fmt.Println(pw_length, total_special_chars, total_nums, total_letters)
	password := generate_pw(pw_length, total_special_chars, total_nums, total_letters)
	fmt.Println(password)
}


func user_pw_specs() (int, int, int, int) {
	var pw_length, total_special_chars, total_nums int 

	fmt.Printf("Enter the minimum password length: ")
	fmt.Scanf("%d", &pw_length)

	fmt.Printf("Enter the total special characters: ")
	fmt.Scanf("%d", &total_special_chars)

	fmt.Printf("Enter the total numbers: ")
	fmt.Scanf("%d", &total_nums)

	total_letters := pw_length - (total_special_chars + total_nums)

	return pw_length, total_special_chars, total_nums, total_letters
}

func generate_pw(pw_length int, total_special_chars int, total_nums int, total_letters int) string {
	password := make([] string, pw_length)

	for len(password) < pw_length {
		password = append(password, generate_random_chars(special_chars, total_special_chars))
		// generate_random_chars(numbers, total_nums)
		// generate_random_chars(letters, total_letters)
	}	
	fmt.Println("here", password)
	return strings.Join(password, " ")
}

func generate_random_chars(chars [9]string, total_chars int) (string) {
	var password string
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(total_chars)

	for total_chars > 0 {
		new_char := chars[index]
		total_chars -= 1
		password += new_char
	}	
	return password
}
