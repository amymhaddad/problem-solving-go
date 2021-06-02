package main

import (
	"fmt"
	"math"
	"strconv"
)

const paint_gallon_coverage = 350

func main() {
	room_square_footage := room_size()
	paint_cans := total_paint_cans(room_square_footage)
	
	fmt.Printf("You'll need to purchase %s cans of paint to cover %s square feet", strconv.Itoa(room_square_footage), paint_cans)
}

func room_size() int {
	var length, width int 
	
	fmt.Printf("Enter the length of the room: ")
	fmt.Scanf("%d", &length)

	fmt.Printf("Enter the width of the room: ")
	fmt.Scanf("%d", &width)
	return length * width
}

func total_paint_cans(room_square_footage int) string{
	paint_gallons := float64(room_square_footage) / float64(paint_gallon_coverage)
	total_paint_cans := math.Ceil(paint_gallons)
	return strconv.FormatFloat(total_paint_cans, 'f', 0, 64)
}
