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

func sumAllRest(numbers ...[]int) (listSum []int) {
	for i := range numbers {
		num := numbers[i]
		if len(num) < 1 {
			num = append(num, 0)
		}
		listSum = append(listSum, sum(num[1:]))
	}
	return
}
