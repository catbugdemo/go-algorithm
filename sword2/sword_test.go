package sword2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJz4(t *testing.T) {
	maxtri := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	t.Run("true", func(t *testing.T) {
		array := findNumberIn2DArray2(maxtri, 14)
		assert.True(t, array)
	})
	t.Run("false", func(t *testing.T) {
		array := findNumberIn2DArray(maxtri, 31)
		assert.False(t, array)
	})
	t.Run("1", func(t *testing.T) {
		m := [][]int{{1, 1}}
		array := findNumberIn2DArray(m, 2)
		assert.False(t, array)
	})
}

func TestLeet41(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		nums := []int{0, 1, 2}
		positive := firstMissingPositive(nums)
		assert.Equal(t, 3, positive)
	})
}

func TestJz7(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	})
}
