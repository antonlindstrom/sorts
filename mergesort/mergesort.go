package mergesort

// Mergesort (top-down)
func Sort(set []int) []int {
	if len(set) <= 1 {
		return set
	}

	var left []int
	var right []int

	middle := len(set) / 2

	// Append all before middle to left
	for _, i := range set[:middle] {
		left = append(left, i)
	}

	// Append all after middle to right
	for _, i := range set[middle:] {
		right = append(right, i)
	}

	// Recurse, yo!
	left = Sort(left)
	right = Sort(right)

	// Jump to merge, which is the does the organizing
	return merge(left, right)
}

// Merging
func merge(left, right []int) (result []int) {

	// Run through all until the sides are empty
	for len(left) > 0 || len(right) > 0 {

		// Compare sides
		if len(left) > 0 && len(right) > 0 {

			if left[0] <= right[0] {

				result = append(result, left[0])
				left = left[1:]

			} else {

				result = append(result, right[0])
				right = right[1:]

			}

		} else if len(left) > 0 {

			result = append(result, left[0])
			left = left[1:]

		} else if len(right) > 0 {

			result = append(result, right[0])
			right = right[1:]

		}

	}

	return result
}
