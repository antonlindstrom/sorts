package bubblesort

// Sort sorts a slice of integers with bubble sort.
func Sort(set []int) []int {
	swapped := false

	for i := 1; i < len(set); i++ {
		if set[i-1] > set[i] {
			a, b := set[i-1], set[i]

			set[i-1] = b
			set[i] = a

			swapped = true
		}
	}

	// Recurse until we don't swap.
	if swapped {
		set = Sort(set)
	}

	return set
}
