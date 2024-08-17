package sum

func Sum(sli []int) int {
	sum := 0
	for _, v := range sli {
		sum += v
	}

	return sum
}

func SumAll(arrs ...[]int) []int {
	sums := []int{}
	for _, s := range arrs {
		sums = append(sums, Sum(s))
	}
	return sums
}

func SumTails(arrs ...[]int) []int {
	sums := []int{}
	for _, s := range arrs {
		if len(s) <= 0 {
			sums = append(sums, 0)
			continue
		}
		sums = append(sums, Sum(s[1:]))
	}
	return sums
}
