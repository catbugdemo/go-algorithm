package main

import "fmt"

func QuickSort(array []int, begin, end int) {
	// 进行切分
	if begin < end {
		loc := partition(array, begin, end)

		// 对左部分进行快排
		QuickSort(array, begin, loc-1)
		// 对右部分进行快排
		QuickSort(array, loc+1, end)
	}
}

// 切分函数，并返回切分元素的下标
func partition(array []int, begin, end int) int {
	i := begin + 1 // 将 array[begin]作为基准数，因此从 array[begin+1] 开始与基准数比较
	j := end

	// 没重合之前
	for i < j {
		if array[i] > array[begin] {
			array[i], array[j] = array[j], array[i] // 交换
			j--
		} else {
			i++
		}
	}

	/* 跳出 for 循环后，i = j。
	 * 此时数组被分割成两个部分   -->  array[begin+1] ~ array[i-1] < array[begin]
	 *                        -->  array[i+1] ~ array[end] > array[begin]
	 * 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置。
	 * 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到最后i = j不满足条件就退出！
	 */
	if array[i] >= array[begin] {
		i--
	}
	array[begin], array[i] = array[i], array[begin]
	return i
}

func main() {
	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(list3, 0, len(list3)-1)
	fmt.Println(list3)
}

func QuickSort1(array []int, begin, end int) {
	if begin < end {
		// 当数组小于 4 时使用直接插入排序
		if end-begin <= 4 {
			InsertSort(array[begin : end+1])
			return
		}

		// 进行切分
		loc := partition(array, begin, end)
		// 对左部分进行快排
		QuickSort1(array, begin, loc-1)
		// 对右部分进行快排
		QuickSort1(array, loc+1, end)
	}
}

// InsertSort 改进：当数组规模小的时候直接插入排序
func InsertSort(list []int) {
	n := len(list)

	// 进行 N-1 轮迭代
	for i := 1; i < n; i++ {
		deal := list[i] // 待排序数
		j := i - 1

		if deal < list[j] {
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1], list[j] = list[j], list[j+1]
			}
			list[j+1] = deal // 结束了，待排序的数插入空位
		}
	}
}

// 尾递归减少程序栈空间占用
func QuickSort3(array []int, begin, end int) {
	for begin < end {
		// 进行切分
		loc := partition(array, begin, end)

		// 哪边元素少先排哪边
		if loc-begin < end-loc {
			QuickSort3(array, begin, loc-1)
			begin = loc + 1
		} else {
			// 先排右边
			QuickSort3(array, loc+1, end)
			end = loc - 1
		}
	}
}
