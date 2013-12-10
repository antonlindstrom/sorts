package shellsort

func Sort(set []int) []int {
	gap := len(set) / 2

	for gap > 0 {
		for i := gap; i < len(set); i++ {
			temp := set[i]
			j := i

			for j >= gap && set[j-gap] > temp {
				set[j] = set[j-gap]
				j -= gap
			}
			set[j] = temp
		}
		gap = gap / 2
	}

	return set
}
