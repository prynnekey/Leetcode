package main

// 统计n以内素数的个数
//
// 暴力解法
func bf(n int) int {
	count := 0

	for i := 1; i <= n; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

// 判断num是否是素数
func isPrime(num int) bool {
	if num == 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// 埃筛法
//
// 如果找到了一个质数，那么它的倍数一定不是质数。
//
// 如 2 为质数 则 2*2=4 2*3=6 ... 都不是质数
func eratosthenes(n int) int {
	count := 0

	// true  表示不是质数
	// false 表示是质数
	isPrime := make([]bool, n)

	for i := 2; i < n; i++ {
		if !isPrime[i] {
			// i是质数
			count++

			// 则i的倍数一定不是质数
			for j := i; i*j < n; j++ {
				isPrime[i*j] = true
			}
		}
	}

	return count
}

/* func main() {
	num := 10000000

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		t := time.Now()
		fmt.Printf("bf(): %v\n", bf(num))
		d := time.Since(t)
		fmt.Printf("暴力解法耗时: %v\n", d)

		wg.Done()
	}()

	go func() {
		t := time.Now()
		fmt.Printf("eratosthenes(): %v\n", eratosthenes(num))
		d := time.Since(t)
		fmt.Printf("埃筛法耗时: %v\n", d)

		wg.Done()
	}()

	wg.Wait()
} */
