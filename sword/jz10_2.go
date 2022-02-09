package sword

const mod = 1e9 + 7

func numWays(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return (numWays(n-1) + numWays(n-2)) % mod
}
