package main

import "fmt"

// 插入排序算法
// 这种排序算法是稳定的

func InsertSort(list []int) {
	n := len(list)

	// 进行 N-1 轮迭代
	for i := 1; i < n; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置

		// 如果第一次比较比左边的已排好的第一个数小，那么进入处理
		if deal < list[j] {
			// 一直往左边找，比待排序大的数都往后挪，腾空给待排序插入
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j] // 位置后移
			}
			// 结束了，插入
			list[j+1] = deal
		}
	}
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	InsertSort(list)
	fmt.Println(list)
}

// 数组中有序性越高，插入排序的新能越高，因为待排序数组有序性越高
// 很少使用冒泡，直接选择，直接插入排序算法，因为在有大量元素的无序数列下，这些算法的效率都很低
