package main

import (
	"fmt"
)

func multiplication_table() {
	max_number := 12
	
	for i:=0; i< max_number; i++ {
		for j:= 0; j<=i; j ++ {
			result := i * j
			fmt.Printf("%d x %d = %d\n", i, j, result)
		}
	}
}

func main() {
	multiplication_table()
}
