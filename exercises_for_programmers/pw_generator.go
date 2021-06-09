/*
Create a program that generates a secure password. Prompt the user for the minimum length, the number of special characters, and the number of numbers. Then generate a password for the user using those inputs.
*/

package main

import (
	"fmt"
	//"math/rand"
)


var letters = [9]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
var numbers = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var special_chars = [9]string{"!", "@", "#", "$", "%", "^", "&", "*", "("}



func main() {
	pw_length, total_special_chars, total_nums, total_letters := user_pw_specs()
	fmt.Println(pw_length, total_special_chars, total_nums, total_letters)
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
