/*
Exercise 8
Write a program to evenly divide pizzas. Prompt for the number of people, the number of pizzas, and the number of slices per pizza. Ensure that the number of pieces comes out even. Display the number of pieces of pizza each person should get. If there are leftovers, show the number of leftover pieces.
Example Output
       How many people? 8
       How many pizzas do you have? 2
	   8 people with 2 pizzas
Each person gets 2 pieces of pizza. There are 0 leftover pieces.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	slices_per_person, leftover_slices := calculate_pizza_slices()
	pizza_distribution := "Each person gets " + slices_per_person + " total slices of pizza."
	total_leftovers := " There are " + leftover_slices + " slices of pizza left over."
	fmt.Println(pizza_distribution + total_leftovers)
}

func calculate_pizza_slices() (string, string){
	var people, pizzas, slices_per_pizza int

	fmt.Printf("How many people? ")
	fmt.Scanf("%d\n", &people)

	fmt.Printf("How many pizzas do you have? ")
	fmt.Scanf("%d\n", &pizzas)

	fmt.Printf("How many slices per pizza? ")
	fmt.Scanf("%d\n", &slices_per_pizza)
	
	total_pizza_slices := pizzas * slices_per_pizza 
	slices_per_person := total_pizza_slices / people
	leftover_slices := total_pizza_slices % people

	return strconv.Itoa(slices_per_person), strconv.Itoa(leftover_slices)
}