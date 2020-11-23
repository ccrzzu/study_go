package dynamic_programming

func Fibonacci(n int) int {
	/*memo := make(map[int]int)
	return fibonacciByDG(memo, n)*/
	return fibonacciByIter(n)
}

func fibonacciByDG(memo map[int]int, n int) int {
	if res, ok := memo[n]; ok {
		return res
	}
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	memo[n] = fibonacciByDG(memo, n-1) + fibonacciByDG(memo, n-2)
	return memo[n]
}

func fibonacciByIter(n int) int {
	//base case
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	prev := 0
	curr := 1
	for i := 3; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

func fibonacciByIter2(n int) int {
	//base case
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	result := make([]int,n+1)
	result[1]=0
	result[2]=1
	for i:=3;i<=n;i++{
		result[i] = result[i-1] + result[i-2]
	}
	return result[n]
}
