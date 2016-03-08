package main

import "fmt"

func main() {
	fmt.Println("The list of the even numbers from 0 to 100 is the following: ")
	for even := 1; even <= 100; even++ {
		if even%2 == 0 {
			fmt.Println(even)
		}
	}
}