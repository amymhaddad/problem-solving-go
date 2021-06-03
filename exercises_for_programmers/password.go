package main

import (
	"fmt"
	//"regexp"
)

const standard_password_length int = 8

func main() {
	//// found, err := regexp.MatchString("(^[1-9]+$)", "123")
	//// fmt.Printf("found=%v err=%v", found, err)

	fmt.Println(weak_password_length("23"))
	fmt.Println(strong_password_length("234523432342"))

}


func weak_password_length(password string) (bool) {
	return len(password) < standard_password_length
}

func strong_password_length(password string) (bool) {
	return len(password) >= standard_password_length
}


