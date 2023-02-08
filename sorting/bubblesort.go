package sorting

func BubbleSort(slice []int) []int {
	size := len(slice)
	swap := true
	for swap {
		swap = false
		for i := 0; i < size-1; i++ {
			if slice[i] > slice[i+1] {
				aux := slice[i]
				slice[i] = slice[i+1]
				slice[i+1] = aux
				swap = true
			}
		}
	}

	return slice
}
