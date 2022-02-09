package sword

import "math"

// 1.可以把这个题转换成一个二维分割
// 2.即找到在 n 范围内能相乘能得到的最大 x,y
// 3.x,y 相乘获得最大乘积
// 4.同时也有规律，即必有两个数相等
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	// 当绳子长为 0 ,1 , 2 的情况
	dp[0], dp[1], dp[2] = 0, 1, 1
	var tmp1, tmp2 int
	for i := 3; i < n+1; i++ {
		// 将绳子分为两部分，两部分自由组合，选择较大的两部分
		dp[i] = dp[i-1]
		for j := 1; j <= i/2; j++ {
			tmp1 = max(dp[j], j)
			tmp2 = max(dp[i-j], i-j)
			dp[i] = max(dp[i], tmp1*tmp2)
		}
	}
	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func cuttingRope2(n int) int {
	if n <= 2 {
		return 1
	}
	if n <= 3 {
		return 2
	}
	// 能分成几份
	parts := n / 3
	another := n % 3
	var res float64
	switch another {
	case 2:
		res = math.Pow(3, float64(parts))
		res *= 2
	case 1:
		res = math.Pow(3, float64(parts-1))
		res *= 4
	default:
		res = math.Pow(3, float64(parts))
	}
	return int(res)
}
