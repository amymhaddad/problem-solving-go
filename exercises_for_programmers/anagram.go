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
	"reflect"
	"strings"
)

//put main on bottom
//shorten var names
func main() {
	fmt.Println("Enter two strings and I'll tell you if they're anagrams.")
	
	word1, word2 := user_input()
	word1_letter_counter := count_letters(word1)
	word2_letter_counter := count_letters(word2)

	anagram := is_anagram(word1_letter_counter, word2_letter_counter)

	//Printf -- i have to add a line break vs Println -- contains a line break
	//Think of string formatting as Python {}
	if anagram {
		fmt.Printf("'%s' and '%s' are anagrams.\n", word1, word2)
		fmt.Println(word1, word2)
		fmt.Println("hello")

	} else {

		fmt.Printf("'%s' and '%s' are not anagrams.\n", word1, word2)
	}
}

func is_anagram(map1 map[rune]int, map2 map[rune]int) bool {
	return reflect.DeepEqual(map1, map2)
}

func user_input() (string, string){
	var word1, word2 string 
	
	fmt.Printf("Enter the first string: ")
	fmt.Scanf("%s", &word1)

	fmt.Printf("Enter the second string: ")
	fmt.Scanf("%s", &word2)

	return word1, word2
}

func count_letters(word string) map[rune]int {
	letter_count := make(map[rune]int)

	for _, letter := range strings.ToLower(word) {
		if _, ok := letter_count[letter]; ok {
			letter_count[letter] += 1
		} else {
			letter_count[letter] = 1
		}
	}
	return letter_count 
}
