package sword

func cuttingRope_2(n int) int {
	dp := make([]int, n+1)
	// 当绳子长为 0 ,1, 2的情况
	dp[0], dp[1], dp[2] = 0, 1, 1
	var tmp1, tmp2 int
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1]
		for j := 0; j < i/2; j++ {
			tmp1 = max(dp[j], j)
			tmp2 = max(dp[i-j], i-j)
			dp[i] = max(dp[i], tmp1*tmp2)
		}
	}
	return dp[n]
}
