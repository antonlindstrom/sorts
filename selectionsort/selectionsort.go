package selectionsort

// Sort sorts a slice of integers with selection sort.
func Sort(set []int) []int {
	min := 0

	for j := 0; j < len(set)-1; j++ {
		min = j

		for i := j + 1; i < len(set); i++ {
			if set[i] < set[min] {
				min = i
			}
		}

		if min != j {
			set[j], set[min] = set[min], set[j]
		}
	}

	return set
}
