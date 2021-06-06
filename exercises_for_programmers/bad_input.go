package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	rate := user_rate()
	fmt.Println(rate)

}

func user_rate() int {
//make htis a func	
	var rate string
	//fmt.Printf("Enter a rate: ")
	//fmt.Scanf("%s", &rate)


	for {

		fmt.Printf("Enter a rate: ")
		fmt.Scanf("%s", &rate)

		valid_rate, _ := regexp.MatchString("[1-9]+", rate)
		if !valid_rate{
			fmt.Printf("Enter a rate: ")
			fmt.Scanf("%s", &rate)
			continue
		} else {
			user_rate, _ := strconv.Atoi(rate)
			return calculate_investment_return(user_rate)
		}

	}
}

func calculate_investment_return(user_rate int) int {
	return 72 / user_rate

}	

