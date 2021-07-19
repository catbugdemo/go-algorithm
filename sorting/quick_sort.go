package sorting

// quickSort 采用了二分递归的思想
// 通过交换操作,找到合适的j,保证
// j 左边的元素全都小于 arr[j]
// j 右边的元素全都大于 arr[j]
// 通过双指针i ,j同时设置一个标志数 arr[low]  一般标志值取数组第一个元素
// 当 arr[i] < arr[low] 指针i从左至右 (左边一定小于arr[low],否则交换)
// 当 arr[j] > arr[low] 指针从右至左

// partition 切分操作
