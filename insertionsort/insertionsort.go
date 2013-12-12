package insertionsort

func Sort(set []int) []int {
	for i := 1; i < len(set); i++ {
		value := set[i]
		hole := i

		for hole > 0 && value < set[hole-1] {
			set[hole] = set[hole-1]
			hole = hole - 1
		}

		set[hole] = value
	}

	return set
}
