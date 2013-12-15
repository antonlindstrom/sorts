package heapsort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	sortedSet := Sort([]int{1})

	if sortedSet[0] != 1 {
		t.Fatal("Could not sort dataset with set 1!")
	}

	sortedSet = Sort([]int{10, 9, 4, 7, 3, 1, 8, 5, 2, 6})
	resultSet := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if !reflect.DeepEqual(sortedSet, resultSet) {
		t.Fatal("Could not sort dataset!")
	}

	sortedSet = Sort([]int{10, 9, 9, 7, 3, 7, 8, 5, 2, 6})
	resultSet = []int{2, 3, 5, 6, 7, 7, 8, 9, 9, 10}

	if !reflect.DeepEqual(sortedSet, resultSet) {
		t.Fatal("Could not sort dataset!")
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort([]int{10, 9, 4, 7, 3, 1, 8, 5, 2, 6})
	}
}
