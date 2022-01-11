package search

import (
	"fmt"
	"testing"
)

func Test46(t *testing.T) {
	t.Run("permute", func(t *testing.T) {
		r := permute([]int{1, 2, 3})
		fmt.Println(r)
	})
}

func Test77(t *testing.T) {
	t.Run("combine", func(t *testing.T) {
		r := combine(1, 1)
		fmt.Println(r)
	})

	t.Run("combine", func(t *testing.T) {
		r := combine(4, 2)
		fmt.Println(r)
	})
}

func Test79(t *testing.T) {
	t.Run("exist", func(t *testing.T) {
		i := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
		b := exist(i, "ABCCE")
		fmt.Println(b)
	})
}
