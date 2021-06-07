/*
Create a small program that contains a list of employee names. Print out the list of names when the program runs the first time. Prompt for an employee name and remove that specific name from the list of names. Display the remaining employees, each on its own line.
*/

package main

import (
	"fmt"
)

func main() {
	employees := []string{"John Smith", "Jackie Jackson", "Chris Jones", "Amanda Cullen", "Jeremy Goodwin"}
	fmt.Printf("There are %d employees:\n", len(employees))
	fmt.Println()

	display_employees(employees)
	fmt.Println()

	var first, last string
	fmt.Printf("Enter an employee to remove: ")
	fmt.Scanf("%s %s\n", &first, &last)
	//create first and last into a single string and pass the single string into
	//function
	fmt.Println(first, last)

// 	remaining_employees := remove_employee(&employees, employee_name)
//
// 	fmt.Printf("There are %d remaining employees:\n", len(remaining_employees))
// 	display_employees(remaining_employees)
 }

func display_employees(employees []string) {
	for _, name := range employees {
		fmt.Println(name)
	}	
}


func remove_employee(employees *[]string, employee_name string) []string {
	for i, name := range *employees {
		if name == employee_name {
			fmt.Println("HERE", name, i)
			*employees = append((*employees)[:i], (*employees)[i+1:]...)
		}
	}
	return *employees
}
