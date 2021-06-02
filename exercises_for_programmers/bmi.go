package main

import (
	"fmt"
)

func main() {
	height, weight := user_height_weight()
	fmt.Println(height, weight)
}

func user_height_weight() (float64, float64) {
	var height, weight int

	fmt.Printf("Enter your height: ")
	fmt.Scanf("%d", &height)

	fmt.Printf("Enter your weight: ")
	fmt.Scanf("%d", &weight)

	user_height := float64(height)
	user_weight := float64(weight)
	return user_height, user_weight

}
