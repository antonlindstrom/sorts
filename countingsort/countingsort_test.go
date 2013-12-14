package countingsort

import "testing"

func TestSort(t *testing.T) {
	sortedSet := Sort([]int{1})

	if sortedSet[0] != 1 {
		t.Fatal("Could not sort dataset with set 1!")
	}

	sortedSet = Sort([]int{10, 9, 4, 7, 3, 1, 8, 5, 2, 6})

	for i := 0; i < 10; i++ {
		if sortedSet[i] != i+1 {
			t.Fatal("Could not sort dataset!")
		}
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort([]int{10, 9, 4, 7, 3, 1, 8, 5, 2, 6})
	}
}
