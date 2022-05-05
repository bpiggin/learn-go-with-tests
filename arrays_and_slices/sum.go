package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberArrays ...[]int) (sums []int) {
	result := []int{}
	for _, numberArray := range numberArrays {
		result = append(result, Sum(numberArray))
	}
	return result
}

func SumAllTails(numberArrays ...[]int) (sums []int) {
	result := []int{}
	for _, numberArray := range numberArrays {
		if len(numberArray) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(numberArray[1:]))
		}
	}
	return result
}
