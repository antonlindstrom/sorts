package countingsort

func Sort(set []int) []int {
	z, min, max := 0, 0, maxNum(set)
	count := make([]int, max-min+1)

	for _, v := range set {
		count[v-min]++
	}

	for i := min; i < max+1; i++ {
		for count[i-min] > 0 {
			set[z] = i + min
			z++
			count[i-min]--
		}
	}

	return set
}

// Find max value in set, we could just assume but
// that's a bit sad when we run over the limit
func maxNum(set []int) int {
	max := 0

	for _, v := range set {
		if v > max {
			max = v
		}
	}

	return max
}
