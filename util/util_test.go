package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbingAsc(t *testing.T) {
	t.Run("冒泡排序：asc", func(t *testing.T) {
		ints := []int{5, 8, 6, 3, 9, 2, 1, 7}

		asc := BubbingAsc(ints)

		assert.Equal(t, 1,asc[0])
		fmt.Println(asc)
	})
}
