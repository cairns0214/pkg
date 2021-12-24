package slice

// QuickSortIntSlice quick sort for int slice
func QuickSortIntSlice(arr []int) {
	quickSortIntSlice(arr, 0, len(arr)-1)
}

func quickSortIntSlice(arr []int, left, right int) {
	if left >= right {
		return
	}

	key := arr[(left+right)/2]
	i, j := left, right
	for {
		for arr[i] < key {
			i++
		}
		for arr[j] > key {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	quickSortIntSlice(arr, left, i-1)
	quickSortIntSlice(arr, j+1, right)
}

// QuickSortInt64Slice quick sort for int64 slice
func QuickSortInt64Slice(arr []int64) {
	quickSortInt64Slice(arr, 0, len(arr)-1)
}

func quickSortInt64Slice(arr []int64, left, right int) {
	if left >= right {
		return
	}

	key := arr[(left+right)/2]
	i, j := left, right
	for {
		for arr[i] < key {
			i++
		}
		for arr[j] > key {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	quickSortInt64Slice(arr, left, i-1)
	quickSortInt64Slice(arr, j+1, right)
}
