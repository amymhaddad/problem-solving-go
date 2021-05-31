package main

import (
	"fmt"
	"math"
)

const paint_gallon_coverage = 350

func main() {
	room_square_footage := room_size()
	//fmt.Println(square_footage)
	paint_cans := total_paint_cans(room_square_footage)
	fmt.Println(paint_cans)
	
}

func room_size() int {
	var length, width int 
	
	fmt.Printf("Enter the length of the room: ")
	fmt.Scanf("%d", &length)

	fmt.Printf("Enter the width of the room: ")
	fmt.Scanf("%d", &width)
	return length * width
}

func total_paint_cans(room_square_footage int) float64 {

	paint_gallons := float64(room_square_footage) / float64(paint_gallon_coverage)
	return math.Ceil(paint_gallons)
}
