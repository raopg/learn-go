package arrays

//ArraySum takes an array and returns its sum
func ArraySum(numbers []int) (sum int) {
	for _, num := range numbers { // Range keyword provides us index and number as an iterator.
		sum += num
	}
	return
}
