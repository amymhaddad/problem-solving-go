package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func user_rate() string {

	var rate string
	fmt.Printf("Enter a rate: ")
	fmt.Scanf("%s", &rate)
	return rate
}

func validate_user_rate(rate string) int {
	for {
		valid_rate, _ := regexp.MatchString("[1-9]+", rate)
		if !valid_rate{
			fmt.Printf("Sorry. '%s' is an invalid rate. Try again.\n", rate)
			rate = user_rate()
			fmt.Println(rate)
			continue
		} else {
			user_investment_rate, _ := strconv.Atoi(rate)
			return calculate_investment_return(user_investment_rate)
		}
	}
}

func calculate_investment_return(user_rate int) int {
	return 72 / user_rate
}	

func main() {
	rate_of_return := user_rate()
	years_to_double_investment := validate_user_rate(rate_of_return)
	fmt.Printf("It'll take %d years to double your investment.\n", years_to_double_investment) 
}