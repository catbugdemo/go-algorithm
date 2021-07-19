package Greedy

//老师想给孩子们分发糖果，有N个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
//
//你需要按照以下要求，帮助老师给这些孩子分发糖果：
//
//每个孩子至少分配到 1 个糖果。
//评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。

//输入：[1,0,2]
//输出：5
//解释：你可以分别给这三个孩子分发 2、1、2 颗糖果。

// candy
// 1.数据初始化
// 2.左遍历 -- 从 1 开始
// 3.如果 ratings[i] > ratings[i-1] && ints[i] <= ints[i-1] : ints[i]=ints[i-1]+1 (全部都设置为ints[i]+1)
// 4.右遍历 -- 从 len -2 开始
// 5.如果 ratings[i] > ratings[i+1] && ints[i] <= ints[i+1] : ints[i]=ints[i+1]+1
// 6.总和
func candy(ratings []int) (ans int) {
	ints := make([]int, len(ratings))
	//初始化数据
	for k, _ := range ratings {
		ints[k] = 1
	}

	//从左开始遍历
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] && ints[i] <= ints[i-1] {
			ints[i] = ints[i-1] + 1
		}
	}

	//从右开始遍历
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && ints[i] <= ints[i+1] {
			ints[i]=ints[i+1]+1
		}
	}
	//总和
	count := 0
	for _, v := range ints {
		count+= v
	}

	return count
}


