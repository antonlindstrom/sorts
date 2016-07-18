package quicksort

// Sort sorts a slice of integers with quicksort.
func Sort(set []int) []int {
	if len(set) <= 1 {
		return set
	}

	// Worst case pivot is to select the leftmost.
	pivot := set[0]
	set = set[1:]

	var less []int
	var greater []int

	for _, i := range set {
		if i <= pivot {
			less = append(less, i)
		} else {
			greater = append(greater, i)
		}
	}

	// Append slices.
	return append(append(Sort(less), pivot), Sort(greater)...)
}
