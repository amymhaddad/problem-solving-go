/*
Quotation marks are often used to denote the start and end of a string. But sometimes we need to print out the quotation marks themselves by using escape characters.
Create a program that prompts for a quote and an author. Display the quotation and author as shown in the example output.
Example Output
What is the quote? These aren't the droids you're looking for. Who said it? Obi-Wan Kenobi
Obi-Wan Kenobi says, "These aren't the droids
you're looking for."
*/

package main

import (
	"bufio"
	"fmt"
	"os"
) 

func quote_and_author() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your quote: ")

	scanner.Scan()

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	quote := scanner.Text()

	var author string
	fmt.Printf("Enter the author of the quote: ")
	fmt.Scanf("%s", &author)
	return quote, author
}

func main() {
	quote, author := quote_and_author()
	fmt.Printf("'%s' said: '%s'", author, quote)
}

