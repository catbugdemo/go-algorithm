package Greedy

//假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
//
//给你一个整数数组flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数n ，能否在不打破种植规则的情况下种入n朵花？能则返回 true ，不能则返回 false。
//
//输入：flowerbed = [1,0,0,0,1], n = 1
//输出：true

//	canPlaceFlowers 判断给的花 n ,能否符合条件
//	从0开始进行种花
//	循环
//	先判断是否存在
//	判断条件	if flowerbed[i-1] ==0 && flowerbed[i]==0 && flowerbed[i+1] ==0
func canPlaceFlowers(flowerbed []int, n int) bool {
	//种植数为0时返回true
	if n == 0 {
		return true
	}
	//判断
	if len(flowerbed) == 1 {
		if n == 1 && flowerbed[0] == 0 {
			return true
		}
		return false
	}

	//从零开始进行种花
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		//最左边和最右边单独判断
		if i == 0 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[0] = 1
				count++
			}
			continue
		}
		//右边
		if i == len(flowerbed)-1 {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 {
				flowerbed[i] = 1
				count++
			}
			continue
		}

		//中间判断
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			count++
		}
	}

	result := false
	if count >= n {
		result = true
	}
	return result
}
