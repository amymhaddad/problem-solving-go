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


//define the struct
func main() {
	pw_length, total_special_chars, total_nums, total_letters := user_pw_specs()
	generate_pw(pw_length, total_special_chars, total_nums, total_letters)
}


func user_pw_specs() (int, int, int, int) {
	//add to the struct	
	var pw_length, total_special_chars, total_nums int 

	fmt.Printf("Enter the minimum password length: ")
	fmt.Scanf("%d", &pw_length)

	fmt.Printf("Enter the total special characters: ")
	fmt.Scanf("%d", &total_special_chars)

	fmt.Printf("Enter the total numbers: ")
	fmt.Scanf("%d", &total_nums)

	total_letters := pw_length - (total_special_chars + total_nums)
	//return a struct that contains each field 
	return pw_length, total_special_chars, total_nums, total_letters
}

//take a struct
//researach when to use byte array
func generate_pw(pw_length int, total_special_chars int, total_nums int, total_letters int) string {
	//instead of string, use byte array. Then I don't have to append -- it's
	//more efficient

	//use var when setting to zero value
	password := []string{}

	for len(password) <= pw_length {
		password = append(password, generate_random_chars(special_chars, total_special_chars))
		password = append(password, generate_random_chars(numbers, total_nums))
		password = append(password, generate_random_chars(letters, total_letters))
	}	
	fmt.Println("New pw: ", password)
	
	return strings.Join(password, "")
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
	//fmt.Println("generate: ", password)
	return password
}
