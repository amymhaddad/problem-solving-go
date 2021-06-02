package main

import (
	"fmt"
)

const bmi_low_boundary float64 = 18.5
const bmi_high_boundary float64 = 25

func main() {
	height, weight := user_height_weight()
//	fmt.Println(height, weight)

	bmi := calculate_bmi(height, weight)
//	fmt.Println(bmi)

	is_normal_weight := bmi_range(bmi)
	fmt.Println(is_normal_weight)
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

func calculate_bmi(height float64, weight(float64)) (float64) {
	return (weight / (height * height)) * 703
}

func bmi_range(bmi float64) (bool) {
	return bmi >= bmi_low_boundary && bmi <= bmi_high_boundary
}

