package BinarySearch

//二分查找只有两种写法

//binarySearch1 写法1：[left,right]
//right是个合法值
//会带来相应的影响
//right 初始值 len(nums) -1
//循环条件是 left < = right ,因为 left == right 中间还能取到
//当 nums[mid] > target, nums[mid]已经判断了 , right = mid -1
func binarySearch1(nums []int, target int) int {
	//1.赋值
	l, r := 0, len(nums)-1
	// 因为right是一个合法的数据
	for l <= r {
		mid := l + (r-l)>>1
		//中间数大于目标数据，目标数在左边，将右边转为mid，mid已经被比较过了
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target { //中间数小于目标函数，目标数据在右边，将左边转为mid,mid已经被比较过了
			l = mid + 1
		} else {
			return mid
		}
	}
	//这个l是target放入后, nums仍能保持顺序的位置
	return l
}

//binarySearch1 写法2：[left,right)
// right是不合法数据,例如[1,1)中间已经没有数了
// mid 虽然已经被判断,但是nums[right]娶不到
func binarySearch2(nums []int, target int) int {
	//1.赋值也有区别,right是一个取不到的数据
	l, r := 0, len(nums)
	//2.循环条件同样不同
	for l < r {
		//中间值赋值相同
		mid := l + (r-l)>>1
		//中间数大于目标数据，目标数在左边，将右边转为mid,中间数取不到nums[right]，所以赋值为mid
		if nums[mid] > target {
			r = mid
		} else if nums[mid] < target { //左边取得到所以要+1
			l = mid + 1
		} else {
			return mid
		}
	}
	return l
}
