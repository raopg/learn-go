package arrays

//ArraySum takes an array and returns its sum
func ArraySum(numbers []int) (sum int) {
	for _, num := range numbers { // Range keyword provides us index and number as an iterator.
		sum += num
	}
	return
}

//SumAll takes a varying number of slices and returns a slice of correspoding sums.
func SumAll(numbersToSum ...[]int) []int { //...syntax is to highlight the a variable number of specified type
	// lengthOfNumbers := len(numbersToSum) // Python syntax! nice
	// sums := make([]int, lengthOfNumbers) //make function can be called to 'make' slices
	// Instead, we can use append syntax
	var sums []int
	for _, nums := range numbersToSum {
		sums = append(sums, ArraySum(nums))
	}

	return sums
}

//SumAllTail function takes a slice of slices and returns the tail sum -> sums of each slice except the first elem.
func SumAllTail(numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, ArraySum(nums[1:])) // Slices can be sliced!
		}
	}

	return sums
}
