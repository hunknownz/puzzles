package fibonacciNumber

func fib(n int) int {
	//return fibRecursion(n)
	return fibLoop(n)
}

func fibLoop(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func fibRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursion(n-1) + fibRecursion(n-2)
}
