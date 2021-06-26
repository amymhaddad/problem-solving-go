package main

import (
   "bufio"
   "fmt"
   "os"
   "reflect"
)
 
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the rate of investment: ")
	scanner.Scan()
	if scanner.Err() != nil {
       fmt.Println("Error: ", scanner.Err())
   }
   rate := scanner.Text()

	//Create a fucntoin to get user input. 
	//Create a fucntion that checks the type of user input and if gets bad input
	//recalls the fucntion to get user input. Put this into a for loop.
	//add TypeOf() to account for int and float --> check if %d works on floats
   rate_type := reflect.TypeOf(rate).Kind()

   if rate_type == reflect.String {
	fmt.Println("string")	
   }

}
