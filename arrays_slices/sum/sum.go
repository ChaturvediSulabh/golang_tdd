package sum

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}
	return
}

func SumAllTrails(numbersToSum ...[]int) (sums []int) {
	for _, nums := range numbersToSum {
		if len(nums) > 0 {
			sums = append(sums, Sum(nums[1:]))
		} else {
			sums = append(sums, 0)
		}
	}
	return
}
