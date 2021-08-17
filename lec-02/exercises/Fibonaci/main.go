package main

import "fmt"

func fibonacci(n int) int {

	a := make([]int, n+1)
	if n < 2 {
		return n
	}
	a[0] = 0
	a[1] = 1
	for i := 2; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	return a[n]
}
func main() {
	fmt.Println(fibonacci(10))
}
