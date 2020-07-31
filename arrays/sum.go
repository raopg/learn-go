package arrays

//ArraySum takes an array and returns its sum
func ArraySum(numbers [5]int) (sum int) {
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return
}
