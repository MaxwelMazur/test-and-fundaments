package list

func sum(numbers []int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return
}

func sumAll(numbers ...[]int) (listSum []int) {
	for i := range numbers {
		listSum = append(listSum, sum(numbers[i]))
	}
	return
}
