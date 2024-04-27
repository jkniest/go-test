package arrays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersToSum ...[]int) (results []int) {
	for _, numbers := range numbersToSum {
		results = append(results, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (results []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			results = append(results, 0)
		} else {
			tail := numbers[1:]
			results = append(results, Sum(tail))
		}
	}

	return
}
