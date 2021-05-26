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

//Is it best practice to use the functions to get the return values and then use the
//main funcion to DO somethign w/the return values?
func main() {

	people := calculate_pizza_slices()
	fmt.Println(people)
}


//Can't ask for input sequentially. Otherwise the questions get asked back to back and I can't add input for each question
func calculate_pizza_slices() string  {
	
	fmt.Println("How many people? ")
	var total_people int
	fmt.Scanln("%d", &total_people)

	fmt.Println("How many pizzas do you have? ")
	var total_pizzas int
	fmt.Scanf("%d", &total_pizzas)
	
	pizza_slices := total_people * total_pizzas
	slices_per_person := total_people / pizza_slices
	leftovers := total_people % pizza_slices

	fmt.Println(slices_per_person, leftovers)

	total_slices := strconv.Itoa(slices_per_person)
	total_leftovers := strconv.Itoa(leftovers)
	
	calculation := "Each person gets " + total_slices + "pieces of pizza. There are " + total_leftovers + "leftover pieces."
	return calculation
 

}		
