package sword2

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return 2
	}
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}
