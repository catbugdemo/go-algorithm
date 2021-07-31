package Pointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test167(t *testing.T) {
	t.Run("twoSun", func(t *testing.T) {
		ints := []int{2, 7, 11, 15}
		sum := twoSum(ints, 9)
		assert.Equal(t, []int{1,2},sum)
	})
	t.Run("twoSun", func(t *testing.T) {
		ints := []int{2,3,4}
		sum := twoSum(ints, 6)
		assert.Equal(t, []int{1,3},sum)
	})

	t.Run("twoSun", func(t *testing.T) {
		ints := []int{-1,0}
		sum := twoSum(ints, -1)
		assert.Equal(t, []int{1,2},sum)
	})

}

func Test88(t *testing.T) {
	t.Run("merge", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2,5,6}
		merge(nums1,3,nums2,3)
	})

	t.Run("merge", func(t *testing.T) {
		nums1 := []int{1}
		nums2 := []int{}
		merge(nums1,1,nums2,0)
	})
}

func Test142(t *testing.T) {
	t.Run("deleteCycle", func(t *testing.T) {
	})
}

func Test76(t *testing.T) {
	t.Run("minWindow", func(t *testing.T) {
		window := minWindow("ADOBECODEBANC", "ABC")
		assert.Equal(t, "BANC",window)
	})
	t.Run("minWindow", func(t *testing.T) {
		window := minWindow("a", "a")
		assert.Equal(t, "a",window)
	})
	t.Run("minWindow", func(t *testing.T) {
		window := minWindow("ab", "a")
		assert.Equal(t, "a",window)
	})
	t.Run("minWindow", func(t *testing.T) {
		window := minWindow("bbaa", "aba")
		assert.Equal(t, "baa",window)
	})}

func Test633(t *testing.T) {
	t.Run("judgeSquareSum", func(t *testing.T) {
		check := judgeSquareSum(5)
		assert.Equal(t, true,check)
	})
	t.Run("judgeSquareSum", func(t *testing.T) {
		check := judgeSquareSum(3)
		assert.Equal(t, false,check)
	})
	t.Run("judgeSquareSum", func(t *testing.T) {
		check := judgeSquareSum(4)
		assert.Equal(t, true,check)
	})
	t.Run("judgeSquareSum", func(t *testing.T) {
		check := judgeSquareSum(2)
		assert.Equal(t, true,check)
	})
	t.Run("judgeSquareSum", func(t *testing.T) {
		check := judgeSquareSum(1)
		assert.Equal(t, true,check)
	})
}

func Test680(t *testing.T) {
	t.Run("validPalindrome", func(t *testing.T) {
		result := validPalindrome("aba")
		assert.Equal(t, true,result)
	})
}

func Test340(t *testing.T) {
	t.Run("lengthOfLongestSubstringKDistinct", func(t *testing.T) {
		result := lengthOfLongestSubstringKDistinct("eceba", 2)
		assert.Equal(t, 3,result)
	})

	t.Run("lengthOfLongestSubstringKDistinct", func(t *testing.T) {
		result := lengthOfLongestSubstringKDistinct("aa", 1)
		assert.Equal(t, 2,result)
	})
	t.Run("lengthOfLongestSubstringKDistinct", func(t *testing.T) {
		result := lengthOfLongestSubstringKDistinct("a", 1)
		assert.Equal(t, 1,result)
	})
	t.Run("lengthOfLongestSubstringKDistinct", func(t *testing.T) {
		result := lengthOfLongestSubstringKDistinct("aa", 2)
		assert.Equal(t, 2,result)
	})
}

func Test189(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		rotate([]int{1,2,3,4,5,6,7},3)
	})
}

func Test283(t *testing.T) {
	t.Run("move_zeroes", func(t *testing.T) {
		moveZeroes([]int{0,1,0,3,12})
	})
}