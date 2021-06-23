/*
Write a quick calculator that prompts for the rate of return on an investment and calculates how many years it will take to double your investment.

Example:
What is the rate of return? ABC
Sorry. That's not a valid input.
What is the rate of return? 4
It will take 18 years to double your initial investment.
*/

package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func getUserRate() string {

	var rate string
	fmt.Printf("Enter a rate: ")
	fmt.Scanf("%s", &rate)
	return rate
}

func validateUserRate(rate string) int {
	for {
		validRate, _ := regexp.MatchString("[1-9]+", rate)
		if !validRate {
			fmt.Printf("Sorry. '%s' is an invalid rate. Try again.\n", rate)
			rate = getUserRate()
			fmt.Println(rate)
		} else {
			userInvestmentRate, _ := strconv.Atoi(rate)
			return userInvestmentRate
		}
	}
}

func calcInvestmentReturn(validRate int) int {

	return 72 / validRate
}

func main() {
	userReturnRate := getUserRate()
	validRate := validateUserRate(userReturnRate)
	yearsToDoubleInvestment := calcInvestmentReturn(validRate)

	fmt.Printf("It'll take %d years to double your investment.\n", yearsToDoubleInvestment)
}
