package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var question string
	fmt.Printf("Enter your question: ")
	fmt.Scanln(&question)

	result := response()

	fmt.Println(result)
}

func response() string {
	//yields a constantly chagning number
	rand.Seed(time.Now().UnixNano())

	possible_answers := []string{"yes", "no", "maybe", "ask again later"}

	index := rand.Intn(len(possible_answers))
	fmt.Println(rand.Intn(index))
	
	return possible_answers[index]

}
