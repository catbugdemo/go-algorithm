package greedy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test455(t *testing.T) {
	t.Run("findContentChildren", func(t *testing.T) {
		g := []int{1, 2}
		s := []int{1, 2, 3}
		count := findContentChildren(g, s)

		assert.Equal(t, 2, count)
	})

	t.Run("findContentChildren", func(t *testing.T) {
		g := []int{10, 9, 8, 7}
		s := []int{5, 6, 7, 8}
		count := findContentChildren(g, s)

		assert.Equal(t, 2, count)
	})

	t.Run("findContentChildren2", func(t *testing.T) {
		g := []int{1, 2}
		s := []int{1, 2, 3}
		count := findContentChildren2(g, s)

		assert.Equal(t, 2, count)
	})
}

func Test135(t *testing.T) {
	t.Run("candy", func(t *testing.T) {
		ints := []int{1, 0, 2} // 1,0,1

		i := candy(ints)

		assert.Equal(t, 5, i)
	})

	t.Run("candy", func(t *testing.T) {
		ints := []int{1, 2, 87, 87, 87, 2, 1} // 0,1,2,0,2,1,0

		i := candy(ints)
		assert.Equal(t, 13, i)
	})

}

func Test435(t *testing.T) {
	t.Run("eraseOverlapIntervals", func(t *testing.T) {
		i := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
		r := eraseOverlapIntervals(i)
		assert.Equal(t, 1, r)
	})

	t.Run("eraseOverlapIntervals", func(t *testing.T) {
		i := [][]int{{1, 2}, {1, 2}, {1, 2}}
		r := eraseOverlapIntervals(i)
		assert.Equal(t, 2, r)
	})
}

func Test605(t *testing.T) {
	t.Run("canPlaceFlowers", func(t *testing.T) {
		i := []int{1, 0, 0, 0, 1}
		flowers := canPlaceFlowers(i, 1)
		assert.Equal(t, true, flowers)
	})
	t.Run("canPlaceFlowers", func(t *testing.T) {
		i := []int{1, 0, 0, 0, 1}
		flowers := canPlaceFlowers(i, 2)
		assert.Equal(t, false, flowers)
	})
}

func Test452(t *testing.T) {
	t.Run("findMinArrowShots", func(t *testing.T) {
		i := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
		shots := findMinArrowShots(i)
		assert.Equal(t, 2, shots)
	})
	t.Run("findMinArrowShots", func(t *testing.T) {
		i := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
		shots := findMinArrowShots(i)
		assert.Equal(t, 4, shots)
	})
	t.Run("findMinArrowShots", func(t *testing.T) {
		i := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
		shots := findMinArrowShots(i)
		assert.Equal(t, 2, shots)
	})
	t.Run("findMinArrowShots", func(t *testing.T) {
		i := [][]int{{1, 2}}
		shots := findMinArrowShots(i)
		assert.Equal(t, 1, shots)
	})
	t.Run("findMinArrowShots", func(t *testing.T) {
		i := [][]int{{2, 3}, {2, 3}}
		shots := findMinArrowShots(i)
		assert.Equal(t, 1, shots)
	})
}

func Test763(t *testing.T) {
	t.Run("partitionLabels", func(t *testing.T) {
		s := "ababcbacadefegdehijhklij"
		labels := partitionLabels(s)
		assert.Equal(t, []int{9, 7, 8}, labels)
	})
	t.Run("partitionLabels2", func(t *testing.T) {
		s := "ababcbacadefegdehijhklij"
		labels := partitionLabels2(s)
		assert.Equal(t, []int{9, 7, 8}, labels)
	})

}

func Test122(t *testing.T) {
	t.Run("maxProfit", func(t *testing.T) {
		i := []int{7, 1, 5, 3, 6, 4}
		profit := maxProfit(i)
		assert.Equal(t, 7, profit)
	})

	t.Run("maxProfit", func(t *testing.T) {
		i := []int{1, 2, 3, 4, 5}
		profit := maxProfit(i)
		assert.Equal(t, 4, profit)
	})
}

func Test406(t *testing.T) {
	t.Run("reconstructQueue", func(t *testing.T) {
		i := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
		queue := reconstructQueue(i)
		assert.Equal(t, [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}, queue)
	})

}

func Test665(t *testing.T) {
	t.Run("checkPossibility", func(t *testing.T) {
		ints := []int{4, 2, 3}
		result := checkPossibility(ints)
		assert.Equal(t, true, result)
	})

	t.Run("checkPossibility", func(t *testing.T) {
		ints := []int{4, 2, 1}
		result := checkPossibility(ints)
		assert.Equal(t, false, result)
	})

	t.Run("checkPossibility", func(t *testing.T) {
		ints := []int{3, 4, 2, 3}
		result := checkPossibility(ints)
		assert.Equal(t, false, result)
	})
}
