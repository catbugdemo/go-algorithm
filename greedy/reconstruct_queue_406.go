package greedy

import (
	"sort"
)

//假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
//
//请你重新构造并返回输入数组people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。

//输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
//输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
//解释：
//编号为 0 的人身高为 5 ，没有身高更高或者相同的人排在他前面。
//编号为 1 的人身高为 7 ，没有身高更高或者相同的人排在他前面。
//编号为 2 的人身高为 5 ，有 2 个身高更高或者相同的人排在他前面，即编号为 0 和 1 的人。
//编号为 3 的人身高为 6 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
//编号为 4 的人身高为 4 ，有 4 个身高更高或者相同的人排在他前面，即编号为 0、1、2、3 的人。
//编号为 5 的人身高为 7 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
//因此 [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] 是重新构造后的队列。
//

// reconstructQueue 排序
// 1.根据 people[i][0] 从大到小进行排序
// 2.根据 people[i][1] 从小到大进行排序
// 这样的话能方便后续进行数据插入，数据是 ：[[7 0] [7 1] [6 1] [5 0] [5 2] [4 4]] 身高优先，比重次之
// 3.按照k值插入到index=k的地方，index之后的往后移动
func reconstructQueue(people [][]int) [][]int {
	//根据 身高升序 权重降序
	sort.Slice(people, func(i, j int) bool {
		return (people[i][0] > people[j][0]) || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})

	// [[7 0] [7 1] [6 1] [5 0] [5 2] [4 4]]
	//按照k值插入到index=k的地方，index之后的往后移动
	for i, p := range people {
		index := p[1]
		copy(people[index+1:i+1], people[index:i+1])
		people[index] = p
	}
	return people
}
