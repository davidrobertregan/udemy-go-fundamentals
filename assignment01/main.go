package main

import "fmt"

func main() {
	nums := []int{}

	for i := 1; i < 11; i++ {
		nums = append(nums, i)
	}

	for _, n := range nums {
		if n % 2 == 0 {
			fmt.Printf("%v is even\n", n)
		} else {
			fmt.Printf("%v is odd\n", n)
		}
	}
}