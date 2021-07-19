package util

// BubbingAsc 冒泡排序 asc
// 两相邻数 循环比较
func BubbingAsc(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1] ,arr[j]
			}
		}
	}
	return arr
}
