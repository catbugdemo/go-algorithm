package main

import "fmt"

// 自底向上构建最大堆，再移除堆元素实现堆排序
func HeapSort(array []int) {
	// 堆的元素数量
	count := len(array)

	// 最底层的叶子节点下标，该节点位置不定，但是该叶子节点右边的节点都是叶子节点
	start := count/2 + 1

	// 最后的元素下标
	end := count - 1

	// 从最底层开始，逐一对节点进行下沉
	for start >= 0 {
		sift(array, start, count)
		start--
	}

	// 下沉结束，开始排序
	for end > 0 {
		// 将堆顶元素与堆尾元素互换，表示移除最大堆元素
		array[end], array[0] = array[0], array[end]
		sift(array, 0, end)
		end--
	}
}

// 下沉操作，需要下沉的元素时 array[start],参数 count 只要用来判断是否到底堆底
func sift(array []int, start, count int) {
	// 父亲节点
	root := start

	// 左儿子
	child := root*2 + 1

	// 如果有下一代
	for child < count {
		// 右儿子比左儿子大，那么要翻转的儿子改为右儿子
		if count-child > 1 && array[child] < array[child+1] {
			child++
		}

		// 父亲节点比儿子小，那么将父亲和儿子位置交换
		if array[root] < array[child] {
			array[root], array[child] = array[child], array[root]
			root = child
			child = root*2 + 1
		} else {
			return
		}
	}
}
func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}

	HeapSort(list)

	// 打印排序后的值
	fmt.Println(list)
}
