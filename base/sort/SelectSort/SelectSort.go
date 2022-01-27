package main

import "fmt"

// 选择排序 -- 每次扫描数列找出最小的数，然后与第一个数交换
// 排除第一个数 从第二个数开始

func SelectSort(list []int) {
	n := len(list)

	// 进行 N -1 次迭代
	for i := 0; i < n-1; i++ {
		min, minIndex := list[i], i
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				min, minIndex = list[j], j
			}
		}

		// 找到最小下标
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

// 算法优化，同时找到最大值
func SelectGoodSort(list []int) {
	n := len(list)

	// 只需要循环一半
	for i := 0; i < n/2; i++ {
		minIndex, maxIndex := i, i
		for j := i + 1; j < n-i; j++ {
			// 找到最大
			if list[j] > list[maxIndex] {
				maxIndex = j
			}
			// 找到最小
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}

		// 完成循环排序
		if maxIndex == i && minIndex != n-i-1 {
			// 如果最大值是开头，最小值不是结尾
			// 现将最大值和尾部交换，然后最小放在开头
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
			list[i], list[minIndex] = list[minIndex], list[i]
		} else if maxIndex == i && minIndex == n-i-1 {
			// 最大值在开头，最小值在结尾,直接交换
			list[minIndex], list[maxIndex] = list[maxIndex], list[minIndex]
		} else {
			// 否则交换
			list[i], list[minIndex] = list[minIndex], list[i]
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
		}
	}
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectGoodSort(list)
	fmt.Println(list)
}
