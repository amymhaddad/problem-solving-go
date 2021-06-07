/*
Create a small program that contains a list of employee names. Print out the list of names when the program runs the first time. Prompt for an employee name and remove that specific name from the list of names. Display the remaining employees, each on its own line.
*/	


package main

import "fmt"

func main() {
	employees := []string{"John Smith", "Jackie Jackson", "Chris Jones", "Amanda Cullen", "Jeremy Goodwin"}
	
	fmt.Println(display_employees(employees))
}

func display_employees(employees slice) {
	
	fmt.Printf("There are %d employees:\n", len(employees))

	for i=0; i<len(employees); i++{
		fmt.Println(employees[i])
	}	
}
