package combsort

// Sort sorts a slice of integers with comb sort.
func Sort(set []int) []int {
	gap := len(set)
	shrink := 1.3
	swapped := false

	for gap > 1 || swapped {
		gap = int(float64(gap) / shrink)

		if gap < 1 {
			gap = 1
		}

		swapped = false

		for i := 0; gap+i < len(set); i++ {
			if set[i]-set[i+gap] > 0 {
				swap := set[i]
				set[i] = set[i+gap]
				set[i+gap] = swap
				swapped = true
			}
		}
	}

	return set
}
