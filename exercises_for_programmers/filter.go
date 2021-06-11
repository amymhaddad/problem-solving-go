/*
Create a program that prompts for a list of numbers, sepa- rated by spaces. Have the program print out a new list con- taining only the even numbers.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func user_input() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a list a numbers, separated by a space: ")
	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
	user_numbers := scanner.Text()
	return user_numbers
}

// func quote_and_author() (string, string) {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Print("Enter your quote: ")
  
// 	scanner.Scan()
  
// 	if scanner.Err() != nil {
// 		fmt.Println("Error: ", scanner.Err())
// 	}
  
// 	quote := scanner.Text()
  
// 	var author string
// 	fmt.Printf("Enter the author of the quote: ")
// 	fmt.Scanf("%s", &author)
// 	return quote, author
//  }
 

func generate_number_range(user_numbers string)  {
	user_nums := strings.Split(user_numbers, " ")
	even_nums := [] int{}

	for _, num := range user_nums {
		digit, _ := strconv.Atoi(num)
		if digit % 2 == 0 {
			even_nums = append(even_nums, digit)
		}
	}
	fmt.Println(even_nums)
}


func main() {
	fmt.Println(user_input())

	// user_numbers := user_input()
	// fmt.Println(user_numbers)
	// number_range := generate_number_range(user_numbers)
	// fmt.Println(user_numbers)
}

