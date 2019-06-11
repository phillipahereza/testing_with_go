package array

// Sum returns the sum of an array of integers
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAll returns a slice of the sum of individual slices
func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

// SumAllTails returns a slice of the sum of the tails of individual slices
func SumAllTails(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		if len(numbers) > 1 {
			sums[i] = Sum(numbers[1:])
		} else {
			sums[i] = 0
		}

	}
	return sums
}
