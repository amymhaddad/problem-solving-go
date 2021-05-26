/*
Create a program that prompts for an input string and dis- plays output that shows the input string and the number of characters the string contains.
Example Output
What is the input string? Homer Homer has 5 characters.
*/


package main

import "fmt"

func main() {
	fmt.Println("Enter a word: ")
	var word string
	fmt.Scanln(&word)

	total_letters := count_chars(word) 
	fmt.Println(total_letters)
}

func count_chars(word string) int{
	total_chars := 0

	for i := 0; i < len(word); i ++ {
		total_chars += 1
	}
	return total_chars
}
