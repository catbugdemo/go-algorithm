package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test215(t *testing.T) {
	t.Run("findKthLargest", func(t *testing.T) {
		ints := []int{3, 2, 1, 5, 6, 4}
		partition(ints,1,5)
	})
}

func Test347(t *testing.T) {
	t.Run("topKFrequent", func(t *testing.T) {
		ints := []int{1, 1, 1, 2, 2, 3}
		frequent := topKFrequent(ints,2)
		assert.Equal(t,[]int{1,2} ,frequent)
	})
}