/*
Quotation marks are often used to denote the start and end of a string. But sometimes we need to print out the quotation marks themselves by using escape characters.
Create a program that prompts for a quote and an author. Display the quotation and author as shown in the example output.
Example Output
What is the quote? These aren't the droids you're looking for. Who said it? Obi-Wan Kenobi
Obi-Wan Kenobi says, "These aren't the droids
you're looking for."
*/

package main

import "fmt"

func quote_and_author() (string, string) {
	var quote, author string
	fmt.Printf("Enter your quote: ")
	fmt.Scanf("%s", &quote)

	fmt.Printf("Enter the author of the quote: ")
	fmt.Scanf("%s", &author)
	return quote, author
}

func main() {
	quote, author := quote_and_author()
	fmt.Printf("'%s' said: '%s'\n", author, quote)
}

