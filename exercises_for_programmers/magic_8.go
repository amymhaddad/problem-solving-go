package main

import (
	"fmt"
	"math/rand"
	"time"
)

func response() string {
	rand.Seed(time.Now().UnixNano())
	possible_answers := []string{"yes", "no", "maybe", "ask again later"}

	index := rand.Intn(len(possible_answers))  
	return possible_answers[index]
}

func main() {
	var question string
	fmt.Printf("Enter your question: ")
	fmt.Scanln(&question)

	result := response()
	fmt.Println(result)
}
