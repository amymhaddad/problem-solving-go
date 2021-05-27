/*
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
	fmt.Println(calculate_pizza_slices())
}

func calculate_pizza_slices() string{
	var total_people, total_slices_per_pizza, total_pizzas int
	
	fmt.Printf("How many people? ")
	fmt.Scanf("%d\n", &total_people)

	fmt.Printf("How many pizzas do you have? ")
	fmt.Scanf("%d\n", &total_pizzas)

	fmt.Printf("How many slices per pizza? ")
	fmt.Scanf("%d\n", &total_slices_per_pizza)

	pizza_slices := total_pizzas * total_slices_per_pizza
	slices_per_person := pizza_slices / total_people
	leftovers := pizza_slices % total_people

	total_slices := strconv.Itoa(slices_per_person)
	total_leftovers := strconv.Itoa(leftovers)
	
	calculation := "Each person gets " + total_slices + " pieces of pizza. There are " + total_leftovers + " leftover pieces."

	return calculation
}		
