package sort

func mergeSort(arr []int) []int {
	size := len(arr)

	if size == 1 {
		return arr
	}

	lArr := mergeSort(arr[:size/2])
	rArr := mergeSort(arr[size/2:])

	lSize := len(lArr)
	rSize := len(rArr)

	i, j := 0, 0

	var merged []int

	for k := 0; k < size; k++ {
		if i == lSize {
			merged = append(merged, rArr[j])
			j++
		} else if j == rSize {
			merged = append(merged, lArr[i])
			i++
		} else {
			if lArr[i] <= rArr[j] {
				merged = append(merged, lArr[i])
				i++
			} else {
				merged = append(merged, rArr[j])
				j++
			}
		}
	}

	return merged
}
