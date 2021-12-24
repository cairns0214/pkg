package slice

// UniqInt Remove duplicate elements in int slice
func UniqInt(arr []int) []int {
	result := make([]int, 0, len(arr))
	temp := make(map[int]struct{}, len(arr))
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// UniqInt64 Remove duplicate elements in int64 slice
func UniqInt64(arr []int64) []int64 {
	result := make([]int64, 0, len(arr))
	temp := make(map[int64]struct{}, len(arr))
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// UniqString Remove duplicate elements in string slice
func UniqString(arr []string) []string {
	result := make([]string, 0, len(arr))
	temp := make(map[string]struct{}, len(arr))
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
