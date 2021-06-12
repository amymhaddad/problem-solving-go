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
	totalSpecialChars int
	totalNums int
	totalLetters int
}

func userPasswordSpecs() (*password, error) {
	//I'm creating an instance of the struct, password.
	//bc I create an instance of the struct here, I return the address of the
	//struct via &user_pw
	var userPassword password
	
	var pwLength, totalSpecialChars, totalNums int
	fmt.Printf("Enter the minimum password length: ")
	fmt.Scanf("%d", &pwLength)

	fmt.Printf("Enter the total special characters: ")
	fmt.Scanf("%d", &totalSpecialChars)

	fmt.Printf("Enter the total numbers: ")
	fmt.Scanf("%d", &totalNums)

	//Added error handling
	if totalSpecialChars < 0 || totalNums < 0 {
		return &userPassword, fmt.Errorf("Invalid number entry.")
	}

	totalLetters := pwLength - (totalSpecialChars + totalNums)

	userPassword.length = pwLength
	userPassword.totalSpecialChars = totalSpecialChars
	userPassword.totalNums = totalNums 
	userPassword.totalLetters = totalLetters
	return &userPassword, nil 
}

//This function takes the struct, user_password
//Use the * to indicate that I'm passing by reference. 
//That's bc I want to reference the data stored in memory associated
//w/user_password
func generate_pw(user_password *password) string {
	var new_password []string
	
	//Use < NOT <= bc all of the function calls run on each iteration. This will cause the program to run 1 too many times
	for len(new_password) <= user_password.length {
		new_password = append(new_password, generate_random_chars(special_chars, user_password.totalSpecialChars))
		new_password = append(new_password, generate_random_chars(numbers, user_password.totalNums))

		if user_password.totalLetters == 0 {
			break 
		} else {
			new_password = append(new_password, generate_random_chars(letters, user_password.totalLetters))
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
		
		total_chars--
		password += new_char
	}	
	return password
}
//IF need to exit immediately: os.exit()
func main() {	
	var password_criteria *password
	var err error

	for {
		password_criteria, err = userPasswordSpecs()
		if err == nil {
			break
		}
		fmt.Println(err, "Please try again.")
	}

	user_password := generate_pw(password_criteria)
	fmt.Printf("The password is %s\n", user_password)
}
