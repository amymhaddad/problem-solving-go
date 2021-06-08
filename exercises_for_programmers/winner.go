/*
Create a program that picks a winner for a contest or prize drawing. Prompt for names of contestants until the user leaves the entry blank. Then randomly select a winner.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func contestants() []string {

	names := []string{}
	for {
		var name string
		fmt.Printf("Enter a name: ")
		fmt.Scanf("%s", &name)
		if name == "" {
			return names
		}
		names = append(names, name)
	}
	return names
}

func select_winner(contestants []string) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(contestants))
	return contestants[index]
}

func remove_winner(contestant_names *[]string, winner string) []string {

	for i, name := range *contestant_names {
		if name == winner {
			*contestant_names = append((*contestant_names)[:i], (*contestant_names)[i+1:]...)
		}
	}
	return *contestant_names
}

func display_contesants(contestant_names []string) {
	for _, name := range contestant_names {
		fmt.Printf(name)
		fmt.Println()
	}
}

func main() {
	contestant_names := contestants()
	winner := select_winner(contestant_names)

	fmt.Printf("The winner is: %s\n", winner)
	remove_winner(&contestant_names, winner)

	fmt.Println("The remaining contestants are: ")
	display_contesants(contestant_names)
}