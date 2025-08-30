package main

import "fmt"

// 地下城房间开锁 房间内价值相同的宝物
// 进入房间后只能进入level更高的房间
// fmt.Printf 输出得到最多的宝物个数
// eg input: 4 7 5 19 5 6 3 1 100 2
// output: 4
// eg input: 5 5 5 5 5
// output: 1

func main() {
	num := 0
	levels := make([]int, 0)
	for {
		n, err := fmt.Scan(&num)
		if n == 0 || err != nil {
			break
		} else {
			levels = append(levels, num)
		}
	}
	res := process(levels)
	fmt.Printf("%d\n", res)
}

func process(levels []int) (res int) {
	n := len(levels)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if levels[j] < levels[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return
}
