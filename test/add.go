package test

func sum(n int) (res int) {
	for i := 1; i <= n; i++ {
		res = res + i
	}
	return res
}
