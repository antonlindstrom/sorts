package heapsort

func Sort(set []int) []int {
	length := len(set)
	end := length-1

	set = heapify(set)

	for end > 0 {
		set[end], set[0] = set[0], set[end]
		end = end-1

		siftDown(set, 0, end)
	}

	return set
}

func heapify(set []int) []int {
	length := len(set)

	start := (length - 2 ) / 2

	for start >= 0 {
		set = siftDown(set, start, length-1)
		start = start-1
	}

	return set
}

func siftDown(set []int, start, end int) []int {
	root := start
	for (root * 2 + 1 <= end) {
		child := root * 2 + 1
		swap  := root

		if set[root] < set[child] {
			swap = child
		}
		if (child+1 <= end && set[swap] < set[child+1]) {
			swap = child + 1
		} 
		if swap != root {
			set[root], set[swap] = set[swap], set[root]
			root = swap
		} else {
			return set
		}

	}
	return set
}
