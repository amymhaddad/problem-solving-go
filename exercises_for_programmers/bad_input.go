package main

import (
	"fmt"
	"regexp"
)

func main() {
	rate := user_rate()
	fmt.Println(rate)
}

func user_rate() int {
	
	var rate string
	fmt.Printf("Enter a rate: ")
	fmt.Scanf("%s", &rate)

	valid_rate, _ := regexp.MatchString("[1-9]+", rate)

	for valid_rate {
		if valid_rate  {
			//Convert the rate into an int, break (see if htat's possible) and call the calculation function 

			fmt.Println("valid")
		} else {
			//ask for user input again
			fmt.Println("invalid")
		}


	}
	

}


