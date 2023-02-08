package sorting

func HeapSort(slice []int) []int {
	var root, node, temp int
	size := len(slice)
	i := size / 2
	for {
		if i > 0 {
			i--
			temp = slice[i]
		} else {
			size--
			if size <= 0 {
				break
			}
			temp = slice[size]
			slice[size] = slice[0]
		}
		root = i
		node = i*2 + 1
		for node < size {
			if (node+1 < size) && (slice[node+1] > slice[node]) {
				node++
			}
			if slice[node] > temp {
				slice[root] = slice[node]
				root = node
				node = root*2 + 1
			} else {
				break
			}
		}
		slice[root] = temp
	}

	return slice
}
