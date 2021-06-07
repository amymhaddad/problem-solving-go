/*
Create a small program that contains a list of employee names. Print out the list of names when the program runs the first time. Prompt for an employee name and remove that specific name from the list of names. Display the remaining employees, each on its own line.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	employees := []string{"John Smith", "Jackie Jackson", "Chris Jones", "Amanda Cullen", "Jeremy Goodwin"}
	//display_employees(employees)

	employee_name := remove_employee(employees)
	fmt.Println(employee_name)
}

func display_employees(employees []string) {
	
	fmt.Printf("There are %d employees:\n", len(employees))

	for _, name := range employees {
		fmt.Println(name)
	}

	//unclear why this indexing doesn't work 
	// for i:=0; i<len(*employees); i++{
	// 	fmt.Println(employees[i])
	// }	
}

func remove_employee(employees []string) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(employees)-1)
	return employees[index]


}
