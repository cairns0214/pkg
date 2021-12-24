package slice

import (
	"fmt"
	"testing"
)

func TestUniqInt(t *testing.T) {
	arr := []int{1, 1, 1, 2, 2, 3, 4, 5, 5, 6, 6, 6, 7, 7, 7}
	res := UniqInt(arr)
	fmt.Println(res)
}

func TestUniqInt64(t *testing.T) {
	arr := []int64{1, 1, 1, 2, 2, 3, 4, 5, 5, 6, 6, 6, 7, 7, 7}
	res := UniqInt64(arr)
	fmt.Println(res)
}

func TestUniqString(t *testing.T) {
	arr := []string{"aaa", "bbb", "aa", "b", "aa", "aab", "aab", "acc"}
	res := UniqString(arr)
	fmt.Println(res)
}
