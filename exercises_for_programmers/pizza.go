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
	 //"strconv"
 )

//Is it best practice to use the functions to get the return values and then use the
//main funcion to DO somethign w/the return values? So in this case I could...
func main() {

	fmt.Println(calculate_pizza_slices())
	
}

func calculate_pizza_slices() int {
	
	// fmt.Println("How many people? ")
	// var total_people int
	// fmt.Scanln(&total_people)
    //
	fmt.Println("How many pizzas do you have? ")
	var total_pizzas int
	fmt.Scanf("%d", &total_pizzas)


//	i :=strconv.Itoa(total_people)
	return total_pizzas

}		
