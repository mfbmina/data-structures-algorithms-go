package sorting

func InsertSort(slice []int) []int {
	size := len(slice)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if slice[i] > slice[j] {
				aux := slice[i]
				slice[i] = slice[j]
				slice[j] = aux
			}
		}
	}

	return slice
}
