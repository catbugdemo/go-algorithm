package sword

func myPow(x float64, n int) float64 {
	var result float64 = 1
	if n > 0 {
		for i := 1; i < n+1; i++ {
			result *= x
		}
	} else if n < 0 {
		for ; n < 0; n++ {
			result /= float64(x)
		}
	} else {
		return 1
	}
	return result
}
