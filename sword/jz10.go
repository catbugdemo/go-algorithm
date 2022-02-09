package sword

func fib(n int) int {
	if n == 0 {
		return 1
	}
	if n == 2 || n == 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % (1e9 + 7)
	}
	return dp[n]
}
