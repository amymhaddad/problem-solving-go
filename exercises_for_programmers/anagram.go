/*
Using functions to abstract the logic away from the rest of
your code makes it easier to read and easier to maintain.
Create a program that compares two strings and determines if the two strings are anagrams. The program should prompt for both input strings and display the output as shown in the example that follows.
Example Output
       Enter two strings and I'll tell you if they
       are anagrams:
       Enter the first string: note
       Enter the second string: tone
"note" and "tone" are anagrams.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	count_letters()
}

func count_letters() {

	word := "tone"

//	counts_per_letter := make(map[int] int)

	for _, letter := range strings.ToLower(word) {
		///HERE: check of the letter is already in the map. If so, then add one to the count.
		//Check what "_, letter" retursn -- buyt update _ to be a variable 
		
	// 	if _, ok := counts_per_letter[letter]; ok {
	// 		counts_per_letter[letter] += 1
	// 	}  else {
        
	// 		counts_per_letter[letter] = 1
	// 	}

	// }
	//fmt.Println(counts_per_letter)


}
