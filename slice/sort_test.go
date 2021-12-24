package slice

import (
	"fmt"
	"testing"
)

func TestQuickSortIntSlice(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	QuickSortIntSlice(arr)
	fmt.Println(arr)
}

func TestQuickSortInt64Slice(t *testing.T) {
	arr := []int64{5, 4, 3, 2, 1}
	QuickSortInt64Slice(arr)
	fmt.Println(arr)
}
