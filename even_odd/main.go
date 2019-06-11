package main

import (
	"fmt"
)


func isEvenOdd(number int) string {
	if number % 2 == 0 {
		return "even"
	}
	return "odd"
}

func main() {

	numbers := []int{0,1,2,3,4,5,6,7,8,9,10}

	for _, number := range numbers {
		fmt.Printf("%d is %s\n", number, isEvenOdd(number))
	}

}