package binarySearch

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	t.Run("binarySearch1", func(t *testing.T) {
		search1 := binarySearch1([]int{3, 4, 5, 6, 7, 8}, 4)
		assert.Equal(t, 1, search1)
	})

	t.Run("binarySearch2", func(t *testing.T) {
		search2 := binarySearch2([]int{3, 4, 5, 6, 7, 8}, 6)
		assert.Equal(t, 3, search2)
	})
}

func Test69(t *testing.T) {
	t.Run("mySqrt", func(t *testing.T) {
		sqrt := mySqrt(4)
		assert.Equal(t, 2, sqrt)
	})

	t.Run("mySqrt", func(t *testing.T) {
		sqrt := mySqrt(8)
		assert.Equal(t, 2, sqrt)
	})
}

func Test34(t *testing.T) {
	t.Run("searchRange", func(t *testing.T) {
		ints := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
		assert.Equal(t, []int{3, 4}, ints)
	})

	t.Run("searchRange", func(t *testing.T) {
		ints := sort.SearchInts([]int{5, 7, 7, 8, 8, 10}, 11)
		fmt.Println(ints)
	})
}

func Test81(t *testing.T) {
	t.Run("search", func(t *testing.T) {
		r := search([]int{2, 5, 6, 0, 0, 1, 2}, 0)
		assert.Equal(t, true, r)
	})

	t.Run("search2", func(t *testing.T) {
		r := search2([]int{2, 5, 6, 0, 0, 1, 2}, 0)
		assert.Equal(t, true, r)
	})

	t.Run("search2", func(t *testing.T) {
		r := search2([]int{2, 5, 6, 0, 0, 1, 2}, 3)
		assert.Equal(t, false, r)
	})

	t.Run("search2", func(t *testing.T) {
		r := search2([]int{1, 0, 1, 1, 1}, 0)
		assert.Equal(t, true, r)
	})
}

func Test540(t *testing.T) {
	t.Run("singleNonDuplicate", func(t *testing.T) {
		ints := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
		result := singleNonDuplicate(ints)
		assert.Equal(t, 2, result)
	})
	t.Run("singleNonDuplicate", func(t *testing.T) {
		ints := []int{3, 3, 7, 7, 10, 11, 11}
		result := singleNonDuplicate(ints)
		assert.Equal(t, 10, result)
	})
}
func TestName(t *testing.T) {
	t.Run("findMedianSortedArrays", func(t *testing.T) {
		ints := []int{1, 2, 3}
		ints2 := []int{1, 2, 3}
		arrays := findMedianSortedArrays(ints, ints2)
		assert.Equal(t, 2, arrays)
	})
}
