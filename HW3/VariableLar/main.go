
package main

import "fmt"

func largest(nums ...int) int {
	var num int
	for _, v := range nums {
		if v > num {
			num = v
		}
	}
	return num
}
func main() {
	value := largest(99, 9, 1, 0, 81, 999, 25, 18)
	fmt.Println(value)
}