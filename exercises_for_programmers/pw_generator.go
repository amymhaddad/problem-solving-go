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

type password struct {
	length int
	total_special_chars int
	total_nums int
	total_letters int
}


func user_pw_specs() *password {
	//I'm creating an instance of the struct, password.
	//bc I create an instance of the struct here, I return the address of the
	//struct via &user_pw
	var user_pw password
	
	var pw_length, total_special_chars, total_nums int
	fmt.Printf("Enter the minimum password length: ")
	fmt.Scanf("%d", &pw_length)

	fmt.Printf("Enter the total special characters: ")
	fmt.Scanf("%d", &total_special_chars)

	fmt.Printf("Enter the total numbers: ")
	fmt.Scanf("%d", &total_nums)

	total_letters := pw_length - (total_special_chars + total_nums)

	user_pw.length = pw_length
	user_pw.total_special_chars = total_special_chars
	user_pw.total_nums = total_nums 
	user_pw.total_letters = total_letters
	return &user_pw
}


func generate_pw(user_password *password) string {
	var new_password []string

	for len(new_password) <= user_password.length{
		new_password = append(new_password, generate_random_chars(special_chars, user_password.total_special_chars))
		new_password = append(new_password, generate_random_chars(numbers, user_password.total_nums))
		if user_password.total_letters == 0 {
			break 
			// return fmt.Errorf("must be greater")
		} else {
			new_password = append(new_password, generate_random_chars(letters, user_password.total_letters))
		}
	 }	
	return strings.Join(new_password, "")
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

func main() {
	// var password_criteria password
	// user_pw_specs(&password_criteria)
	password_criteria := user_pw_specs()
	user_password, err := generate_pw(password_criteria)
	fmt.Printf("The password is %s", user_password)
}
