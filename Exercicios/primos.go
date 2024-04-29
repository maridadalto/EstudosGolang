package main

import (
	"fmt"
	"math"
)

func Primos(nums []int) []int {
	var primos []int
	for _, num := range nums {
		if num <= 1 {
			continue
		}
		if num == 2 {
			primos = append(primos, num)
		}
		if num%2 == 0 {
			continue
		}
		primo := true
		for i := 3; i <= int(math.Sqrt(float64(num))); i += 2 {
			if num%i == 0 {
				primo = false
				break
			}
			if primo {
				primos = append(primos, num)
			}

		}

	}
	return primos
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	primos := Primos(numbers)
	fmt.Println(primos)
}
