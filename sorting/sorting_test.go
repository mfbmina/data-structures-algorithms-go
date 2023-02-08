package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sortTests = []struct {
	input    []int
	expected []int
	name     string
}{
	//Sorted slice
	{
		input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		name:     "Sorted Unsigned",
	},
	//Reversed slice
	{
		input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		name:     "Reversed Unsigned",
	},
	//Sorted slice
	{
		input:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		name:     "Sorted Signed",
	},
	//Reversed slice
	{
		input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
		expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		name:     "Reversed Signed",
	},
	//Reversed slice, even length
	{
		input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
		expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		name:     "Reversed Signed #2",
	},
	//Random order with repetitions
	{
		input:    []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
		expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
		name:     "Random order Signed",
	},
	//Single-entry slice
	{
		input:    []int{1},
		expected: []int{1},
		name:     "Singleton",
	},
	// Empty slice
	{
		input:    []int{},
		expected: []int{},
		name:     "Empty Slice",
	},
}

func TestBubbleSort(t *testing.T) {
	for _, sT := range sortTests {
		r := BubbleSort(sT.input)
		assert.Equal(t, sT.expected, r)
	}
}

func TestInsertSort(t *testing.T) {
	for _, sT := range sortTests {
		r := InsertSort(sT.input)
		assert.Equal(t, sT.expected, r)
	}
}

func TestHeapSort(t *testing.T) {
	for _, sT := range sortTests {
		r := HeapSort(sT.input)
		assert.Equal(t, sT.expected, r)
	}
}

func BenchmarkSorting(b *testing.B) {
	slice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10}

	b.Run("BubbleSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BubbleSort(slice)
		}
	})

	b.Run("InsertSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			InsertSort(slice)
		}
	})

	b.Run("HeapSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			InsertSort(slice)
		}
	})
}
