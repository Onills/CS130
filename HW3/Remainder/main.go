package main

import "fmt"

func main() {
	var ssum int
	var lsum int
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&ssum)
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&lsum)
	var remainder = lsum % ssum
	fmt.Println("The remainder is:", remainder)
}