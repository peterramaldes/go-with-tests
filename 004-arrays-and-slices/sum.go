package arrays_and_slices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(arr ...[]int) []int {
	var result []int
	for _, v := range arr {
		result = append(result, Sum(v))
	}
	return result
}
