package main

import "fmt"

// 归并排序
// 切割合并

// 自顶向下
func MergeSort(array []int, begin int, end int) {
	// 元素数量大于 1 时才进入递归
	if end-begin > 1 {
		// 将数组一分为二，分为 array[begin,mid) 和 [mid,high)
		mid := begin + (end-begin+1)/2

		// 先将左边排序好
		MergeSort(array, begin, mid)

		// 再将右边排序好
		MergeSort(array, mid, end)

		// 两个有序数组进行合并
		merge(array, begin, mid, end)
	}
}

// 归并操作
func merge(array []int, begin, mid, end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := mid - begin         // 左边数组的长度
	rightSize := end - mid          // 右边数组的长度
	newSize := leftSize + rightSize // 辅助数组长度
	result := make([]int, 0, newSize)

	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := array[begin+l] // 左边数组的元素
		rValue := array[mid+r]   // 右边数组的元素

		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else {
			result = append(result, rValue)
			r++
		}
	}

	// 将剩下的元素追加
	result = append(result, array[begin+l:mid]...)
	result = append(result, array[mid+r:end]...)

	// 将辅助元素复制回原数组
	for i := 0; i < newSize; i++ {
		array[begin+i] = result[i]
	}
	return
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	//MergeSort(list, 0, len(list))
	MergeSort2(list, 0, len(list))
	fmt.Println(list)
}

// 自底向上
func MergeSort2(array []int, begin, end int) {
	// 步数为 1 开始
	step := 1

	// 范围大于 step 的数组才可以进入归并
	for end-begin > step {
		// 从头到尾对数组进行归并排序
		for i := begin; i < end; i += step << 1 {
			var lo = i
			var mid = lo + step
			var hi = lo + (step << 1)

			if mid > end {
				return
			}

			if hi > end {
				hi = end
			}

			merge(array, lo, mid, hi)
		}
		// 步长翻倍
		step <<= 1
	}
}

// 改进
// 1.对于小规模数组，使用直接插入排序
// 2.原地排序，节约掉辅助数组空间的占用
