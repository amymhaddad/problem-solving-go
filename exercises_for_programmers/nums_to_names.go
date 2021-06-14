/*
Write a program that converts a number from 1 to 12 to the corresponding month. Prompt for a number and display the corresponding calendar month, with 1 being January and 12 being December. For any value outside that range, display an appropriate error message.
*/

package main

import (
	"fmt"
)

var months = map[int]string {1: "January", 2: "February", 3: "March", 4: 
"April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: 
"October", 11: "November", 12: "December"}

func user_number() (int, error) {
	var number int
	start_month := 1 
	end_month := 12

	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &number)

	if number < start_month || number > end_month {
		return number, fmt.Errorf("Invalid number")
	}
	return number, nil
}

func main() {
	for {
		calendar_num, err := user_number()
		if err == nil {
			fmt.Printf("The number of the month is %s", months[calendar_num])
			break 
		} else {
			fmt.Println("Please try again.")
		}	
	}
 }
