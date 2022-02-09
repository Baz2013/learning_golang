package leetcode_cn

import (
	"fmt"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	data := []int{-4, -1, 0, 3, 10}
	r := sortedSquares(data)
	fmt.Println(r)

	data1 := []int{-7, -3, 2, 3, 11}
	r1 := sortedSquares(data1)
	fmt.Println(r1)
}
