package sword

import "sort"

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func getLeastNumbers2(arr []int, k int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr[:k]
}

func quickSort(arr []int, begin, end int) {
	if begin < end {
		loc := partition(arr, begin, end)

		// 左排，右排
		quickSort(arr, begin, loc-1)
		quickSort(arr, loc+1, end)
	}
}

// 切分函数
func partition(array []int, begin, end int) int {
	// 设置左，右移动
	i, j := begin+1, end

	// 左右切分
	for i < j {
		// 如果左边点大于比较点，array[i] 交换
		if array[i] > array[begin] {
			array[i], array[j] = array[j], array[i]
			j--
		} else {
			i++
		}
	}

	// 此时i==j，比较完成，分为两个范围
	// array[begin+1]~ array[i-1] < array[begin]
	// array[begin] < array[i+1]~ array[end]
	// 比较 array[i] 和  array[begin]
	if array[i] >= array[begin] {
		i--
	}
	array[i], array[begin] = array[begin], array[i]
	return i
}
